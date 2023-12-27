<template>
  <div class="right">
    <div class="top"></div>
    <div class="center"></div>
    <div class="bottom"></div>
  </div>
</template>

<script lang="ts"></script>

<style scoped lang="scss">
.right {
  .top {
    height: 380px;
    width: 480px;
    background: url("@/assets/images/screen/dataScreen-main-rt.png") no-repeat;
    background-size: cover;
    margin-top: 10px;
  }

  .center {
    height: 300px;
    width: 480px;
    background: url("@/assets/images/screen/dataScreen-main-rc.png") no-repeat;
    background-size: cover;
    margin-top: 20px;
  }

  .bottom {
    height: 300px;
    width: 480px;
    background: url("@/assets/images/screen/dataScreen-main-rb.png") no-repeat;
    background-size: cover;
    margin-top: 20px;
  }
}
</style>
