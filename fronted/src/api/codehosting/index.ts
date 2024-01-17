import request from "@/utils/request.ts";
import {
  CodeHostingListVoResponseData,
  CodeHostingVoResponseData,
  CreateCodeHostingBo,
} from "@/api/codehosting/type.ts";
import { BooleanResponseData, PageQuery } from "@/api/type.ts";

enum API {
  CREATE_URL = "/gateway/core/api/v1/codehosting/create",
  CODEHOSTING_LIST_URL = "/gateway/core/api/v1/codehosting/list",
  CODEHOSTING_DELETE_URL = "/gateway/core/api/v1/codehosting/delete",
}

export const reqCreateCodeHosting = (data: CreateCodeHostingBo) =>
  request.post<any, CodeHostingVoResponseData>(API.CREATE_URL, data);

export const reqCodeHostingList = (pageQuery: PageQuery) =>
  request.get<any, CodeHostingListVoResponseData>(API.CODEHOSTING_LIST_URL, {
    params: {
      ...pageQuery,
    },
  });

export const reqCodeHostingDelete = (id: number) =>
  request.get<any, BooleanResponseData>(`${API.CODEHOSTING_DELETE_URL}/${id}`);