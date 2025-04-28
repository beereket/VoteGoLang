<template>
  <div class="petitions-page">
    <h1>📋 Active Petitions</h1>
    <div class="petition-list">
      <div v-for="petition in petitions" :key="petition.id" class="petition-card">
        <h2>{{ petition.title }}</h2>
        <p>{{ petition.description.slice(0, 100) }}...</p>
        <router-link :to="`/petitions/${petition.id}`">
          <button>View Details</button>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import api from "@/services/api";

export default {
  name: "PetitionList",
  data() {
    return {
      petitions: [],
    };
  },
  async mounted() {
    const res = await api.get('/petitions');
    this.petitions = res.data;
  }
};
</script>

<style scoped>
.petitions-page {
  padding: 30px;
}
.petition-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 40px;
}
.petition-card {
  background: white;
  border-radius: 10px;
  padding: 20px;
  width: 300px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}
button {
  background-color: #3498db;
  color: white;
  border: none;
  padding: 10px 15px;
  border-radius: 5px;
  cursor: pointer;
}
button:hover {
  background-color: #2980b9;
}
</style>
