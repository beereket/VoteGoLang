<template>
  <div class="candidate-management">
    <h1>👑 Manage Candidates for {{ election.name }}</h1>

    <button @click="showCreateForm = true" class="add-btn">➕ Add Candidate</button>

    <div v-if="candidates.length === 0" class="empty-state">
      No candidates yet.
    </div>

    <div v-else class="candidate-list">
      <div v-for="candidate in candidates" :key="candidate.id" class="candidate-item">
        <div>
          <strong>{{ candidate.name }}</strong> ({{ candidate.party }}) — {{ candidate.region }} — {{ candidate.type }}
        </div>
        <div class="actions">
          <button @click="startEdit(candidate)">✏️ Edit</button>
          <button @click="deleteCandidate(candidate.id)">🗑️ Delete</button>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <div v-if="showCreateForm || editingCandidate" class="modal">
      <div class="modal-content">
        <h2>{{ editingCandidate ? 'Edit Candidate' : 'Add Candidate' }}</h2>
        <form @submit.prevent="submitForm">
          <input v-model="form.name" type="text" placeholder="Name" required />
          <input v-model="form.party" type="text" placeholder="Party" required />
          <input v-model="form.region" type="text" placeholder="Region" required />
          <input v-model="form.age" type="number" placeholder="Age" required />
          <select v-model="form.type" required>
            <option disabled value="">Select Type</option>
            <option value="president">President</option>
            <option value="deputy">Deputy</option>
          </select>

          <div class="modal-actions">
            <button type="submit">{{ editingCandidate ? 'Save Changes' : 'Create Candidate' }}</button>
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
  name: "CandidateManagement",
  data() {
    return {
      election: {},
      candidates: [],
      showCreateForm: false,
      editingCandidate: null,
      form: {
        name: '',
        party: '',
        region: '',
        age: '',
        type: '',
      }
    }
  },
  methods: {
    async fetchElectionAndCandidates() {
      try {
        const electionId = this.$route.params.id;
        const res = await api.get(`/admin/elections/${electionId}/full`);
        this.election = res.data.election;
        this.candidates = res.data.candidates || [];
      } catch (err) {
        console.error('❌ Failed to load election or candidates', err);
        Swal.fire('Error', 'Failed to load election or candidates', 'error');
      }
    },
    async submitForm() {
      const electionId = this.$route.params.id;
      try {
        if (this.editingCandidate) {
          await api.put(`/admin/candidates/update/${this.editingCandidate.id}`, this.form);
          Swal.fire('Updated!', 'Candidate updated successfully.', 'success');
        } else {
          await api.post('/admin/candidates/create', {
            ...this.form,
            election_id: electionId,
          });
          Swal.fire('Created!', 'Candidate created successfully.', 'success');
        }
        this.cancelEdit();
        this.fetchElectionAndCandidates();
      } catch (err) {
        console.error('❌ Failed to save candidate', err);
        Swal.fire('Error', 'Failed to save candidate.', 'error');
      }
    },
    startEdit(candidate) {
      this.editingCandidate = candidate;
      this.form = { ...candidate };
    },
    cancelEdit() {
      this.editingCandidate = null;
      this.showCreateForm = false;
      this.form = { name: '', party: '', region: '', age: '', type: '' };
    },
    async deleteCandidate(id) {
      const confirm = await Swal.fire({
        title: 'Are you sure?',
        text: "This will permanently delete the candidate.",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#3085d6',
        confirmButtonText: 'Yes, delete it!'
      });

      if (confirm.isConfirmed) {
        try {
          await api.delete(`/admin/candidates/delete/${id}`);
          this.fetchElectionAndCandidates();
          Swal.fire('Deleted!', 'Candidate has been deleted.', 'success');
        } catch (err) {
          console.error('❌ Failed to delete candidate', err);
          Swal.fire('Error', 'Failed to delete candidate.', 'error');
        }
      }
    }
  },
  mounted() {
    this.fetchElectionAndCandidates();
  }
}
</script>

<style scoped>
.candidate-management {
  padding: 30px;
  font-family: 'Poppins', sans-serif;
}
h1 {
  text-align: center;
  margin-bottom: 20px;
}
.add-btn {
  margin-bottom: 20px;
}
button {
  padding: 10px;
  background: #3498db;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  margin: 5px;
}
button:hover {
  background: #2980b9;
}
.candidate-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}
.candidate-item {
  background: #f9f9f9;
  padding: 15px;
  border-radius: 12px;
  display: flex;
  justify-content: space-between;
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
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
}
.empty-state {
  text-align: center;
  margin-top: 50px;
  color: #888;
}
</style>