import vue from "@vitejs/plugin-vue";
import path from "path";
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";
import { ConfigEnv, UserConfigExport } from "vite";
import { viteMockServe } from "vite-plugin-mock";

export default ({ command }: ConfigEnv): UserConfigExport => {
  return {
    base: "./",
    plugins: [
      vue(),
      createSvgIconsPlugin({
        iconDirs: [path.resolve(process.cwd(), "src/assets/icons")],
        symbolId: "icon-[dir]-[name]",
      }),
      viteMockServe({
        localEnabled: command === "serve",
      }),
    ],
    resolve: { alias: { "@": path.resolve("./src") } },
    css: {
      preprocessorOptions: {
        scss: {
          javascriptEnabled: true,
          additionalData: '@import "./src/styles/variable.scss";',
        },
      },
    },
  };
};
