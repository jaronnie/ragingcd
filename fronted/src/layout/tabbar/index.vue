<template>
  <div class="tabbar_container">
    <div class="tabbar_left"><Left></Left></div>
    <div class="tabbar_right"><Right></Right></div>
  </div>
</template>

<script setup lang="ts">
import Left from "./left/index.vue";
import Right from "./right/index.vue";
</script>

<style scoped lang="scss">
.tabbar_container {
  width: 100%;
  height: 100%;
  display: flex;
  // 一左一右
  justify-content: space-between;

  .tabbar_left {
    // 上下居中
    align-items: center;
    margin-left: 20px;
    display: flex;
  }
  .tabbar_right {
    // 上下居中
    align-items: center;
    display: flex;
  }
}
</style>
