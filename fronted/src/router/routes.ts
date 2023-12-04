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
    path: "/acl",
    component: () => import("@/layout/index.vue"),
    name: "acl",
    meta: {
      title: "权限管理",
      icon: "User",
    },
    redirect: "/acl/user",
    children: [
      {
        path: "/acl/user",
        component: () => import("@/views/acl/user/index.vue"),
        name: "user",
        meta: {
          title: "用户管理",
          icon: "User",
        },
      },
      {
        path: "/acl/permission",
        component: () => import("@/views/acl/permission/index.vue"),
        name: "permission",
        meta: {
          title: "权限管理",
          icon: "User",
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
