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

const preYearDate: Date = dayjs().subtract(11, 'month').startOf('month');
const preYearStr: string = preYearDate.format('YYYY-MM-DD');
const pre3MonthDate: Date = dayjs().subtract(2, 'month').startOf('month');
const pre3MonthStr: string = pre3MonthDate.format('YYYY-MM-DD');

function refreshCalendar(x) {
    if (x.matches) {
      // 手机端仅展示三个月
      paintCalendar(pre3MonthStr, 3);
    } else {
      // 电脑短展示一年
      paintCalendar(preYearStr, 12);
    }
}
 
function paintCalendar(startTime : string, dateRange : int) {
  cal.paint(
    {
      data: { source: calData.value, x: 'date', y: 'value' },
      date: {
        locale: { 'name': 'zh-cn', 'weekStart': 1 },
        start: new Date(startTime),
      },
      range: dateRange,
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
  var mwMatchMedia = window.matchMedia("(max-width: 1100px)")
  refreshCalendar(mwMatchMedia)
  mwMatchMedia.addListener(refreshCalendar)
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
    /*display: none;*/
  }
}
</style>
