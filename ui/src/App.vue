<script setup lang="ts">
import Toast from './components/Toast.vue';
import { toastMessage } from './utils/toast';
</script>

<template>
  <router-view />
  <Toast 
    v-if="toastMessage" 
    :message="toastMessage.message" 
    :type="toastMessage.type" 
    :duration="toastMessage.duration"
  />
</template>

<style>
</style>
