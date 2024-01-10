<template>
  <div class="container">
    <div class="content">
      <div class="text">页面找不到了...</div>
      <img
        src="@/assets/images/error_images/404.svg"
        style="display: block; margin: 0 auto"
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
}

.content {
  text-align: center;
}

.text {
  font-size: 30px;
}
</style>
