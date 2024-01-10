<template>
  <div>
    <el-dialog :model-value="visible" title="添加托管" @open="beforeOpenCallback" :before-close="beforeClose">
      <el-form label-width="80px" style="width: 80%" ref="ruleFormRef" :rules="rules" :model="createCodeHostingForm">
        <el-form-item label="名称" prop="name" required>
          <el-input placeholder="请输入名称" v-model="createCodeHostingForm.name" />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-radio-group size="large" v-model="createCodeHostingForm.type">
            <el-radio-button v-for="type in types" :key="type" :label="type">
              {{ type }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <template v-if="createCodeHostingForm.type != 'github'">
          <el-form-item label="地址" prop="url" required>
            <el-input placeholder="请输入地址" v-model="createCodeHostingForm.url" />
          </el-form-item>
        </template>
        <template v-if="createCodeHostingForm.type !== 'local'">
          <el-form-item label="用户名" prop="username" required>
            <el-input placeholder="请输入用户名" v-model="createCodeHostingForm.username" />
          </el-form-item>
        </template>
        <template v-if="createCodeHostingForm.type !== 'local'">
          <el-form-item label="Token" prop="token" required>
            <el-input placeholder="请输入 Token" v-model="createCodeHostingForm.token" />
          </el-form-item>
        </template>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancel">取消</el-button>
          <el-button @click="confirm" :loading="loading">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import { ElMessage, FormInstance, FormRules } from "element-plus";
import type { CreateCodeHostingBo } from "@/api/codehosting/type.ts";
import { reqCreateCodeHosting } from "@/api/codehosting";

type HeaderProps = {
  visible: boolean;
};
defineProps<HeaderProps>();

// 收集托管信息
let createCodeHostingForm = reactive<CreateCodeHostingBo>({
  name: "",
  type: "github",
  url: "",
  username: "",
  token: "",
});

// 控制按钮 loading 变量, 加载效果
let loading = ref(false);

// 表单校验
const ruleFormRef = ref<FormInstance>();
const rules = reactive<FormRules>({});

const emits = defineEmits(["cancelAdd", "confirmAdd"]);

const cancel = () => {
  emits("cancelAdd");
};

const confirm = async () => {
  try {
    await ruleFormRef.value.validate();
  } catch (error) {
    return;
  }

  loading.value = true;

  // 发送请求
  let result: any = await reqCreateCodeHosting({
    name: createCodeHostingForm.name,
    type: createCodeHostingForm.type,
    url: createCodeHostingForm.url,
    username: createCodeHostingForm.username,
    token: createCodeHostingForm.token,
  });
  if (result.code === 200) {
    ElMessage({
      type: "success",
      message: "添加成功",
    });
    emits("confirmAdd");
  } else {
    ElMessage({
      type: "error",
      message: result.message,
    });
  }

  loading.value = false;
};

const beforeClose = () => {
  emits("cancelAdd");
};

const types = ["github", "gitlab", "gitea", "local"];

const beforeOpenCallback = () => {
  // 清除表单数据以及校验结果
  createCodeHostingForm.username = "";
  createCodeHostingForm.token = "";
  createCodeHostingForm.name = "";
  createCodeHostingForm.url = "";
  ruleFormRef.value?.clearValidate([
    "username",
    "type",
    "token",
    "name",
    "url",
  ]);
};
</script>

<style scoped lang="scss"></style>
