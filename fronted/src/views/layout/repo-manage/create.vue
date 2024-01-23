<template>
  <div>
    <el-dialog
      :model-value="visible"
      title="添加仓库"
      @open="beforeOpenCallback"
      :before-close="beforeClose"
    >
      <el-form
        label-width="80px"
        style="width: 80%"
        ref="ruleFormRef"
        :rules="rules"
        :model="createRepoForm"
      >
        <el-form-item label="名称" prop="name" required>
          <el-input placeholder="请输入名称" v-model="createRepoForm.name" />
        </el-form-item>

        <el-form-item label="托管" prop="codeHostingName" required>
          <el-select
            v-model="createRepoForm.codeHostingName"
            placeholder="Select"
            style="width: 240px"
          >
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
              :disabled="item.disabled"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="地址" required>
          <el-autocomplete
            v-model="createRepoForm.url"
            :fetch-suggestions="querySearchAsync"
            placeholder="Please input"
            @select="handleSelect"
          />
        </el-form-item>
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
import { onMounted, reactive, ref } from "vue";
import { ElMessage, FormInstance, FormRules } from "element-plus";
import { CreateRepoBo } from "@/api/repo/type.ts";
import { reqCreateRepo } from "@/api/repo/index.ts";

type HeaderProps = {
  visible: boolean;
};
defineProps<HeaderProps>();

let createRepoForm = reactive<CreateRepoBo>({
  url: "",
  codeHostingName: "",
  codeHostingId: 0,
  name: "",
});

// 控制按钮 loading 变量, 加载效果
let loading = ref(false);

// 表单校验
const ruleFormRef = ref<FormInstance>();
const rules = reactive<FormRules>({});

const emits = defineEmits(["cancelAdd", "confirmAdd"]);
const options = [
  {
    value: "Option1",
    label: "Option1",
  },
  {
    value: "Option2",
    label: "Option2",
    disabled: true,
  },
  {
    value: "Option3",
    label: "Option3",
  },
  {
    value: "Option4",
    label: "Option4",
  },
  {
    value: "Option5",
    label: "Option5",
  },
];

interface LinkItem {
  value: string;
  link: string;
}

const links = ref<LinkItem[]>([]);

const loadAll = () => {
  return [
    { value: "vue", link: "https://github.com/vuejs/vue" },
    { value: "element", link: "https://github.com/ElemeFE/element" },
    { value: "cooking", link: "https://github.com/ElemeFE/cooking" },
    { value: "mint-ui", link: "https://github.com/ElemeFE/mint-ui" },
    { value: "vuex", link: "https://github.com/vuejs/vuex" },
    { value: "vue-router", link: "https://github.com/vuejs/vue-router" },
    { value: "babel", link: "https://github.com/babel/babel" },
    { value: "vue2", link: "https://github.com/vuejs/vue" },
    { value: "element2", link: "https://github.com/ElemeFE/element" },
    { value: "cooking2", link: "https://github.com/ElemeFE/cooking" },
    { value: "mint-ui2", link: "https://github.com/ElemeFE/mint-ui" },
    { value: "vuex2", link: "https://github.com/vuejs/vuex" },
    { value: "vue-router2", link: "https://github.com/vuejs/vue-router" },
    { value: "babel2", link: "https://github.com/babel/babel" },
  ];
};

let timeout: ReturnType<typeof setTimeout>;
const querySearchAsync = (queryString: string, cb: (arg: any) => void) => {
  const results = queryString
    ? links.value.filter(createFilter(queryString))
    : links.value;

  clearTimeout(timeout);
  timeout = setTimeout(() => {
    cb(results);
  }, 500 * Math.random());
};
const createFilter = (queryString: string) => {
  return (restaurant: LinkItem) => {
    return restaurant.value.toLowerCase().includes(queryString.toLowerCase());
  };
};
const handleSelect = (item: Record<string, any>) => {
  console.log(item);
};

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
  let result: any = await reqCreateRepo({
    url: createRepoForm.url,
    codeHostingName: createRepoForm.codeHostingName,
    codeHostingId: createRepoForm.codeHostingId,
    name: createRepoForm.name,
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

const beforeOpenCallback = () => {
  // 清除表单数据以及校验结果
  createRepoForm.name = "";
  ruleFormRef.value?.clearValidate(["name"]);
};

onMounted(() => {
  links.value = loadAll();
});
</script>

<style scoped lang="scss"></style>
