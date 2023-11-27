<template>
  <div>
    <el-button type="primary" :icon="Plus">Primary</el-button>

    <h1>我是根 App</h1>

    <el-pagination
      :page-size="20"
      :pager-count="11"
      layout="prev, pager, next"
      :total="1000"
    />

    <SvgIcon name="setting" width="50px" height="50px"></SvgIcon>
  </div>
</template>

<script setup lang="ts">
import { Plus } from "@element-plus/icons-vue";

// 测试 mock 接口
import { onMounted } from "vue";
import { reqLogin, reqUserInfo } from "@/api/user";

onMounted(() => {
  reqLogin({
    username: "admin",
    password: "123456",
  });
  reqUserInfo();
});
</script>

<style scoped lang="scss">
div {
  h1 {
    color: $color;
  }
}
</style>
