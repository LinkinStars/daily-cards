<script setup lang="ts">
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import InfiniteLoading from "v3-infinite-loading";
import "v3-infinite-loading/lib/style.css";
import { getCardsPage, Card } from "@/api/card";
import { uploadAvatar } from "@/api/site";
import CalHeatmap from "@/components/CalHeatmap.vue";
import useClipboard from "vue-clipboard3";
import { showSuccess } from "@/utils/toast";

const router = useRouter();
const route = useRoute();
const { toClipboard } = useClipboard();

const cards = ref<Card[]>([]);
let page = 1;
let pageSize = 20;
const load = async ($state) => {
  try {
    const resp = await getCardsPage(
      page,
      route.query.q as string,
      route.query.d as string
    );
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

const isLogin = ref(false);
if (localStorage.getItem("accessToken")) {
  isLogin.value = true;
}

const jumpCardDetailPage = (id: number) => {
  router.push({ name: "card-detail", params: { id: id } });
};

const jumpPostPage = () => {
  router.push({ name: "user-card-post", params: { id: 0 } });
};

const jumpDayCardPage = async (date: string) => {
  router.push({ name: "card-page", query: { d: date } });
};

const jumpCardStatPage = async () => {
  router.push({ name: "card-stat" });
};

const shareCard = async (id: number) => {
  const url = window.location.protocol + "//" + window.location.host + "/card/" + id;
  try {
    await toClipboard(url);
    showSuccess("链接已复制到剪切板");
  } catch (e) {
    console.error(e);
  }
};

let siteInfo = JSON.parse(localStorage.getItem('siteInfo') || '{}');
const nickname = siteInfo.nickname;
const avatar = siteInfo.avatar;

const tryToUploadAvatar = async (e: any) => {
  const file = e.target.files[0];
  const formData = new FormData();
  formData.append("file", file);
  const res = await uploadAvatar(formData);
  if (res.code === 200) {
    location.reload();
  }
};

const handleAvatarClick = () => {
  if (!isLogin.value) {
    return;
  }
  const input = document.getElementById("upload") as HTMLInputElement;
  input.click();
};
</script>

<template>
  <div class="min-h-screen bg-base-200">
    <div class="container mx-auto px-4 py-6 max-w-6xl">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- 左侧主内容区 -->
        <div class="lg:col-span-2 space-y-4">
          <!-- 卡片列表 -->
          <div v-for="card in cards" :key="card.id" class="card bg-base-100 shadow-md hover:shadow-xl transition-shadow">
            <div class="card-body p-6">
              <!-- 卡片内容 -->
              <div class="markdown-content cursor-pointer" v-html="card.content" v-highlight @click="jumpCardDetailPage(card.id)"></div>

              <!-- 互动按钮 -->
              <div class="flex items-center gap-6 pt-3 text-base-content/60 border-t border-base-300 mt-4">
                <!-- 查看次数 -->
                <div class="flex items-center gap-2" v-if="isLogin">
                  <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                  <span class="text-sm">{{ card.pv || '' }}</span>
                </div>

                <!-- 分享 -->
                <button class="flex items-center gap-2 hover:text-success transition-colors group" @click="shareCard(card.id)">
                  <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 group-hover:fill-success/20" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                  </svg>
                  <span class="text-sm">分享</span>
                </button>

                <!-- 查看详情按钮 -->
                <button class="btn btn-sm btn-ghost text-primary ml-auto" @click="jumpCardDetailPage(card.id)">
                  查看详情
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                  </svg>
                </button>
              </div>
            </div>
          </div>

          <!-- 无限滚动加载 -->
          <div class="flex justify-center py-8">
            <InfiniteLoading @infinite="load">
              <template #complete>
                <div class="flex items-center justify-center gap-2 text-base-content/40">
                  <span class="w-12 h-px bg-base-content/20"></span>
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="currentColor" viewBox="0 0 20 20">
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                  </svg>
                  <span class="w-12 h-px bg-base-content/20"></span>
                </div>
              </template>
              <template #error>
                <div class="text-error/60 text-sm">加载失败，请重试</div>
              </template>
            </InfiniteLoading>
          </div>
        </div>

        <!-- 右侧边栏 -->
        <div class="lg:col-span-1 space-y-4">
          <!-- 日历热力图卡片 -->
          <div class="card bg-base-100 shadow-md sticky top-20">
            <div class="card-body">
              <div class="flex items-center gap-2 mb-2">
                <h3 class="card-title text-base">打卡日历</h3>
                <button 
                  class="btn btn-circle btn-xs btn-ghost flex items-center justify-center tooltip tooltip-bottom" 
                  data-tip="查看统计"
                  @click="jumpCardStatPage"
                >
                  <svg width="14" height="14" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M5.81836 6.72729V14H13.0911" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
                    <path d="M4 24C4 35.0457 12.9543 44 24 44V44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C16.598 4 10.1351 8.02111 6.67677 13.9981" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
                    <path d="M24.005 12L24.0038 24.0088L32.4832 32.4882" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
                  </svg>
                </button>
              </div>
              <CalHeatmap @clickBox="jumpDayCardPage" />
              <div class="divider my-2"></div>
              <div class="flex items-center justify-end gap-3">
                <div class="avatar cursor-pointer" @click="handleAvatarClick">
                  <div class="w-10 h-10 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2">
                    <img :src="avatar" alt="avatar" />
                  </div>
                </div>
                <input id="upload" type="file" accept="image/png, image/jpeg" @change="tryToUploadAvatar" style="display: none" />
                <p class="font-medium">{{ nickname }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 浮动创建按钮（仅登录用户可见） -->
    <button 
      v-if="isLogin" 
      class="btn btn-circle btn-primary btn-lg fixed bottom-8 right-8 shadow-2xl z-40"
      @click="jumpPostPage()"
    >
      <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
      </svg>
    </button>
  </div>
</template>

<style scoped>
:deep(.markdown-content) {
  @apply text-left;
}

:deep(.markdown-content) img {
  @apply w-full rounded-lg my-4;
}

:deep(.markdown-content) ul {
  @apply list-disc ml-5 space-y-2 my-4;
}

:deep(.markdown-content) ol {
  @apply list-decimal ml-5 space-y-2 my-4;
}

:deep(.markdown-content) p {
  @apply mb-3 leading-relaxed;
}

:deep(.markdown-content) a {
  @apply text-primary hover:underline;
}

:deep(.markdown-content) code {
  @apply bg-base-200 px-1.5 py-0.5 rounded text-sm font-mono;
}

:deep(.markdown-content) pre {
  @apply bg-base-300 p-4 rounded-lg overflow-x-auto my-4;
}

:deep(.markdown-content) blockquote {
  @apply border-l-4 border-primary pl-4 py-2 my-4 bg-primary/5 rounded-r-lg italic;
}
</style>
