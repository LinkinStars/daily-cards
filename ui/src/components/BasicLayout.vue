<template>
  <LinkCorner v-if="!hideLinkCorner"></LinkCorner>
  
  <!-- 顶部导航栏 - Twitter 风格 -->
  <header class="navbar bg-base-100 border-b border-base-300 fixed top-0 z-50 px-4 md:px-8 shadow-sm">
    <div class="navbar-start">
      <div class="flex items-center gap-3 cursor-pointer" @click="clearQueryAndJumpCardPage">
        <img src="/icon/favicon-60.png" alt="logo" class="w-10 h-10 rounded-xl ring ring-primary ring-offset-base-100 ring-offset-1" />
        <span class="text-xl md:text-2xl font-bold text-primary">{{ headerTitle }}</span>
      </div>
    </div>
    
    <div class="navbar-center hidden md:flex">
      <div class="join">
        <input 
          type="text" 
          v-model="searchText" 
          placeholder="搜索卡片..." 
          class="input input-bordered join-item w-80 focus:outline-none"
          @keyup.enter="jumpCardPage()"
        />
        <button class="btn btn-primary join-item" @click="jumpCardPage()">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </button>
      </div>
    </div>

    <div class="navbar-end">
      <!-- 移动端菜单 -->
      <div class="dropdown dropdown-end md:hidden">
        <label tabindex="0" class="btn btn-ghost btn-circle">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </label>
        <ul tabindex="0" class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow-lg bg-base-100 rounded-box w-52 border border-base-300">
          <li>
            <div class="form-control">
              <input 
                type="text" 
                v-model="searchText" 
                placeholder="搜索卡片..." 
                class="input input-bordered input-sm w-full"
                @keyup.enter="jumpCardPage()"
              />
            </div>
          </li>
          <li v-if="isLogin"><a @click="handleLogout">退出登录</a></li>
          <li v-else><a @click="jumpToLoginPage">登录</a></li>
        </ul>
      </div>
      
      <!-- 桌面端用户按钮 -->
      <div v-if="isLogin" class="dropdown dropdown-end hidden md:block">
        <label tabindex="0" class="btn btn-ghost btn-circle avatar">
          <div class="w-10 rounded-full ring ring-primary ring-offset-base-100 ring-offset-1">
            <img :src="userAvatar" alt="avatar" />
          </div>
        </label>
        <ul tabindex="0" class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow-lg bg-base-100 rounded-box w-40 border border-base-300">
          <li><a @click="handleLogout">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
            </svg>
            退出登录
          </a></li>
        </ul>
      </div>
      
      <!-- 未登录时显示登录按钮 -->
      <button 
        v-else
        class="btn btn-ghost btn-sm hidden md:flex gap-2"
        @click="jumpToLoginPage"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1" />
        </svg>
        登录
      </button>
    </div>
  </header>

  <!-- 主内容区域 -->
  <main class="mt-16 bg-base-200">
    <router-view :key="route.query" />
  </main>
</template>

<script lang="ts" setup>
import { ref, computed } from "vue";
import LinkCorner from "./LinkCorner.vue";
import { getSiteInfo } from "../api/site";
import { useRouter, useRoute } from "vue-router";
const router = useRouter();
const route = useRoute();

const headerTitle = ref("Daily Cards");
const hideLinkCorner = ref(true);
const searchText = ref("");
const isLogin = ref(!!localStorage.getItem("accessToken"));
const userAvatar = ref("/icon/favicon-60.png");

const setupSiteInfo = async () => {
  const resp = await getSiteInfo();
  if (resp.code === 200 && resp.data.site_name.length > 0) {
    headerTitle.value = resp.data.site_name;
    hideLinkCorner.value = resp.data.hide_link_corner;
    localStorage.setItem("siteInfo", JSON.stringify(resp.data));
    
    // Update avatar from site info
    if (resp.data.avatar) {
      userAvatar.value = resp.data.avatar;
    }
    
    // Update login state based on server response
    isLogin.value = resp.data.is_login;
    
    // If token exists but server says not logged in, clear token
    let token = localStorage.getItem("accessToken");
    if (token && !resp.data.is_login) {
      localStorage.removeItem("accessToken");
      isLogin.value = false;
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

const handleLogout = () => {
  localStorage.removeItem("accessToken");
  isLogin.value = false;
  router.push({ name: "card-page" });
};

setupSiteInfo();
</script>
