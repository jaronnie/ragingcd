<template>
  <div>
    <el-dialog
      :model-value="visible"
      :title="`确定删除 ${codeHostingVo.name} 吗`"
      style="width: 428px; height: 168px"
      :before-close="beforeClose"
    >
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="cancel">取消</el-button>
          <el-button @click="confirm(codeHostingVo.id)" type="primary">
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ElMessage } from "element-plus";
import { reqCodeHostingDelete } from "@/api/codehosting";
import { CodeHostingVo } from "@/api/codehosting/type.ts";

type HeaderProps = {
  visible: boolean;
  codeHostingVo: CodeHostingVo;
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
  let result: any = await reqCodeHostingDelete(id);
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
