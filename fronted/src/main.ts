import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import { createApp } from "vue";
import App from "@/App.vue";
// svg 配置
import "virtual:svg-icons-register";
import SvgIcon from "@/components/SvgIcon/index.vue";

const app = createApp(App);
app.use(ElementPlus);
app.component("SvgIcon", SvgIcon);

// 获取环境配置信息
console.log(import.meta.env);

app.mount("#app");
