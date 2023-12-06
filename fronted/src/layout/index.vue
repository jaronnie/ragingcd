<template>
  <div class="layout_container">
    <!-- 左侧菜单 -->
    <div class="layout_slider" :class="{ fold: layoutStore.fold }">
      <Logo :showLogoName="!layoutStore.fold"></Logo>
      <el-scrollbar class="scrollbar">
        <el-menu
          background-color="$base-menu-background"
          text-color="white"
          :router="true"
          :collapse="layoutStore.fold"
          :collapse-transition="layoutStore.fold"
        >
          <Menu :menuList="menus"></Menu
        ></el-menu>
      </el-scrollbar>
    </div>
    <!-- 顶部导航 -->
    <div class="layout_tabbar" :class="{ fold: layoutStore.fold }">
      <Tabbar></Tabbar>
    </div>
    <!-- 内容展示 -->
    <div class="layout_main" :class="{ fold: layoutStore.fold }">
      <Main></Main>
    </div>
  </div>
</template>

<script setup lang="ts">
import Logo from "./logo/index.vue";
import Menu from "./menu/index.vue";
import Main from "./main/index.vue";
import Tabbar from "./tabbar/index.vue";
import useLayoutStore from "@/store/modules/layout";

const layoutStore = useLayoutStore();

// TODO: get menus from backend
const menus = [
  {
    path: "/layout/dashboard",
    meta: {
      title: "仪表盘",
      icon: "Odometer",
    },
  },
  {
    path: "/screen",
    meta: {
      title: "数据大屏",
      icon: "Monitor",
    },
  },
  {
    path: "/layout/admin",
    name: "admin",
    meta: {
      title: "超级管理员",
      icon: "User",
    },
    children: [
      {
        path: "/layout/admin/user",
        name: "user",
        meta: {
          title: "用户管理",
          icon: "UserFilled",
        },
      },
      {
        path: "/layout/admin/permission",
        name: "permission",
        meta: {
          title: "权限管理",
          icon: "Operation",
        },
      },
    ],
  },
  {
    path: "/layout/gitlog/",
    name: "gitlog",
    meta: {
      title: "更新记录",
      icon: "Flag",
    },
  },
];
</script>

<style lang="scss" scoped>
.layout_container {
  width: 100%;
  height: 100vh;

  .layout_slider {
    width: $base-menu-width;
    height: 100vh;
    background-color: $base-menu-background;
    transition: all 0.1s;

    &.fold {
      width: $base-menu-fold-width;
    }

    .scrollbar {
      width: 100%;
      height: calc(100vh - $base-menu-log-height);
      .el-menu {
        border-right: none;
      }
    }
  }

  .layout_tabbar {
    // 相对于视口（浏览器窗口）进行定位，而不是相对于其父元素
    position: fixed;
    width: calc(100% - $base-menu-width);
    height: $base-tabbar-height;
    top: 0px;
    left: $base-menu-width;
    transition: all 0.1s;

    &.fold {
      width: calc(100vw - $base-menu-fold-width);
      left: $base-menu-fold-width;
    }
  }

  .layout_main {
    position: absolute;
    width: calc(100% - $base-menu-width);
    height: calc(100vh - $base-tabbar-height);
    left: $base-menu-width;
    top: $base-tabbar-height;
    // 设置内边距
    padding: 20px;
    // 滚动条
    overflow: auto;
    transition: all 0.1s;

    &.fold {
      width: calc(100vw - $base-menu-fold-width);
      left: $base-menu-fold-width;
    }
  }
}
</style>
