<template>
  <div>
    <el-card>
      <el-button type="primary" icon="Plus">添加 sku</el-button>
      <el-table border style="margin: 10px 0">
        <el-table-column prop="uuid" label="序号" width="80" align="center" />
        <el-table-column prop="name" label="sku 名称" />
      </el-table>
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :page-sizes="[100, 200, 300, 400]"
        :background="true"
        layout="prev, pager, next, jumper, ->,  total, sizes"
        :total="total"
      />
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

// 当前第几页
const currentPage = ref<number>(1);

// 定义每页多少条数据
const pageSize = ref<number>(10);

//一共有多个数据
const total = ref<number>(20);
</script>

<style scoped></style>
