// 管理菜单是否收缩的仓库
import { defineStore } from "pinia";
import type { LayoutState } from "./types/type";

const useLayoutStore = defineStore("Layout", {
  state: (): LayoutState => {
    return {
      fold: false,
      refresh: false,
    };
  },
});

export default useLayoutStore;
