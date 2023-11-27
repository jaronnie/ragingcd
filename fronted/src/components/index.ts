// 对外暴露插件对象注册为全局组件
import SvgIcon from "./SvgIcon/index.vue";
import Pagination from "./Pagination/index.vue";

const allGlobalComponents = { SvgIcon, Pagination };

export default {
  install(app: { component: (arg0: string, arg1: any) => void }) {
    Object.keys(allGlobalComponents).forEach((key) => {
      app.component(key, allGlobalComponents[key]);
    });
  },
};
