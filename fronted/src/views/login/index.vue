<template>
  <div class="login_container">
    <el-row>
      <el-col :span="12" :xs="0"></el-col>
      <el-col :span="12" :xs="24">
        <el-form class="login_form">
          <h1>你好世界!</h1>
          <h2>by jaronnie</h2>
          <el-form-item>
            <el-input
              :prefix-icon="User"
              v-model="loginForm.username"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-input
              type="password"
              :prefix-icon="Lock"
              showPassword
              v-model="loginForm.password"
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

let userStore = useUserStore();

let loginForm = reactive({
  username: "",
  password: "",
});

// 获取路由器
let $router = useRouter();

// 控制按钮 loading 变量, 加载效果
let loading = ref(false);

// 登录按钮回调
const login = async () => {
  loading.value = true;
  try {
    await userStore.userLogin(loginForm);
    $router.push("/");
    ElNotification({
      type: "success",
      message: "登录成功",
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
