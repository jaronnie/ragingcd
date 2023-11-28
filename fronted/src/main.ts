import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import { createApp } from "vue";
import App from "@/App.vue";
// svg 配置
import "virtual:svg-icons-register";

const app = createApp(App);
app.use(ElementPlus);

// 引入全局对象组件
import globalComponent from "@/components";
app.use(globalComponent);

// 引入全局样式
import "@/styles/index.scss";

// 注册模板路由
import router from "@/router";
app.use(router);

// 状态管理
import pinia from "@/store";
app.use(pinia);

// 获取环境配置信息
console.log(import.meta.env);

app.mount("#app");
