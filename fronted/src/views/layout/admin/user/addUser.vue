<template>
  <div>
    <el-dialog
      :model-value="visible"
      title="添加用户"
      @open="beforeOpenCallback"
      :before-close="beforeClose"
    >
      <el-form
        label-width="80px"
        style="width: 80%"
        ref="ruleFormRef"
        :model="addUserForm"
        :rules="rules"
      >
        <el-form-item label="用户名" prop="username" required>
          <el-input placeholder="请输入用户名" v-model="addUserForm.username" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input placeholder="请输入邮箱" v-model="addUserForm.email" />
        </el-form-item>
        <el-form-item label="密码" prop="password" required>
          <el-input placeholder="请输入密码" v-model="addUserForm.password" />
        </el-form-item>
        <el-form-item label="用户头像" prop="avatar" required>
          <el-upload
            class="avatar-uploader"
            action="/api/v1.0/system/oss_file/upload"
            :headers="{ Authorization: tokenPrefix + ' ' + GET_TOKEN() }"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :on-progress="handleAvatarProgress"
            :before-upload="beforeAvatarUpload"
          >
            <img
              v-if="addUserForm.avatar.url"
              :src="addUserForm.avatar.url"
              style="width: 178px; height: 178px"
            />
            <el-icon
              v-else
              class="avatar-uploader-icon"
              :class="{ 'uploading-icon': uploading }"
              ><uploadFilled
            /></el-icon>
          </el-upload>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancel">取消</el-button>
          <el-button @click="confirm" type="primary"> 确定 </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import type { FormInstance, UploadProps } from "element-plus";
import { ElMessage, FormRules } from "element-plus";
import type { AddUserBo } from "@/api/user/type.ts";
import { reqPublicKey, reqUserAdd } from "@/api/user";
import { encrypt } from "@/utils/crypto.ts";
import { PublicKeyVoResponseData } from "@/api/user/type.ts";

type HeaderProps = {
  visible: boolean;
};
defineProps<HeaderProps>();

const emits = defineEmits(["cancelAdd", "confirmAdd"]);

// 收集用户信息
let addUserForm = reactive<AddUserBo>({
  username: "",
  password: "",
  email: "",
  avatar: {
    url: "",
    uploadTime: "",
    filename: "",
  },
});

// 表单校验
const ruleFormRef = ref<FormInstance>();

const uploading = ref(false);
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

const validateAvatar = (_: any, value: any, callback: any) => {
  if (value.url === "") {
    callback(new Error("请上传 LOGO"));
  }
  callback();
};

const rules = reactive<FormRules<AddUserBo>>({
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
  avatar: [
    {
      required: true,
      trigger: "blur",
      validator: validateAvatar,
    },
  ],
});

const handleAvatarSuccess: UploadProps["onSuccess"] = (response) => {
  addUserForm.avatar = response.data;
  ElMessage({
    type: "success",
    message: "上传成功",
  });
  uploading.value = false;
};

const handleAvatarProgress = () => {
  uploading.value = true;
};

const beforeAvatarUpload: UploadProps["beforeUpload"] = (rawFile) => {
  if (rawFile.size / 1024 / 1024 > 5) {
    ElMessage({
      type: "error",
      message: "图片不能大于 5MB",
    });
  }
  return true;
};

const cancel = () => {
  emits("cancelAdd");
};

const beforeClose = () => {
  emits("cancelAdd");
};

const beforeOpenCallback = () => {
  // 清除表单数据以及校验结果
  addUserForm.username = "";
  addUserForm.password = "";
  addUserForm.avatar.url = "";
  addUserForm.email = "";
  ruleFormRef.value?.clearValidate(["username", "password", "avatar", "email"]);
};

const confirm = async () => {
  try {
    await ruleFormRef.value.validate();
  } catch (error) {
    return;
  }

  // 获取 public key
  let publicKey: string;
  const publicKeyResult: PublicKeyVoResponseData = await reqPublicKey();
  if (publicKeyResult.code == 200) {
    publicKey = publicKeyResult.data.publicKey;
  } else {
    ElMessage({
      type: "error",
      message: publicKeyResult.message,
    });
    return;
  }

  // 发送请求
  let result: any = await reqUserAdd({
    username: addUserForm.username,
    password: encrypt(addUserForm.password, publicKey),
    avatar: addUserForm.avatar,
    email: addUserForm.email,
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
};
</script>

<style scoped>
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}

@keyframes rotate {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.avatar-uploader-icon.uploading-icon {
  animation: rotate 1s linear infinite;
}
</style>
