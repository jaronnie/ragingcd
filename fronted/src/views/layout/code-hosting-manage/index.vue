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
          <template #default="{ row }">
            <el-button
              size="small"
              :icon="Delete"
              circle
              type="danger"
              @click="deleteCodeHostingFunc(row)"
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
      @cancelAdd="cancelAddCodeHosting"
      @confirmAdd="confirmAddCodeHosting"
    ></CreateCodeHosting>

    <DeleteCodeHosting
      v-model:visible="viewState.deleteCodeHosting.visible"
      :codeHostingVo="viewState.deleteCodeHosting.codeHostingVo"
      @cancelDelete="cancelDelete"
      @confirmDelete="confirmDelete"
    ></DeleteCodeHosting>
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

import DeleteCodeHosting from "./deleteCodeHosting.vue";

// 是否弹出 dialog 页面
const viewState = reactive({
  addCodeHosting: { visible: false },
  deleteCodeHosting: {
    codeHostingVo: {} as CodeHostingVo,
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

const deleteCodeHostingFunc = (codeHostingVo: CodeHostingVo) => {
  viewState.deleteCodeHosting.visible = true;
  viewState.deleteCodeHosting.codeHostingVo = codeHostingVo;
};

const cancelAddCodeHosting = () => {
  viewState.addCodeHosting.visible = false;
};

const confirmAddCodeHosting = () => {
  viewState.addCodeHosting.visible = false;
  reqCodeHostingListFunc();
};

const cancelDelete = () => {
  viewState.deleteCodeHosting.visible = false;
};

const confirmDelete = () => {
  viewState.deleteCodeHosting.visible = false;
  reqCodeHostingListFunc();
};

onMounted(() => {
  reqCodeHostingListFunc();
});
</script>

<style scoped></style>
