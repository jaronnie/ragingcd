// 二次封装 axios
import axios, { AxiosError } from "axios";
import { GET_TOKEN } from "./token";

const request = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_API,
  timeout: 5000, // 5s
});

export const tokenPrefix = "Bearer";

// 添加请求与响应拦截器
request.interceptors.request.use((config) => {
  config.headers.set("Authorization", tokenPrefix + " " + GET_TOKEN());
  return config;
});

request.interceptors.response.use(
  (response) => {
    console.log(response.data);
    return response.data;
  },
  (error) => {
    // 失败回调
    let message = "";
    const status = error.response.status;
    switch (status) {
      case 404:
        message = "接口 404 异常";
        break;
      default:
        message = "后端服务异常";
        break;
    }
    return Promise.reject(new AxiosError(message));
  },
);

// 对外暴露
export default request;
