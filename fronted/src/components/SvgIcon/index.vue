<template>
  <div>
    <svg :style="{ width, height }">
      <use :xlink:href="prefix + name" :fill="color"></use>
    </svg>
  </div>
</template>

<script setup lang="ts">
// 接收父组件传过来的参数
defineProps({
  prefix: {
    type: String,
    default: "#icon-",
  },
  name: String,
  color: {
    type: String,
  },
  width: {
    type: String,
    default: "",
  },
  height: {
    type: String,
    default: "16px",
  },
});
</script>

<style lang="scss" scoped></style>
