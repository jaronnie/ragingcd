import axios from "axios";

const request_github = axios.create({
  timeout: 5000, // 5s
});

request_github.interceptors.response.use((response) => {
  return response.data;
});
// 对外暴露
export default request_github;
