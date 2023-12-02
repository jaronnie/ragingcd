<template>
  <div class="login_container">
    <el-row>
      <el-col :span="12" :xs="0"></el-col>
      <el-col :span="12" :xs="24">
        <el-form
          class="login_form"
          ref="ruleFormRef"
          :model="loginInput"
          :rules="rules"
        >
          <h1>你好世界!</h1>
          <h2>by jaronnie</h2>
          <el-form-item prop="username">
            <el-input
              :prefix-icon="User"
              v-model="loginInput.username"
            ></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              type="password"
              :prefix-icon="Lock"
              showPassword
              v-model="loginInput.password"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button
              class="login_btn"
              type="primary"
              size="default"
              :loading="loading"
              @click="login"
              >登录</el-button
            >
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { User, Lock } from "@element-plus/icons-vue";
import { reactive, ref } from "vue";
import useUserStore from "@/store/modules/user";
import { useRouter } from "vue-router";
import { ElNotification } from "element-plus";
import { getTime } from "@/utils/time";
import type { FormInstance, FormRules } from "element-plus";
import type { loginForm } from "@/api/user/type";

let userStore = useUserStore();

// 收集用户输入
const loginInput = reactive<loginForm>({
  username: "",
  password: "",
});

// 表单校验
const ruleFormRef = ref<FormInstance>();
// 自定义校验
const validateUsername = (_: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请输入用户名"));
  }
  callback();
};
const validatePassword = (_: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请输入密码"));
  }
  callback();
};
const rules = reactive<FormRules<loginForm>>({
  username: [
    {
      required: true,
      message: "请输入用户名",
      trigger: "blur",
      validator: validateUsername,
    },
  ],
  password: [
    {
      required: true,
      message: "请输入密码",
      trigger: "blur",
      validator: validatePassword,
    },
  ],
});

// 获取路由器
let $router = useRouter();

// 控制按钮 loading 变量, 加载效果
let loading = ref(false);

// 登录按钮回调
const login = async () => {
  try {
    await ruleFormRef.value.validate();
  } catch (error) {
    return;
  }

  loading.value = true;
  try {
    await userStore.userLogin(loginInput);
    $router.push("/home");
    ElNotification({
      type: "success",
      message: "欢迎回来",
      title: `HI, ${getTime()}好`,
    });
    loading.value = false;
  } catch (error) {
    loading.value = false;
    ElNotification({
      type: "error",
      message: error.message,
    });
  }
};
</script>

<style scoped lang="scss">
.login_container {
  width: 100%;
  height: 100vh;
  background: url("@/assets/images/background.jpg") no-repeat;
  background-size: cover;
}
.login_form {
  position: relative;
  width: 80%;
  top: 30vh;
  background: url("@/assets/images/login_form.png") no-repeat;
  background-size: cover;
  padding: 40px;
}

h1 {
  color: white;
  font-size: 40px;
}

h2 {
  font-size: 20px;
  color: white;
  margin: 20px 0px;
}

.login_btn {
  width: 100%;
}
</style>
