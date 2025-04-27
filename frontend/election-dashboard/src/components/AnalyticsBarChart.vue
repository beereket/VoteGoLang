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
  name: 'AnalyticsBarChart',
  components: { Bar },
  props: {
    data: {
      type: Array,
      required: true
    }
  },
  computed: {
    chartData() {
      return {
        labels: this.data.map(item => item.party),
        datasets: [
          {
            label: 'Votes',
            data: this.data.map(item => item.votes),
            backgroundColor: 'rgba(54, 162, 235, 0.6)',
          },
        ],
      }
    },
    chartOptions() {
      return {
        responsive: true,
        scales: {
          y: {
            beginAtZero: true,
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
