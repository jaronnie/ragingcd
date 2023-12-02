<template>
  <div>
    <h1>你好 {{ userStore.username }}</h1>
  </div>
</template>

<script setup lang="ts">
import useUserStore from "@/store/modules/user";
import { onMounted } from "vue";
import { useRouter } from "vue-router";

// 获取路由器
let $router = useRouter();

let userStore = useUserStore();
onMounted(async () => {
  try {
    await userStore.userInfo();
    $router.push("/home");
  } catch (error) {
    $router.push("/login");
  }
});
</script>

<style scoped></style>
