import request from "@/utils/request";
import {
  AddUserBo,
  LoginBo,
  LoginVoResponseData,
  PublicKeyVoResponseData,
  UserVoResponseData,
  UserListVoResponseData,
  RegisterUserBo,
} from "./type";
import type { pageQuery } from "@/api/type.ts";

enum API {
  LOGIN_URL = "/v1.0/system/user/login",
  LOGOUT_URL = "/v1.0/system/user/logout",
  USER_INFO_URL = "/v1.0/system/user/info",
  USER_LIST_URL = "/v1.0/system/user_manage/list",
  USER_ADD_URL = "/v1.0/system/user_manage/add",
  PUBLIC_KEY_URL = "/v1.0/system/user/public-key",
  USER_DELETE_URL = "/v1.0/system/user_manage/delete",
  USER_REGISTER_URL = "/v1.0/system/user/register",
  USER_REGISTER_SEND_MAIL = "/v1.0/system/user/send_email",
}

export const reqLogin = (data: LoginBo) =>
  request.post<any, LoginVoResponseData>(API.LOGIN_URL, data);

export const reqLogout = () => request.get<any, any>(API.LOGOUT_URL);

export const reqUserInfo = () =>
  request.get<any, UserVoResponseData>(API.USER_INFO_URL);

export const reqPublicKey = () =>
  request.get<any, PublicKeyVoResponseData>(API.PUBLIC_KEY_URL);

export const reqUserList = (params: pageQuery) =>
  request.get<any, UserListVoResponseData>(API.USER_LIST_URL, { params });

export const reqUserAdd = (data: AddUserBo) =>
  request.post<any, UserListVoResponseData>(API.USER_ADD_URL, data);

export const reqUserRegister = (data: RegisterUserBo) =>
  request.post<any, UserVoResponseData>(API.USER_REGISTER_URL, data);

export const reqUserDelete = (userId: number) =>
  request.get<any, boolean>(`${API.USER_DELETE_URL}/${userId}`);

export const reqUserRegisterSendMail = (mail: string, username: string) =>
  request.get<boolean>(`${API.USER_REGISTER_SEND_MAIL}/${mail}`, {
    params: { username },
  });
