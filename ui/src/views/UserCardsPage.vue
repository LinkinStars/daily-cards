<script setup lang="ts">
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import InfiniteLoading from "v3-infinite-loading";
import "v3-infinite-loading/lib/style.css";
import { getCardsPage, Card } from "@/api/card";
import CalHeatmap from "@/components/CalHeatmap.vue"

const router = useRouter();
const route = useRoute();

const cards = ref<Card[]>([]);
let page = 1;
let pageSize = 20;
const load = async ($state) => {
  try {
    const resp = await getCardsPage(page, route.query.q as string, route.query.d as string);
    cards.value.push(...resp.data.cards);
    if (resp.data.cards.length < pageSize) $state.complete();
    else {
      $state.loaded();
    }
    page++;
  } catch (error) {
    $state.error();
  }
};

const jumpCardDetailPage = (id: number) => {
  router.push({
    name: "user-card-detail",
    params: {
      id: id,
    },
  });
};

const jumpPostPage = () => {
  router.push({
    name: "user-card-post",
    params: { id: 0 },
  });
};

const jumpDayCardPage = async (date : string) => {
  router.push({ name: "user-card-page", query: { d: date }});
};
</script>

<template>
  <div class="card-list-bg">
    <div class="card-list">
      <CalHeatmap @clickBox="jumpDayCardPage" />
      <div class="card-list-item" v-for="card in cards" :key="card.id">
        <div class="card-content" v-html="card.content" v-highlight></div>
        <hr />
        <div class="card-footer">
          <span>📅 {{ card.created_at }} 👀 {{ card.pv }}</span>
          <span @click="jumpCardDetailPage(card.id)">详情</span>
        </div>
      </div>
      <InfiniteLoading
        @infinite="load"
        :slots="{ complete: '没有更多卡片了', error: '加载失败' }"
      />
    </div>
    <button class="post-card-btn" @click="jumpPostPage()">+</button>
  </div>
</template>

<style scoped>
.card-list-bg {
  display: flex;
  flex-direction: row;
  justify-content: center;
  
  text-align: center;
  /**min-height: 100%;**/

  min-height: 100vh;
  background: linear-gradient(-45deg, #ee7752, #e73c7e, #23a6d5, #23d5ab);
  background-size: 400% 400%;
}

.card-list {
  padding: 20px 20px;
  margin-top: 60px;
  max-width: 1200px;
  width: calc(100% - 40px);
}
.post-card-btn {
  position: fixed;
  bottom: 80px;
  right: 30px;
  border-radius: 50%;
  border: none;
  width: 60px;
  height: 60px;
  font-size: 40px;
  background-color: #1fc6e6;
  color: #f7f9fe;
  box-shadow: 10px 10px 20px rgba(0, 0, 0, 0.3),
    -10px -10px 20px rgba(255, 255, 255, 0.3);
}

.card-list-item {
  display: flex;
  flex-direction: column;
  gap: 5px;
  font-weight: 300;
  /*width: 100%;*/
  padding: 10px;
  margin: 0 auto 10px auto;
  text-align: center;
  background: #ffffff;
  border-radius: 10px;
}

.card-list-item hr {
  display: flex;
  position: relative;
  margin: 8px 0 0 0;
  border: 1px dashed #4259ef23;
  width: 100%;
}

.card-footer {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: baseline;
}

:deep(.card-content) {
  height: 100%;
  width: 100%;
  word-wrap: break-word;
  text-align: left;
}
</style>

<style>
.card-content > ul {
  margin-block-start: 0em;
  margin-block-end: 0em;
  padding-inline-start: 20px;
}
</style>
