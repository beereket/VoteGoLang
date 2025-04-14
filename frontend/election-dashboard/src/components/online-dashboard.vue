<script setup>

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
        <div class="card">
          <p class="card-arrow">Teldeniya</p>
          <p class="card-position">Mahanuwara</p>

          <span class="winning">
            <p>SLPP</p>
            <p>72.83%</p>
          </span>

          <p class="card-name">Anura Dissaskad</p>
        </div>

        <div class="card">
          <p class="card-arrow">Teldeniya</p>
          <p class="card-position">Mahanuwara</p>

          <span class="winning">
            <p>SLPP</p>
            <p>72.83%</p>
          </span>

          <p class="card-name">Anura Dissaskad</p>
        </div>

        <div class="card">
          <p class="card-arrow">Teldeniya</p>
          <p class="card-position">Mahanuwara</p>

          <span class="winning">
            <p>SLPP</p>
            <p>72.83%</p>
          </span>

          <p class="card-name">Anura Dissaskad</p>
        </div>
        <div class="card">
          <p class="card-arrow">Teldeniya</p>
          <p class="card-position">Mahanuwara</p>

          <span class="winning">
            <p>SLPP</p>
            <p>72.83%</p>
          </span>

          <p class="card-name">Anura Dissaskad</p>
        </div>
        <div class="card">
          <p class="card-arrow">Teldeniya</p>
          <p class="card-position">Mahanuwara</p>

          <span class="winning">
            <p>SLPP</p>
            <p>72.83%</p>
          </span>

          <p class="card-name">Anura Dissaskad</p>
        </div>
        <div class="card">
          <p class="card-arrow">Teldeniya</p>
          <p class="card-position">Mahanuwara</p>

          <span class="winning">
            <p>SLPP</p>
            <p>72.83%</p>
          </span>

          <p class="card-name">Anura Dissaskad</p>
        </div>

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

</style>