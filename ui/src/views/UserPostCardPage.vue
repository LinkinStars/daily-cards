<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { addCard, updateCard, getCard } from "@/api/card.ts";
import { saveDraft, readDraft, removeDraft } from "@/utils/post-draft.ts";
import { showError } from "@/utils/toast";
import Vditor from 'vditor';
import 'vditor/dist/index.css';

const vditor = ref<Vditor | null>(null);

const route = useRoute();
const router = useRouter();

const cardID = ref(+route.params.id);
const content = ref("");
const pv = ref(0);

// Confirm modal state
const showLoadModal = ref(false);

let nowDate = new Date();
let date = {
  year: nowDate.getFullYear(),
  month: nowDate.getMonth() + 1,
  date: nowDate.getDate(),
};

const currentDate = ref(date.year + "-");
if (date.month < 10) {
  currentDate.value += "0" + date.month + "-";
} else {
  currentDate.value += date.month + "-";
}
if (date.date < 10) {
  currentDate.value += "0" + date.date;
} else {
  currentDate.value += date.date;
}
const today = currentDate.value;

const getCardInfo = async () => {
  if (cardID.value > 0) {
    loadCardInfo(cardID.value)
    return
  }
  let draft = readDraft()
  if (draft) {
    content.value = draft
  }
};

const tryLoadPreCard = () => {
  if (vditor.value!.getValue().trim().length > 0) {
    showLoadModal.value = true;
    return;
  }
  doLoadPreCard();
};

const doLoadPreCard = async () => {
  showLoadModal.value = false;
  await loadCardInfo(0);
  pv.value = 0;
  currentDate.value = today;
};

const loadCardInfo = async (id : number) => {
    const res = await getCard(id);
    if (res.code != 200) {
      showError("未查询到数据");
      return
    }
    content.value = res.data.original_text;
    pv.value = res.data.pv;
    currentDate.value = res.data.created_at;
    vditor.value!.setValue(content.value);
}

const postButtonFlag = ref(false);
const postCard = async () => {
  content.value = vditor.value!.getValue();
  postButtonFlag.value = true;
  if (cardID.value > 0) {
    const res = await updateCard(cardID.value, content.value, currentDate.value);
    if (res.code === 200) {
      removeDraft()
      router.push({
        name: "card-detail",
        params: { id: cardID.value },
      });
    }
  } else {
    const res = await addCard(content.value);
    if (res.code === 200) {
      removeDraft()
      router.push({ name: "card-page" });
    }
  }
  postButtonFlag.value = false;
};

const addContent = (c: string) => {
  if (content.value.length === 0) {
    content.value = c;
    return;
  }
  content.value += "\n" + c;
};

const jumpCardPage = async () => {
  router.push({ name: "card-page" });
};

const inputPost = () => {
  saveDraft(content.value)
};

onMounted(() => {
  vditor.value = new Vditor('vditor', {
    height: 'calc(100vh - 340px)',
    minHeight: 400,
    toolbar: ['headings', 'bold', 'italic', 'strike', '|', 'line', 'quote', 'list', 'ordered-list', 'check', 'outdent', 'indent', 'code', 'inline-code', '|', 'insert-after', 'insert-before', 'undo', 'redo', 'link', 'table', 'upload'],
    toolbarConfig: {
      pin: true,
    },
    preview: {
      hljs: {
        style: 'github-dark-dimmed'
      },
      actions: []
    },
    cache: {
      id: 'post-draft',
      enable: true,
    },
    tab: '  ',
    cdn: 'https://unpkg.com/vditor@3.11.2',
    after: () => {
      getCardInfo();
    },
  });
});
</script>

<template>
  <div class="bg-base-200 pt-6 pb-6" style="min-height: calc(100vh - 64px);">
    <div class="container mx-auto px-4 max-w-7xl">
      <!-- 顶部工具栏 -->
      <div class="card bg-base-100 shadow-md mb-4">
        <div class="card-body p-4">
          <div class="flex flex-wrap justify-between items-center gap-4">
            <div class="flex items-center gap-4">
              <button class="btn btn-ghost btn-sm" @click="jumpCardPage()">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
                </svg>
                返回
              </button>
              
              <div class="divider divider-horizontal mx-0"></div>
              
              <div class="flex items-center gap-2">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-base-content/70" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                <input 
                  type="date" 
                  v-model="currentDate" 
                  class="input input-bordered input-sm w-40"
                />
              </div>
              
              <div v-if="pv > 0" class="badge badge-ghost gap-2">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
                {{ pv }} 次查看
              </div>
            </div>

            <div class="flex gap-2">
              <button class="btn btn-sm btn-outline" @click="tryLoadPreCard()">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                </svg>
                加载最新
              </button>
              <button 
                class="btn btn-sm btn-primary"
                :class="{ 'loading': postButtonFlag }"
                :disabled="postButtonFlag" 
                @click="postCard()"
              >
                <svg v-if="!postButtonFlag" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                {{ cardID > 0 ? '保存修改' : '发布卡片' }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 编辑器区域 -->
      <div class="card bg-base-100 shadow-md">
        <div class="card-body p-0">
          <div id="vditor" class="rounded-lg overflow-hidden" />
        </div>
      </div>

      <!-- 提示信息 -->
      <div class="alert shadow-lg mt-4 mb-0">
        <div>
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info flex-shrink-0 w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
          <span>内容会自动保存为草稿，刷新页面不会丢失</span>
        </div>
      </div>
    </div>
  </div>

  <!-- 加载确认弹窗 -->
  <dialog :class="['modal', { 'modal-open': showLoadModal }]">
    <div class="modal-box">
      <h3 class="font-bold text-lg">确认加载</h3>
      <p class="py-4">当前已经有内容，确定要重新加载吗？</p>
      <div class="modal-action">
        <button class="btn" @click="showLoadModal = false">取消</button>
        <button class="btn btn-primary" @click="doLoadPreCard()">确定加载</button>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop" @click="showLoadModal = false">
      <button>close</button>
    </form>
  </dialog>
</template>

<style>
/* Vditor 编辑器自定义样式 */
.vditor {
  @apply border-0;
}

.vditor-toolbar {
  @apply bg-base-200 border-b border-base-300;
}

.vditor-toolbar--pin {
  @apply bg-base-200;
}

.vditor-content {
  @apply bg-base-100;
}
</style>
