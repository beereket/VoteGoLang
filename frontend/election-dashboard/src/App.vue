<template>
  <div id="app">
    <online-dashboard/>
  </div>
</template>

<script>

import OnlineDashboard from "@/components/online-dashboard.vue";

export default {
  name: 'App',
  components: {
    OnlineDashboard
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
