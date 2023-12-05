<template>
  <el-icon style="margin-right: 10px" @click="expandOrFold">
    <component :is="layoutStore.fold ? 'Expand' : 'Fold'"></component>
  </el-icon>
  <el-breadcrumb separator-icon="ArrowRight">
    <el-breadcrumb-item
      v-for="(item, index) in $router.matched"
      :key="index"
      v-show="!item.meta.hidden"
      :to="item.path"
    >
      <!-- <el-icon><component :is="item.meta.icon"></component></el-icon> -->
      <span style="margin: 0px 2px">{{ item.meta.title }}</span>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script setup lang="ts">
import useLayoutStore from "@/store/modules/layout";
import { useRoute } from "vue-router";

let $router = useRoute();

let layoutStore = useLayoutStore();

// 控制菜单的折叠或者展开
const expandOrFold = () => {
  layoutStore.fold = !layoutStore.fold;
};
</script>

<style scoped lang="scss"></style>
