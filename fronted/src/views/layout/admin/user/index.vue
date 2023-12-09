<template>
  <div>
    <el-card>
      <el-button type="primary" icon="Plus">添加用户</el-button>
      <el-table :data="userListRes" style="margin: 10px 0">
        <el-table-column prop="username" label="用户名" />
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
        :page-sizes="[10, 20, 30, 40, 50]"
        :background="true"
        layout="prev, pager, next, jumper, ->,  total, sizes"
        :total="total"
      />
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { reqUserList } from "@/api/user";
import { userInfo } from "@/api/user/type.ts";

const userListRes = ref<userInfo[]>([]);
const total = ref<number>();
try {
  reqUserList().then((response) => {
    userListRes.value = response.data.rows;
    total.value = response.data.total;
  });
} catch (error) {
  /* 错误处理 */
}

// 当前第几页
const currentPage = ref<number>(1);

// 定义每页多少条数据
const pageSize = ref<number>(10);

</script>

<style scoped></style>
