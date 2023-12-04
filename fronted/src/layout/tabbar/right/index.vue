<template>
  <el-button
    size="small"
    icon="Refresh"
    circle
    @click="refreshMain"
  ></el-button>
  <el-button
    size="small"
    icon="FullScreen"
    circle
    @click="fullscreen"
  ></el-button>
  <el-button size="small" icon="Setting" circle></el-button>
  <img
    src="/logo.png"
    style="height: 30px; width: 30px; margin-left: 10px; margin-right: 10px"
  />
  <el-dropdown>
    <span class="el-dropdown-link">
      admin
      <el-icon class="el-icon--right">
        <arrow-down />
      </el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item>退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import useLayoutStore from "@/store/modules/layout";
import screenfull from "screenfull";

const layoutStore = useLayoutStore();
const refreshMain = () => {
  // 点击按钮值更改
  layoutStore.refresh = !layoutStore.refresh;
};

const fullscreen = () => {
  screenfull.toggle();
};
</script>

<style scoped lang="scss">
.tabbar_right {
  // 上下居中
  align-items: center;
  display: flex;
}
</style>
