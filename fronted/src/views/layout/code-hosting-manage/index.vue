<template>
  <el-card style="margin-top: 10px">
    <el-button type="primary" icon="Plus" @click="addCodeHostingFunc"
      >添加托管</el-button
    >
    <el-table style="margin: 10px 0" border>
      <el-table-column prop="uuid" label="uuid" />
      <el-table-column prop="name" label="名称" />
      <el-table-column prop="type" label="类型" />
      <el-table-column prop="url" label="地址" />
      <el-table-column label="操作">
        <template>
          <el-button
            size="small"
            :icon="Delete"
            circle
            type="danger"
            @click="deleteCodeHostingFunc()"
          ></el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>

  <AddCodeHosting
    v-model:visible="viewState.addCodeHosting.visible"
    @cancelAdd="cancelAddUser"
    @confirmAdd="confirmAddUser"
  ></AddCodeHosting>
</template>

<script setup lang="ts">
import { Delete } from "@element-plus/icons-vue";
import { reactive } from "vue";
import AddCodeHosting from "./addCodeHosting.vue";

// 是否弹出 dialog 页面
const viewState = reactive({
  addCodeHosting: { visible: false },
  deleteCodeHosting: {
    visible: false,
  },
});

const addCodeHostingFunc = () => {
  viewState.addCodeHosting.visible = true;
};

const deleteCodeHostingFunc = () => {
  viewState.deleteCodeHosting.visible = true;
};

const cancelAddUser = () => {
  viewState.addCodeHosting.visible = false;
};

const confirmAddUser = () => {
  viewState.addCodeHosting.visible = false;
};
</script>

<style scoped></style>
