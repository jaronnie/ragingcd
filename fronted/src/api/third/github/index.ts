import request_github from "@/utils/request_github";
import { getCommitListReq, getCommitListResItem } from "./type";

enum API {
  GET_COMMIT_LIST_URL = "https://api.github.com/repos/jaronnie/ragingcd/commits",
}

export const getCommitList = (params: getCommitListReq) =>
  request_github.get<any, getCommitListResItem[]>(API.GET_COMMIT_LIST_URL, {
    params,
  });
