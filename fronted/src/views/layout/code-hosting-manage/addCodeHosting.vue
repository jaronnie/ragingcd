<template>
  <div>
    <el-dialog
      :model-value="visible"
      title="添加托管"
      :before-close="beforeClose"
    >
      <el-form
        label-width="80px"
        style="width: 80%"
        ref="ruleFormRef"
        :rules="rules"
      >
        <el-form-item label="名称" prop="name" required>
          <el-input placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="url" prop="url" required>
          <el-input placeholder="请输入 url" />
        </el-form-item>
        <el-form-item label="用户名" prop="username" required>
          <el-input placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="Token" prop="token" required>
          <el-input placeholder="请输入 token" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancel">取消</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";
import { FormRules } from "element-plus";

type HeaderProps = {
  visible: boolean;
};
defineProps<HeaderProps>();

const rules = reactive<FormRules>({});

const emits = defineEmits(["cancelAdd", "confirmAdd"]);

const cancel = () => {
  emits("cancelAdd");
};

const beforeClose = () => {
  emits("cancelAdd");
};
</script>

<style scoped>
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
