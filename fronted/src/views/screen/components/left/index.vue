<template>
  <div class="left">
    <div class="top">
      <LeftTop></LeftTop>
    </div>
    <div class="center"></div>
    <div class="bottom"></div>
  </div>
</template>

<script lang="ts">
export default {
  name: "Left",
};
</script>

<script setup lang="ts">
import LeftTop from "@/views/screen/components/left/leftTop/index.vue";
</script>

<style scoped lang="scss">
.left {
  .top {

  }

  .center {
    height: 300px;
    width: 480px;
    background: url("@/assets/images/screen/dataScreen-main-lc.png") no-repeat;
    background-size: cover;
    margin-top: 20px;
  }

  .bottom {
    height: 300px;
    width: 480px;
    background: url("@/assets/images/screen/dataScreen-main-lb.png") no-repeat;
    background-size: cover;
    margin-top: 20px;
  }
}
</style>
