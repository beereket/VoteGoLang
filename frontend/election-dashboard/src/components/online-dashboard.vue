<script setup>
import { ref } from 'vue'

// Только один выбранный ID
const selectedCard = ref(null)

const cardData = ref([
  {
    id: 1,
    name: 'John Doe',
    region: 'Teldeniya',
    position: 'Mahanuwara',
    party: 'SLPP',
    percent: '72.83%',
    type: 'President',
    avatar: 'https://ui-avatars.com/api/?name=John+Doe'
  },
  {
    id: 2,
    name: 'Jane Smith',
    region: 'Kandy',
    position: 'Central Province',
    party: 'UNP',
    percent: '58.42%',
    type: 'Deputy',
    avatar: 'https://ui-avatars.com/api/?name=Jane+Smith'
  },
  {
    id: 3,
    name: 'Ali Khan',
    region: 'Galle',
    position: 'Southern Province',
    party: 'JVP',
    percent: '66.12%',
    type: 'Session Deputy',
    avatar: 'https://ui-avatars.com/api/?name=Ali+Khan'
  },
  {
    id: 4,
    name: 'Maria Ivanova',
    region: 'Colombo',
    position: 'Western Province',
    party: 'SLFP',
    percent: '81.22%',
    type: 'Petition',
    avatar: 'https://ui-avatars.com/api/?name=Maria+Ivanova'
  }
])

const selectCard = (id) => {
  selectedCard.value = id
}

const isSelected = (id) => selectedCard.value === id

</script>

<!--President, Deputat, Session Deputat, Petition-->

<template>
  <div class="container">
    <header class="header">
      <img src="../assets/logoWh.svg" alt="logo" width="130" class="logo">
      <p>Home</p>
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

      <div>

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
  position: fixed;
  top: 30px;

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

  filter: invert(100%) sepia(0%) saturate(7500%) hue-rotate(13deg) brightness(101%) contrast(106%);
}

.dashboard {

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

  background: #D9D9D9;
  border-radius: 19px;

  text-align: left;
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


</style>