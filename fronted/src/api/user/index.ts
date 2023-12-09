import request from "@/utils/request";
import { loginForm, loginResponseData, userInfoResponseData, userListResponseData } from "./type";

enum API {
  LOGIN_URL = "/v1.0/user/login",
  USERINFO_URL = "/v1.0/user/info",
  USERLIST_URL = "/v1.0/user/list",
}

export const reqLogin = (data: loginForm) =>
  request.post<any, loginResponseData>(API.LOGIN_URL, data);

export const reqUserInfo = () =>
  request.get<any, userInfoResponseData>(API.USERINFO_URL);

export const reqUserList = () =>
  request.get<any, userListResponseData>(API.USERLIST_URL);
