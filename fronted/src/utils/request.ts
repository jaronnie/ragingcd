// 二次封装 axios
import axios from "axios";
import { ElMessage } from "element-plus";
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
    return response.data;
  },
  (error) => {
    // 失败回调
    let message = "";
    const status = error.response.status;
    switch (status) {
      case "404":
        message = "请求地址错误";
        break;
      default:
        message = "网络错误";
        break;
    }
    // 提示错误消息
    ElMessage({
      type: "error",
      message: message,
    });
    return Promise.reject(error);
  },
);

// 对外暴露
export default request;
