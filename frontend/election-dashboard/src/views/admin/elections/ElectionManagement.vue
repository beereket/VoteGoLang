<template>
  <div class="elections-management-page">
    <h1>🗳️ Elections Management</h1>

    <!-- Create New Election -->
    <div class="form-section">
      <h2>Create New Election</h2>
      <form @submit.prevent="createElection">
        <input v-model="newElection.name" type="text" placeholder="Election Name" required />
        <input v-model="newElection.description" type="text" placeholder="Election Description" required />
        <input v-model="newElection.election_date" type="date" required />
        <button type="submit">Create Election</button>
      </form>
    </div>

    <hr />

    <div class="list-section">
      <h2>📋 Existing Elections</h2>

      <div v-if="elections.length === 0" class="empty-state">No elections found.</div>

      <div v-else class="elections-list">
        <div v-for="election in elections" :key="election.id" class="election-item">
          <div>
            <strong>{{ election.name }}</strong> - {{ election.description }} ({{ formatDate(election.election_date) }})
          </div>

          <div class="actions">
            <button @click="startEdit(election)">✏️ Edit</button>
            <button @click="deleteElection(election.id)">🗑️ Delete</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Edit Modal -->
    <div v-if="editingElection" class="modal">
      <div class="modal-content">
        <h2>Edit Election</h2>
        <form @submit.prevent="updateElection">
          <input v-model="editingElection.name" type="text" required />
          <input v-model="editingElection.description" type="text" required />
          <input v-model="editingElection.election_date" type="date" required />
          <div class="modal-actions">
            <button type="submit">Save Changes</button>
            <button type="button" @click="cancelEdit">Cancel</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import api from '@/services/api';

export default {
  name: "ElectionsManagement",
  data() {
    return {
      elections: [],
      newElection: {
        name: '',
        description: '',
        election_date: ''
      },
      editingElection: null,
    }
  },
  methods: {
    async fetchElections() {
      try {
        const res = await api.get('/admin/elections');
        this.elections = res.data;
      } catch (err) {
        console.error('❌ Failed to load elections', err);
      }
    },
    async createElection() {
      try {
        await api.post('/admin/elections/create', this.newElection);
        this.newElection = { name: '', description: '', election_date: '' };
        this.fetchElections();
      } catch (err) {
        console.error('❌ Failed to create election', err);
      }
    },
    async deleteElection(id) {
      if (confirm('Are you sure you want to delete this election?')) {
        try {
          await api.delete(`/admin/elections/delete/${id}`);
          this.fetchElections();
        } catch (err) {
          console.error('❌ Failed to delete election', err);
        }
      }
    },
    startEdit(election) {
      this.editingElection = { ...election };
    },
    cancelEdit() {
      this.editingElection = null;
    },
    async updateElection() {
      try {
        await api.put(`/admin/elections/update/${this.editingElection.id}`, this.editingElection);
        this.editingElection = null;
        this.fetchElections();
      } catch (err) {
        console.error('❌ Failed to update election', err);
      }
    },
    formatDate(dateStr) {
      const date = new Date(dateStr);
      return date.toLocaleDateString();
    }
  },
  mounted() {
    this.fetchElections();
  }
}
</script>

<style scoped>
.elections-management-page {
  padding: 30px;
  font-family: 'Poppins', sans-serif;
  min-height: 100vh;
  background: #f7f9fb;
}
h1, h2 {
  text-align: center;
  color: #333;
}
form {
  display: flex;
  flex-direction: column;
  gap: 10px;
  max-width: 400px;
  margin: 0 auto;
}
input {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 8px;
}
button {
  padding: 10px;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
}
button:hover {
  background-color: #2980b9;
}
.elections-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
  margin-top: 20px;
}
.election-item {
  background: white;
  padding: 15px;
  border-radius: 12px;
  box-shadow: 0 4px 10px rgba(0,0,0,0.1);
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.actions button {
  margin-left: 10px;
}
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  backdrop-filter: blur(3px);
  display: flex;
  align-items: center;
  justify-content: center;
}
.modal-content {
  background: white;
  padding: 30px;
  border-radius: 12px;
  width: 400px;
}
.modal-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}
.empty-state {
  text-align: center;
  color: #888;
}
</style>
