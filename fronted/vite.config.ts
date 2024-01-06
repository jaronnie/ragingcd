import vue from "@vitejs/plugin-vue";
import path from "path";
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";
import { ConfigEnv, UserConfigExport } from "vite";
import { viteMockServe } from "vite-plugin-mock";

export default ({ command }: ConfigEnv): UserConfigExport => {
  const prodMock = false;
  const devMock = false;
  // const env = loadEnv(mode, process.cwd());
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
        localEnabled: command === "serve" && devMock,
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
    // 代理跨域
    server: {
      proxy: {
        "/gateway/stdb": {
          target: "http://127.0.0.1:8082/",
          // 需要代理跨域
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/gateway\/stdb/, ""),
        },
        "/gateway/core": {
          target: "http://127.0.0.1:81",
          // 需要代理跨域
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/gateway\/core/, "apps.core"),
        },
      },
    },
  };
};
