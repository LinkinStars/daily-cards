<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { getCardsStat } from "@/api/card";
import { getSiteInfo } from "../api/site";
import getBackgroundColor from "@/utils/bg-color.ts";
import useClipboard from "vue-clipboard3";
import html2canvas from "html2canvas-pro";
import QrcodeVue from 'qrcode.vue'
import VueEasyLightbox from 'vue-easy-lightbox';
import { useRouter } from "vue-router";
import CalHeatmap from '@/components/CalHeatmap.vue';
import dayjs from "dayjs";
import { showError, showSuccess } from '@/utils/toast';

const router = useRouter();

const { toClipboard } = useClipboard();
const cardInfo = reactive({
  created_at: "",
  nickname: "",
  avatar: "",
  checked_yearly: 0,
  checked_total: 0,
  day_total: 0,
});
// 获取今天的时间格式为 2021-01-01
const today = dayjs();
const todayStr = today.format("YYYY-MM-DD");
cardInfo.created_at = todayStr;

const showTip = ref(false);
const showCapture = ref(false);
const showTooltip = ref(true);
const imgsRef = ref<string | string[]>([])

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
   return window.location.protocol + "//" + window.location.host + "/card/stat";
}

const getCardByID = async () => {
  const res = await getSiteInfo();
  if (res.code == 200) {
    cardInfo.nickname = res.data.nickname;
    cardInfo.avatar = res.data.avatar;
  }
};
getCardByID();

const capture = async (showPop : boolean) => {
  try {
    const element = document.querySelector("#card-bg") as HTMLElement;
    if (!element) {
      console.error('[StatCard] element not found');
      showError('未找到卡片元素');
      return;
    }
    
    // 隐藏提示文字
    showTooltip.value = false;
    
    // 等待DOM更新
    await new Promise(resolve => setTimeout(resolve, 100));
    
    console.log('[StatCard] generating image with html2canvas');
    
    // 使用 html2canvas - 对 SVG 元素（如 cal-heatmap）支持更好
    const canvas = await html2canvas(element, {
      scale: 1,
      useCORS: true,
      allowTaint: true,
      backgroundColor: null,
    });
    
    // 恢复提示文字
    showTooltip.value = true;
    
    console.log('[StatCard] image generated successfully');
    
    if (showPop) {
      // 微信等不允许下载，只能生成进行长按保存
      const url = canvas.toDataURL('image/png');
      imgsRef.value = url;
      showCapture.value = true;
    } else {
      // 正常浏览器可以下载
      canvas.toBlob((blob) => {
        if (!blob) {
          showError('生成图片失败');
          return;
        }
        const a = document.createElement('a');
        const url = window.URL.createObjectURL(blob);
        const filename = 'card_' + cardInfo.created_at + '_stat.png';
        a.href = url;
        a.download = filename;
        a.click();
        window.URL.revokeObjectURL(url);
      });
    }
  } catch (err) {
    showTooltip.value = true;
    console.error('[StatCard] image generation error:', err);
    showError('生成图片失败，请稍后重试');
  }
};

let siteInfo = JSON.parse(localStorage.getItem('siteInfo') || '{}');
const showQrcode = siteInfo.show_qrcode;

const setCardsStat = async () => {
  const resp = await getCardsStat();
  if (resp.code != 200) {
    return;
  }
  cardInfo.checked_yearly = resp.data.checked_days.length;
  cardInfo.checked_total = resp.data.checked_total;
  cardInfo.day_total = resp.data.day_total;
  if (cardInfo.day_total == 0) {
    cardInfo.day_total = 1;
  }
}
setCardsStat();
</script>

<template>
  <div class="w-full flex justify-center">
    <div class="flex gap-4 items-center" style="max-width: calc(100vw - 2rem);">
      <!-- 卡片主体 -->
      <div id="card-bg" :style="getBackgroundColor(dayjs().day())" class="rounded-3xl shadow-2xl p-4">
        <div class="card bg-base-100/95 backdrop-blur-md">
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
                <span class="text-sm text-base-content/50">{{ cardInfo.created_at }}</span>
              </div>
              <p class="text-sm text-base-content/60 ml-auto">近一年打卡记录</p>
            </div>

            <!-- 日历热力图 -->
            <div class="w-full max-w-full overflow-hidden mb-4 bg-white rounded-lg p-2">
              <CalHeatmap :months="12" />
            </div>

            <!-- 统计数据和二维码 -->
            <div class="flex items-center gap-6">
              <div class="grid grid-cols-4 gap-6 flex-1">
              <div class="flex items-center gap-1.5">
                <svg width="20" height="20" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-primary"><g clip-path="url(#icon-4bd80d63353e4f38)"><path d="M42 20V39C42 40.6569 40.6569 42 39 42H9C7.34315 42 6 40.6569 6 39V9C6 7.34315 7.34315 6 9 6H30" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M16 20L26 28L41 7" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></g><defs><clipPath id="icon-4bd80d63353e4f38"><rect width="48" height="48" fill="#FFF"/></clipPath></defs></svg>
                <div class="text-sm">
                  <p class="text-base-content/70">近一年</p>
                  <p class="text-xl font-bold text-primary">{{ cardInfo.checked_yearly }}</p>
                </div>
              </div>
              
              <div class="flex items-center gap-1.5">
                <svg width="20" height="20" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-secondary"><path d="M14 24L15.25 25.25M44 14L24 34L22.75 32.75" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M4 24L14 34L34 14" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
                <div class="text-sm">
                  <p class="text-base-content/70">总打卡</p>
                  <p class="text-xl font-bold text-secondary">{{ cardInfo.checked_total }}</p>
                </div>
              </div>
              
              <div class="flex items-center gap-1.5">
                <svg width="20" height="20" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-accent"><path d="M5.81836 6.72729V14H13.0911" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M4 24C4 35.0457 12.9543 44 24 44V44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C16.598 4 10.1351 8.02111 6.67677 13.9981" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M24.005 12L24.0038 24.0088L32.4832 32.4882" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
                <div class="text-sm">
                  <p class="text-base-content/70">已坚持</p>
                  <p class="text-xl font-bold text-accent">{{ cardInfo.day_total }}</p>
                </div>
              </div>
              
              <div class="flex items-center gap-1.5">
                <svg width="20" height="20" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-info"><path d="M30.2972 18.7786C30.2972 18.7786 27.0679 27.8809 25.5334 29.47C23.9988 31.0591 21.4665 31.1033 19.8774 29.5687C18.2882 28.0341 18.244 25.5019 19.7786 23.9127C21.3132 22.3236 30.2972 18.7786 30.2972 18.7786Z" fill="currentColor" stroke="currentColor" stroke-width="4" stroke-linejoin="round"/><path d="M38.8492 38.8492C42.6495 35.049 45 29.799 45 24C45 12.402 35.598 3 24 3C12.402 3 3 12.402 3 24C3 29.799 5.35051 35.049 9.15076 38.8492" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M24 4V8" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M38.8454 11.1421L35.7368 13.6593" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M42.5223 27.2328L38.6248 26.333" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M5.47737 27.2328L9.37485 26.333" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M9.15463 11.142L12.2632 13.6593" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
                <div class="text-sm">
                  <p class="text-base-content/70">百分比</p>
                  <p class="text-xl font-bold text-info">{{ (cardInfo.checked_total/cardInfo.day_total*100).toFixed(2) }}%</p>
                </div>
              </div>
            </div>

            <!-- 二维码 -->
            <div v-if="showQrcode">
              <qrcode-vue :value="shareURL()" :size="60"></qrcode-vue>
            </div>
          </div>
          </div>
        </div>
      </div>

      <!-- 右侧操作按钮 -->
      <div class="flex flex-col gap-4">
        <button 
          @click="copyURL()" 
          class="group relative btn btn-circle btn-lg border-0 bg-base-100 shadow-xl hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 tooltip tooltip-left flex items-center justify-center"
          data-tip="复制分享链接"
        >
          <div class="absolute inset-0 rounded-full bg-gradient-to-br from-blue-400 to-cyan-300 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 relative z-10 text-blue-500 group-hover:text-white transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
          </svg>
        </button>
        
        <button 
          @click="capture(false)" 
          class="group relative btn btn-circle btn-lg border-0 bg-base-100 shadow-xl hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 tooltip tooltip-left flex items-center justify-center"
          data-tip="生成卡片图片"
        >
          <div class="absolute inset-0 rounded-full bg-gradient-to-br from-purple-400 to-pink-300 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 relative z-10 text-purple-500 group-hover:text-white transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
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
</style>
