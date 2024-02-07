<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { getCardsStat } from "@/api/card";
import { getSiteInfo } from "../api/site";
import getBackgroundColor from "@/utils/bg-color.ts";
import useClipboard from "vue-clipboard3";
import html2canvas from "html2canvas";
import QrcodeVue from 'qrcode.vue'
import VueEasyLightbox from 'vue-easy-lightbox';
import { useRouter } from "vue-router";
import CalHeatmap from '@/components/CalHeatmap.vue';
import dayjs from "dayjs";

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
const imgsRef = ref([])

const copyURL = async () => {
  var url = shareURL()
  try {
    await toClipboard(url);
    showTip.value = true;
   setTimeout(() => {
      showTip.value = false;
    }, 5000);
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
  const element = document.querySelector("#card-bg");
  html2canvas(element, {
    scale: 2, // 设置截图分辨率为原来的2倍
    useCORS: true, // 允许使用跨域图片
    proxy: window.location.protocol + "//" + window.location.host,
    allowTaint: true, // 允许截图跨域图片
    backgroundColor: null,
  }).then((canvas) => {
    // 微信不允许下载，只能生成进行长按保存
    if (showPop) {
      let url = canvas.toDataURL('image/png')
      imgsRef.value = url
      showCapture.value = true
    } else {
      // 正常浏览器应该可以下载
      canvas.toBlob(function(blob){
          var a = document.createElement('a')
          var url = window.URL.createObjectURL(blob)
          var filename = 'card_' + cardInfo.created_at + '_stat' + '.png'
          a.href = url
          a.download = filename
          a.click()
          // 当图片文件加载完成释放这个url
          window.URL.revokeObjectURL(url)
      })
    }
    //document.querySelector('#result').innerHTML = `<img src="${url}" alt="海报"  width="100%" height="100%"/>`
    //const image = canvas
    //  .toDataURL("image/png")
    //  .replace("image/png", "image/octet-stream");
    //const link = document.createElement("a");
    //link.download = "screenshot.png";
    //link.href = image;
    //link.click();
  });
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
  <div class="card-border">
    <div id="card-bg" :style="getBackgroundColor(dayjs().day())">
    <div class="card">
      <div class="user-info">
        <img :src="cardInfo.avatar" @click="capture(true)" />
        <div class="user-nickname">
          <p>{{ cardInfo.nickname }} 近一年打卡记录</p>
        </div>
      </div>
      <CalHeatmap />
      <div class="card-stat-grand">
        <div class="card-stat-col">
          <div class="card-stat-desc">
            <svg width="24" height="24" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><g clip-path="url(#icon-4bd80d63353e4f38)"><path d="M42 20V39C42 40.6569 40.6569 42 39 42H9C7.34315 42 6 40.6569 6 39V9C6 7.34315 7.34315 6 9 6H30" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M16 20L26 28L41 7" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></g><defs><clipPath id="icon-4bd80d63353e4f38"><rect width="48" height="48" fill="#FFF"/></clipPath></defs></svg>
            <p style="margin-left: 5px;">近一年：<em>{{ cardInfo.checked_yearly }}</em>次</p>
          </div>
          <div class="card-stat-desc">
            <svg width="24" height="24" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M14 24L15.25 25.25M44 14L24 34L22.75 32.75" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M4 24L14 34L34 14" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
            <p style="margin-left: 5px;">总打卡：<em>{{ cardInfo.checked_total }}</em>次</p>
          </div>
          <div class="card-stat-desc">
            <svg width="24" height="24" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M5.81836 6.72729V14H13.0911" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M4 24C4 35.0457 12.9543 44 24 44V44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C16.598 4 10.1351 8.02111 6.67677 13.9981" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M24.005 12L24.0038 24.0088L32.4832 32.4882" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
            <p style="margin-left: 5px;">已坚持：<em>{{ cardInfo.day_total }}</em>天</p>
          </div>
          <div class="card-stat-desc">
            <svg width="24" height="24" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M30.2972 18.7786C30.2972 18.7786 27.0679 27.8809 25.5334 29.47C23.9988 31.0591 21.4665 31.1033 19.8774 29.5687C18.2882 28.0341 18.244 25.5019 19.7786 23.9127C21.3132 22.3236 30.2972 18.7786 30.2972 18.7786Z" fill="#333" stroke="#333" stroke-width="4" stroke-linejoin="round"/><path d="M38.8492 38.8492C42.6495 35.049 45 29.799 45 24C45 12.402 35.598 3 24 3C12.402 3 3 12.402 3 24C3 29.799 5.35051 35.049 9.15076 38.8492" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M24 4V8" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M38.8454 11.1421L35.7368 13.6593" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M42.5223 27.2328L38.6248 26.333" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M5.47737 27.2328L9.37485 26.333" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M9.15463 11.142L12.2632 13.6593" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
            <p style="margin-left: 5px;">百分比：<em>{{ (cardInfo.checked_total/cardInfo.day_total*100).toFixed(2) }}%</em></p>
          </div>
        </div>
        <div v-if="showQrcode" class="share-qrcode">
          <qrcode-vue :value="shareURL()" :size="50"></qrcode-vue>
        </div>
      </div>
      <hr />
      <div class="card-footer">
        <span>📅 {{ cardInfo.created_at }}</span>
        <svg @click="copyURL()" v-touch:hold="(event) => { capture(false) }" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"
          height="1.5em" width="1.5em">
          <path 
            d="M404 344a75.9 75.9 0 0 0-60.208 29.7l-163.923-93.036a75.693 75.693 0 0 0 0-49.328L343.792 138.3a75.937 75.937 0 1 0-13.776-28.976L163.3 203.946a76 76 0 1 0 0 104.108l166.717 94.623A75.991 75.991 0 1 0 404 344Zm0-296a44 44 0 1 1-44 44 44.049 44.049 0 0 1 44-44ZM108 300a44 44 0 1 1 44-44 44.049 44.049 0 0 1-44 44Zm296 164a44 44 0 1 1 44-44 44.049 44.049 0 0 1-44 44Z"
            />
        </svg>
      </div>
    </div>
  </div>
  </div>
  <div v-if="showTip" class="copy-tip">
    已复制链接到剪切板<br />长按保存图片<br/>也可点击头像后长按保存
  </div>
  <vue-easy-lightbox
      :visible="showCapture"
      :moveDisabled="true"
      :imgs="imgsRef"
      @hide="showCapture = false"
    ></vue-easy-lightbox>
  <div id="result" v-show="showCapture" @click="showCapture = false">
  </div>
</template>

<style scoped>
.card-border {
  padding: 10px;
  width: calc(100% - 40px);
}
#card-bg {
  width: calc(100% - 40px);
  padding: 20px;
  box-shadow: -1px 15px 30px -12px black;
  border-radius: 10px;
}

.card {
  padding: 10px;
  backdrop-filter: blur(16px) saturate(180%);
  -webkit-backdrop-filter: blur(16px) saturate(180%);
  background-color: rgba(255, 255, 255, 0.75);
  border-radius: 12px;
  border: 1px solid rgba(209, 213, 219, 0.3);
}

.user-info {
  display: flex;
  flex-direction: row;
  align-items: center;
}

.user-info>img {
  border-radius: 50%;
  width: 3em;
  height: 3em;
  margin-right: 0.75em;
}

.user-nickname {
  font-size: 1.2em;
}

.card-content {
  text-align: left;
}

:deep(.card-content) {
  height: 100%;
  width: 100%;
  word-wrap: break-word;
  text-align: left;
}

hr {
  display: flex;
  position: relative;
  margin: 8px 0 8px 0;
  border: 1px dashed #4259ef23;
  width: 100%;
}

.card-footer {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: baseline;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.copy-tip {
  position: fixed;
  bottom: 100px;
  left: 50%;
  transform: translateX(-50%);
  padding: 10px;
  border-radius: 5px;
  background-color: #eee;
}

.share-qrcode {
  margin-right: 0px;
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
}
.card-stat-grand {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
}
.card-stat-col {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 25px;
}
@media screen and (max-width: 768px) {
  .card-stat-col {
    flex-direction: column;
    align-items: flex-start;
    gap: 0px;
  }
}
.card-stat-desc {
  display: flex;
  flex-direction: row;
  align-items: center;
}
.card-stat-desc > p > em {
  font-size: 1.5rem;
}
</style>

<style>
.card-content>ul {
  margin-block-start: 0em;
  margin-block-end: 0em;
  padding-inline-start: 20px;
}
img {
  width: 100%;
}
blockquote {
    margin: 0 0 1rem;
    padding: 1rem;
    border-left: 0.25rem solid #49b1f5;
    background-color: rgba(73,177,245,0.1);
    color: #6a737d;
}
blockquote > :last-child {
    margin-bottom: 0 !important;
}
blockquote > p {
    margin: 0 0 16px;
}
</style>
