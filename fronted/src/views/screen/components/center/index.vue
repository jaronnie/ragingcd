<template>
  <div class="center">
    <div class="bottom"></div>
  </div>
</template>

<script lang="ts">
export default {
  // eslint-disable-next-line vue/no-reserved-component-names
  name: "Center",
};
</script>

<style scoped lang="scss">
.center {
  .bottom {
    background: url("@/assets/images/screen/dataScreen-main-cb.png") no-repeat;
    background-size: cover;
    width: 900px;
    height: 216px;
    margin-left: 30px;
    //margin-top: calc(100vh + 35px);
    margin-top: 814px;
  }
}
</style>
