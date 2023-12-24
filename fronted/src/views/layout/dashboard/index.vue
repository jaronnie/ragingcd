<template>
  <div>
    <el-card style="height: 128px; margin-top: -10px">
      <div class="card">
        <div class="card_left">
          <img
            :src="userStore.avatar"
            style="
              height: 120px;
              width: 120px;
              border-radius: 50%;
              margin-top: -15px;
            "
          />
          <div class="card_left_title">
            <h1>你好 {{ userStore.username }}, 欢迎回来!</h1>
            <h3 class="small_title">{{ setting.title }}</h3>
          </div>
        </div>
        <div class="card_right"></div>
      </div>
    </el-card>

    <div class="welcome_img">
      <img
        src="@/assets/images/welcome.png"
        style="height: 600px; width: 600px"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import useUserStore from "@/store/modules/user.ts";
import { onMounted } from "vue";
import setting from "@/setting";

// import { useRouter } from "vue-router";

// // 获取路由器
// let $router = useRouter();

let userStore = useUserStore();
onMounted(async () => {
  console.log("dashboard");
});
</script>

<style scoped lang="scss">
.welcome_img {
  margin-top: 10px;
  display: flex;
  justify-content: center;
}
.card {
  display: flex;

  .card_left {
    display: flex;
    width: 50%;
    height: auto;
  }

  .card_left_title {
    h1 {
      margin-left: 10px;
      margin-top: 30px;
      font-size: 25px;
    }

    .small_title {
      margin-left: 10px;
      margin-top: 15px;
      font-style: italic;
      color: skyblue;
    }
  }
}
</style>
