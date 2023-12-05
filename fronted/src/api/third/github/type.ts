export interface getCommitListReq {
  page: number;
  per_page: number;
}

export interface getCommitListResItem {
  commit: commit;
}

export interface commit {
  author: author;
  message: string;
}

export interface author {
  name: string;
  email: string;
  date: string;
}
