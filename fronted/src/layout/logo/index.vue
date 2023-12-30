<template>
  <div class="logo">
    <img
      :src="setting.logo.path"
      style="width: 40px; height: 40px"
      v-if="setting.logo.enabled"
      @click="goHome"
    />
    <template v-if="showLogoName">
      <p>{{ setting.title }}</p></template
    >
  </div>
</template>

<script setup lang="ts">
import setting from "@/setting";
import { useRouter } from "vue-router";

let $router = useRouter();

// 获取父组件传递过来的全部路由
type HeaderProps = {
  showLogoName: boolean;
};
defineProps<HeaderProps>();

const goHome = async () => {
  await $router.push("/");
};
</script>

<style lang="scss" scoped>
.logo {
  width: 100%;
  height: $base-menu-log-height;
  // 飘起来
  display: flex;
  // 竖直方向对齐
  align-items: center;
  // 文字颜色
  color: white;
  padding: 10px;

  img {
    width: 50px;
    height: 50px;
  }

  p {
    font-size: $base-menu-logo-size;
    margin-left: 10px;
  }
}
</style>
