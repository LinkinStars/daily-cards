<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { userLogin } from "../api/user";

const router = useRouter();

const username = ref("");
const password = ref("");
const isLoading = ref(false);
const errorMessage = ref("");

const login = async () => {
  if (!username.value || !password.value) {
    errorMessage.value = "请输入用户名和密码";
    return;
  }

  isLoading.value = true;
  errorMessage.value = "";

  try {
    const resp = await userLogin(username.value, password.value);
    if (resp.code === 200) {
      localStorage.setItem("accessToken", resp.data.token);
      router.push({ name: "card-page" });
    } else {
      errorMessage.value = "登录失败，请检查用户名和密码";
    }
  } catch (error) {
    errorMessage.value = "登录失败，请稍后重试";
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <div class="h-screen flex items-center justify-center bg-base-200 overflow-hidden">
    <div class="card w-full max-w-md bg-base-100 shadow-2xl m-4">
      <div class="card-body">
        <!-- Logo 和标题 -->
        <div class="text-center mb-6">
          <div class="avatar">
            <div class="w-20 h-20">
              <img src="/icon/favicon-60.png" alt="logo" />
            </div>
          </div>
          <h1 class="text-3xl font-bold mt-4 text-primary">日拱一卒</h1>
          <p class="text-base-content/70 mt-2">记录每一天的成长</p>
        </div>

        <!-- 错误提示 -->
        <div v-if="errorMessage" class="alert alert-error shadow-lg mb-4">
          <div>
            <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current flex-shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span>{{ errorMessage }}</span>
          </div>
        </div>

        <!-- 登录表单 -->
        <form @submit.prevent="login" class="space-y-4">
          <div class="form-control">
            <label class="label">
              <span class="label-text">用户名</span>
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none text-base-content/70">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
              </div>
              <input
                type="text"
                v-model="username"
                placeholder="请输入用户名"
                class="input input-bordered w-full pl-10"
                required
              />
            </div>
          </div>

          <div class="form-control">
            <label class="label">
              <span class="label-text">密码</span>
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none text-base-content/70">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </div>
              <input
                type="password"
                v-model="password"
                placeholder="请输入密码"
                class="input input-bordered w-full pl-10"
                required
                @keyup.enter="login()"
              />
            </div>
          </div>

          <div class="form-control mt-6">
            <button
              type="submit"
              class="btn btn-primary w-full"
              :disabled="isLoading"
            >
              <span v-if="isLoading" class="loading loading-spinner loading-sm"></span>
              {{ isLoading ? '登录中...' : '登录' }}
            </button>
          </div>
        </form>

        <!-- 底部提示 -->
        <div class="divider mt-6">Daily Cards</div>
        <div class="text-center text-sm text-base-content/70">
          <p>坚持每天记录，见证成长轨迹</p>
        </div>
      </div>
    </div>
  </div>
</template>
