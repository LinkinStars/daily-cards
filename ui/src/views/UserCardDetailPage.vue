<script setup lang="ts">
import { ref } from "vue";
import Card from "../components/Card.vue";
import { getCard, getCardWithOffset } from "../api/card";
import { useRoute, useRouter } from "vue-router";
import { deleteCard } from "@/api/card.ts";

const route = useRoute();
const router = useRouter();

const cardID = ref(+route.params.id);

const deleteCardByID = async () => {
  if (!window.confirm("确定要删除吗？")) {
    console.log("取消");
    return;
  }
  const res = await deleteCard(cardID.value);
  if (res.code === 200) {
    alert("删除成功");
    router.push({ name: "user-card-page" });
  }
  return;
};

const jumpUpdateCardPage = async () => {
  router.push({
    name: "user-card-post",
    params: {
      id: cardID.value,
    },
  });
};

const otherCardDetailPage = async (offset: number) => {
  const res = await getCardWithOffset(cardID.value, offset);
  if (res.code == 200) {
    router.push({
      name: "user-card-detail",
      params: {
        id: res.data.id,
      },
    });
  } else {
    alert("已无数据");
  }
};

const jumpCardPage = async () => {
  router.push({ name: "user-card-page" });
};

const getCardByID = async () => {
  const res = await getCard(cardID.value);
  if (res.code != 200) {
    alert("未查询到数据");
    return
  }
  if (!res.data.is_login) {
    router.push({ name: "user-login" });
  }
};
getCardByID();
</script>

<template>
  <div class="card-detail-bg">
    <div class="card-detail">
      <div style="height: 60px"></div>
      <Card :id="cardID" />
      <div style="height: 20px"></div>
      <div class="card-detail-btn">
        <button @click="otherCardDetailPage(-1)">上一张</button>
        <button @click="jumpCardPage()">返回</button>
        <button @click="jumpUpdateCardPage()">修改</button>
        <button @click="deleteCardByID()">删除</button>
        <button @click="otherCardDetailPage(+1)">下一张</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.card-detail-bg {
  display: flex;
  flex-direction: row;
  justify-content: center;
  background: linear-gradient(45deg, #c4ddfd 0%, #ffffff 100%);
  width: 100%;
  min-height: 100%;
}
.card-detail {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: calc(100% - 60px);
  max-width: 800px;
  width: 100%;
}
.card-detail-btn {
  display: flex;
  gap: 10px;
}
@media screen and (max-width: 768px) {
  .card-detail-btn {
    margin-bottom: 160px;
  }
}
.card-detail-btn button {
  display: inline-flex;
  justify-content: center;
  align-items: center;
  line-height: 1;
  height: 32px;
  white-space: nowrap;
  cursor: pointer;
  color: #409eff;
  text-align: center;
  box-sizing: border-box;
  outline: none;
  transition: 0.1s;
  font-weight: 500;
  user-select: none;
  vertical-align: middle;
  -webkit-appearance: none;
  background-color: #ecf5ff;
  border: 1px solid #dcdfe6;
  border-color: #a0cfff;
  padding: 8px 15px;
  border-radius: 4px;
}
</style>
