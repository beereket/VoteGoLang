<template>
  <Bar :data="data" :options="chartOptions" />
</template>

<script>
import Bar  from 'vue-chartjs';
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale
} from 'chart.js';

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale);

export default {
  name: 'PetitionBarChart',
  components: { Bar },
  props: {
    data: {
      type: Object,
      required: true
    }
  },
  computed: {
    chartOptions() {
      return {
        responsive: true,
        plugins: {
          legend: {
            position: 'bottom',
          },
        },
        scales: {
          y: {
            beginAtZero: true
          }
        }
      }
    }
  }
};
</script>
