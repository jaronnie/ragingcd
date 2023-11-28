export interface loginForm {
  username?: string;
  password?: string;
}

interface dataType {
  token?: string;
  message?: string;
}

export interface loginResponseData {
  code?: number;
  message?: string;
  data?: dataType;
}

export interface userInfo {
  userId?: number;
  avatar?: string;
  username?: string;
}
