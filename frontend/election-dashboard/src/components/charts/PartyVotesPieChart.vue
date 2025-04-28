<template>
  <Pie :data="chartData" :options="chartOptions" />
</template>

<script>
import { Pie } from 'vue-chartjs'
import {
  Chart as ChartJS,
  ArcElement,
  Tooltip,
  Legend
} from 'chart.js'

import ChartDataLabels from "chartjs-plugin-datalabels"

ChartJS.register(ArcElement, Tooltip, Legend, ChartDataLabels)
export default {
  name: "PartyVotesPieChart",
  components: { Pie },
  props: {
    data: {
      type: Object,  // ← we fixed it to Object earlier
      required: true
    }
  },
  computed: {
    chartData() {
      return this.data;
    },
    chartOptions() {
      return {
        responsive: true,
        plugins: {
          legend: {
            position: 'bottom',
          },
          datalabels: {
            formatter: (value, ctx) => {
              const sum = ctx.chart.data.datasets[0].data.reduce((a, b) => a + b, 0);
              const percentage = ((value / sum) * 100).toFixed(1) + '%';
              return percentage;
            },
            color: '#fff',
            font: {
              weight: 'bold',
              size: 16
            }
          }
        },
      }
    }
  }
}
</script>
