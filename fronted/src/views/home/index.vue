<template>
  <div>
    <h1>你好 {{ userStore.username }}</h1>
  </div>
</template>

<script setup lang="ts">
import useUserStore from "@/store/modules/user";
let userStore = useUserStore();
userStore.userInfo();
</script>

<style scoped></style>
