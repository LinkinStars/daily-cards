<script setup lang="ts">
import { ref, reactive } from "vue";
import { getCard } from "../api/card";
import getBackgroundColor from "@/utils/bg-color.ts";
import useClipboard from "vue-clipboard3";
import html2canvas from "html2canvas";
import QrcodeVue from 'qrcode.vue'
import VueEasyLightbox from 'vue-easy-lightbox';

const { toClipboard } = useClipboard();
const props = defineProps<{ id: number }>();
const cardInfo = reactive({
  id: 0,
  created_at: "",
  content: "",
  nickname: "",
  avatar: "",
});
const showShareTip = ref(false);
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
   return window.location.protocol + "//" + window.location.host + "/card/" + props.id;
}

const getCardByID = async () => {
  const res = await getCard(props.id);
  if (res.code != 200) {
    alert("未查询到数据");
  } else {
    cardInfo.id = res.data.id;
    cardInfo.created_at = res.data.created_at;
    cardInfo.content = res.data.content;
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
          var filename = 'card_' + cardInfo.created_at + '_' + cardInfo.id + '.png'
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
</script>

<template>
  <div class="card-border">
    <div id="card-bg" :style="getBackgroundColor(props.id)">
    <div class="card">
      <div class="user-info">
        <img :src="cardInfo.avatar" @click="capture(true)" />
        <div class="user-nickname">
          <p>{{ cardInfo.nickname }}</p>
        </div>
      </div>
      <div class="card-content" v-html="cardInfo.content"></div>
      <div v-if="showQrcode" class="share-qrcode">
        <qrcode-vue :value="shareURL()" :size="40"></qrcode-vue>
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
