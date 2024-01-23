import request from "@/utils/request.ts";
import {
  RepoListVoResponseData,
  RepoVoResponseData,
  CreateRepoBo,
} from "@/api/repo/type.ts";
import { BooleanResponseData, PageQuery } from "@/api/type.ts";

enum API {
  CREATE_URL = "/gateway/core/api/v1/repo/create",
  Repo_LIST_URL = "/gateway/core/api/v1/repo/list",
  Repo_DELETE_URL = "/gateway/core/api/v1/repo/delete/",
}

export const reqCreateRepo = (data: CreateRepoBo) =>
  request.post<any, RepoVoResponseData>(API.CREATE_URL, data);

export const reqRepoList = (pageQuery: PageQuery) =>
  request.get<any, RepoListVoResponseData>(API.Repo_LIST_URL, {
    params: {
      ...pageQuery,
    },
  });

export const reqRepoDelete = (RepoId: number) =>
  request.get<any, BooleanResponseData>(`${API.Repo_DELETE_URL}/${RepoId}`);
