<template>
  <div class="auth-container">
    <div class="auth-box">
      <h2>Register 📝</h2>
      <form @submit.prevent="register">
        <input v-model="username" type="text" placeholder="Username" required />
        <input v-model="fullName" type="text" placeholder="Full Name" required />
        <input v-model="birthDate" type="date" placeholder="Birth Date" required />
        <input v-model="address" type="text" placeholder="Address" required />
        <input v-model="password" type="password" placeholder="Password" required />
        <button type="submit" class="primary-btn">Register</button>
        <p class="link-text">
          Already have an account?
          <router-link to="/login">Login</router-link>
        </p>
      </form>
    </div>
  </div>
</template>

<script>
import api from '@/services/api';
import Swal from 'sweetalert2';

export default {
  name: "RegisterPage",
  data() {
    return {
      username: '',
      fullName: '',
      birthDate: '',
      address: '',
      password: ''
    };
  },
  methods: {
    async register() {
      try {
        await api.post('/users/create', {
          username: this.username,
          user_full_name: this.fullName,
          birth_date: this.birthDate,
          address: this.address,
          password: this.password,
        });
        Swal.fire('✅ Success', 'Account created! Now login.', 'success');
        this.$router.push('/login');
      } catch (err) {
        Swal.fire('❌ Error', err.response?.data || 'Registration failed', 'error');
      }
    }
  }
};
</script>

<style scoped>
.auth-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #ffecd2, #fcb69f);
}

.auth-box {
  background: white;
  padding: 40px;
  border-radius: 16px;
  width: 100%;
  max-width: 450px;
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
  background-color: #e67e22;
  color: white;
  border: none;
  border-radius: 8px;
  font-weight: bold;
  cursor: pointer;
}

.primary-btn:hover {
  background-color: #d35400;
}

.link-text {
  margin-top: 15px;
  font-size: 14px;
}
.link-text a {
  color: #e67e22;
  font-weight: bold;
  text-decoration: none;
}
</style>
