import { RouteRecordRaw, RouteRecordRedirect } from "vue-router";

const constantRoute: (RouteRecordRaw | RouteRecordRedirect)[] = [
  {
    path: "/",
    meta: {
      hidden: true,
    },
    redirect: "/layout",
  },
  {
    path: "/login",
    component: () => import("@/views/login/index.vue"),
    name: "login",
  },
  {
    path: "/register",
    component: () => import("@/views/register/index.vue"),
    name: "register",
  },
  {
    path: "/layout",
    component: () => import("@/layout/index.vue"),
    name: "layout",
    redirect: "/layout/dashboard",
    children: [
      {
        path: "/layout/dashboard",
        component: () => import("@/views/layout/dashboard/index.vue"),
      },
      {
        path: "/layout/admin",
        name: "admin",
        redirect: "/layout/admin/user",
        children: [
          {
            path: "/layout/admin/user",
            component: () => import("@/views/layout/admin/user/index.vue"),
            name: "user",
          },
        ],
      },
      {
        path: "/layout/code-hosting-manage",
        component: () => import("@/views/layout/code-hosting-manage/index.vue"),
        name: "code-hosting-manage",
      },
      {
        path: "/layout/gitlog",
        component: () => import("@/views/layout/gitlog/index.vue"),
        name: "gitlog",
      },
    ],
  },
  {
    path: "/screen",
    component: () => import("@/views/screen/index.vue"),
    name: "screen",
  },
  {
    path: "/404",
    component: () => import("@/views/404/index.vue"),
    name: "404",
    meta: {
      title: "404",
      hidden: true,
    },
  },
  {
    path: "/:pathMatch(.*)*",
    redirect: "/404",
    name: "Any",
    meta: {
      title: "任意",
      hidden: true,
    },
  },
];

export default constantRoute;
