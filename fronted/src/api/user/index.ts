import request from "@/utils/request";
import {
  addUserBo,
  loginForm,
  loginResponseData,
  userInfoResponseData,
  userListResponseData,
} from "./type";
import type { pageQuery } from "@/api/type.ts";

enum API {
  LOGIN_URL = "/v1.0/user/login",
  LOGOUT_URL = "/v1.0/user/logout",
  USER_INFO_URL = "/v1.0/user/info",
  USER_LIST_URL = "/v1.0/user/list",
  USER_ADD_URL = "/v1.0/user/add",
}

export const reqLogin = (data: loginForm) =>
  request.post<any, loginResponseData>(API.LOGIN_URL, data);

export const reqLogout = () => request.get<any, any>(API.LOGOUT_URL);

export const reqUserInfo = () =>
  request.get<any, userInfoResponseData>(API.USER_INFO_URL);

export const reqUserList = (params: pageQuery) =>
  request.get<any, userListResponseData>(API.USER_LIST_URL, { params });

export const reqUserAdd = (data: addUserBo) =>
  request.post<any, userListResponseData>(API.USER_ADD_URL, data);
