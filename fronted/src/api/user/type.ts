export interface loginForm {
  username?: string;
  password?: string;
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

export interface userInfo {
  userId?: number;
  avatar?: string;
  username?: string;
}
