import axios, { AxiosInstance } from "axios";

const defaultCallApiTimeout = 1000 * 30;

const host = "http://localhost:8080";

const client: AxiosInstance = axios.create({
  baseURL: `${host}/api`,
  headers: {
    "Content-Type": "application/json;charset=UTF-8",
  },
  timeout: defaultCallApiTimeout,
});

export default client;
