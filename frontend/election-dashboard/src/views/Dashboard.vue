<template>
  <div class="dashboard">
    <header>
      <h1>Admin Dashboard</h1>
      <p>Welcome back, Admin! 💪</p>
    </header>

    <div class="summary-cards">
      <div v-for="card in summaryCards" :key="card.title" class="card">
        <div class="card-value">{{ card.value }}</div>
        <div class="card-title">{{ card.title }}</div>
      </div>
    </div>

    <!-- Charts -->
    <div class="charts">
      <div class="chart-container" v-if="votesPerDayData">
        <h2>Votes Per Day 📈</h2>
        <BarChart :chart-data="votesPerDayData" />
      </div>

      <div class="chart-container" v-if="usersPerWeekData">
        <h2>User Registrations Per Week 📅</h2>
        <BarChart :chart-data="usersPerWeekData" />
      </div>
    </div>

    <!-- Analytics -->
    <div class="charts">
      <div class="chart-container" v-if="topCandidatesData">
        <h2>Top Candidates 🏆</h2>
        <PartyVotesPieChart :data="topCandidatesData" />
      </div>

      <div class="chart-container" v-if="partyAnalyticsData">
        <h2>Party Analytics 🌟</h2>
        <BarChart :chart-data="partyAnalyticsData" />
      </div>
    </div>
  </div>
</template>

<script>
import api from '@/services/api';
import BarChart from "@/components/BarChart.vue";
import PartyVotesPieChart from "@/components/PartyVotesPieChart.vue";

export default {
  name: 'DashboardPage',
  components: {BarChart, PartyVotesPieChart},
  data() {
    return {
      summaryCards: [],
      votesPerDayData: null,
      usersPerWeekData: null,
      topCandidatesData: null,
      partyAnalyticsData: null,
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
          {title: 'Top Candidates', value: dashboardRes.data.top_candidates.length},
        ];

        const votesRes = await api.get('/admin/dashboard/votes-per-day');
        this.votesPerDayData = {
          labels: votesRes.data.map(d => d.date.split('T')[0]),
          datasets: [{
            label: 'Votes',
            data: votesRes.data.map(d => d.count),
            backgroundColor: 'skyblue',
          }]
        };

        const usersRes = await api.get('/admin/dashboard/users-per-week');
        this.usersPerWeekData = {
          labels: usersRes.data.map(d => d.year_week),
          datasets: [{
            label: 'New Users',
            data: usersRes.data.map(d => d.count),
            backgroundColor: 'lightgreen',
          }]
        };

        const topCandidatesRes = await api.get('/admin/analytics/top-candidates');
        this.topCandidatesData = {
          labels: topCandidatesRes.data.map(d => d.candidate_name),
          datasets: [{
            label: 'Votes',
            data: topCandidatesRes.data.map(d => d.total_votes),
            backgroundColor: ['#3490dc', '#9561e2', '#f66d9b', '#38c172', '#ffed4a', '#e3342f', '#6cb2eb'],
          }]
        };

        const partyAnalyticsRes = await api.get('/admin/analytics/party');
        this.partyAnalyticsData = {
          labels: partyAnalyticsRes.data.map(d => d.party_name),
          datasets: [{
            label: 'Total Votes',
            data: partyAnalyticsRes.data.map(d => d.total_votes),
            backgroundColor: 'coral',
          }]
        };

      } catch (err) {
        console.error('❌ Failed to load dashboard data', err);
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
  background: #f0f2f5;
  min-height: 100vh;
}

header {
  margin-bottom: 30px;
}

h1 {
  font-size: 36px;
  color: #333;
}

p {
  color: #777;
}

.summary-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 20px;
  margin-bottom: 40px;
}

.card {
  background: #fff;
  border-radius: 15px;
  padding: 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: 0.3s;
}

.card:hover {
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
  transform: translateY(-3px);
}

.card-value {
  font-size: 28px;
  font-weight: bold;
  margin-bottom: 8px;
  color: #222;
}

.card-title {
  color: #777;
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
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

h2 {
  font-size: 20px;
  margin-bottom: 15px;
  color: #555;
}
</style>
