export interface CreateCodeHostingBo {
  name: string;
  type: string;
  url: string;
  username: string;
  token: string;
}

export interface CodeHostingVo {
  id: number;
  uuid: string;
  name: string;
  type: string;
  url: string;
  username: string;
}

export interface CodeHostingVoResponseData {
  code?: number;
  message?: string;
  data?: CodeHostingVo;
}

export interface CodeHostingListVo {
  total?: number;
  rows?: CodeHostingVo[];
}

export interface CodeHostingListVoResponseData {
  code?: number;
  message?: string;
  data?: CodeHostingListVo;
}
