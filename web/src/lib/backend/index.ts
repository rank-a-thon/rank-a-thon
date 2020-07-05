import axios from 'axios';
import { getMe, saveMe } from '../../data/me';
import Router from 'next/router';

let API_URL = process.env.API_URL;

function getFullEndpoint(name: string): string {
  if (window !== undefined) {
    // Client only code
    API_URL = `${location.protocol}//${window.location.host}/api/`;
  }
  return `${API_URL}${name}`;
}

export async function makeBackendRequest(
  method: 'post' | 'get' | 'put' | 'patch' | 'delete',
  endpoint: string,
  data?: any,
) {
  return await axios({
    method: method,
    url: getFullEndpoint(endpoint),
    data: data,
  });
}

export async function makeAuthedBackendRequest(
  method: 'post' | 'get' | 'put' | 'patch' | 'delete',
  endpoint: string,
  data?: any,
) {
  const me = getMe();
  const access_token = me.access_token;
  try {
    const response = await axios({
      method: method,
      headers: { Authorization: `Bearer ${access_token}` },
      url: getFullEndpoint(endpoint),
      data: data,
    });
    return response;
  } catch (err) {
    if (err.response.status === 400) {
      // Auth expiry - refresh token
      const refresh_token = me.refresh_token;
      try {
        const response = await axios({
          method: 'post',
          headers: { Authorization: `Bearer ${refresh_token}` },
          url: getFullEndpoint('v1/token/refresh'),
          data: { refresh_token: refresh_token },
        });
        console.log(response);
        saveMe(response.data);
        // TODO: we can do better
        return makeAuthedBackendRequest(method, endpoint, data);
      } catch {
        throw 'Unable to refresh JWT token';
      }
    } else if (err.response.status === 401) {
      Router.push('/login');
    }
  }
}
