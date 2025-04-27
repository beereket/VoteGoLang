<script>
import api from '@/services/api';
import Swal from 'sweetalert2';

export default {
  name: 'CandidatesPage',
  data() {
    return {
      candidates: [],
      showCreateModal: false,
      showEditModal: false,
      newCandidate: {
        name: '',
        photo: '',
        education: '',
        age: null,
        party: '',
        region: '',
        type: '',
      },
      editCandidate: {},
      currentPage: 1,
      itemsPerPage: 5,
      searchQuery: '',
      sortKey: 'name',
      sortOrder: 'asc',
    };
  },
  computed: {
    filteredCandidates() {
      if (!this.searchQuery) {
        return this.candidates;
      }
      const query = this.searchQuery.toLowerCase();
      return this.candidates.filter(c =>
          c.name.toLowerCase().includes(query) ||
          c.party.toLowerCase().includes(query) ||
          c.region.toLowerCase().includes(query)
      );
    },
    sortedCandidates() {
      const sorted = [...this.filteredCandidates];
      sorted.sort((a, b) => {
        let valA = a[this.sortKey]?.toString().toLowerCase();
        let valB = b[this.sortKey]?.toString().toLowerCase();

        if (valA < valB) return this.sortOrder === 'asc' ? -1 : 1;
        if (valA > valB) return this.sortOrder === 'asc' ? 1 : -1;
        return 0;
      });
      return sorted;
    },
    paginatedCandidates() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.sortedCandidates.slice(start, end);
    },
  },
  methods: {
    async fetchCandidates() {
      try {
        const response = await api.get('/candidates');
        this.candidates = response.data;
      } catch (error) {
        console.error('❌ Failed to fetch candidates', error);
      }
    },
    openCreateForm() {
      this.showCreateModal = true;
    },
    closeCreateForm() {
      this.showCreateModal = false;
      this.resetForm();
    },
    resetForm() {
      this.newCandidate = {
        name: '',
        photo: '',
        education: '',
        age: null,
        party: '',
        region: '',
        type: '',
      };
    },
    async createCandidate() {
      if (!this.newCandidate.name || !this.newCandidate.type || !this.newCandidate.age) {
        Swal.fire({
          icon: 'warning',
          title: '❗ Please fill required fields!',
        });
        return;
      }
      try {
        await api.post('/admin/candidates/create', this.newCandidate);
        this.fetchCandidates();
        this.closeCreateForm();
        Swal.fire({
          icon: 'success',
          title: '✅ Candidate created!',
          showConfirmButton: false,
          timer: 1500
        });
      } catch (error) {
        console.error('❌ Failed to create candidate', error);
      }
    },
    openEditForm(candidate) {
      this.showEditModal = true;
      this.editCandidate = { ...candidate };
    },
    closeEditForm() {
      this.showEditModal = false;
      this.editCandidate = {};
    },
    async updateCandidate() {
      if (!this.editCandidate.name || !this.editCandidate.type || !this.editCandidate.age) {
        Swal.fire({
          icon: 'warning',
          title: '❗ Please fill required fields!',
        });
        return;
      }
      try {
        await api.put(`/admin/candidates/update/${this.editCandidate.id}`, this.editCandidate);
        this.fetchCandidates();
        this.closeEditForm();
        Swal.fire({
          icon: 'success',
          title: '✅ Candidate updated!',
          showConfirmButton: false,
          timer: 1500
        });
      } catch (error) {
        console.error('❌ Failed to update candidate', error);
      }
    },
    async deleteCandidate(id) {
      if (confirm('❗ Are you sure you want to delete this candidate?')) {
        try {
          await api.delete(`/admin/candidates/delete/${id}`);
          this.fetchCandidates();
          Swal.fire({
            icon: 'success',
            title: '✅ Candidate deleted!',
            showConfirmButton: false,
            timer: 1500
          });
        } catch (error) {
          console.error('❌ Failed to delete candidate', error);
        }
      }
    },
    nextPage() {
      if (this.currentPage * this.itemsPerPage < this.sortedCandidates.length) {
        this.currentPage++;
      }
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
      }
    },
    sortBy(key) {
      if (this.sortKey === key) {
        this.sortOrder = this.sortOrder === 'asc' ? 'desc' : 'asc';
      } else {
        this.sortKey = key;
        this.sortOrder = 'asc';
      }
    },
  },
  mounted() {
    this.fetchCandidates();
  }
};
</script>

<template>
  <div class="page-container">
    <div class="header">
      <h1>Candidates</h1>
      <button class="add-button" @click="openCreateForm">+ Add Candidate</button>
    </div>

    <div class="search-bar">
      <input
          type="text"
          v-model="searchQuery"
          placeholder="🔎 Search candidates..."
      />
    </div>

    <div class="sort-buttons">
      <button @click="sortBy('name')">
        Sort by Name ({{ sortOrder === 'asc' ? '↑' : '↓' }})
      </button>
      <button @click="sortBy('party')">
        Sort by Party ({{ sortOrder === 'asc' ? '↑' : '↓' }})
      </button>
    </div>

    <table class="candidates-table">
      <thead>
      <tr>
        <th>Name</th>
        <th>Party</th>
        <th>Region</th>
        <th>Age</th>
        <th>Votes</th>
        <th>Actions</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="candidate in paginatedCandidates" :key="candidate.id">
        <td>{{ candidate.name }}</td>
        <td>{{ candidate.party }}</td>
        <td>{{ candidate.region }}</td>
        <td>{{ candidate.age }}</td>
        <td>{{ candidate.votes }}</td>
        <td>
          <button @click="openEditForm(candidate)">Edit</button>
          <button @click="deleteCandidate(candidate.id)">Delete</button>
        </td>
      </tr>
      </tbody>
    </table>

    <div class="pagination">
      <button @click="prevPage" :disabled="currentPage === 1">Previous</button>
      <span>Page {{ currentPage }}</span>
      <button @click="nextPage" :disabled="currentPage * itemsPerPage >= sortedCandidates.length">Next</button>
    </div>

    <!-- Create Modal -->
    <div v-if="showCreateModal" class="modal-overlay">
      <div class="modal">
        <h2>Create Candidate</h2>
        <input v-model="newCandidate.name" placeholder="Name" />
        <input v-model="newCandidate.photo" placeholder="Photo URL" />
        <input v-model="newCandidate.education" placeholder="Education" />
        <input v-model.number="newCandidate.age" placeholder="Age" type="number" />
        <input v-model="newCandidate.party" placeholder="Party" />
        <input v-model="newCandidate.region" placeholder="Region" />
        <input v-model="newCandidate.type" placeholder="Type (president, deputy)" />
        <div class="modal-actions">
          <button @click="createCandidate">Submit</button>
          <button @click="closeCreateForm">Cancel</button>
        </div>
      </div>
    </div>

    <!-- Edit Modal -->
    <div v-if="showEditModal" class="modal-overlay">
      <div class="modal">
        <h2>Edit Candidate</h2>
        <input v-model="editCandidate.name" placeholder="Name" />
        <input v-model="editCandidate.photo" placeholder="Photo URL" />
        <input v-model="editCandidate.education" placeholder="Education" />
        <input v-model.number="editCandidate.age" placeholder="Age" type="number" />
        <input v-model="editCandidate.party" placeholder="Party" />
        <input v-model="editCandidate.region" placeholder="Region" />
        <input v-model="editCandidate.type" placeholder="Type (president, deputy)" />
        <div class="modal-actions">
          <button @click="updateCandidate">Save</button>
          <button @click="closeEditForm">Cancel</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.page-container {
  padding: 20px;
}
.header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}
.add-button {
  background-color: #4CAF50;
  color: white;
  padding: 10px 18px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
.add-button:hover {
  background-color: #45a049;
}
.candidates-table {
  width: 100%;
  border-collapse: collapse;
}
.candidates-table th, .candidates-table td {
  border: 1px solid #ddd;
  padding: 8px;
}
.candidates-table th {
  background-color: #f2f2f2;
  text-align: left;
}
.modal-overlay {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}
.modal {
  background: white;
  padding: 30px;
  border-radius: 8px;
  width: 400px;
  max-width: 90%;
}
.modal input {
  display: block;
  width: 100%;
  margin: 10px 0;
  padding: 8px;
}
.modal-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}
.search-bar {
  margin-bottom: 20px;
  text-align: center;
}
.search-bar input {
  width: 300px;
  padding: 10px;
  font-size: 16px;
}

.sort-buttons {
  margin-bottom: 20px;
  text-align: center;
}
.sort-buttons button {
  margin: 0 5px;
  padding: 8px 14px;
  background-color: #2196F3;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
.sort-buttons button:hover {
  background-color: #1976D2;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
  gap: 10px;
}
.pagination button {
  background-color: #4CAF50;
  color: white;
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
.pagination button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
</style>
