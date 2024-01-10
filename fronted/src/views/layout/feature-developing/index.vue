<template>
  <div class="container">
    <div class="content">
      <div class="text">功能开发中, 请等待上线...</div>
      <img
        src="@/assets/images/error_images/feature-developing.svg"
        style="display: block; margin: 0 auto; width: 600px; height: 600px"
      />
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped lang="scss">
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;

  .content {
    text-align: center;
  }

  .text {
    font-size: 30px;
  }
}
</style>
