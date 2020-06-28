import axios from 'axios';

const API_URL = process.env.API_URL;

function getFullEndpoint(name: string): string {
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
