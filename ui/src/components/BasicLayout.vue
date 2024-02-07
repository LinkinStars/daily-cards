<template>
  <LinkCorner v-if="!hideLinkCorner"></LinkCorner>
  <div class="card-list-header">
    <img src="/icon/favicon-60.png" alt="logo" style="pointer-events: none" />
    <span @click="jumpCardPage()">{{ headerTitle }}</span>
    <div class="carr-list-header-jumper" v-touch:tap="clearQueryAndJumpCardPage" v-touch:longtap="jumpToLoginPage"></div>
    <input type="text" v-model="searchText" placeholder="搜索" @keyup.enter.native="jumpCardPage()"/>
  </div>
  <div class="main-container">
    <router-view :key="route.query" />
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import LinkCorner from "./LinkCorner.vue";
import { getSiteInfo } from "../api/site";
import { useRouter, useRoute } from "vue-router";
const router = useRouter();
const route = useRoute();

const headerTitle = ref("Daily Cards");
const hideLinkCorner = ref(true);
const searchText = ref("")

const setupSiteInfo = async () => {
  const resp = await getSiteInfo();
  if (resp.code === 200 && resp.data.site_name.length > 0) {
    headerTitle.value = resp.data.site_name;
    hideLinkCorner.value = resp.data.hide_link_corner;
    localStorage.setItem("siteInfo", JSON.stringify(resp.data));
    // 如果 token 存在且 token 过期, 直接删除并跳转登录
    let token = localStorage.getItem("accessToken");
    if (token && !resp.data.is_login) {
      localStorage.removeItem("accessToken");
      router.push({ name: "user-login" });
    }
  }
};

const clearQueryAndJumpCardPage = () => {
  searchText.value = "";
  jumpCardPage();
};

const jumpCardPage = async () => {
  router.push({ name: "card-page", query: { q: searchText.value }});
};

const jumpToLoginPage = async () => {
  router.push({ name: "user-login" });
};
setupSiteInfo();
</script>

<style>
.card-list-header {
  background-color: #f7f9fe;
  width: 100%;
  height: 60px;
  display: flex;
  flex-direction: row;
  /* justify-content: space-between; */
  align-items: center;
  position: fixed;
  z-index: 99;
  overflow: hidden;
}

.card-list-header span {
  color: #363636;
  font-size: 24px;
  font-weight: bold;
  line-height: 1;
  /* width: 100%; */
  text-align: left;
  margin-left: 12px;
}

.card-list-header input {
  width: 20%;
  height: 32px;
  border-radius: 16px;
  border: 1px solid #eaeaea;
  padding: 0 12px;
  margin-left: 12px;
}

@media (max-width: 768px) {
  .card-list-header input {
    display: none;
  }
}

.card-list-header img {
  margin-left: 24px;
  width: 32px;
  height: 32px;
}

@media screen and (max-width: 768px) {
  .main-container {
    height: 100%;
    overflow: auto;
  }
}

.carr-list-header-jumper {
  width: 32px;
  height: 32px;
  position: absolute;
  margin-left: 24px;
}

.main-container {
  overflow: auto;
  height: 100%;
  width: 100%;
}
</style>
