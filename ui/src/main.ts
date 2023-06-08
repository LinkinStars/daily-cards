import { createApp } from 'vue'
import router from './router'
import './style.css'
import App from './App.vue'
import Vue3TouchEvents from "vue3-touch-events";
import hljs from "highlight.js";
import 'highlight.js/styles/github-dark-dimmed.css';

import * as dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
import * as localeData from 'dayjs/plugin/localeData';
dayjs.extend(localeData);
dayjs.locale('zh-cn');


const app = createApp(App)
app.use(Vue3TouchEvents)
app.directive("highlight", function(el) {
  let blocks = el.querySelectorAll("pre code");
  blocks.forEach(block => {
    hljs.highlightElement(block);
  });
});
app.config.globalProperties.$hljs = hljs;
app.config.globalProperties.dayjs=dayjs
app.use(router)
app.mount('#app')
