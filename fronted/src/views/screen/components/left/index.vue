<template>
  <div class="left">
    <div class="top"></div>
    <div class="center"></div>
    <div class="bottom"></div>
  </div>
</template>

<script lang="ts">
export default {
  name: "Left",
};
</script>

<style scoped lang="scss">
.left {
  .top {
    height: 380px;
    width: 480px;
    background: url("@/assets/images/screen/dataScreen-main-lt.png") no-repeat;
    background-size: cover;
    margin-top: 10px;
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
