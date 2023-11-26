import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import { createApp } from "vue";
import App from "@/App.vue";

const app = createApp(App);
app.use(ElementPlus);

// 获取环境配置信息
console.log(import.meta.env);

// svg 配置
import "virtual:svg-icons-register";

app.mount("#app");
