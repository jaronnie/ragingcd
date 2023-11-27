import request from "@/utils/request";
import { loginForm, loginResponseData, userInfo } from "./type";

enum API {
  LOGIN_URL = "/user/login",
  USERINFO_URL = "/user/info",
}

export const reqLogin = (data: loginForm) =>
  request.post<any, loginResponseData>(API.LOGIN_URL, data);

export const reqUserInfo = () => request.get<any, userInfo>(API.USERINFO_URL);
