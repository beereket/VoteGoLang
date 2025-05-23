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
            <router-link :to="`/admin/elections/${election.id}/candidates`" class="manage-btn">Candidates
            </router-link>
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
import Swal from 'sweetalert2';

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
        const payload = {
          ...this.newElection,
          election_date: new Date(this.newElection.election_date).toISOString()
        };
        await api.post('/admin/elections/create', payload);
        this.newElection = { name: '', description: '', election_date: '' };
        this.fetchElections();

        Swal.fire({
          icon: 'success',
          title: 'Election Created!',
          text: '✅ Your election has been successfully created.',
          timer: 2000,
          showConfirmButton: false,
        });

      } catch (err) {
        console.error('❌ Failed to create election', err);

        // ❌ Error SweetAlert
        Swal.fire('Error', err.response?.data || 'Something went wrong!', 'error');
      }
    },

    async deleteElection(id) {
      if (confirm('Are you sure you want to delete this election?')) {
        try {
          await api.delete(`/admin/elections/delete/${id}`);
          this.fetchElections();
          Swal.fire('Deleted!', 'Election has been deleted.', 'success');
        } catch (err) {
          console.error('❌ Failed to delete election', err);
          Swal.fire('Error', err.response?.data || 'Failed to delete!', 'error');
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
        const payload = {
          ...this.editingElection,
          election_date: new Date(this.editingElection.election_date).toISOString()
        };
        await api.put(`/admin/elections/update/${this.editingElection.id}`, payload);
        this.editingElection = null;
        this.fetchElections();

        // ✅ Success SweetAlert
        Swal.fire({
          icon: 'success',
          title: 'Election Updated!',
          text: '✅ Election changes saved successfully!',
          timer: 2000,
          showConfirmButton: false,
        });

      } catch (err) {
        console.error('❌ Failed to update election', err);

        // ❌ Error SweetAlert
        Swal.fire('Error', err.response?.data || 'Failed to update election.', 'error');
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
button:hover, .manage-btn:hover {
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
.manage-btn {
  padding: 8px 12px;
  background-color: #27ae60;
  color: white;
  border-radius: 8px;
  text-decoration: none;
  font-weight: bold;
  margin-left: 8px;
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
