<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { addCard, updateCard, getCard } from "@/api/card.ts";
import { saveDraft, readDraft, removeDraft } from "@/utils/post-draft.ts";
import Vditor from 'vditor';
import 'vditor/dist/index.css';

const vditor = ref<Vditor | null>(null);

const route = useRoute();
const router = useRouter();

const cardID = ref(+route.params.id);
const content = ref("");
const pv = ref(0);

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

const loadPreCard = async () => {
  if (vditor.value!.getValue().length > 0) {
    if (!window.confirm("当前已经有内容，确定要重新加载吗？")) {
      return;
    }
  }
  await loadCardInfo(0)
  // 当加载上一次卡片的时候只需要加载内容，pv 的统计不需要加载
  pv.value = 0
  currentDate.value = today
};

const loadCardInfo = async (id : number) => {
    const res = await getCard(id);
    if (res.code != 200) {
      alert("未查询到数据");
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
        name: "user-card-detail",
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
    height: 400,
    toolbar: ['headings', 'bold', 'italic', 'strike', '|', 'line', 'quote', 'list', 'ordered-list', 'check', 'outdent', 'indent', 'code', 'inline-code', '|', 'insert-after', 'insert-before', 'undo', 'redo', 'link'],
    toolbarConfig: {
      pin: true,
    },
    preview: {
      hljs: {
        style: 'dracula'
      },
      actions: []
    },
    cache: {
      id: 'post-draft',
      enable: true,
    },
    after: () => {
      getCardInfo();
    },
  });
});
</script>

<template>
  <div class="card-edit-container-bg">
    <div class="card-edit-container">
      <div style="height: 20px"></div>
      <div><input class="card-edit-date-input" v-model="currentDate"/> {{ "PV:" + pv }}</div>
      <div style="height: 20px"></div>
      <div class="card-editor">
        <!-- <textarea autofocus v-model="content" rows="20" @input="inputPost"></textarea> -->
        <div id="vditor" style="height: 200px;" />
      </div>
      <div style="height: 20px"></div>
      <div class="card-edit-btn">
        
        <button @click="jumpCardPage()">返回</button>
        <button @click="loadPreCard()">加载最新</button>
        <button :disabled="postButtonFlag" @click="postCard()">{{ cardID > 0 ? '修改' : '发布' }}</button>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.card-edit-container-bg {
  display: flex;
  flex-direction: row;
  justify-content: center;

  background: linear-gradient(-45deg, #ee7752, #e73c7e, #23a6d5, #23d5ab);
  background-size: 400% 400%;

  width: 100%;
  min-height: 100%;
}

.card-edit-container {
  padding: 60px 20px 0 20px;
  min-height: calc(100% - 60px);
  max-width: 800px;
  width: calc(100% - 40px);

  align-items: center;
  flex-direction: column;

  min-height: calc(100% - 60px);
}
.card-edit-date-input {
  background-color: transparent;
  border: none;
  width: 90px;
}

@media screen and (max-width: 768px) {
  .card-edit-container {
    padding: 60px 20px 0 20px;
  }
}

.card-editor {
  width: 100%;
  text-align: left;
}

.card-editor textarea {
  display: inline-block;
  width: 100%;
  height: 100%;
  vertical-align: top;
  box-sizing: border-box;
  border: none;
  resize: none;
  outline: none;
  font-size: 14px;
  padding: 20px;
}

.card-edit-btn {
  display: flex;
  flex-direction: row;
  justify-content: space-between !important;
}

.card-edit-btn > button {
  display: inline-flex;
  justify-content: center;
  align-items: center;
  line-height: 1;
  height: 32px;
  white-space: nowrap;
  cursor: pointer;
  color: #409eff;
  text-align: center;
  box-sizing: border-box;
  outline: none;
  transition: 0.1s;
  font-weight: 500;
  user-select: none;
  vertical-align: middle;
  -webkit-appearance: none;
  background-color: #ecf5ff;
  border: 1px solid #dcdfe6;
  border-color: #a0cfff;
  padding: 8px 15px;
  border-radius: 4px;
}
</style>
