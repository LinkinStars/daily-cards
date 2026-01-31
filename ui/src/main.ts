import { createApp } from 'vue'
import router from './router'
import './style.css'
import 'vditor/dist/index.css'
import 'highlight.js/styles/github-dark-dimmed.css'
import App from './App.vue'
import Vue3TouchEvents from "vue3-touch-events"
import hljs from "highlight.js"
import { registerSW } from 'virtual:pwa-register'

// 注册 Service Worker
const updateSW = registerSW({
  onNeedRefresh() {
    if (confirm('有新版本可用，是否刷新页面？')) {
      updateSW(true)
    }
  },
  onOfflineReady() {
    console.log('应用已可离线使用')
  },
})

const app = createApp(App)
app.use(Vue3TouchEvents)
app.directive("highlight", function(el) {
  let blocks = el.querySelectorAll("pre code");
  blocks.forEach(block => {
    hljs.highlightElement(block);
  });
});
app.config.globalProperties.$hljs = hljs;
app.use(router)
app.mount('#app')
