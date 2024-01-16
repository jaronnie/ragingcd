export interface CreateSshBo {
  name: string;
  type: string;
  ip: string;
  port: number;
  username: string;
  password: string;
  private_key: string;
}

export interface SshVo {
  id: number;
  uuid: string;
  name: string;
  type: string;
  ip: string;
  port: number;
  username: string;
}

export interface SshVoResponseData {
  code?: number;
  message?: string;
  data?: SshVo;
}

export interface CodeHostingListVo {
  total?: number;
  rows?: SshVo[];
}

export interface SshListVoResponseData {
  code?: number;
  message?: string;
  data?: CodeHostingListVo;
}
