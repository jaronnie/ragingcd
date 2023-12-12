// 路由鉴权
import router from "@/router";
// 获取项目配置
import setting from "@/setting";
import nprogress from "nprogress";
// 引入进度条样式
import "nprogress/nprogress.css";
// 去掉加载时的小球
nprogress.configure({ showSpinner: false });

import useUserStore from "@/store/modules/user";
import { ElMessage } from "element-plus";

// 全局前置守卫
router.beforeEach(async (to, _from, next) => {
  // to: 你将要访问的路由
  // from: 你从哪个路由而来
  // next: 路由的放行函数

  // 获取用户仓库
  const userStore = useUserStore();

  // 获取 token 判断用户是否登录
  const token = userStore.token;

  if (token) {
    // 有 token
    // 不能再访问
    if (to.path === "/login") {
      if (userStore.username != "") {
        next("/");
      } else {
        // token 过期了
        await userStore.userLogout();
        next();
      }
    } else {
      // 在守卫这里发请求获取用户信息
      if (userStore.username == "") {
        try {
          await userStore.userInfo();
          next();
        } catch (error) {
          // token 过期
          // 用户手动修改本地存储的 token
          // 直接退出登录
          ElMessage({
            type: "error",
            message: error.message,
          });
          await userStore.userLogout();
          // 回到 login 并带上参数
          next({
            path: "/login",
            query: {
              redirect: to.path,
            },
          });
        }
      }
      next();
    }
  } else {
    if (to.path === "/login") {
      // 放行
      next();
    } else {
      // 用户未登录, 任意路由将跳转到 login, 并且带上参数
      next({
        path: "/login",
        query: {
          redirect: to.path,
        },
      });
    }
  }

  // 进度条
  nprogress.start();
  // 放行
  next();
});

// 全局后置守卫
router.afterEach((to) => {
  nprogress.done();

  // 跳转之后, 操作 dom 修改标题
  document.title = setting.title + "-" + to.meta.title;
});
