<template>
  <div>
    <el-card>
      <el-button type="primary" icon="Plus" @click="addUserFunc"
        >添加用户</el-button
      >
      <el-table :data="userListRes" style="margin: 10px 0" border>
        <el-table-column prop="username" label="用户名" />
        <el-table-column label="用户头像">
          <template #="{ row }">
            <el-image :src="row.avatar" style="height: 40px; width: 40px" />
          </template>
        </el-table-column>
        <el-table-column label="Operations">
          <template #default>
            <el-button link type="primary" size="small">查看</el-button>
            <el-button link type="primary" size="small">更新</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        @current-change="reqUserListFunc"
        @size-change="reqUserListFunc"
        :page-sizes="[10, 20, 30, 40, 50]"
        :background="true"
        layout="prev, pager, next, jumper, ->,  total, sizes"
        :total="total"
        prev-text="上一页"
        next-text="下一页"
      />
    </el-card>

    <addUser
      v-model:visible="viewState.addUser.visible"
      @cancel="cancelAddUser"
      @confirm="confirmAddUser"
    ></addUser>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from "vue";
import { reqUserList } from "@/api/user";
import { userInfo, userListResponseData } from "@/api/user/type.ts";
import { pageQuery } from "@/api/type.ts";
import addUser from "./addUser.vue";

const userListRes = ref<userInfo[]>([]);
const total = ref<number>();
// 当前第几页
const currentPage = ref<number>(1);

// 定义每页多少条数据
const pageSize = ref<number>(10);

// 是否弹出添加用户的 dialog 页面
const viewState = reactive({
  addUser: { visible: false },
});

const reqUserListFunc = async (page = 1) => {
  currentPage.value = page;
  const query: pageQuery = {
    pageNum: currentPage.value,
    pageSize: pageSize.value,
  };
  const response: userListResponseData = await reqUserList(query);
  if (response.code == 200) {
    userListRes.value = response.data.rows;
    total.value = response.data.total;
  }
};

const addUserFunc = () => {
  viewState.addUser.visible = true;
};

const cancelAddUser = () => {
  viewState.addUser.visible = false;
};

const confirmAddUser = () => {
  viewState.addUser.visible = false;
  reqUserListFunc();
};

onMounted(() => {
  reqUserListFunc();
});
</script>

<style scoped></style>
