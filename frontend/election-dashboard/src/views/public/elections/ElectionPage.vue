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
      <div v-for="candidate in candidates" :key="candidate.id" class="result-card">
        <p><strong>{{ candidate.name }}</strong> ({{ candidate.party }}) — {{ candidate.votes }} votes</p>
      </div>
    </div>
  </div>
</template>

<script>
import api from '@/services/api';

export default {
  name: "ElectionDetailPage",
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
    }
  },
  methods: {
    async fetchElection() {
      const id = this.$route.params.id;
      const res = await api.get(`/elections/${id}`);
      this.election = res.data.election;
      this.candidates = res.data.candidates;
      this.voted = res.data.alreadyVoted;
    },
    async vote(candidateId) {
      this.loading = true;
      try {
        await api.post('/votes/cast', {
          candidate_id: candidateId,
          candidate_type: this.election.name.includes('Presidential') ? 'president' : 'deputy'
        });
        this.voted = true;
      } catch (err) {
        alert('⚠️ ' + (err.response?.data || 'Failed to vote'));
      }
      this.loading = false;
    }
  },
  mounted() {
    this.fetchElection();
  }
}
</script>

<style scoped>
.candidates {
  display: flex;
  flex-wrap: wrap;
}
.election-detail-page {
  padding: 30px;
  text-align: center;
}
.candidate-card, .result-card {
  background: #f9f9f9;
  margin: 10px auto;
  padding: 20px;
  width: 400px;
  border-radius: 12px;
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
</style>
