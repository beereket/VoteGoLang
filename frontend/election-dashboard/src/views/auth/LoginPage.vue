<template>
  <div class="login-container">
    <div class="login-box">
      <h2 class="login-title">Admin Login</h2>
      <form @submit.prevent="login">
        <div class="form-group">
          <label>Username</label>
          <input
              v-model="username"
              type="text"
              required
              placeholder="Enter your username"
          />
        </div>
        <div class="form-group">
          <label>Password</label>
          <input
              v-model="password"
              type="password"
              required
              placeholder="Enter your password"
          />
        </div>
        <button type="submit" class="login-button">Login</button>
      </form>
    </div>
  </div>
</template>

<script>
import api from '@/services/api';
import { EventBus } from '@/services/eventBus.js';

export default {
  name: "LoginPage",
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

        // Emit login event
        EventBus.$emit('logged-in');

        this.$router.replace('/admin/dashboard');
      } catch (error) {
        alert('❌ Login failed! Check your credentials.');
      }
    }
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #f3f4f6;
}

.login-box {
  background-color: #ffffff;
  padding: 40px 30px;
  border-radius: 10px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  width: 100%;
  max-width: 400px;
}

.login-title {
  font-size: 26px;
  font-weight: bold;
  margin-bottom: 25px;
  text-align: center;
  color: #333333;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  font-weight: 600;
  margin-bottom: 8px;
  color: #555555;
}

.form-group input {
  width: 90%;
  padding: 10px 12px;
  font-size: 16px;
  border: 1px solid #cccccc;
  border-radius: 6px;
  outline: none;
  transition: border-color 0.3s;
}

.form-group input:focus {
  border-color: #4a90e2;
}

.login-button {
  width: 100%;
  padding: 12px;
  background-color: #4a90e2;
  border: none;
  border-radius: 6px;
  color: white;
  font-weight: 600;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.login-button:hover {
  background-color: #3b7dd8;
}
</style>
