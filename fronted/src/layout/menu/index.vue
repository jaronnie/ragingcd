<template>
  <template v-for="item in menuList" :key="item.path">
    <el-menu-item v-if="!item.children && !item.meta.hidden" :index="item.path">
      <el-icon><component :is="item.meta.icon"></component></el-icon>
      <template #title>
        <span>{{ item.meta.title }}</span>
      </template>
    </el-menu-item>

    <el-sub-menu
      v-if="item.children && item.children.length >= 1 && !item.meta.hidden"
      :index="item.path"
      :key="item.path"
      teleported
      popper-class="menu-popo"
    >
      <template #title>
        <el-icon><component :is="item.meta.icon"></component></el-icon>
        <span>{{ item.meta.title }}</span>
      </template>
      <Menu :menuList="item.children"></Menu>
    </el-sub-menu>
  </template>
</template>

<script setup lang="ts">
// 获取父组件传递过来的全部路由
type HeaderProps = {
  menuList: any[];
};
defineProps<HeaderProps>();
</script>

<script lang="ts">
export default {
  // eslint-disable-next-line vue/no-reserved-component-names
  name: "Menu",
};
</script>

<style scoped lang="scss">
:global(.menu-popo) {
  background: $base-menu-background !important;
}
</style>
