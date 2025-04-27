<script setup>
import { ref } from 'vue'
import BarChart from "@/components/BarChart.vue"
import router from "@/router";

const selectedCard = ref(null)

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('role')
  router.replace('/login')
}

const cardData = ref([
  // your candidate data...
])

const selectCard = (id) => {
  selectedCard.value = id
}

const isSelected = (id) => selectedCard.value === id
</script>


<template>
  <div class="container">
    <header class="header">
      <img src="../assets/logoWh.svg" alt="logo" width="130" class="logo">
      <p>Home</p>

      <router-link to="users">USERS</router-link>
      <router-link to="candidates">Candidates</router-link>

      <button @click="logout" class="bg-red-500 hover:bg-red-600 py-2 px-4 rounded-lg">
        Logout
      </button>

    </header>

    <div class="dashboard">
      <h3 class="card-title">Latest Released Results</h3>

      <div class="card-container">
        <div
            v-for="card in cardData"
            :key="card.id"
            class="card"
            :class="{ selected: isSelected(card.id) }"
            @click="selectCard(card.id)"
        >
          <!--          <img :src="card.avatar" alt="avatar" style="width: 50px; margin: 15px; border-radius: 50%" />-->
          <p class="card-arrow">{{ card.region }}</p>
          <p class="card-position">{{ card.position }}</p>

          <span class="winning">
            <p>{{ card.party }}</p>
            <p>{{ card.percent }}</p>
          </span>

          <p class="card-name">{{ card.name }}</p>
        </div>
      </div>
    </div>

    <div class="result-scores">
      <h3 class="result-title">Vote Results - Cumulative</h3>

      <div v-if="selectedCard" class="chart-container">
        <BarChart
            :name="cardData.find(c => c.id === selectedCard)?.name"
            :percent="cardData.find(c => c.id === selectedCard)?.percent"
        />
      </div>

    </div>
  </div>

</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap');

.container {
  margin: 0;
  padding: 130px;
}

.header {
  position: absolute;
  top: 10px;

  display: flex;
  align-items: center;
}

.header p {
  margin-left: 400px;

  font-family: 'Poppins', sans-serif;
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 45px;

  color: #000000;


}

.logo {
  color: black;
  margin-left: 10px;
  //filter: invert(76%) sepia(21%) saturate(5911%) hue-rotate(350deg) brightness(94%) contrast(97%);
  filter: invert(35%) sepia(0%) saturate(0%) hue-rotate(78deg) brightness(94%) contrast(95%);
}

.dashboard {
  margin-top: -100px;
}

.card-title {
  text-align: left;
  font-size: 24px;
}

.card-container {
  display: flex;
  flex-wrap: wrap;

  gap: 30px;
}

.card {
  width: 334px;
  height: 188px;
  box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;

  background: #e8e8e8;
  border-radius: 19px;

  text-align: left;

  transition: 0.7s;
}

.card:hover {
  transition: 0.5s;
  transform: scale(102.5%);
  border: 2px solid gray;
}


.card-arrow {
  font-family: 'Poppins', sans-serif;
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  line-height: 30px;

  margin-left: 15px;
  color: #000000;
}

.card-position {

  font-family: 'Poppins', sans-serif;
  font-style: normal;
  font-weight: 300;
  font-size: 18px;
  line-height: 30px;

  margin-left: 15px;
  margin-top: -25px;
}

.winning {
  margin: 15px;

  font-weight: 600;

  display: flex;
  justify-content: space-between;
}

.card-name {
  margin-top: -20px;
  margin-left: 15px;
}

.card.selected {
  border: 2px solid #800020;
  transition: 0.4s in ease-in-out;
}

.result-container {

}
.result-title {
  text-align: left;
  font-size: 24px;

  margin-top: 50px;
}




</style>