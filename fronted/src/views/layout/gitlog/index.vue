<template>
  <div class="gitlog_timeline_container">
    <el-timeline>
      <template v-for="item in result_show" :key="item">
        <el-timeline-item
          :timestamp="format_ISO8601(item.commit.author.date)"
          placement="top"
        >
          <div class="custom-icon-container">
            <img :src="item.committer.avatar_url" class="custom-icon" />
          </div>
          <el-card>
            <h2>
              {{ item.commit.author.name }}({{ item.commit.author.email }})
            </h2>
            <br />
            <div class="gitlog_timeline_item_message">
              <p>{{ item.commit.message }}</p>
              <Svg-Icon name="successFilled"></Svg-Icon>
            </div>
          </el-card>
        </el-timeline-item>
      </template>
    </el-timeline>

    <el-button
      v-if="result_show.length >= 10"
      v-show="!(result.length < 10)"
      @click="readmore"
      type="primary"
      class="gitlog_readmore"
      >查看更多</el-button
    >
  </div>
</template>

<script setup lang="ts">
// 请求获取数据
import { ref } from "vue";
import { getCommitList } from "@/api/third/github";
import { format_ISO8601 } from "@/utils/time.ts";
import type {
  getCommitListReq,
  getCommitListResItem,
} from "@/api/third/github/type.ts";

const page = ref(1);

const query: getCommitListReq = {
  page: page.value,
  per_page: 10,
};

const result = ref<getCommitListResItem[]>([]);
const result_show = ref<getCommitListResItem[]>([]);
try {
  getCommitList(query).then((response) => {
    result.value = response;
    result_show.value = response;
  });
} catch (error) {
  /* 错误处理 */
}

// read more
const readmore = () => {
  page.value = page.value + 1;
  query.page = page.value;
  try {
    getCommitList(query).then((response) => {
      result.value = response;
      result_show.value = result_show.value.concat(response);
    });
  } catch (error) {
    /* 错误处理 */
  }
};
</script>

<style scoped lang="scss">
h2 {
  font-size: 15px;
  font-weight: bold;
}
.gitlog_readmore {
  margin-left: 28px;
}
.gitlog_timeline_container {
  margin-left: 50px;
}
.custom-icon-container {
  position: absolute;
  top: 45px;
  left: -50px;
}
.custom-icon {
  width: 50px;
  height: 50px;
}
.gitlog_timeline_item_message {
  display: flex;
}
</style>
