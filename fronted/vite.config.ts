import vue from "@vitejs/plugin-vue";
import path from "path";
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";
import { ConfigEnv, UserConfigExport } from "vite";
import { viteMockServe } from "vite-plugin-mock";

export default ({ command }: ConfigEnv): UserConfigExport => {
  const prodMock = true;
  return {
    base: "./",
    plugins: [
      vue(),
      createSvgIconsPlugin({
        iconDirs: [path.resolve(process.cwd(), "src/assets/icons")],
        symbolId: "icon-[dir]-[name]",
      }),
      viteMockServe({
        mockPath: "./mock/", //mock文件地址
        localEnabled: command === "serve",
        prodEnabled: command !== "serve" && prodMock,
        injectCode: ` import { setupProdMockServer } from './mockProdServer'; setupProdMockServer(); `,
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
