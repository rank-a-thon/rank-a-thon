import axios from 'axios';

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
