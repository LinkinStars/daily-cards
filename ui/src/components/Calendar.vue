<script setup lang="ts">
import { ref } from 'vue'
import { getCardsStat, CardStat } from "@/api/card";

const cardsStat = ref<CardStat[]>([]);

const setCardsStat = async () => {
  const resp = await getCardsStat();
  if (resp.code === 200) {
    cardsStat.value.push(...resp.data);
  }
}
setCardsStat();

const getDayClass = (checked : boolean) => {
  if (checked) {
    return "calendar-day calendar-day-checked"
  } else {
    return "calendar-day"
  }
}

let nowDate = new Date();
const curMonth = ref(nowDate.getMonth() + 1)

var currentDate = new Date();
currentDate.setDate(1);
var firstDayOfWeek = currentDate.getDay();
console.log("当前月份第一天是星期" + firstDayOfWeek);
</script>

<template>
<div class="calendar-bg">
  <div class="calendar-month">{{ curMonth }} 月</div>
  <div class="calendar-day-grid">
    <div :class="getDayClass(false)" v-for="index of firstDayOfWeek-1"></div>
    <div :class="getDayClass(card.checked)" v-for="card in cardsStat">{{ card.day }}</div>
  </div>
</div>
</template>

<style scoped>
.calendar-bg {
  background: #eceef0;
  padding: 10px;
  border-radius: 10px;
  margin-bottom: 10px;
  min-height: 230px;
}
.calendar-month {
  margin-bottom: 5px;
}
.calendar-day-grid {
  width: 100%;
  display: grid;
  grid-template-columns: repeat(7, 14.28%);
  grid-row-gap: 2px;
  grid-column-gap: 1px;
  justify-items: center;
  align-items: center;
}
.calendar-day {
  width: 30px;
  height: 30px;
  line-height: 30px;
  text-align: center;
  /**flex: 1;
  flex-basis: 14%;**/
}
.calendar-day-checked {
  border-radius: 50%; /* 将边框半径设置为元素宽度和高度的一半，即可画出一个圆形 */
  background-color: #029FFB; /* 设置背景颜色 */
  color: #FFF;
}
</style>
