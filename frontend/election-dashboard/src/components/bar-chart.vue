<!-- components/BarChart.vue -->
<template>
  <Bar :data="chartData" :options="chartOptions" />
</template>

<script>
import { Bar } from 'vue-chartjs'
import {
  Chart as ChartJS,
  BarElement,
  CategoryScale,
  LinearScale,
  Tooltip,
  Legend,
} from 'chart.js'

ChartJS.register(BarElement, CategoryScale, LinearScale, Tooltip, Legend)

export default {
  name: 'BarChart',
  components: { Bar },
  props: {
    name: {
      type: String,
      required: true
    },
    percent: {
      type: String,
      required: true
    }
  },
  computed: {
    chartData() {
      return {
        labels: [this.name],
        datasets: [
          {
            label: 'Votes (%)',
            data: [parseFloat(this.percent)],
            backgroundColor: 'lightblue',
          },
        ],
      }
    },
    chartOptions() {
      return {
        responsive: true,
        scales: {
          y: {
            min: 0,
            max: 100,
            ticks: {
              stepSize: 10,
            },
          },
        },
      }
    }
  }
}
</script>
