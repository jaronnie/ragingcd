<template>
  <div>
    <el-dialog
      :model-value="visible"
      :title="`确定删除 ${repoVo.name} 吗`"
      style="width: 428px; height: 168px"
      :before-close="beforeClose"
    >
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancel">取消</el-button>
          <el-button @click="confirm(repoVo.id)" type="primary">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ElMessage } from "element-plus";
import { reqRepoDelete } from "@/api/repo";
import { RepoVo } from "@/api/repo/type.ts";

type HeaderProps = {
  visible: boolean;
  repoVo: RepoVo;
};
defineProps<HeaderProps>();

const emits = defineEmits(["cancelDelete", "confirmDelete"]);

const cancel = () => {
  emits("cancelDelete");
};

const beforeClose = () => {
  emits("cancelDelete");
};

const confirm = async (id: number) => {
  // 调用删除接口
  let result: any = await reqRepoDelete(id);
  if (result.code === 200) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
  } else {
    ElMessage({
      type: "error",
      message: result.message,
    });
  }
  emits("confirmDelete");
};
</script>
