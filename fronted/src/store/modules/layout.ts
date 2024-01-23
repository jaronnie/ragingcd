// 管理菜单是否收缩的仓库
import { defineStore } from "pinia";
import type { LayoutState } from "./types/type";
import { GET_FOLD, SET_FOLD } from "@/utils/storage.ts";

const useLayoutStore = defineStore("Layout", {
  state: (): LayoutState => {
    return {
      fold: GET_FOLD(),
      refresh: false,
    };
  },
  actions: {
    ExpandOrFold() {
      SET_FOLD();
      this.fold = GET_FOLD();
    },
  },
});

export default useLayoutStore;
