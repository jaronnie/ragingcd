<template>
  <el-drawer
    :model-value="visible"
    title="I am terminal"
    :before-close="handleClose"
    direction="btt"
  >
  </el-drawer>
</template>

<script lang="ts" setup>
import { ElMessageBox } from "element-plus";

type HeaderProps = {
  visible: boolean;
};
defineProps<HeaderProps>();

const emits = defineEmits(["closeTerminal"]);

const handleClose = (done: () => void) => {
  ElMessageBox.confirm("Are you sure you want to close this?")
    .then(() => {
      done();
      emits("closeTerminal");
    })
    .catch(() => {
      // catch error
    });
};
</script>
