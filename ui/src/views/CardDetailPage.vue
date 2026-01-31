<script setup lang="ts">
import { ref } from "vue"
import Card from '../components/Card.vue';
import { useRoute, useRouter } from 'vue-router'
import { getCard, getCardWithOffset, deleteCard } from "@/api/card";
import { showSuccess, showError, showInfo } from "@/utils/toast";

const route = useRoute()
const router = useRouter()
const cardID = ref(+route.params.id)

const isLogin = ref(false);
if (localStorage.getItem("accessToken")) {
  isLogin.value = true;
}

// Confirm modal state
const showDeleteModal = ref(false);

const confirmDelete = () => {
  showDeleteModal.value = true;
};

const deleteCardByID = async () => {
  showDeleteModal.value = false;
  const res = await deleteCard(cardID.value);
  if (res.code === 200) {
    showSuccess("删除成功");
    router.push({ name: "card-page" });
  }
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
      name: "card-detail",
      params: {
        id: res.data.id,
      },
    });
  } else {
    showInfo("已无数据");
  }
};

const jumpCardPage = async () => {
  router.push({ name: "card-page" });
};

const getCardByID = async () => {
  if (!isLogin.value) return;
  
  const res = await getCard(cardID.value);
  if (res.code != 200) {
    showError("未查询到数据");
    return
  }
  if (!res.data.is_login) {
    router.push({ name: "user-login" });
  }
};
getCardByID();
</script>

<template>
  <div class="bg-base-200 pt-8 pb-8" style="min-height: calc(100vh - 64px);">
    <div class="container mx-auto">
      <Card :id="cardID" />
      
      <!-- 操作按钮（仅登录用户可见） -->
      <div v-if="isLogin" class="mt-6 flex flex-wrap gap-2 justify-center">
        <button @click="otherCardDetailPage(-1)" class="btn btn-outline btn-primary btn-sm">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
          上一张
        </button>
        
        <button @click="jumpCardPage()" class="btn btn-outline btn-sm">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
          </svg>
          返回
        </button>
        
        <button @click="jumpUpdateCardPage()" class="btn btn-primary btn-sm">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
          </svg>
          修改
        </button>
        
        <button @click="confirmDelete()" class="btn btn-error btn-sm">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
          删除
        </button>
        
        <button @click="otherCardDetailPage(+1)" class="btn btn-outline btn-primary btn-sm">
          下一张
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </button>
      </div>
    </div>
  </div>

  <!-- 删除确认弹窗 -->
  <dialog :class="['modal', { 'modal-open': showDeleteModal }]">
    <div class="modal-box">
      <h3 class="font-bold text-lg">确认删除</h3>
      <p class="py-4">确定要删除这张卡片吗？此操作不可恢复。</p>
      <div class="modal-action">
        <button class="btn" @click="showDeleteModal = false">取消</button>
        <button class="btn btn-error" @click="deleteCardByID()">删除</button>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop" @click="showDeleteModal = false">
      <button>close</button>
    </form>
  </dialog>
</template>

<style scoped>
</style>
