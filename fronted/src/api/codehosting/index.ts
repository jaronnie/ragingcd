import request from "@/utils/request.ts";
import {
  CodeHostingListVoResponseData,
  CodeHostingVoResponseData,
  CreateCodeHostingBo
} from "@/api/codehosting/type.ts";
import { PageQuery } from "@/api/type.ts";

enum API {
  CREATE_URL = "/gateway/core/api/v1/codehosting/create",
  CODEHOSTING_LIST_URL = "/gateway/core/api/v1/codehosting/list",
}

export const reqCreateCodeHosting = (data: CreateCodeHostingBo) =>
  request.post<any, CodeHostingVoResponseData>(API.CREATE_URL, data);

export const reqCodeHostingList = (pageQuery: PageQuery) =>
  request.get<any, CodeHostingListVoResponseData>(API.CODEHOSTING_LIST_URL, {
    params: {
      ...pageQuery,
    },
  });
