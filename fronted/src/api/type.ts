export interface PageQuery {
  pageNum: number;
  pageSize: number;
}

export interface OssFileVo {
  filename: string;
  uploadTime: string;
  url: string;
}

export interface BooleanResponseData {
  code?: number;
  message?: string;
  data?: boolean;
}
