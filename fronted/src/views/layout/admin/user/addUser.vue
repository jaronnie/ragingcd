<template>
  <div>
    <el-dialog :model-value="visible" title="添加用户">
      <el-form label-width="80px" style="width: 80%">
        <el-form-item label="用户名">
          <el-input placeholder="请输入用户名" v-model="addUserForm.username" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input placeholder="请输入密码" v-model="addUserForm.password" />
        </el-form-item>
        <el-form-item label="用户头像">
          <el-upload
            class="avatar-uploader"
            action="/api/v1.0/system/oss_file/upload"
            :headers="{ Authorization: tokenPrefix + ' ' + GET_TOKEN() }"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
          >
            <img
              v-if="imageUrl"
              :src="imageUrl"
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
import type { UploadProps } from "element-plus";
import { ElMessage } from "element-plus";
import type { addUserBo } from "@/api/user/type.ts";
import { reqUserAdd } from "@/api/user";
import { GET_TOKEN } from "@/utils/token.ts";
import { tokenPrefix } from "@/utils/request.ts";

type HeaderProps = {
  visible: boolean;
};
defineProps<HeaderProps>();

const emits = defineEmits(["cancel", "confirm"]);

const imageUrl = ref("");

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

const handleAvatarSuccess: UploadProps["onSuccess"] = (response) => {
  imageUrl.value = response.data.url;
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

const confirm = async () => {
  // 发送请求
  let result: any = await reqUserAdd(addUserForm);
  if (result.code == 200) {
    ElMessage({
      type: "success",
      message: "添加成功",
    });
    emits("confirm");
  } else {
    ElMessage({
      type: "error",
      message: result.data.message,
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
