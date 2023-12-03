// 管理菜单是否收缩的仓库
import { defineStore } from "pinia";
import type { MenuState } from "./types/type";

const userMenuStore = defineStore("Menu", {
  state: (): MenuState => {
    return {
      fold: false,
    };
  },
});

export default userMenuStore;
