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
  const { access_token, refresh_token } = getMe();
  try {
    const response = await axios({
      method: method,
      headers: { Authorization: `Bearer ${access_token}` },
      url: getFullEndpoint(endpoint),
      data: data,
    });
    return response;
  } catch (err) {
    if (err.response.status === 401) {
      // Auth expiry - refresh token
      try {
        const response = await axios({
          method: 'post',
          headers: { Authorization: `Bearer ${refresh_token}` },
          url: getFullEndpoint('v1/token/refresh'),
          data: { refresh_token },
        });
        console.log(response);
        saveMe(response.data);
        // TODO: we can do better
        return makeAuthedBackendRequest(method, endpoint, data);
      } catch {
        throw err;
      }
    } else if (err.response.status === 401) {
      Router.push('/login');
      return;
    }
    // Rethrow error so that call site can handle
    throw err;
  }
}

export async function sendAuthedFormData(
  method: 'post' | 'get' | 'put' | 'patch' | 'delete',
  endpoint: string,
  formData: any,
) {
  const { access_token, refresh_token } = getMe();
  try {
    const response = await axios({
      method: method,
      headers: {
        Authorization: `Bearer ${access_token}`,
        'content-type': 'multipart/form-data',
      },
      url: getFullEndpoint(endpoint),
      data: formData,
    });
    return response;
  } catch (err) {
    console.log(err.response);
    if (err.response.status === 400) {
      // Auth expiry - refresh token
      try {
        const response = await axios({
          method: 'post',
          headers: { Authorization: `Bearer ${refresh_token}` },
          url: getFullEndpoint('v1/token/refresh'),
          data: { refresh_token },
        });
        console.log(response);
        saveMe(response.data);
        // TODO: we can do better
        return sendAuthedFormData(method, endpoint, formData);
      } catch {
        throw err;
      }
    } else if (err.response.status === 401) {
      Router.push('/login');
      return;
    }
    // Rethrow error so that call site can handle
    throw err;
  }
}
