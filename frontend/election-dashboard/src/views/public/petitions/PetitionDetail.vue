<template>
  <div class="petition-detail">
    <h1>{{ petition.title }}</h1>
    <p>{{ petition.description }}</p>

    <div class="vote-buttons">
      <button @click="vote('favor')" :disabled="hasVoted">Vote In Favor 👍</button>
      <button @click="vote('against')" :disabled="hasVoted">Vote Against 👎</button>
    </div>

    <div v-if="showChart" class="chart-container">
      <PartyVotesPieChart :data="chartData" />
    </div>

<!--    <div v-if="hasVoted" class="voted-message">-->
<!--      ✅ Thanks for voting!-->
<!--    </div>-->
  </div>
</template>

<script>
import api from "@/services/api";
import PartyVotesPieChart from "@/components/charts/PartyVotesPieChart.vue";
import Swal from "sweetalert2";


export default {
  name: "PetitionDetail",
  components: { PartyVotesPieChart },
  data() {
    return {
      petition: {},
      hasVoted: false,
    };
  },
  computed: {
    chartData() {
      return {
        labels: ["In Favor", "Against"],
        datasets: [
          {
            label: 'Votes',
            data: [this.petition.vote_in_favor, this.petition.vote_against],
            backgroundColor: ['#36A2EB', '#FF6384'],
            hoverOffset: 10,
          }
        ]
      };
    },
    showChart() {
      return this.petition.vote_in_favor !== undefined;
    }
  },
  methods: {
    async fetchPetition() {
      try {
        const res = await api.get(`/petitions/${this.$route.params.id}`);
        this.petition = res.data;
      } catch (err) {
        console.error("❌ Failed to fetch petition", err);
      }
    },
    async vote(type) {
      try {
        await api.post('/petitions/vote', {
          petition_id: this.petition.id,
          vote_type: type,
        });

        this.hasVoted = true;
        await this.fetchPetition();

        Swal.fire({
          icon: 'success',
          title: 'Thank you!',
          text: '✅ Your vote has been recorded!',
          timer: 2000,
          showConfirmButton: false,
        });
      } catch (err) {
        console.error(err);
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: err.response?.data || '❌ Something went wrong!',
        });
      }
    }

  },
  mounted() {
    this.fetchPetition();
  }
};
</script>


<style scoped>
.petition-detail {
  padding: 40px;
}
.vote-buttons {
  margin-top: 20px;
  display: flex;
  gap: 20px;
  justify-content: center;
}
button {
  padding: 10px 20px;
  border-radius: 8px;
  border: none;
  cursor: pointer;
  font-weight: bold;
}
button:first-child {
  background-color: #2ecc71;
  color: white;
}
button:last-child {
  background-color: #e74c3c;
  color: white;
}
button:disabled {
  background-color: grey;
  cursor: not-allowed;
}
.chart-container {
  max-width: 400px;
  padding-top: 100px;
  margin: auto;
}
.voted-message {
  margin-top: 20px;
  font-size: 18px;
  color: green;
}
</style>
