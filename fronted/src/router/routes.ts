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
    meta: {
      title: "登录",
      hidden: true,
    },
  },
  {
    path: "/layout",
    component: () => import("@/layout/index.vue"),
    name: "layout",
    meta: {
      title: "layout",
      hidden: true,
    },
    redirect: "/layout/dashboard",
    children: [
      {
        path: "/layout/dashboard",
        component: () => import("@/views/layout/dashboard/index.vue"),
        meta: {
          title: "仪表盘",
          icon: "Odometer",
        },
      },
      {
        path: "/layout/admin",
        name: "admin",
        meta: {
          title: "超级管理员",
          icon: "User",
        },
        redirect: "/layout/admin/user",
        children: [
          {
            path: "/layout/admin/user",
            component: () => import("@/views/layout/admin/user/index.vue"),
            name: "user",
            meta: {
              title: "用户管理",
              icon: "UserFilled",
            },
          },
          {
            path: "/layout/admin/permission",
            component: () =>
              import("@/views/layout/admin/permission/index.vue"),
            name: "permission",
            meta: {
              title: "权限管理",
              icon: "Operation",
            },
          },
        ],
      },
      {
        path: "/layout/product",
        name: "product",
        meta: {
          title: "产品管理",
          icon: "Product",
        },
        redirect: "/layout/product/attr",
        children: [
          {
            path: "/layout/product/attr",
            component: () => import("@/views/layout/product/attr/index.vue"),
            name: "attr",
            meta: {
              title: "属性管理",
              icon: "UserFilled",
            },
          },
          {
            path: "/layout/product/spu",
            component: () => import("@/views/layout/product/spu/index.vue"),
            name: "spu",
            meta: {
              title: "spu 管理",
              icon: "Operation",
            },
          },
          {
            path: "/layout/product/sku",
            component: () => import("@/views/layout/product/sku/index.vue"),
            name: "sku",
            meta: {
              title: "sku 管理",
              icon: "Operation",
            },
          },
          {
            path: "/layout/product/trademark",
            component: () =>
              import("@/views/layout/product/trademark/index.vue"),
            name: "trademark",
            meta: {
              title: "trademark 管理",
              icon: "Operation",
            },
          },
        ],
      },
      {
        path: "/layout/gitlog/",
        component: () => import("@/views/layout/gitlog/index.vue"),
        name: "gitlog",
        meta: {
          title: "更新记录",
          icon: "Flag",
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
      icon: "Monitor",
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
