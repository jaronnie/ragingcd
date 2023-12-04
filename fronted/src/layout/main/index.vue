<template>
  <router-view v-slot="{ Component }">
    <transition name="fade">
      <!-- 渲染 layout 一级路由组件的子路由 -->
      <component :is="Component" v-if="!hidden" />
    </transition>
  </router-view>
</template>

<script setup lang="ts">
import { nextTick, ref, watch } from "vue";
import useLayoutStore from "@/store/modules/layout";
const layoutStore = useLayoutStore();

// watch layoutStore.refresh 的更改
// 开始不需要隐藏
let hidden = ref(false);

watch(
  () => layoutStore.refresh,
  () => {
    hidden.value = true;
    nextTick(() => {
      hidden.value = false;
    });
  },
);
</script>

<style lang="scss" scoped>
.fade-enter-from {
  opacity: 0;
  transform: scale(0);
}

.fade-enter-to {
  opacity: 1;
  transform: scale(1);
}

.fade-enter-active {
  transition: all 0.5s;
}
</style>
