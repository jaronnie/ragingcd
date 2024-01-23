// 封装本地存储数据与读取数据方法
export const SET_TOKEN = (token: string) => {
  localStorage.setItem("TOKEN", token);
};

export const GET_TOKEN = (): string => {
  return localStorage.getItem("TOKEN") as string;
};

export const REMOVE_TOKEN = () => {
  localStorage.removeItem("TOKEN");
};

export const SET_FOLD = () => {
  let fold: string = "";
  if (GET_FOLD()) {
    fold = "false";
  } else {
    fold = "true";
  }
  localStorage.setItem("FOLD", fold);
};

export const GET_FOLD = (): boolean => {
  const fold = localStorage.getItem("FOLD");
  return fold == "true";
};
