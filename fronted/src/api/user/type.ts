import { OssFileVo } from "@/api/type.ts";

export interface loginForm {
  username?: string;
  password: string;
}

export interface addUserBo {
  username?: string;
  password?: string;
  avatar?: OssFileVo;
}

interface dataType {
  token?: string;
}

export interface loginResponseData {
  code?: number;
  message?: string;
  data?: dataType;
}

export interface userInfoResponseData {
  code?: number;
  message?: string;
  data?: userInfo;
}

export interface userListResponseData {
  code?: number;
  message?: string;
  data?: userPageResult;
}

export interface userPageResult {
  total?: number;
  rows?: userInfo[];
}

export interface userInfo {
  userId?: number;
  avatar?: string;
  username?: string;
}

export interface publicKeyResponseData {
  code?: number;
  message?: string;
  data?: publicKeyVo;
}

export interface publicKeyVo {
  type: string;
  publicKey: string;
}
