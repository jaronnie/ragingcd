<template>
  <el-button circle size="small" @click="gotoGithub">
    <Svg-Icon name="gitHub"></Svg-Icon>
  </el-button>
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
    :src="userStore.avatar"
    style="
      height: 30px;
      width: 30px;
      margin-left: 10px;
      margin-right: 10px;
      border-radius: 50%;
    "
  />
  <el-dropdown>
    <span class="el-dropdown-link">
      {{ userStore.username }}
      <el-icon class="el-icon--right">
        <arrow-down />
      </el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import useUserStore from "@/store/modules/user";
import useLayoutStore from "@/store/modules/layout";
import screenfull from "screenfull";
import { useRouter } from "vue-router";
import { ElMessage } from "element-plus";

const userStore = useUserStore();
const layoutStore = useLayoutStore();

const gotoGithub = () => {
  window.open("https://github.com/jaronnie/ragingcd", "_blank");
};

const refreshMain = () => {
  // 点击按钮值更改
  layoutStore.refresh = !layoutStore.refresh;
};

const fullscreen = () => {
  screenfull.toggle();
};

// 退出登录
const $router = useRouter();
const logout = async () => {
  try {
    await userStore.userLogout();
  } catch (error) {
    ElMessage({
      type: "error",
      message: error,
    });
  }
  await $router.push("/login");
};
</script>

<style scoped lang="scss">
.tabbar_right {
  // 上下居中
  align-items: center;
  display: flex;
}
</style>
