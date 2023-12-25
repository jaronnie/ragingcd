<template>
  <el-card style="height: 120px; margin-top: -10px">
    <el-form :inline="true">
      <el-form-item label="用户名">
        <el-input
          placeholder="请输入用户名"
          v-model="searchUserForm.username"
        ></el-input>
      </el-form-item>
      <el-form-item label="邮箱">
        <el-input
          placeholder="请输入邮箱"
          v-model="searchUserForm.email"
        ></el-input>
      </el-form-item>
    </el-form>
    <el-button type="primary" icon="Search" @click="searchUserFunc"
      >查询</el-button
    >
    <el-button type="primary" icon="Refresh" @click="resetSearchUserFunc"
      >重置</el-button
    >
  </el-card>
  <div>
    <el-card style="margin-top: 10px">
      <el-button type="primary" icon="Plus" @click="addUserFunc"
        >添加用户</el-button
      >
      <el-table :data="userListRes" style="margin: 10px 0" border>
        <el-table-column label="用户头像">
          <template #="{ row }">
            <el-image
              :src="
                row.avatar !== '' && row.avatar !== null
                  ? row.avatar
                  : '/logo.png'
              "
              style="height: 40px; width: 40px; border-radius: 50%"
            />
          </template>
        </el-table-column>
        <el-table-column prop="username" label="用户名" />
        <el-table-column prop="email" label="邮箱" />
        <el-table-column label="操作">
          <template #default="{ row }">
            <el-button
              size="small"
              :icon="Delete"
              circle
              type="danger"
              @click="deleteUserFunc(row)"
            ></el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        @current-change="reqUserListFunc"
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
      @cancelAdd="cancelAddUser"
      @confirmAdd="confirmAddUser"
    ></addUser>

    <deleteUser
      v-model:visible="viewState.deleteUser.visible"
      :userInfo="viewState.deleteUser.userInfo"
      @cancelDelete="cancelDeleteUser"
      @confirmDelete="confirmDeleteUser"
    ></deleteUser>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from "vue";
import { reqUserList } from "@/api/user";
import {
  UserVo as User,
  UserListVoResponseData,
  SearchUserQuery,
} from "@/api/user/type.ts";
import { PageQuery } from "@/api/type.ts";
import addUser from "./addUser.vue";
import deleteUser from "./deleteUser.vue";
import { Delete } from "@element-plus/icons-vue";

const userListRes = ref<User[]>([]);
const total = ref<number>();
// 当前第几页
const currentPage = ref<number>(1);

// 定义每页多少条数据
const pageSize = ref<number>(10);

// 是否弹出添加用户的 dialog 页面
const viewState = reactive({
  addUser: { visible: false },
  deleteUser: {
    visible: false,
    userInfo: {} as User,
  },
});

// 收集查询用户时填写的表单
let searchUserForm = reactive<SearchUserQuery>({
  username: "",
  email: "",
});

const reqUserListFunc = async (page = 1) => {
  currentPage.value = page;
  const query: PageQuery = {
    pageNum: currentPage.value,
    pageSize: pageSize.value,
  };
  const response: UserListVoResponseData = await reqUserList(
    query,
    searchUserForm,
  );
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

// 删除 user
const deleteUserFunc = (userInfo: User) => {
  viewState.deleteUser.visible = true;
  viewState.deleteUser.userInfo = userInfo;
};

const cancelDeleteUser = () => {
  viewState.deleteUser.visible = false;
};

const confirmDeleteUser = () => {
  viewState.deleteUser.visible = false;
  reqUserListFunc();
};

// ================================== 查询 user 业务 ==================================
const searchUserFunc = () => {
  reqUserListFunc();
};

const resetSearchUserFunc = () => {
  searchUserForm.email = "";
  searchUserForm.username = "";
  reqUserListFunc();
};

onMounted(() => {
  reqUserListFunc();
});
</script>

<style scoped lang="scss"></style>
