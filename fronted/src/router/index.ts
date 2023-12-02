// 通过 vue-router 实现模板路由配置
import { createRouter, createWebHistory } from "vue-router";
import constantRoute from "./routes.ts";

const router = createRouter({
  // 路由模式
  history: createWebHistory(),
  // 滚动行为
  scrollBehavior() {
    return {
      left: 0,
      top: 0,
    };
  },
  routes: constantRoute,
});

export default router;
