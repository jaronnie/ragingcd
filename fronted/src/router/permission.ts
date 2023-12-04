// 路由鉴权
import router from "@/router";
import nprogress from "nprogress";
// 引入进度条样式
import "nprogress/nprogress.css";

console.log("permission");

// 全局前置守卫
router.beforeEach((_to, _from, next) => {
  // to: 你将要访问的路由
  // from: 你从哪个路由而来
  // next: 路由的放行函数

  // 进度条
  nprogress.start();
  // 放行
  next();
});

// 全局后置守卫
router.afterEach(() => {
  nprogress.done();
});
