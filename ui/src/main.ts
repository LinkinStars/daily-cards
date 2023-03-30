import { createApp } from 'vue'
import router from './router'
import './style.css'
import App from './App.vue'
import Vue3TouchEvents from "vue3-touch-events";
import hljs from "highlight.js";
import 'highlight.js/styles/github-dark-dimmed.css';


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
