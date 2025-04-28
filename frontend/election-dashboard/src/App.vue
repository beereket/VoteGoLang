<template>
  <div id="app" :class="{ 'with-header': showHeader }">

    <AdminHeader v-if="isAdminPage" />
    <UserHeader v-else-if="showHeader" />

    <router-view/>
  </div>
</template>

<script>

import AdminHeader from "@/components/headers/AdminHeader.vue";
import UserHeader from "@/components/headers/UserHeader.vue";

export default {
  name: "App",
  components: {
    AdminHeader,
    UserHeader
  },
  computed: {
    showHeader() {
      const noHeaderRoutes = ['/login', '/register'];
      return !noHeaderRoutes.includes(this.$route.path);
    },
    isAdminPage() {
      return this.$route.path.startsWith('/admin');
    }
  }
}
</script>

<style>
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  background: #f3f4f6;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

router-view {
  flex: 1;
}
.with-header {
  margin-top: 70px;
}
</style>
