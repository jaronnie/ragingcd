declare module "*.vue" {
  import { ComponentOptions } from "vue";
  const component: ComponentOptions;
  export default component;
}

declare module "nprogress";
declare module "luxon";
declare module "element-plus/dist/locale/zh-cn.mjs";
