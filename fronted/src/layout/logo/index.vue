<template>
  <div class="logo">
    <img :src="setting.logo.path" v-if="setting.logo.enabled" />
    <p>{{ setting.title }}</p>
  </div>
</template>

<script setup lang="ts">
import setting from "@/setting";
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
    widows: 50px;
    height: 50px;
  }

  p {
    font-size: $base-menu-logo-size;
    margin-left: 10px;
  }
}
</style>
