// 用户相关
import { defineStore } from "pinia";
// 引入 api
import { reqLogin, reqLogout, reqPublicKey, reqUserRegister } from "@/api/user";

import type {
  LoginBo,
  LoginVoResponseData,
  PublicKeyVoResponseData,
  RegisterUserBo,
  UserVoResponseData,
} from "@/api/user/type";
import type { UserState } from "./types/type";
import { SET_TOKEN, GET_TOKEN, REMOVE_TOKEN } from "@/utils/token";
import { reqUserInfo } from "@/api/user";

// 引入路由
import constantRoute from "@/router/routes";
import { encrypt } from "@/utils/crypto.ts";

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
    async userLogin(data: LoginBo) {
      // 获取 public key
      let publicKey: string;
      const publicKeyResult: PublicKeyVoResponseData = await reqPublicKey();
      if (publicKeyResult.code == 200) {
        publicKey = publicKeyResult.data.publicKey;
      } else {
        return Promise.reject(new Error(publicKeyResult.message));
      }

      const result: LoginVoResponseData = await reqLogin({
        username: data.username,
        password: <string>encrypt(data.password, publicKey),
      });
      if (result.code == 200) {
        this.token = result.data.token as string;
        // 本地存储
        SET_TOKEN(this.token);
        return "ok";
      } else {
        return Promise.reject(new Error(result.message));
      }
    },
    async userRegister(data: RegisterUserBo) {
      // 获取 public key
      let publicKey: string;
      const publicKeyResult: PublicKeyVoResponseData = await reqPublicKey();
      if (publicKeyResult.code == 200) {
        publicKey = publicKeyResult.data.publicKey;
      } else {
        return Promise.reject(new Error(publicKeyResult.message));
      }

      const result: UserVoResponseData = await reqUserRegister({
        username: data.username,
        password: <string>encrypt(data.password, publicKey),
        email: data.email,
        verifyCode: data.verifyCode,
      });
      if (result.code == 200) {
        return "ok";
      } else {
        return Promise.reject(new Error(result.message));
      }
    },
    async userInfo() {
      const res: UserVoResponseData = await reqUserInfo();
      if (res.code === 200) {
        this.username = res.data.username;
        this.avatar = res.data.avatar;
        console.log(this.avatar);
        if (this.avatar == "" || this.avatar == null) {
          this.avatar = "/logo.png";
        }
        return "ok";
      } else {
        return Promise.reject(new Error(res.message));
      }
    },
    async userLogout() {
      this.token = "";
      this.avatar = "";
      this.username = "";
      REMOVE_TOKEN();
      try {
        const res = await reqLogout();
        if (res.code === 200) {
          return "ok";
        }
      } catch (error) {
        return Promise.reject(error.message);
      }
    },
  },
  getters: {},
});

export default useUserStore;
