<template>
  <el-dialog
    :model-value="visible"
    title="添加托管"
    @open="beforeOpenCallback"
    :before-close="beforeClose"
  >
    <el-form
      label-width="80px"
      style="width: 80%"
      ref="ruleFormRef"
      :rules="rules"
      :model="addCodeHostingForm"
    >
      <el-form-item label="名称" prop="name" required>
        <el-input placeholder="请输入名称" v-model="addCodeHostingForm.name" />
      </el-form-item>
      <el-form-item label="类型" prop="type">
        <el-radio-group size="large" v-model="checkedType">
          <el-radio-button v-for="type in types" :key="type" :label="type">
            {{ type }}
          </el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="url" prop="url" required>
        <el-input placeholder="请输入 url" v-model="addCodeHostingForm.url" />
      </el-form-item>
      <el-form-item label="用户名" prop="username" required>
        <el-input
          placeholder="请输入用户名"
          v-model="addCodeHostingForm.username"
        />
      </el-form-item>
      <el-form-item label="Token" prop="token" required>
        <el-input
          placeholder="请输入 token"
          v-model="addCodeHostingForm.token"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="cancel">取消</el-button>
        <el-button @click="confirm">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import { FormInstance, FormRules } from "element-plus";
import { AddCodeHostingBo } from "@/api/codehosting/type.ts";

type HeaderProps = {
  visible: boolean;
};
defineProps<HeaderProps>();

// 收集托管信息
let addCodeHostingForm = reactive<AddCodeHostingBo>({
  name: "",
  type: "",
  url: "",
  username: "",
  token: "",
});

// 表单校验
const ruleFormRef = ref<FormInstance>();
const rules = reactive<FormRules>({});

const emits = defineEmits(["cancelAdd", "confirmAdd"]);

const cancel = () => {
  emits("cancelAdd");
};

const confirm = () => {
  emits("confirmAdd");
};

const beforeClose = () => {
  emits("cancelAdd");
};

const types = ["github", "gitlab", "gitea", "local"];

const checkedType = ref("github");

const beforeOpenCallback = () => {
  // 清除表单数据以及校验结果
  addCodeHostingForm.username = "";
  addCodeHostingForm.token = "";
  addCodeHostingForm.name = "";
  addCodeHostingForm.url = "";
  addCodeHostingForm.type = "";
  ruleFormRef.value?.clearValidate([
    "username",
    "token",
    "name",
    "url",
    "type",
  ]);
};
</script>

<style scoped lang="scss"></style>
