<template>
  <div class="results-with-news">
    <div class="news-section">
      <h2>📰 Latest News</h2>
      <ul>
        <li v-for="newsItem in news" :key="newsItem.id">
          <img
              v-if="newsItem.photo"
              :src="newsItem.photo"
              alt="News Photo"
              class="news-image"
          />
          <div class="news-content">
            <strong>{{ newsItem.title }}</strong><br />
            <small>{{ formatDate(newsItem.created_at) }}</small>
          </div>
        </li>
      </ul>
    </div>

    <div class="results-section">
      <h2>📊 Voting Results</h2>

      <div class="tabs">
        <button @click="activeTab = 'elections'" :class="{ active: activeTab === 'elections' }">
          Elections
        </button>
        <button @click="activeTab = 'petitions'" :class="{ active: activeTab === 'petitions' }">
          Petitions
        </button>
      </div>

      <div v-if="activeTab === 'elections'" class="section">
        <div v-for="election in elections" :key="election.id" class="card" @click="fetchElectionResults(election.id)">
          {{ election.name }}
        </div>

        <div v-if="selectedElectionResults.length" class="election-results">
          <h3>Results for {{ selectedElectionName }}</h3>
          <ul>
            <li v-for="r in selectedElectionResults" :key="r.name">
              {{ r.name }} ({{ r.party }}) — {{ r.votes }} votes
            </li>
          </ul>
        </div>
      </div>

      <div v-else class="section">
        <div v-for="petition in petitions" :key="petition.id" class="card" @click="fetchPetitionResults(petition.id)">
          {{ petition.title }}
        </div>

        <div v-if="petitionResults">
          <h3>Results:</h3>
          <p>✅ In Favor: {{ petitionResults.vote_in_favor }}</p>
          <p>❌ Against: {{ petitionResults.vote_against }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import api from "@/services/api";

export default {
  data() {
    return {
      news: [],
      elections: [],
      petitions: [],
      activeTab: 'elections',
      selectedElectionResults: [],
      selectedElectionName: '',
      petitionResults: null
    };
  },
  async mounted() {
    const newsRes = await api.get('/news');
    this.news = newsRes.data;

    const electionsRes = await api.get('/elections');
    this.elections = electionsRes.data;

    const petitionsRes = await api.get('/petitions');
    this.petitions = petitionsRes.data;
  },
  methods: {
    async fetchElectionResults(electionId) {
      const res = await api.get(`/elections/${electionId}/results`);
      const election = this.elections.find(e => e.id === electionId);
      this.selectedElectionName = election ? election.name : '';
      this.selectedElectionResults = res.data;
      this.petitionResults = null;
    },
    async fetchPetitionResults(petitionId) {
      const res = await api.get(`/petitions/${petitionId}/results`);
      this.petitionResults = res.data;
      this.selectedElectionResults = [];
    },
    formatDate(dateStr) {
      const date = new Date(dateStr);
      return date.toLocaleDateString();
    }
  }
};
</script>

<style scoped>
.results-with-news {
  display: flex;
  flex-wrap: wrap;
  padding: 30px;
  gap: 20px;
}

.news-section, .results-section {
  flex: 1;
  min-width: 400px;
  background: #f7f9fb;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 10px rgba(0,0,0,0.1);
}

.news-section h2, .results-section h2 {
  text-align: center;
  margin-bottom: 20px;
}

.news-section ul {
  list-style: none;
  padding: 0;
}

.news-section li {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  background: #fff;
  padding: 10px;
  border-radius: 8px;
}

.news-image {
  width: 80px;
  height: 80px;
  object-fit: cover;
  border-radius: 8px;
  margin-right: 15px;
}

.news-content {
  flex-grow: 1;
}

.tabs {
  display: flex;
  justify-content: center;
  gap: 10px;
  margin-bottom: 20px;
}

.tabs button {
  padding: 10px 20px;
  border: none;
  background: #ddd;
  cursor: pointer;
  border-radius: 5px;
}

.tabs button.active {
  background: #4a90e2;
  color: white;
}

.section {
  margin-top: 10px;
}

.card {
  background: #ffffff;
  padding: 15px;
  margin: 10px 0;
  border-radius: 8px;
  cursor: pointer;
  transition: 0.2s;
}

.card:hover {
  background: #ececec;
}

.election-results {
  margin-top: 20px;
  background: #fff;
  padding: 20px;
  border-radius: 12px;
}
</style>