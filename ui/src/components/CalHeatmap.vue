<script setup lang="ts">
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn';
import localeData from 'dayjs/plugin/localeData';
dayjs.extend(localeData);
dayjs.locale('zh-cn');

import { ref, defineEmits, defineProps } from 'vue';
import CalHeatmap from 'cal-heatmap';
import Tooltip from 'cal-heatmap/plugins/Tooltip';
import LegendLite from 'cal-heatmap/plugins/LegendLite';
import CalendarLabel from 'cal-heatmap/plugins/CalendarLabel';
import 'cal-heatmap/cal-heatmap.css';

const props = defineProps({
  months: {
    type: Number,
    default: 3  // 默认3个月，适合侧边栏
  }
});

const emit = defineEmits(['clickBox']);
const cal: CalHeatmap = new CalHeatmap();

function monthFormat(date) { return dayjs(date).format("MMMM") }

function refreshCalendar(x) {
    // 使用传入的月份数
    let duration = props.months;
    let dateRange = props.months;
    
    // 手机端根据原始月份调整
    if (x.matches) {
      duration = Math.min(props.months, 3);  // 手机端最多3个月
      dateRange = Math.min(props.months, 3);
    }

    // 当超过 20 号之后，如果这个月有 31 天会导致显示不全，所以需要额外显示下个月多出来的几天
    if (dayjs().day() > 20) {
      duration = duration - 1;
    }
    let startTime = dayjs().subtract(duration, 'month').startOf('month').add(1, 'month').format('YYYY-MM-DD');
    paintCalendar(startTime, dateRange);
}
 
function paintCalendar(startTime : string, dateRange : number) {
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
          range: ['#e5e7eb', '#166b34', '#37a446', '#4dd05a'],
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
          text: () => ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
          padding: [25, 0, 0, 0],
        },
      ],
    ]
  );
  cal.on('click', (event, timestamp, value) => {
    const date = dayjs(timestamp).format('YYYY-MM-DD');
    emit('clickBox', date)
  });
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
  <div id="heatmap" class="bg-white overflow-hidden">
    <div id="cal-heatmap" class="bg-white"></div>
  </div>
</template>

<style>
#heatmap {
  background-color: white !important;
  width: 100%;
  overflow: hidden;
}
#cal-heatmap {
  background-color: white !important;
  width: fit-content;
  max-width: 100%;
}
#cal-heatmap .ch-container {
  background-color: white !important;
}
/* 确保日历svg背景为白色 */
#cal-heatmap svg {
  background-color: white !important;
  display: block;
}
/* 隐藏SVG溢出部分 */
#cal-heatmap svg rect,
#cal-heatmap svg text {
  fill-opacity: 1;
}
</style>
