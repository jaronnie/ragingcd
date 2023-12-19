<template>
  <div>
    <el-dialog
      :model-value="visible"
      title="添加用户"
      draggable
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
            :before-upload="beforeAvatarUpload"
          >
            <img
              v-if="addUserForm.avatar.url"
              :src="addUserForm.avatar.url"
              style="width: 178px; height: 178px"
            />
            <el-icon v-else class="avatar-uploader-icon"
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
import type { addUserBo } from "@/api/user/type.ts";
import { reqPublicKey, reqUserAdd } from "@/api/user";
import { GET_TOKEN } from "@/utils/token.ts";
import { tokenPrefix } from "@/utils/request.ts";
import { encrypt } from "@/utils/crypto.ts";
import { publicKeyResponseData } from "@/api/user/type.ts";

type HeaderProps = {
  visible: boolean;
};
defineProps<HeaderProps>();

const emits = defineEmits(["cancel", "confirm"]);

// 收集用户信息
let addUserForm = reactive<addUserBo>({
  username: "",
  password: "",
  avatar: {
    url: "",
    uploadTime: "",
    filename: "",
  },
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

const validateAvatar = (_: any, value: any, callback: any) => {
  if (value.url === "") {
    callback(new Error("请上传 LOGO"));
  }
  callback();
};

const rules = reactive<FormRules<addUserBo>>({
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
  avatar: [
    {
      required: true,
      message: "请上传 LOGO",
      trigger: "blur",
      validator: validateAvatar,
    },
  ],
});

const handleAvatarSuccess: UploadProps["onSuccess"] = (response) => {
  addUserForm.avatar = response.data;
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
  emits("cancel");
};

const beforeClose = () => {
  emits("cancel");
};

const beforeOpenCallback = () => {
  // 清除表单数据以及校验结果
  addUserForm.username = "";
  addUserForm.password = "";
  addUserForm.avatar.url = "";
  ruleFormRef.value?.clearValidate(["username", "password", "avatar"]);
};

const confirm = async () => {
  try {
    await ruleFormRef.value.validate();
  } catch (error) {
    return;
  }

  // 获取 public key
  let publicKey: string;
  const publicKeyResult: publicKeyResponseData = await reqPublicKey();
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
  });
  console.log(result);
  if (result.code === 200) {
    ElMessage({
      type: "success",
      message: "添加成功",
    });
    emits("confirm");
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
</style>
