import { name } from "../../package.json";

const accessTokenKey = `${name}-accessToken`;
const refreshTokenKey = `${name}-refreshToken`;

export function getLocalStorageAccessToken() {
  return localStorage.getItem(accessTokenKey);
}
export function setLocalStorageAccessToken(token: string) {
  localStorage.setItem(accessTokenKey, token);
}
export function delLocalStorageAccessToken() {
  localStorage.removeItem(accessTokenKey);
}

export function getLocalStorageRefreshToken() {
  return localStorage.getItem(refreshTokenKey);
}
export function setLocalStorageRefreshToken(token: string) {
  localStorage.setItem(refreshTokenKey, token);
}
export function delLocalStorageRefreshToken() {
  localStorage.removeItem(refreshTokenKey);
}
