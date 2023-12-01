// 通过 vue-router 实现模板路由配置
import { createRouter, createWebHistory } from "vue-router";

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
  routes: [
    {
      path: "/login",
      component: () => import("@/views/login/index.vue"),
      name: "login",
    },
    {
      path: "/",
      component: () => import("@/layout/index.vue"),
      name: "layout",
    },
    {
      path: "/404",
      component: () => import("@/views/404/index.vue"),
      name: "404",
    },
    {
      path: "/:pathMatch(.*)*",
      redirect: "/404",
      name: "Any",
    },
  ],
});

export default router;
