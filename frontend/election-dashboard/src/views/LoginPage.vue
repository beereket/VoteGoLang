<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="w-full max-w-md bg-white rounded-lg shadow-md p-8">
      <h2 class="text-2xl font-bold mb-6 text-center text-gray-800">Admin Login</h2>
      <form @submit.prevent="login">
        <div class="mb-4">
          <label class="block text-gray-700 mb-2">Username</label>
          <input
              v-model="username"
              type="text"
              required
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400"
              placeholder="Enter your username"
          />
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 mb-2">Password</label>
          <input
              v-model="password"
              type="password"
              required
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400"
              placeholder="Enter your password"
          />
        </div>
        <button
            type="submit"
            class="w-full bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-lg transition duration-300"
        >
          Login
        </button>
      </form>
    </div>
  </div>
</template>

<script>
import api from '@/services/api'; // Axios service

export default {
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    async login() {
      try {
        const response = await api.post('/users/login', {
          username: this.username,
          password: this.password
        });

        localStorage.setItem('token', response.data.token);
        localStorage.setItem('role', response.data.role);

        setTimeout(() => {
          // 🔥 Only redirect if not already on dashboard
          if (this.$route.path !== '/admin/dashboard') {
            this.$router.replace('/admin/dashboard');
          }
        }, 100); // 100 milliseconds
      } catch (error) {
        alert('❌ Login failed! Check your credentials.');
      }
    }

  }
};
</script>

<style scoped>
</style>
