import { RouteRecordRaw, RouteRecordRedirect } from "vue-router";

const constantRoute: (RouteRecordRaw | RouteRecordRedirect)[] = [
  {
    path: "/login",
    component: () => import("@/views/login/index.vue"),
    name: "login",
    meta: {
      title: "登录",
      hidden: true,
    },
  },
  {
    path: "/",
    component: () => import("@/layout/index.vue"),
    name: "layout",
    meta: {
      title: "",
      hidden: true,
    },
    redirect: "/home",
    children: [
      {
        path: "/home",
        component: () => import("@/views/home/index.vue"),
        name: "home",
        meta: {
          title: "首页",
          icon: "HomeFilled",
        },
      },
    ],
  },
  {
    path: "/screen",
    component: () => import("@/views/screen/index.vue"),
    name: "screen",
    meta: {
      title: "数据大屏",
      icon: "Odometer",
    },
  },
  {
    path: "/admin",
    component: () => import("@/layout/index.vue"),
    name: "admin",
    meta: {
      title: "超级管理员",
      icon: "User",
    },
    redirect: "/admin/user",
    children: [
      {
        path: "/admin/user",
        component: () => import("@/views/admin/user/index.vue"),
        name: "user",
        meta: {
          title: "用户管理",
          icon: "UserFilled",
        },
      },
      {
        path: "/admin/permission",
        component: () => import("@/views/admin/permission/index.vue"),
        name: "permission",
        meta: {
          title: "权限管理",
          icon: "Operation",
        },
      },
    ],
  },
  {
    path: "/gitlog",
    component: () => import("@/layout/index.vue"),
    name: "gitlog",
    meta: {
      title: "",
    },
    children: [
      {
        path: "/gitlog/",
        component: () => import("@/views/gitlog/index.vue"),
        name: "gitlog",
        meta: {
          title: "更新记录",
          icon: "Flag",
        },
      },
    ],
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
