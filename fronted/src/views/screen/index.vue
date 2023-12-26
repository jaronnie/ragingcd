<template>
  <div class="container">
    <!--数据大屏展示内容区域-->
    <div class="screen" ref="screen">
      <!--顶部-->
      <div class="top">
        <Top></Top>
      </div>
      <!--底部-->
      <div class="bottom">
        <div class="left"></div>
        <div class="center"></div>
        <div class="right"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import Top from "./components/top/index.vue";

// 定义大屏缩放比例
const getScale = (w = 1920, h = 1080) => {
  const ww = window.innerWidth / w;
  const wh = window.innerHeight / h;
  return ww < wh ? ww : wh;
};

// 监听窗口的变化
window.onresize = () => {
  screen.value.style.transform = `scale(${getScale()}) translate(-50%,-50%)`;
};

onMounted(() => {
  screen.value.style.transform = `scale(${getScale()}) translate(-50%,-50%)`;
});

let screen = ref();
</script>

<style scoped lang="scss">
.container {
  width: 100vw;
  height: 100vh;
  background: url("@/assets/images/screen/bg.png") no-repeat;
  background-size: cover;

  .screen {
    position: fixed;
    width: 1920px;
    height: 1080px;
    left: 50%;
    top: 50%;
    transform-origin: left top;

    .top {
      width: 100%;
      height: 40px;
    }

    .bottom {
      display: flex;

      .right {
        flex: 1;
      }

      .left {
        flex: 1;
      }

      .center {
        flex: 2;
      }
    }
  }
}
</style>
