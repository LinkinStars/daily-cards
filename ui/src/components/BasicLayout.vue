<template>
  <div class="card-list-header" @click="jumpCardPage()">
    <img src="/icon/favicon-60.png" alt="logo"/>
    <span>{{ headerTitle }}</span>
  </div>
  <div class="main-container">
    <router-view />
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { getSiteInfo } from "../api/site";
import { useRouter } from "vue-router";
const router = useRouter();

const headerTitle = ref("Daily Cards")

const setupSiteInfo = async () => {
  const resp = await getSiteInfo();
  if (resp.code === 200 && resp.data.site_name.length > 0) {
    headerTitle.value = resp.data.site_name;
    localStorage.setItem("siteInfo", JSON.stringify(resp.data));
  }
};

const jumpCardPage = async () => {
  let token = localStorage.getItem('accessToken')
  if (token) {
    router.push({ name: "user-card-page" });
  } else {
    router.push({ name: "card-page" });
  }
};
setupSiteInfo()
</script>

<style>
.card-list-header{
  background-color: #F7F9FE;
  width: 100%;
  height: 60px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  position: fixed;
  z-index: 9999;
  overflow: hidden;
}

.card-list-header span {
  color: #363636;
  font-size: 24px;
  font-weight: bold;
  line-height: 1;
  width: 100%;
  text-align: left;
  margin-left: 12px;
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

.main-container {
  overflow: auto;
  height: 100%;
  width: 100%;
}
</style>
