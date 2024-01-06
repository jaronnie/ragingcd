<template>
  <div>
    <el-card style="margin-top: 10px">
      <el-button type="primary" icon="Plus" @click="addCodeHostingFunc"
        >添加托管</el-button
      >
      <el-table style="margin: 10px 0" border :data="codeHostingListRes">
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
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        @current-change="reqCodeHostingListFunc"
        :page-sizes="[10, 20, 30, 40, 50]"
        :background="true"
        layout="prev, pager, next, jumper, ->,  total, sizes"
        :total="total"
        prev-text="上一页"
        next-text="下一页"
      />
    </el-card>

    <CreateCodeHosting
      v-model:visible="viewState.addCodeHosting.visible"
      @cancelAdd="cancelAddUser"
      @confirmAdd="confirmAddUser"
    ></CreateCodeHosting>
  </div>
</template>

<script setup lang="ts">
import { Delete } from "@element-plus/icons-vue";
import { onMounted, reactive, ref } from "vue";
import CreateCodeHosting from "./createCodeHosting.vue";
import { PageQuery } from "@/api/type.ts";
import { reqCodeHostingList } from "@/api/codehosting";
import {
  CodeHostingListVoResponseData,
  CodeHostingVo,
} from "@/api/codehosting/type.ts";

// 是否弹出 dialog 页面
const viewState = reactive({
  addCodeHosting: { visible: false },
  deleteCodeHosting: {
    visible: false,
  },
});

const codeHostingListRes = ref<CodeHostingVo[]>([]);
const total = ref<number>();
// 当前第几页
const currentPage = ref<number>(1);

// 定义每页多少条数据
const pageSize = ref<number>(10);

const reqCodeHostingListFunc = async (page = 1) => {
  currentPage.value = page;
  const query: PageQuery = {
    pageNum: currentPage.value,
    pageSize: pageSize.value,
  };
  const response: CodeHostingListVoResponseData =
    await reqCodeHostingList(query);
  if (response.code == 200) {
    codeHostingListRes.value = response.data.rows;
    total.value = response.data.total;
  }
};

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
  reqCodeHostingListFunc();
};

onMounted(() => {
  reqCodeHostingListFunc();
});
</script>

<style scoped></style>
