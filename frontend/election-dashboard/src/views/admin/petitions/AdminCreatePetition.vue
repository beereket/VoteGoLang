<template>
  <div class="create-petition-page">
    <h1>Create New Petition 📝</h1>

    <form @submit.prevent="submitPetition" class="form-container">
      <div class="form-group">
        <label>Title:</label>
        <input v-model="form.title" type="text" required placeholder="Enter petition title" />
      </div>

      <div class="form-group">
        <label>Description:</label>
        <textarea v-model="form.description" required placeholder="Enter petition description"></textarea>
      </div>

      <div class="form-group">
        <label>Photo URL:</label>
        <input v-model="form.photo" type="text" placeholder="Optional: Upload or paste photo link" />
      </div>

      <div class="form-group">
        <label>Deadline:</label>
        <input v-model="form.deadline" type="date" />
      </div>

      <button type="submit" :disabled="loading">
        {{ loading ? 'Creating...' : 'Create Petition' }}
      </button>
    </form>

    <div v-if="message" class="message">{{ message }}</div>
  </div>
</template>

<script>
import api from '@/services/api'; // your Axios wrapper

export default {
  name: "AdminCreatePetition",
  data() {
    return {
      form: {
        title: '',
        description: '',
        photo: '',
        deadline: ''
      },
      loading: false,
      message: ''
    };
  },
  methods: {
    async submitPetition() {
      this.loading = true;
      try {
        const payload = {
          title: this.form.title,
          description: this.form.description,
          photo: this.form.photo,
          deadline: this.form.deadline ? new Date(this.form.deadline).toISOString() : null
        };
        await api.post('/admin/petitions/create', payload);
        this.message = "✅ Petition created successfully!";
        this.form = { title: '', description: '', photo: '', deadline: '' };
      } catch (err) {
        console.error(err);
        this.message = "❌ Failed to create petition.";
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>

<style scoped>
.create-petition-page {
  max-width: 600px;
  margin: 40px auto;
  padding: 20px;
  background: white;
  border-radius: 10px;
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.1);
}

h1 {
  text-align: center;
  margin-bottom: 20px;
  color: #333;
}

.form-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
}

input, textarea {
  padding: 10px;
  border-radius: 8px;
  border: 1px solid #ccc;
  outline: none;
}

textarea {
  min-height: 120px;
}

button {
  padding: 12px;
  border: none;
  border-radius: 8px;
  background: #3498db;
  color: white;
  cursor: pointer;
  font-weight: bold;
  transition: background 0.3s;
}

button:hover {
  background: #2980b9;
}

.message {
  margin-top: 20px;
  text-align: center;
  font-weight: bold;
}
</style>
