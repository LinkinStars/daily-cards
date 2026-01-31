<script setup lang="ts">
import { ref, reactive } from "vue";
import { getCard } from "../api/card";
import getBackgroundColor from "@/utils/bg-color.ts";
import useClipboard from "vue-clipboard3";
import { toPng } from 'html-to-image';
import QrcodeVue from 'qrcode.vue'
import VueEasyLightbox from 'vue-easy-lightbox';
import { useRouter } from "vue-router";
import { showError, showSuccess } from '@/utils/toast';

const router = useRouter();

const { toClipboard } = useClipboard();
const props = defineProps<{ id: number }>();
const cardInfo = reactive({
  id: 0,
  created_at: "",
  content: "",
  nickname: "",
  avatar: "",
});
const showTip = ref(false);
const showCapture = ref(false);
const showTooltip = ref(true);
const imgsRef = ref([])

const copyURL = async () => {
  var url = shareURL()
  try {
    await toClipboard(url);
    showSuccess("链接已复制到剪切板");
  } catch (e) {
    console.error(e);
  }
};

const shareURL = () => {
   return window.location.protocol + "//" + window.location.host + "/card/" + props.id;
}

const getCardByID = async () => {
  const res = await getCard(props.id);
  if (res.code != 200) {
    router.push({ name: "not-found" });
  } else {
    cardInfo.id = res.data.id;
    cardInfo.created_at = res.data.created_at;
    cardInfo.content = res.data.content;
    cardInfo.nickname = res.data.nickname;
    cardInfo.avatar = res.data.avatar;
  }
  if (cardInfo.id === 0) {
    router.push({ name: "not-found" });
  }
};
getCardByID();

const capture = async (showPop : boolean) => {
  try {
    const element = document.querySelector("#card-bg") as HTMLElement;
    if (!element) {
      console.error('[Card] element not found');
      showError('未找到卡片元素');
      return;
    }
    
    // 隐藏提示文字
    showTooltip.value = false;
    // 等待DOM更新
    await new Promise(resolve => setTimeout(resolve, 50));
    
    console.log('[Card] generating image with html-to-image');
    const dataUrl = await toPng(element, {
      quality: 0.95,
      pixelRatio: 1,
      skipFonts: true,
      cacheBust: true,
    });
    
    // 恢复提示文字
    showTooltip.value = true;
    
    console.log('[Card] image generated successfully');
    
    if (showPop) {
      console.log('[Card] showing popup with image');
      imgsRef.value = dataUrl;
      showCapture.value = true;
    } else {
      console.log('[Card] attempting to download image');
      const a = document.createElement('a');
      const filename = 'card_' + cardInfo.created_at + '_' + cardInfo.id + '.png';
      console.log('[Card] downloading as:', filename);
      a.href = dataUrl;
      a.download = filename;
      a.click();
    }
  } catch (err) {
    showTooltip.value = true;
    console.error('[Card] image generation error:', err);
    showError('生成图片失败，请稍后重试');
  }
};

let siteInfo = JSON.parse(localStorage.getItem('siteInfo') || '{}');
const showQrcode = siteInfo.show_qrcode;
</script>

<template>
  <div class="w-full flex justify-center px-4">
    <div class="flex flex-col lg:flex-row gap-4 items-center w-full" style="max-width: 850px;">
      <!-- 卡片主体 -->
      <div id="card-bg" :style="getBackgroundColor(props.id)" class="rounded-3xl shadow-2xl p-4 w-full lg:w-auto" style="max-width: 770px;">
        <div class="card bg-base-100 bg-opacity-95 backdrop-blur-md">
          <div class="card-body p-6">
            <!-- 用户信息 -->
            <div class="flex items-center gap-3">
              <div class="avatar cursor-pointer flex-shrink-0" @click="capture(true)">
                <div class="w-12 h-12 rounded-full">
                  <img :src="cardInfo.avatar" alt="avatar" />
                </div>
              </div>
              <div class="flex flex-col justify-center">
                <h2 class="text-base font-bold leading-tight">{{ cardInfo.nickname }}</h2>
                <span class="text-sm text-base-content/50">{{ cardInfo.created_at.split(' ')[0] }}</span>
              </div>
            </div>

            <!-- 卡片内容 -->
            <div class="markdown-content text-left" v-html="cardInfo.content" v-highlight></div>

            <!-- 二维码 -->
            <div v-if="showQrcode" class="flex justify-end">
              <div class="p-1.5 bg-base-200/50 rounded-lg backdrop-blur-sm">
                <qrcode-vue :value="shareURL()" :size="50"></qrcode-vue>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 操作按钮 - 桌面端右侧垂直排列，移动端底部水平排列 -->
      <div class="flex lg:flex-col flex-row gap-3 lg:gap-4">
        <button 
          @click="copyURL()" 
          class="group relative btn btn-circle btn-md lg:btn-lg border-0 bg-base-100 shadow-xl hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 tooltip tooltip-top lg:tooltip-left flex items-center justify-center"
          data-tip="复制分享链接"
        >
          <div class="absolute inset-0 rounded-full bg-gradient-to-br from-blue-400 to-cyan-300 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 lg:h-6 lg:w-6 relative z-10 text-blue-500 group-hover:text-white transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
          </svg>
        </button>
        
        <button 
          @click="capture(false)" 
          class="group relative btn btn-circle btn-md lg:btn-lg border-0 bg-base-100 shadow-xl hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 tooltip tooltip-top lg:tooltip-left flex items-center justify-center"
          data-tip="生成卡片图片"
        >
          <div class="absolute inset-0 rounded-full bg-gradient-to-br from-purple-400 to-pink-300 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 lg:h-6 lg:w-6 relative z-10 text-purple-500 group-hover:text-white transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
        </button>
      </div>
    </div>
  </div>

  <!-- 图片预览 -->
  <vue-easy-lightbox
    :visible="showCapture"
    :moveDisabled="true"
    :imgs="imgsRef"
    @hide="showCapture = false"
  ></vue-easy-lightbox>
</template>

<style scoped>
:deep(.markdown-content) img {
  @apply w-full rounded-lg my-4;
}

:deep(.markdown-content) ul {
  @apply list-disc ml-5 space-y-2;
}

:deep(.markdown-content) ol {
  @apply list-decimal ml-5 space-y-2;
}

:deep(.markdown-content) blockquote {
  @apply border-l-4 border-primary pl-4 py-2 my-4 bg-primary/5 rounded-r-lg italic;
}

:deep(.markdown-content) code {
  @apply bg-base-200 px-1.5 py-0.5 rounded text-sm font-mono;
}

:deep(.markdown-content) pre {
  @apply bg-base-300 p-4 rounded-lg overflow-x-auto my-4;
}

:deep(.markdown-content) pre code {
  @apply bg-transparent p-0;
}

:deep(.markdown-content) a {
  @apply text-primary hover:underline;
}

:deep(.markdown-content) h1,
:deep(.markdown-content) h2,
:deep(.markdown-content) h3 {
  @apply font-bold mt-6 mb-3;
}

:deep(.markdown-content) p {
  @apply mb-4 leading-relaxed;
}
</style>
