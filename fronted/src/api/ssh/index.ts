import request from "@/utils/request.ts";
import {
  SshListVoResponseData,
  SshVoResponseData,
  CreateSshBo,
} from "@/api/ssh/type.ts";
import { BooleanResponseData, PageQuery } from "@/api/type.ts";

enum API {
  CREATE_URL = "/gateway/core/api/v1/ssh/create",
  SSH_LIST_URL = "/gateway/core/api/v1/ssh/list",
  SSH_DELETE_URL = "/gateway/core/api/v1/ssh/delete/",
}

export const reqCreateSsh = (data: CreateSshBo) =>
  request.post<any, SshVoResponseData>(API.CREATE_URL, data);

export const reqSshList = (pageQuery: PageQuery) =>
  request.get<any, SshListVoResponseData>(API.SSH_LIST_URL, {
    params: {
      ...pageQuery,
    },
  });

export const reqSshDelete = (sshId: number) =>
  request.get<any, BooleanResponseData>(`${API.SSH_DELETE_URL}/${sshId}`);
