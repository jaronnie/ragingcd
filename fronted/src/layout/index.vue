<template>
  <div class="layout_container">
    <!-- 左侧菜单 -->
    <div class="layout_slider">
      <Logo></Logo>
    </div>
    <!-- 顶部导航 -->
    <div class="layout_tabbar">顶部导航</div>
    <!-- 内容展示 -->
    <div class="layout_main">
      <p style="height: 10000px; background-color: red">内容展示</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import Logo from "./logo/index.vue";
</script>

<style lang="scss" scoped>
.layout_container {
  width: 100%;
  height: 100vh;

  .layout_slider {
    width: $base-menu-width;
    height: 100vh;
    background-color: $base-menu-background;
  }

  .layout_tabbar {
    // 相对于视口（浏览器窗口）进行定位，而不是相对于其父元素
    position: fixed;
    width: calc(100% - $base-menu-width);
    height: $base-tabbar-height;
    background-color: blueviolet;
    top: 0px;
    left: $base-menu-width;
  }

  .layout_main {
    position: absolute;
    width: calc(100% - $base-menu-width);
    background-color: bisque;
    height: calc(100vh - $base-tabbar-height);
    left: $base-menu-width;
    top: $base-tabbar-height;
    // 设置内边距
    padding: 20px;
    // 滚动条
    overflow: auto;
  }
}
</style>
