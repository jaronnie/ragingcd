export interface CreateRepoBo {
  name: string;
  codeHostingId: number;
  codeHostingName: string;
  url: string;
}

export interface RepoVo {
  id: number;
  uuid: string;
  name: string;
  type: string;
  ip: string;
  port: number;
  username: string;
}

export interface RepoVoResponseData {
  code?: number;
  message?: string;
  data?: RepoVo;
}

export interface RepoListVo {
  total?: number;
  rows?: RepoVo[];
}

export interface RepoListVoResponseData {
  code?: number;
  message?: string;
  data?: RepoListVo;
}
