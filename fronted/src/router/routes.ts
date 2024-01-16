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
        meta: {
          title: "仪表盘",
        },
      },
      {
        path: "/layout/admin",
        name: "admin",
        redirect: "/layout/admin/user",
        meta: {
          title: "超级管理员",
        },
        children: [
          {
            path: "/layout/admin/user",
            component: () => import("@/views/layout/admin/user/index.vue"),
            name: "user",
            meta: {
              title: "用户管理",
            },
          },
        ],
      },
      {
        path: "/layout/code-hosting-manage",
        component: () => import("@/views/layout/code-hosting-manage/index.vue"),
        name: "code-hosting-manage",
        meta: {
          title: "代码托管平台管理",
        },
      },
      {
        path: "/layout/repo-manage",
        component: () => import("@/views/layout/feature-developing/index.vue"),
        name: "repo-manage",
        meta: {
          title: "仓库管理",
        },
      },
      {
        path: "/layout/project-manage",
        component: () => import("@/views/layout/feature-developing/index.vue"),
        name: "project-manage",
        meta: {
          title: "项目管理",
        },
      },
      {
        path: "/layout/ssh-manage",
        component: () => import("@/views/layout/ssh-manage/index.vue"),
        name: "ssh-manage",
        meta: {
          title: "SSH 管理",
        },
      },
      {
        path: "/layout/gitlog",
        component: () => import("@/views/layout/gitlog/index.vue"),
        name: "gitlog",
        meta: {
          title: "更新记录",
        },
      },
      {
        path: "/layout/404",
        component: () => import("@/views/layout/404/index.vue"),
        name: "404",
        meta: {
          title: "404",
          hidden: true,
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
    },
  },
  {
    path: "/:pathMatch(.*)*",
    redirect: "/layout/404",
    name: "Any",
    meta: {
      title: "任意",
      hidden: true,
    },
  },
];

export default constantRoute;
