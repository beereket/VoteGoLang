<template>
  <div class="dashboard">
    <header>
      <h1>Admin Dashboard</h1>
      <p>Welcome back, Admin! 💪</p>
    </header>

    <!-- Summary Cards -->
    <div class="summary-cards">
      <div v-for="card in summaryCards" :key="card.title" class="card">
        <div class="card-value">{{ card.value }}</div>
        <div class="card-title">{{ card.title }}</div>
      </div>
    </div>

    <div class="charts">
      <div class="chart-container">
        <h2>Votes Per Day 📈</h2>
        <BarChart :chart-data="votesPerDayData" :chart-options="barChartOptions" />
      </div>
      <div class="chart-container">
        <h2>User Registrations Per Week 📅</h2>
        <BarChart :chart-data="usersPerWeekData" :chart-options="barChartOptions" />
      </div>
    </div>

    <div class="charts">
      <div class="chart-container">
        <h2>Top Candidates 🏆</h2>
        <PieChart :chart-data="topCandidatesData" :chart-options="pieChartOptions" />
      </div>
      <div class="chart-container">
        <h2>Party Analytics 🌟</h2>
        <BarChart :chart-data="partyAnalyticsData" :chart-options="barChartOptions" />
      </div>
    </div>
  </div>
</template>

<script>
import BarChart from '@/components/BarChart.vue';
import PieChart from '@/components/PartyVotesPieChart.vue';
import api from '@/services/api';

export default {
  name: 'DashboardPage',
  components: { BarChart, PieChart },
  data() {
    return {
      summaryCards: [],
      votesPerDayData: null,
      usersPerWeekData: null,
      topCandidatesData: null,
      partyAnalyticsData: null,
    }
  },
  computed: {
    barChartOptions() {
      return {
        responsive: true,
        animation: {
          duration: 1500,
          easing: 'easeInOutCubic'
        },
        hover: {
          mode: 'index',
          intersect: false,
          animationDuration: 400
        },
        plugins: {
          legend: {
            position: 'bottom',
            labels: {
              color: '#333',
              font: { size: 14, weight: 'bold' }
            }
          },
          tooltip: {
            backgroundColor: '#2d3748',
            titleColor: '#fff',
            bodyColor: '#fff',
            padding: 10,
            cornerRadius: 8,
          }
        },
        scales: {
          y: {
            beginAtZero: true
          }
        }
      }
    },
    pieChartOptions() {
      return {
        responsive: true,
        animation: {
          duration: 1500,
          easing: 'easeInOutCubic'
        },
        plugins: {
          legend: {
            position: 'bottom',
            labels: {
              color: '#333',
              font: { size: 14, weight: 'bold' }
            }
          },
          tooltip: {
            backgroundColor: '#2d3748',
            titleColor: '#fff',
            bodyColor: '#fff',
            padding: 10,
            cornerRadius: 8,
          }
        }
      }
    }
  },
  methods: {
    async fetchDashboardData() {
      try {
        const dashboardRes = await api.get('/admin/dashboard');
        this.summaryCards = [
          {title: 'Total Users', value: dashboardRes.data.total_users},
          {title: 'Total Votes', value: dashboardRes.data.total_votes},
          {title: 'Total Candidates', value: dashboardRes.data.total_candidates},
          {title: 'Top Candidate', value: dashboardRes.data.top_candidates[0]?.name || 'N/A'},
        ];

        const votesRes = await api.get('/admin/dashboard/votes-per-day');
        this.votesPerDayData = {
          labels: votesRes.data.map(d => new Date(d.date).toLocaleDateString()),
          datasets: [{
            label: 'Votes',
            data: votesRes.data.map(d => d.count),
            backgroundColor: '#4CAF50',
          }]
        };

        const usersRes = await api.get('/admin/dashboard/users-per-week');
        this.usersPerWeekData = {
          labels: usersRes.data.map(d => d.year_week),
          datasets: [{
            label: 'New Users',
            data: usersRes.data.map(d => d.count),
            backgroundColor: '#36A2EB',
          }]
        };

        const topCandidatesRes = await api.get('/admin/analytics/top-candidates');
        this.topCandidatesData = {
          labels: topCandidatesRes.data.map(d => d.candidate_name),
          datasets: [{
            label: 'Votes',
            data: topCandidatesRes.data.map(d => d.total_votes),
            backgroundColor: [
              '#3490dc', '#9561e2', '#f66d9b', '#38c172', '#ffed4a', '#ff9f40', '#00ACC1'
            ],
            hoverOffset: 12
          }]
        };

        const partyAnalyticsRes = await api.get('/admin/analytics/party');
        this.partyAnalyticsData = {
          labels: partyAnalyticsRes.data.map(d => d.party_name),
          datasets: [{
            label: 'Total Votes',
            data: partyAnalyticsRes.data.map(d => d.total_votes),
            backgroundColor: '#FF6384',
            hoverBackgroundColor: '#FF9F40'
          }]
        };
      } catch (err) {
        console.error('Failed to load dashboard data', err);
      }
    }
  },
  mounted() {
    this.fetchDashboardData();
  }
};
</script>

<style scoped>
.dashboard {
  padding: 50px;
  background: linear-gradient(to right, #e0eafc, #cfdef3);
  min-height: 100vh;
}

header {
  margin-bottom: 30px;
  text-align: center;
}

h1 {
  font-size: 36px;
  font-weight: bold;
  color: #333;
}

p {
  color: #666;
  font-size: 18px;
}

.summary-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 20px;
  margin: 40px 0;
}

.card {
  background: #ffffff;
  padding: 20px;
  border-radius: 15px;
  box-shadow: 0 8px 24px rgba(0,0,0,0.1);
  transition: 0.3s;
  text-align: center;
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 28px rgba(0,0,0,0.15);
}

.card-value {
  font-size: 30px;
  font-weight: bold;
  margin-bottom: 8px;
  color: #333;
}

.card-title {
  color: #888;
  font-size: 16px;
}

.charts {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(380px, 1fr));
  gap: 30px;
  margin-bottom: 50px;
}

.chart-container {
  background: #fff;
  padding: 20px;
  border-radius: 15px;
  box-shadow: 0 8px 24px rgba(0,0,0,0.1);
}
</style>
