<template>
  <div>
    <div id="chart-container"></div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from "vue";
import * as echarts from "echarts";

onMounted(() => {
  const chartContainer = document.getElementById("chart-container");
  const chart = echarts.init(chartContainer);

  // 在这里配置和渲染你的Echarts图表
  const options = {
    title: {
      text: "ECharts 入门示例",
    },
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
  chart.setOption(options);
});
</script>

<style lang="scss" scoped>
#chart-container {
  width: 100%;
  height: 300px;
}
</style>
