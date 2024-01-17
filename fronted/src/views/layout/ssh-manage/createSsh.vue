<template>
  <div>
    <el-dialog
      :model-value="visible"
      title="添加 ssh"
      @open="beforeOpenCallback"
      :before-close="beforeClose"
    >
      <el-form
        label-width="80px"
        style="width: 80%"
        ref="ruleFormRef"
        :rules="rules"
        :model="createSshForm"
      >
        <el-form-item label="名称" prop="name" required>
          <el-input placeholder="请输入名称" v-model="createSshForm.name" />
        </el-form-item>
        <el-form-item label="地址" prop="ip" required>
          <el-input placeholder="请输入地址" v-model="createSshForm.ip" />
        </el-form-item>
        <el-form-item label="端口" prop="port" required>
          <el-input placeholder="请输入端口" v-model="createSshForm.port" />
        </el-form-item>
        <el-form-item label="用户名" prop="username" required>
          <el-input
            placeholder="请输入用户名"
            v-model="createSshForm.username"
          />
        </el-form-item>
        <el-form-item label="登录方式" prop="type">
          <el-radio-group size="large" v-model="createSshForm.type">
            <el-radio-button v-for="type in types" :key="type" :label="type">
              {{ type }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <template v-if="createSshForm.type == 'password'">
          <el-form-item label="密码" prop="password" required>
            <el-input
              placeholder="请输入密码"
              v-model="createSshForm.password"
            />
          </el-form-item>
        </template>
        <template v-if="createSshForm.type == 'private_key'">
          <el-form-item label=" 私钥" prop="private_key" required>
            <el-input
              placeholder="请输入私钥"
              v-model="createSshForm.private_key"
            />
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
import { reqCreateSsh } from "@/api/ssh";
import { CreateSshBo } from "@/api/ssh/type.ts";

type HeaderProps = {
  visible: boolean;
};
defineProps<HeaderProps>();

// 收集托管信息
let createSshForm = reactive<CreateSshBo>({
  ip: "",
  type: "password",
  port: 22,
  name: "",
  username: "",
  password: "",
  private_key: "",
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
  let result: any = await reqCreateSsh({
    ip: createSshForm.ip,
    port: createSshForm.port,
    name: createSshForm.name,
    type: createSshForm.type,
    username: createSshForm.username,
    password: createSshForm.password,
    private_key: createSshForm.private_key,
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

const types = ["password", "private_key"];

const beforeOpenCallback = () => {
  createSshForm.username = "";
  createSshForm.name = "";
  createSshForm.ip = "";
  createSshForm.port = 22;
  ruleFormRef.value?.clearValidate([
    "username",
    "type",
    "name",
    "password",
    "private_key",
    "ip",
  ]);
};
</script>

<style scoped lang="scss"></style>
