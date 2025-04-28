<template>
  <nav :class="['header', { 'shrink': isShrunk }]">
    <div class="logo">VoteGolang 🗳️</div>
    <ul class="nav-links">
      <li><router-link to="/">Live Results</router-link></li>
      <li><router-link to="/elections">Elections</router-link></li>
      <li><router-link to="/petitions">Petitions</router-link></li>
      <li v-if="isAdmin"><router-link to="/admin/dashboard">Admin Dashboard</router-link></li>

      <li v-if="isAuthenticated">
        <button @click="logout">Logout</button>
      </li>
      <li v-else>
        <router-link to="/login">Login</router-link>
      </li>
    </ul>
  </nav>
</template>

<script>
import { EventBus } from '@/services/eventBus.js';

export default {
  name: 'UserHeader',
  data() {
    return {
      isShrunk: false,
      isAuthenticated: !!localStorage.getItem('token'),
    };
  },
  computed: {
    isAdmin() {
      return localStorage.getItem('role') === 'admin';
    }
  },
  methods: {
    logout() {
      localStorage.removeItem('token');
      localStorage.removeItem('role');
      localStorage.removeItem('user_id');
      this.isAuthenticated = false;
      this.$router.replace('/login');
    },
    handleScroll() {
      this.isShrunk = window.scrollY > 50;
    }
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll);

    EventBus.$on('logged-in', () => {
      this.isAuthenticated = true;
    });
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll);
    EventBus.$off('logged-in');
  }
};
</script>

<style scoped>
.header {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 70px;
  background: #2ecc71;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 30px;
  font-family: 'Poppins', sans-serif;
  z-index: 1000;
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
  transition: all 0.3s ease;
}

.header.shrink {
  height: 50px;
  padding: 0 20px;
  background: #27ae60;
}

.logo {
  font-size: 22px;
  color: white;
  font-weight: bold;
  transition: font-size 0.3s ease;
}

.header.shrink .logo {
  font-size: 18px;
}

.nav-links {
  list-style: none;
  display: flex;
  gap: 20px;
  margin: 0;
  padding: 40px;
}

.nav-links a, .nav-links button {
  text-decoration: none;
  color: white;
  font-weight: 500;
  background: none;
  border: none;
  cursor: pointer;
  font-size: 16px;
  transition: 0.3s;
}

.nav-links a:hover, .nav-links button:hover {
  text-decoration: underline;
}
</style>
