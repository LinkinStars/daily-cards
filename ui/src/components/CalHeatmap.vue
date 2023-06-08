<script setup lang="ts">
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn';
import localeData from 'dayjs/plugin/localeData';
dayjs.extend(localeData);
dayjs.locale('zh-cn');

import { ref } from 'vue';
import CalHeatmap from 'cal-heatmap';
import Tooltip from 'cal-heatmap/plugins/Tooltip';
import LegendLite from 'cal-heatmap/plugins/LegendLite';
import CalendarLabel from 'cal-heatmap/plugins/CalendarLabel';
import 'cal-heatmap/cal-heatmap.css';

const cal: CalHeatmap = new CalHeatmap();

function monthFormat(date) { return dayjs(date).format("MMMM") }

function paintCalendar() {
  const today: Date = new Date(); // 获取今天的日期
  const nextMonth: Date = new Date(today.getFullYear() - 1, today.getMonth() + 2, 1); // 获取下一个月的第一天
  const formattedDate: string = nextMonth.toISOString().slice(0, 10); // 将日期格式化为 'YYYY-MM-DD'
  console.log(formattedDate); // 输出：'2019-07-01'

  cal.paint(
    {
      data: { source: calData.value, x: 'date', y: 'value' },
      date: {
        locale: { 'name': 'zh-cn', 'weekStart': 1 },
        start: new Date(formattedDate),
      },
      range: 12,
      scale: {
        color: {
          //type: 'threshold',
          range: ['#8E949E', '#166b34', '#37a446', '#4dd05a'],
          //domain: [10, 20, 30],
          domain: [0, 5, 10, 15, 20, 25, 30],
          type: 'ordinal',
        },
      },
      domain: {
        type: 'month',
        gutter: 4,
        label: { text: monthFormat, textAlign: 'start', position: 'top' },
      },
      subDomain: {
        type: 'ghDay',
        radius: 2,
        width: 15,
        height: 15,
        gutter: 4,
      },
    },
    [
      [
        Tooltip,
        {
          text: function (date, value, dayjsDate) {
            return (
              dayjs(date).format('YYYY-MM-DD') + " " + (value ? '已打卡' : '未打卡')
            );
          },
        },
      ],
      [
        CalendarLabel,
        {
          width: 30,
          textAlign: 'start',
          text: () => dayjs.weekdaysShort(),
          padding: [25, 0, 0, 0],
        },
      ],
    ]
  );
}

import { getCardsStat } from "@/api/card";
interface CheckData {
  date: string;
  value: number;
}

const calData = ref<CheckData[]>([]);
const setCardsStat = async () => {
  const resp = await getCardsStat();
  if (resp.code != 200) {
    return;
  }
  for (let i = 0; i < resp.data.checked_days.length; i++) {
    const date = resp.data.checked_days[i];
    const value = 15;
    calData.value.push({ date, value });
  }
  paintCalendar()
}
setCardsStat();

</script>

<template>
  <div id="heatmap" class="heatmap-bg">
    <div id="cal-heatmap"></div>
  </div>
</template>

<style>
.heatmap-bg {
  display: flex;
  justify-content: center;
  /* background: #eceef0; */
  background: #ffffff;
  padding: 10px;
  border-radius: 10px;
  margin-bottom: 10px;
}

@media (max-width: 1100px) {
  #heatmap {
    display: none;
  }
}
</style>
