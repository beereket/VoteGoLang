<template>
  <div class="election-detail-page">
    <h1>{{ election.name }}</h1>
    <p>{{ election.description }}</p>

    <div v-if="isOngoing">
      <h2>Candidates 🧑‍💼</h2>
      <div class="candidates">
        <div v-for="candidate in candidates" :key="candidate.id" class="candidate-card">
          <p><strong>{{ candidate.name }}</strong> ({{ candidate.party }})</p>
          <button
              :disabled="voted || loading"
              @click="vote(candidate.id)"
              class="vote-button">
            {{ voted ? '✅ Voted' : 'Vote' }}
          </button>
        </div>
      </div>
    </div>

    <div v-else>
      <h2>Results 🏆</h2>

      <div class="chart-container" v-if="chartData.datasets[0].data.length">
        <PartyVotesPieChart :data="chartData" />
      </div>

      <div
          v-for="candidate in candidates"
          :key="candidate.id"
          :class="['result-card', { 'winner': candidate.id === winnerId }]">
        <p>
          <strong>{{ candidate.name }}</strong> ({{ candidate.party }}) —
          {{ candidate.votes }} votes ({{ getPercentage(candidate.votes) }}%)
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import api from '@/services/api';
import Swal from 'sweetalert2';
import PartyVotesPieChart from '@/components/charts/PartyVotesPieChart.vue';

export default {
  name: "ElectionDetail",
  components: { PartyVotesPieChart },
  data() {
    return {
      election: {},
      candidates: [],
      voted: false,
      loading: false
    };
  },
  computed: {
    isOngoing() {
      return new Date(this.election.election_date) >= new Date();
    },
    chartData() {
      return {
        labels: this.candidates.map(c => c.name),
        datasets: [{
          label: 'Votes',
          data: this.candidates.map(c => c.votes),
          backgroundColor: [
            '#FF6384', '#36A2EB', '#FFCE56', '#4BC0C0',
            '#9966FF', '#FF9F40', '#FFCD56', '#4DD0E1'
          ],
          hoverOffset: 10,
        }]
      };
    },
    winnerId() {
      if (!this.candidates.length) return null;
      return this.candidates.reduce((max, c) => c.votes > max.votes ? c : max, this.candidates[0]).id;
    },
    totalVotes() {
      return this.candidates.reduce((sum, c) => sum + c.votes, 0);
    }
  },
  methods: {
    async fetchElection() {
      try {
        const id = this.$route.params.id;
        const res = await api.get(`/elections/${id}`);
        this.election = res.data.election;
        this.candidates = res.data.candidates;
        this.voted = res.data.alreadyVoted;
      } catch (err) {
        console.error("❌ Failed to fetch election:", err);
        Swal.fire('Error', 'Failed to load election details', 'error');
      }
    },
    async vote(candidateId) {
      this.loading = true;
      try {
        await api.post('/votes/cast', {
          candidate_id: candidateId,
          candidate_type: this.election.name.includes('Presidential') ? 'president' : 'deputy'
        });
        this.voted = true;
        await this.fetchElection();
        Swal.fire({
          icon: 'success',
          title: 'Vote Cast!',
          text: '✅ Your vote has been recorded!',
          timer: 2000,
          showConfirmButton: false,
        });
      } catch (err) {
        console.error("❌ Failed to vote:", err);
        Swal.fire('Error', err.response?.data || 'Something went wrong', 'error');
      }
      this.loading = false;
    },
    getPercentage(votes) {
      if (this.totalVotes === 0) return 0;
      return ((votes / this.totalVotes) * 100).toFixed(1);
    }
  },
  mounted() {
    this.fetchElection();
  }
}
</script>

<style scoped>
.election-detail-page {
  padding: 30px;
  text-align: center;
}

.candidates {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 20px;
  margin-top: 20px;
}

.candidate-card, .result-card {
  background: #f9f9f9;
  margin: 10px auto;
  padding: 20px;
  width: 400px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.winner {
  border: 2px solid gold;
  background: #fff8dc;
}

.vote-button {
  margin-top: 10px;
  padding: 10px 20px;
  background: #2ecc71;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
}

.vote-button:disabled {
  background: #bbb;
}

.chart-container {
  margin: 40px auto;
  max-width: 500px;
}
</style>
