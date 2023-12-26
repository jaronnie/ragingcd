<template>
  <div class="top">
    <div class="left">
      <span class="lbtn" @click="goHome">首页</span>
    </div>
    <div class="center">
      <div class="title">jaronnie 运营平台数据大屏</div>
    </div>
    <div class="right">
      <span class="rbtn">统计报告</span>
      <span class="time">当前时间 {{ time }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { ref, onMounted, onBeforeUnmount } from "vue";
import moment from "moment";

let $router = useRouter();

let time = ref(moment().format("YYYY-MM-DD HH:MM:ss"));
let timer = ref(0);
const goHome = async () => {
  await $router.push("/");
};

onMounted(() => {
  setInterval(() => {
    time.value = moment().format("YYYY-MM-DD HH:MM:ss");
  }, 1000);
});

onBeforeUnmount(() => {
  clearInterval(timer.value);
});
</script>

<style scoped lang="scss">
.top {
  display: flex;
  width: 100%;
  height: 40px;

  .left {
    flex: 1;
    background: url("@/assets/images/screen/dataScreen-header-left-bg.png")
      no-repeat;
    background-size: cover;

    .lbtn {
      float: right;
      width: 128px;
      height: 40px;
      background-size: 100% 100%;
      background: url("@/assets/images/screen/dataScreen-header-btn-bg-l.png")
        no-repeat;
      font-size: 20px;
      line-height: 40px; // TODO: 关键点, 待理解
      text-align: center;
      color: #29fcff;
    }
  }

  .center {
    flex: 2;
    .title {
      background: url("@/assets/images/screen/dataScreen-header-center-bg.png")
        no-repeat;
      background-size: cover;
      height: 82px;
      width: 100%;
      line-height: 82px;
      text-align: center;
      color: #29fcff;
      font-size: 30px;
    }
  }

  .right {
    flex: 1;
    background: url("@/assets/images/screen/dataScreen-header-right-bg.png")
      no-repeat;
    background-size: cover;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .rbtn {
      float: left;
      width: 150px;
      height: 40px;
      background-size: 100% 100%;
      background: url("@/assets/images/screen/dataScreen-header-btn-bg-r.png")
        no-repeat;
      font-size: 20px;
      line-height: 40px;
      text-align: center;
      color: #29fcff;
    }

    .time {
      color: #29fcff;
      font-size: 20px;
      margin-right: 30px;
    }
  }
}
</style>
