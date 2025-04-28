<template>
  <div class="election-list-page">
    <h1 class="title">🗳️ Elections</h1>

    <div v-if="elections.length" class="election-cards">
      <div v-for="election in elections" :key="election.id" class="election-card">
        <h2>{{ election.name }}</h2>
        <p>{{ election.description }}</p>
        <p><strong>Date:</strong> {{ formatDate(election.election_date) }}</p>
        <p><strong>Status:</strong>
          <span :class="{'status-ongoing': isOngoing(election), 'status-ended': !isOngoing(election)}">
            {{ isOngoing(election) ? '🟢 Ongoing' : '🔴 Ended' }}
          </span>
        </p>
        <router-link :to="'/elections/' + election.id">
          <button class="enter-button">View Details</button>
        </router-link>
      </div>
    </div>

    <div v-else class="no-elections">
      <p>No elections available.</p>
    </div>
  </div>
</template>

<script>
import api from '@/services/api';

export default {
  name: "ElectionListPage",
  data() {
    return {
      elections: []
    };
  },
  methods: {
    async fetchElections() {
      const res = await api.get('/elections');
      this.elections = res.data;
    },
    formatDate(date) {
      return new Date(date).toLocaleDateString();
    },
    isOngoing(election) {
      return new Date(election.election_date) >= new Date();
    }
  },
  mounted() {
    this.fetchElections();
  }
}
</script>

<style scoped>
.election-list-page {
  padding: 40px;
  text-align: center;
}
.title {
  font-size: 36px;
  margin-bottom: 30px;
}
.election-cards {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: center;
}
.election-card {
  width: 300px;
  background: #fff;
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  transition: 0.3s;
}
.election-card:hover {
  transform: translateY(-5px);
}
.enter-button {
  margin-top: 10px;
  background: #3498db;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
}
.status-ongoing {
  color: green;
}
.status-ended {
  color: red;
}
</style>
