<script setup lang="ts">
import * as dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
import * as localeData from 'dayjs/plugin/localeData';
dayjs.extend(localeData);
dayjs.locale('zh-cn');

import { onMounted } from 'vue';
import CalHeatmap from 'cal-heatmap';
import Tooltip from 'cal-heatmap/plugins/Tooltip';
import LegendLite from 'cal-heatmap/plugins/LegendLite';
import CalendarLabel from 'cal-heatmap/plugins/CalendarLabel';
import 'cal-heatmap/cal-heatmap.css';

const cal: CalHeatmap = new CalHeatmap();

function monthFormat(date) { return dayjs(date).format("MMMM")}

function paintCalendar() {
const calData = [
  { date: '2023-01-01', value: 1 },
  { date: '2023-01-02', value: 2 },
  { date: '2023-01-03', value: 3 },
  { date: '2023-01-04', value: 4 },
  { date: '2023-01-05', value: 5 },
  { date: '2023-01-06', value: 6 },
  { date: '2023-01-07', value: 7 },
  { date: '2023-01-08', value: 8 },
  { date: '2023-01-09', value: 9 },
  { date: '2023-01-10', value: 20 },
  { date: '2023-01-11', value: 20 },
  { date: '2023-01-12', value: 20 },
  { date: '2023-01-13', value: 20 },
  { date: '2023-01-14', value: 20 },
  { date: '2023-01-15', value: 20 },
  { date: '2023-01-16', value: 20 },
  { date: '2023-01-17', value: 20 },
  { date: '2023-01-18', value: 20 },
  { date: '2023-01-19', value: 20 },
  { date: '2023-01-20', value: 20 },
  { date: '2023-01-21', value: 20 },
  { date: '2023-01-22', value: 20 },
  { date: '2023-01-23', value: 20 },
  { date: '2023-01-24', value: 20 },
  { date: '2023-01-25', value: 20 },
  { date: '2023-01-26', value: 20 },
  { date: '2023-01-27', value: 20 },
  { date: '2023-01-28', value: 20 },
  { date: '2023-01-29', value: 20 },
  { date: '2023-01-30', value: 20 },
  { date: '2023-01-31', value: 20 },
  { date: '2023-02-01', value: 20 },
  { date: '2023-02-02', value: 20 },
  { date: '2023-02-03', value: 20 },
  { date: '2023-02-04', value: 20 },
  { date: '2023-02-05', value: 20 },
  { date: '2023-02-06', value: 20 },
  { date: '2023-02-07', value: 20 },
  { date: '2023-02-08', value: 20 },
  { date: '2023-02-09', value: 20 },
  { date: '2023-02-10', value: 20 },
  { date: '2023-02-11', value: 20 },
  { date: '2023-02-12', value: 20 },
  { date: '2023-02-13', value: 20 },
  { date: '2023-02-14', value: 20 },
  { date: '2023-02-15', value: 20 },
  { date: '2023-02-16', value: 20 },
  { date: '2023-02-17', value: 20 },
  { date: '2023-02-18', value: 49 },
  { date: '2023-02-19', value: 50 },
  { date: '2023-02-20', value: 51 },
  { date: '2023-02-21', value: 52 },
  { date: '2023-06-08', value: 15 },
];

  cal.paint(
  {
    data: { source: calData, x: 'date', y: 'value' },
    date: { 
      locale: { name:'zh-cn', weekStart: 1 },
      start: new Date('2023-01-01'),
    },
    range: 12,
    scale: {
      color: {
        //type: 'threshold',
        range: ['#8E949E',  '#166b34', '#37a446', '#4dd05a'],
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

onMounted(() => paintCalendar());

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
  background: #eceef0;
  padding: 10px;
  border-radius: 10px;
  margin-bottom: 10px;
}
@media (max-width: 1000px) {
  #heatmap {
    display: none;
  }
}
</style>
