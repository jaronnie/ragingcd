import request from "@/utils/request";
import {
  loginForm,
  loginResponseData,
  userInfoResponseData,
  userListResponseData,
} from "./type";
import type { pageQuery } from "@/api/type.ts";

enum API {
  LOGIN_URL = "/v1.0/user/login",
  USER_INFO_URL = "/v1.0/user/info",
  USER_LIST_URL = "/v1.0/user/list",
}

export const reqLogin = (data: loginForm) =>
  request.post<any, loginResponseData>(API.LOGIN_URL, data);

export const reqUserInfo = () =>
  request.get<any, userInfoResponseData>(API.USER_INFO_URL);

export const reqUserList = (params: pageQuery) =>
  request.get<any, userListResponseData>(API.USER_LIST_URL, { params });
