<script>
import api from "@/services/api";
import { ref, onMounted } from "vue";
import AnalyticsBarChart from "@/components/AnalyticsBarChart.vue";

export default {
  name: "DashboardAnalytics",
  components: { AnalyticsBarChart },
  setup() {
    const stats = ref({
      totalUsers: 0,
      totalVotes: 0,
      totalPetitions: 0,
      totalCandidates: 0,
    });

    const partyData = ref([]);

    const fetchDashboardData = async () => {
      try {
        const [usersRes, votesRes, petitionsRes, candidatesRes, partyRes] = await Promise.all([
          api.get('/admin/stats/users'),
          api.get('/admin/stats/votes'),
          api.get('/admin/stats/petitions'),
          api.get('/admin/stats/candidates'),
          api.get('/admin/stats/party-votes')
        ]);

        stats.value.totalUsers = usersRes.data.total;
        stats.value.totalVotes = votesRes.data.total;
        stats.value.totalPetitions = petitionsRes.data.total;
        stats.value.totalCandidates = candidatesRes.data.total;
        partyData.value = partyRes.data;
      } catch (error) {
        console.error('❌ Failed to fetch dashboard stats', error);
      }
    };

    onMounted(() => {
      fetchDashboardData();
    });

    return {
      stats,
      partyData
    };
  }
};
</script>

<template>
  <div class="analytics-page">
    <h1>📊 Dashboard Analytics</h1>

    <div class="stats-cards">
      <div class="card">👥 Users<br/>{{ stats.totalUsers }}</div>
      <div class="card">🗳️ Votes<br/>{{ stats.totalVotes }}</div>
      <div class="card">📝 Petitions<br/>{{ stats.totalPetitions }}</div>
      <div class="card">🎯 Candidates<br/>{{ stats.totalCandidates }}</div>
    </div>

    <div class="charts">
      <h2>Votes Per Party</h2>
      <AnalyticsBarChart v-if="partyData.length > 0" :data="partyData"/>
      <p v-else>No party vote data available.</p>
    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600;700&display=swap');

.analytics-page {
  font-family: 'Poppins', sans-serif;
  padding: 40px;
  background: linear-gradient(to right, #ece9e6, #ffffff);
  min-height: 100vh;
}

h1 {
  font-size: 36px;
  margin-bottom: 40px;
  color: #333;
  font-weight: 700;
  text-align: center;
}

.stats-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 25px;
  margin-bottom: 50px;
}

.card {
  background: white;
  padding: 30px;
  border-radius: 16px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.08);
  text-align: center;
  transition: all 0.3s ease;
  font-size: 24px;
  font-weight: 600;
  color: #555;
  cursor: pointer;
}

.card:hover {
  transform: translateY(-8px) scale(1.02);
  background: linear-gradient(to right, #a8edea, #fed6e3);
  color: #222;
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.12);
}

.charts {
  background: white;
  padding: 40px;
  border-radius: 16px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.08);
  margin-top: 30px;
}

h2 {
  font-size: 28px;
  color: #333;
  margin-bottom: 25px;
  font-weight: 700;
  text-align: center;
}

@media (max-width: 768px) {
  .analytics-page {
    padding: 20px;
  }
  .charts {
    padding: 20px;
  }
  .card {
    padding: 20px;
    font-size: 20px;
  }
}
</style>
