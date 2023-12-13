export interface getCommitListReq {
  page: number;
  per_page: number;
}

export interface getCommitListResItem {
  commit: commit;
  committer: committer;
}

export interface commit {
  author: author;
  message: string;
}

export interface committer {
  avatar_url: string;
}

export interface author {
  name: string;
  email: string;
  date: string;
  avatar_url: string;
}
