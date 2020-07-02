import axios from 'axios';
import { getMe } from '../../data/me';

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
  const access_token = getMe().access_token;
  return await axios({
    method: method,
    headers: { Authorization: `Bearer ${access_token}` },
    url: getFullEndpoint(endpoint),
    data: data,
  });
}
