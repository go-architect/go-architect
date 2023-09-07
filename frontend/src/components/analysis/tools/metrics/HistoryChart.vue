<script lang="ts" setup>
import { Line } from 'vue-chartjs'
import { Chart as ChartJS } from 'chart.js'
import { Title, Tooltip, Legend, PointElement, LineElement, CategoryScale, LinearScale } from 'chart.js';
import {onMounted, reactive} from "vue";

const props = defineProps(['data'])

const localData = reactive({
  chartOptions: {
    responsive: true,
    interaction: {
      mode: 'index' as 'index',
      intersect: false,
    },
    scales: {
      y: {
        type: 'linear',
        min: 0,
        display: true,
        position: 'left',
      },
      y1: {
        type: 'linear',
        min: 0,
        max: 100,
        display: 'auto',
        position: 'right',
        grid: {
          drawOnChartArea: false, // only want the grid lines for one axis to show up
        },
      }
    }
  },
})

ChartJS.register(Title, Tooltip, Legend, LineElement, PointElement, CategoryScale, LinearScale)
</script>

<template>
  <div>
    <Line
        id="types-chart"
        :options="localData.chartOptions"
        :data="data"
    />
  </div>
</template>

<style scoped>

</style>
