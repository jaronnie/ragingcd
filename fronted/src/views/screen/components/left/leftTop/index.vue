<template>
  <div class="leftTop">
    <div class="user_count_title_container">
      <p style="color: white; font-size: 20px">实时用户数量统计</p>
      <p class="title"></p>
    </div>
    <div class="user_count_container">
      <span v-for="(item, index) in userCountStr" :key="index">{{ item }}</span>
    </div>
    <div class="charts" ref="charts" style="width: 500px; height: 280px"></div>
  </div>
</template>

<script lang="ts">
export default {
  name: "LeftTop",
};
</script>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import * as echarts from "echarts";

const userCountStr = ref("1000000人");
let charts = ref();

onMounted(() => {
  let myCharts = echarts.init(charts.value);

  let option = {
    tooltip: {},
    legend: {
      data: ["销量"],
    },
    xAxis: {
      data: ["衬衫", "羊毛衫", "雪纺衫", "裤子", "高跟鞋", "袜子"],
    },
    yAxis: {},
    series: [
      {
        name: "销量",
        type: "bar",
        data: [5, 20, 36, 10, 10, 20],
      },
    ],
  };
  myCharts.setOption(option);
});
</script>

<style scoped lang="scss">
.leftTop {
  height: 380px;
  width: 480px;
  background: url("@/assets/images/screen/dataScreen-main-lt.png") no-repeat;
  background-size: cover;
  margin-top: 10px;

  .user_count_title_container {
    margin-left: 24px;

    .title {
      background: url("@/assets/images/screen/dataScreen-title.png");
      background-size: cover;
      height: 10px;
      width: 100px;
      margin-top: 10px;
      padding-left: 24px;
    }
  }

  .user_count_container {
    display: flex;
    margin-top: 18px;
    padding: 10px;
    justify-content: center;

    span {
      height: 55px;
      width: 55px;
      text-align: center;
      line-height: 55px;
      background: url("@/assets/images/screen/total.png") no-repeat;
      color: #29fcff;
    }
  }
}
</style>
