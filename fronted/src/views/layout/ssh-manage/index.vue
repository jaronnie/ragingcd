<template>
  <div>
    <el-card style="margin-top: 10px">
      <el-button type="primary" icon="Plus" @click="addSshFunc"
        >添加 ssh</el-button
      >
      <el-table style="margin: 10px 0" border :data="sshListRes">
        <el-table-column prop="uuid" label="uuid" />
        <el-table-column prop="name" label="名称" />
        <el-table-column prop="type" label="类型" />
        <el-table-column prop="ip" label="地址" />
        <el-table-column prop="port" label="端口" />
        <el-table-column prop="username" label="用户名" />
        <el-table-column label="操作">
          <template #default="{ row }">
            <el-button circle size="small" @click="openTerminalFunc">
              <Svg-Icon name="terminal"></Svg-Icon>
            </el-button>
            <el-button
              size="small"
              :icon="Delete"
              circle
              type="danger"
              @click="deleteSshFunc(row)"
            ></el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        @current-change="reqSshList"
        :page-sizes="[10, 20, 30, 40, 50]"
        :background="true"
        layout="prev, pager, next, jumper, ->,  total, sizes"
        :total="total"
        prev-text="上一页"
        next-text="下一页"
      />
    </el-card>

    <CreateSsh
      v-model:visible="viewState.addSsh.visible"
      @cancelAdd="cancelAddSsh"
      @confirmAdd="confirmAddSsh"
    ></CreateSsh>

    <DeleteSsh
      v-model:visible="viewState.deleteSsh.visible"
      :sshVo="viewState.deleteSsh.sshVo"
      @cancelDelete="cancelDelete"
      @confirmDelete="confirmDelete"
    ></DeleteSsh>

    <Terminal
      v-model:visible="viewState.openTerminal.visible"
      @closeTerminal="closeTerminal"
    ></Terminal>
  </div>
</template>

<script setup lang="ts">
import { Delete } from "@element-plus/icons-vue";
import { onMounted, reactive, ref } from "vue";
import CreateSsh from "./createSsh.vue";
import DeleteSsh from "./deleteSsh.vue";
import Terminal from "./terminal.vue";
import { PageQuery } from "@/api/type.ts";
import { SshListVoResponseData, SshVo } from "@/api/ssh/type.ts";
import { reqSshList } from "@/api/ssh";

// 是否弹出 dialog 页面
const viewState = reactive({
  addSsh: { visible: false },
  deleteSsh: {
    visible: false,
    sshVo: {} as SshVo,
  },
  openTerminal: { visible: false },
});

const sshListRes = ref<SshVo[]>([]);
const total = ref<number>();
// 当前第几页
const currentPage = ref<number>(1);

// 定义每页多少条数据
const pageSize = ref<number>(10);

const reqSshListFunc = async (page = 1) => {
  currentPage.value = page;
  const query: PageQuery = {
    pageNum: currentPage.value,
    pageSize: pageSize.value,
  };
  const response: SshListVoResponseData = await reqSshList(query);
  if (response.code == 200) {
    sshListRes.value = response.data.rows;
    total.value = response.data.total;
  }
};

const addSshFunc = () => {
  viewState.addSsh.visible = true;
};

const deleteSshFunc = (sshVo: SshVo) => {
  viewState.deleteSsh.visible = true;
  viewState.deleteSsh.sshVo = sshVo;
};

const cancelAddSsh = () => {
  viewState.addSsh.visible = false;
};

const confirmAddSsh = () => {
  viewState.addSsh.visible = false;
  reqSshListFunc();
};

const cancelDelete = () => {
  viewState.deleteSsh.visible = false;
};

const confirmDelete = () => {
  viewState.deleteSsh.visible = false;
  reqSshListFunc();
};

const openTerminalFunc = () => {
  viewState.openTerminal.visible = true;
};

const closeTerminal = () => {
  viewState.openTerminal.visible = false;
};

onMounted(() => {
  reqSshListFunc();
});
</script>

<style scoped></style>
