// 用户相关
import { defineStore } from "pinia";
// 引入 api
import { reqLogin } from "@/api/user";

import type {
  loginForm,
  loginResponseData,
  userInfoResponseData,
} from "@/api/user/type";
import type { UserState } from "./types/type";
import { SET_TOKEN, GET_TOKEN, REMOVE_TOKEN } from "@/utils/token";
import { reqUserInfo } from "@/api/user";

// 引入路由
import constantRoute from "@/router/routes";

const useUserStore = defineStore("User", {
  state: (): UserState => {
    return {
      token: GET_TOKEN(), // 登录成功存储 token
      username: "",
      avatar: "",
      menuRoutes: constantRoute,
    };
  },
  actions: {
    async userLogin(data: loginForm) {
      const result: loginResponseData = await reqLogin(data);
      if (result.code == 200) {
        this.token = result.data.token as string;
        // 本地存储
        SET_TOKEN(this.token);
        return "ok";
      } else {
        return Promise.reject(new Error(result.message));
      }
    },
    async userInfo() {
      const res: userInfoResponseData = await reqUserInfo();
      if (res.code === 200) {
        this.username = res.data.username;
        this.avatar = res.data.avatar;
        return "ok";
      } else {
        return Promise.reject(new Error(res.message));
      }
    },
    userLogout() {
      // TODO 发送接口给服务端
      this.token = "";
      this.avatar = "";
      this.username = "";
      REMOVE_TOKEN();
    },
  },
  getters: {},
});

export default useUserStore;
