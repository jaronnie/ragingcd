<template>
  <div class="register_container">
    <el-row>
      <el-col :span="12" :xs="0"></el-col>
      <el-col :span="12" :xs="24">
        <el-form
          class="register_form"
          ref="ruleFormRef"
          :model="registerInput"
          :rules="rules"
        >
          <h1>用户注册</h1>
          <el-form-item prop="username">
            <el-input
              v-model="registerInput.username"
              placeholder="请输入用户名"
            ></el-input>
          </el-form-item>
          <el-form-item prop="email">
            <el-input
              v-model="registerInput.email"
              placeholder="请输入邮箱"
            ></el-input>
          </el-form-item>
          <el-form-item prop="verifyCode">
            <div class="verify_email">
              <el-input
                v-model="registerInput.verifyCode"
                placeholder="请输入验证码"
              ></el-input>
              <el-button
                type="primary"
                style="margin-left: 15px"
                @click="getVerifyCodeFunc"
                :loading="loadingSendEmail"
                :disabled="
                  registerInput.username == '' ||
                  registerInput.email == '' ||
                  countdown > 0
                "
              >
                {{
                  countdown > 0 ? `${countdown} 秒` : "获取验证码"
                }}</el-button
              >
            </div>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              type="password"
              showPassword
              v-model="registerInput.password"
              placeholder="请输入密码"
              autocomplete="new-password"
              @keydown.enter="register"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button
              class="register_btn"
              type="primary"
              size="default"
              :loading="loadingRegister"
              @click="register"
              >立即注册</el-button
            >
          </el-form-item>
          <el-form-item>
            <el-button
              class="register_btn"
              type="primary"
              size="default"
              @click="returnLoginFunc"
              >返回登录</el-button
            >
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import type { FormInstance, FormRules } from "element-plus";
import { RegisterUserBo } from "@/api/user/type";
import { reqUserRegisterSendMail } from "@/api/user";
import useUserStore from "@/store/modules/user.ts";
import { ElNotification } from "element-plus";
import { useRouter } from "vue-router";
import { BooleanResponseData } from "@/api/type.ts";

let userStore = useUserStore();

// 获取路由器
let $router = useRouter();

// 收集用户输入
const registerInput = reactive<RegisterUserBo>({
  email: "",
  password: "",
  username: "",
  verifyCode: "",
});

// 表单校验
const ruleFormRef = ref<FormInstance>();
// 自定义校验
const validateUsername = (_: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请输入用户名"));
  } else if (value.length < 3 || value.length > 20) {
    callback(new Error("用户名长度必须在3到20个字符之间"));
  } else {
    callback();
  }
};
const validatePassword = (_: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请输入密码"));
  } else if (!/^(?=.*[a-zA-Z])(?=.*\d).{6,}$/.test(value)) {
    callback(new Error("密码必须包含字母和数字，且长度不少于6位"));
  } else {
    callback();
  }
};

const validateEmail = (_: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请输入邮箱"));
  } else if (!/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/.test(value)) {
    callback(new Error("请输入有效的邮箱地址"));
  } else {
    callback();
  }
};

const validateVerifyCode = (_: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("请输入验证码"));
  }
  callback();
};

const rules = reactive<FormRules<RegisterUserBo>>({
  username: [
    {
      required: true,
      trigger: "blur",
      validator: validateUsername,
    },
  ],
  password: [
    {
      required: true,
      trigger: "blur",
      validator: validatePassword,
    },
  ],
  email: [
    {
      required: true,
      trigger: "blur",
      validator: validateEmail,
    },
  ],
  verifyCode: [
    {
      required: true,
      trigger: "blur",
      validator: validateVerifyCode,
    },
  ],
});

// 控制按钮 loading 变量, 加载效果
let loadingRegister = ref(false);
let loadingSendEmail = ref(false);

const getVerifyCodeFunc = async () => {
  loadingSendEmail.value = true;
  let result: BooleanResponseData = await reqUserRegisterSendMail(
    registerInput.email,
    registerInput.username,
  );
  loadingSendEmail.value = false;

  if (result.code == 200) {
    startCountdown();
  } else {
    ElNotification({
      type: "error",
      message: result.message,
    });
  }
};

// 注册按钮回调
const register = async () => {
  try {
    await ruleFormRef.value.validate();
  } catch (error) {
    return;
  }

  loadingRegister.value = true;

  // 注册
  try {
    await userStore.userRegister(registerInput);
    ElNotification({
      type: "success",
      message: "注册成功",
    });
    await $router.push("/login");
  } catch (error) {
    ElNotification({
      type: "error",
      message: error.message,
    });
  }
  loadingRegister.value = false;
};

// 按钮倒计时
const countdown = ref(0);
const startCountdown = () => {
  countdown.value = 60;
  const timer = setInterval(() => {
    if (countdown.value > 0) {
      countdown.value--;
    } else {
      clearInterval(timer);
    }
  }, 1000);
};

const returnLoginFunc = async () => {
  await $router.push("/login");
};

// 初始化输入框的值
onMounted(() => {
  registerInput.username = "";
  registerInput.password = "";
});
</script>

<style scoped lang="scss">
.register_container {
  width: 100%;
  height: 100vh;
  background: url("@/assets/images/background.jpg") no-repeat;
  background-size: cover;
}

.register_form {
  position: relative;
  width: 500px;
  background: url("@/assets/images/login_form.png") no-repeat;
  background-size: cover;
  padding: 50px 50px;
  left: 135px;
  align-items: center;
  top: 21vh;
}

h1 {
  color: white;
  font-size: 40px;
  margin: 30px 0px;
}

.verify_email {
  display: flex;
  flex-grow: 1;
}

.register_btn {
  width: 100%;
}
</style>
