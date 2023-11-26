<template>
  <div>
    <el-button type="primary" :icon="Plus">Primary</el-button>

    <el-pagination
      :page-size="20"
      :pager-count="11"
      layout="prev, pager, next"
      :total="1000"
    />

    <SvgIcon name="setting" width="200px" height="500px"></SvgIcon>
  </div>
</template>

<script setup lang="ts">
import { Plus } from "@element-plus/icons-vue";
</script>

<style scoped></style>
