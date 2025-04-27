<script>
import { Pie } from 'vue-chartjs'
import {
  Chart as ChartJS,
  ArcElement,
  Tooltip,
  Legend
} from 'chart.js'

ChartJS.register(ArcElement, Tooltip, Legend)

export default {
  name: "PartyVotesPieChart",
  components: { Pie },
  props: {
    data: {
      type: Object,
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
        },
      };
    }
  }
}
</script>

<template>
  <Pie :data="chartData" :options="chartOptions" />
</template>
