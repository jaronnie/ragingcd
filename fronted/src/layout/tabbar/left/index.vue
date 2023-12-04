<template>
  <el-icon style="margin-right: 10px" @click="expandOrFold">
    <component :is="menuStore.fold ? 'Fold' : 'Expand'"></component>
  </el-icon>
  <el-breadcrumb separator-icon="ArrowRight">
    <el-breadcrumb-item
      v-for="(item, index) in $router.matched"
      :key="index"
      v-show="!item.meta.hidden"
      :to="item.path"
    >
      <el-icon><component :is="item.meta.icon"></component></el-icon>
      <span style="margin: 0px 2px">{{ item.meta.title }}</span>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script setup lang="ts">
import userMenuStore from "@/store/modules/menu";
import { useRoute } from "vue-router";

let $router = useRoute();

let menuStore = userMenuStore();

// 控制菜单的折叠或者展开
const expandOrFold = () => {
  menuStore.fold = !menuStore.fold;
  console.log(menuStore.fold);
};
</script>

<style scoped lang="scss"></style>
