<template>
  <el-icon style="margin-right: 10px" @click="expandOrFold">
    <component :is="menuStore.fold ? 'Fold' : 'Expand'"></component>
  </el-icon>
  <el-breadcrumb separator-icon="ArrowRight">
    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>权限管理</el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script setup lang="ts">
import userMenuStore from "@/store/modules/menu";

let menuStore = userMenuStore();

// 控制菜单的折叠或者展开
const expandOrFold = () => {
  menuStore.fold = !menuStore.fold;
  console.log(menuStore.fold);
};
</script>

<style scoped lang="scss"></style>
