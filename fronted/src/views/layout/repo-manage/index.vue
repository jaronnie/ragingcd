<template>
  <div>
    <el-card style="margin-top: 10px">
      <el-button type="primary" icon="Plus" @click="addRepoFunc"
        >添加仓库</el-button
      >
      <el-table style="margin: 10px 0" border :data="repoListRes">
        <el-table-column prop="uuid" label="uuid" />
        <el-table-column prop="name" label="名称" />
        <el-table-column prop="url" label="地址" />
        <el-table-column label="操作">
          <template #default="{ row }">
            <el-button
              size="small"
              :icon="Delete"
              circle
              type="danger"
              @click="deleteRepoFunc(row)"
            ></el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        @current-change="reqRepoListFunc"
        :page-sizes="[10, 20, 30, 40, 50]"
        :background="true"
        layout="prev, pager, next, jumper, ->,  total, sizes"
        :total="total"
        prev-text="上一页"
        next-text="下一页"
      />
    </el-card>

    <CreateRepo
      v-model:visible="viewState.addRepo.visible"
      @cancelAdd="cancelAddRepo"
      @confirmAdd="confirmAddRepo"
    ></CreateRepo>

    <DeleteRepo
      v-model:visible="viewState.deleteRepo.visible"
      :repoVo="viewState.deleteRepo.repoVo"
      @cancelDelete="cancelDelete"
      @confirmDelete="confirmDelete"
    ></DeleteRepo>
  </div>
</template>

<script setup lang="ts">
import { Delete } from "@element-plus/icons-vue";
import { onMounted, reactive, ref } from "vue";
import CreateRepo from "./create.vue";
import DeleteRepo from "./delete.vue";
import { PageQuery } from "@/api/type.ts";
import { reqRepoList } from "@/api/repo/index.ts";
import { RepoListVoResponseData, RepoVo } from "@/api/repo/type.ts";

// 是否弹出 dialog 页面
const viewState = reactive({
  addRepo: { visible: false },
  deleteRepo: {
    repoVo: {} as RepoVo,
    visible: false,
  },
});

const repoListRes = ref<RepoVo[]>([]);
const total = ref<number>();
// 当前第几页
const currentPage = ref<number>(1);

// 定义每页多少条数据
const pageSize = ref<number>(10);

const reqRepoListFunc = async (page = 1) => {
  currentPage.value = page;
  const query: PageQuery = {
    pageNum: currentPage.value,
    pageSize: pageSize.value,
  };
  const response: RepoListVoResponseData = await reqRepoList(query);
  if (response.code == 200) {
    repoListRes.value = response.data.rows;
    total.value = response.data.total;
  }
};

const addRepoFunc = () => {
  viewState.addRepo.visible = true;
};

const deleteRepoFunc = (repoVo: RepoVo) => {
  viewState.deleteRepo.visible = true;
  viewState.deleteRepo.repoVo = repoVo;
};

const cancelAddRepo = () => {
  viewState.addRepo.visible = false;
};

const confirmAddRepo = () => {
  viewState.addRepo.visible = false;
  reqRepoListFunc();
};

const cancelDelete = () => {
  viewState.deleteRepo.visible = false;
};

const confirmDelete = () => {
  viewState.deleteRepo.visible = false;
  reqRepoListFunc();
};

onMounted(() => {
  reqRepoListFunc();
});
</script>

<style scoped></style>
