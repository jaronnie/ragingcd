// 用户相关
import { defineStore } from "pinia";
// 引入 api
import { reqLogin } from "@/api/user";

import type { loginForm } from "@/api/user/type";

const useUserStore = defineStore("User", {
  state: () => {
    return {
      token: localStorage.getItem("TOKEN"), // 登录成功存储 token
    };
  },
  actions: {
    async userLogin(data: loginForm) {
      const result = await reqLogin(data);
      if (result.code == 200) {
        this.token = result.data.token;
        // 本地存储
        localStorage.setItem("TOKEN", result.data.token);
        return "ok";
      } else {
        return Promise.reject(new Error(result.data.message));
      }
    },
  },
  getters: {},
});

export default useUserStore;
