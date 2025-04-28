<template>
  <div class="auth-container">
    <div class="auth-box">
      <h2>Login 🔑</h2>
      <form @submit.prevent="login">
        <input v-model="username" type="text" placeholder="Username" required />
        <input v-model="password" type="password" placeholder="Password" required />
        <button type="submit" class="primary-btn">Login</button>
        <p class="link-text">
          Don't have an account?
          <router-link to="/register">Register</router-link>
        </p>
      </form>
    </div>
  </div>
</template>

<script>
import api from '@/services/api';
import { EventBus } from '@/services/eventBus';

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
        const res = await api.post('/users/login', {
          username: this.username,
          password: this.password
        });

        localStorage.setItem('token', res.data.token);
        localStorage.setItem('role', res.data.role);
        localStorage.setItem('user_id', res.data.user_id);

        EventBus.$emit('logged-in');
        if (this.$route.path !== '/') {
          this.$router.replace('/');
        }
      } catch (err) {
        alert('❌ Login failed. Check your credentials!');
      }
    }
  }
};
</script>

<style scoped>
.auth-container {
  margin: 0;

  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #74ebd5, #acb6e5);
}

.auth-box {
  background: white;
  padding: 40px;
  border-radius: 16px;
  width: 100%;
  max-width: 400px;
  box-shadow: 0 10px 25px rgba(0,0,0,0.1);
  text-align: center;
}

.auth-box h2 {
  margin-bottom: 20px;
  color: #333;
}

form {
  display: flex;
  flex-direction: column;
}

input {
  margin-bottom: 15px;
  padding: 12px;
  border: 1px solid #ccc;
  border-radius: 8px;
  font-size: 16px;
}

.primary-btn {
  padding: 12px;
  background-color: #4a90e2;
  color: white;
  border: none;
  border-radius: 8px;
  font-weight: bold;
  cursor: pointer;
}

.primary-btn:hover {
  background-color: #3b7dd8;
}

.link-text {
  margin-top: 15px;
  font-size: 14px;
}
.link-text a {
  color: #4a90e2;
  font-weight: bold;
  text-decoration: none;
}
</style>
