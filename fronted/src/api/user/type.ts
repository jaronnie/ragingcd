import { OssFileVo } from "@/api/type.ts";

export interface LoginBo {
  username: string;
  password: string;
}

export interface RegisterUserBo {
  username: string;
  password: string;
  email: string;
  verifyCode: string;
}

export interface AddUserBo {
  username: string;
  password: string;
  avatar: OssFileVo;
}

interface LoginVo {
  token?: string;
}

export interface LoginVoResponseData {
  code?: number;
  message?: string;
  data?: LoginVo;
}

export interface UserVoResponseData {
  code?: number;
  message?: string;
  data?: UserVo;
}

export interface UserListVoResponseData {
  code?: number;
  message?: string;
  data?: UserListVo;
}

export interface UserListVo {
  total?: number;
  rows?: UserVo[];
}

export interface UserVo {
  id: number;
  avatar: string;
  username: string;
}

export interface PublicKeyVoResponseData {
  code?: number;
  message?: string;
  data?: PublicKeyVo;
}

export interface PublicKeyVo {
  type: string;
  publicKey: string;
}
