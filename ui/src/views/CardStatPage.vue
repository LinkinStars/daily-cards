<script setup lang="ts">
import StatCard from '../components/StatCard.vue';
</script>

<template>
  <div class="card-detail-bg">
    <div class="card-detail">
      <StatCard />
    </div>
  </div>
</template>

<style scoped>
.card-detail-bg {
  display: flex;
  flex-direction: row;
  justify-content: center;
  background: linear-gradient(45deg, #c4ddfd 0%, #ffffff 100%);
  width: 100%;
  min-height: 100%;
}
.card-detail {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: calc(100% - 60px);
  max-width: 1200px;
  width: 100%;
}
</style>
