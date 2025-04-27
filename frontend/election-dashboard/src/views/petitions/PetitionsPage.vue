<script>
import { ref, computed, onMounted } from "vue";
import api from "@/services/api";
import Swal from "sweetalert2";

export default {
  name: "PetitionsPage",
  setup() {
    const petitions = ref([]);
    const newPetition = ref({
      title: '',
      description: '',
      photo: ''
    });
    const editingPetition = ref(null);
    const searchQuery = ref('');
    const currentPage = ref(1);
    const petitionsPerPage = 6;

    const fetchPetitions = async () => {
      try {
        const res = await api.get('/petitions');
        petitions.value = res.data;
      } catch (err) {
        console.error("❌ Failed to fetch petitions", err);
      }
    };

    const savePetition = async () => {
      if (!newPetition.value.title || !newPetition.value.description) {
        Swal.fire('Error', 'Please fill in all fields', 'error');
        return;
      }

      if (editingPetition.value) {
        try {
          await api.put(`/admin/petitions/update/${editingPetition.value}`, newPetition.value);
          Swal.fire('✅ Success', 'Petition updated!', 'success');
          fetchPetitions();
          resetForm();
        } catch (err) {
          Swal.fire('❌ Failed', 'Could not update petition', 'error');
          console.error(err);
        }
      } else {
        try {
          await api.post('/admin/petitions/create', newPetition.value);
          Swal.fire('✅ Success', 'Petition created!', 'success');
          fetchPetitions();
          resetForm();
        } catch (err) {
          Swal.fire('❌ Failed', 'Could not create petition', 'error');
          console.error(err);
        }
      }
    };

    const deletePetition = async (id) => {
      const confirm = await Swal.fire({
        title: 'Are you sure?',
        text: "You won't be able to revert this!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonText: 'Yes, delete it!'
      });
      if (confirm.isConfirmed) {
        try {
          await api.delete(`/admin/petitions/delete/${id}`);
          Swal.fire('Deleted!', 'Petition deleted.', 'success');
          fetchPetitions();
        } catch (err) {
          Swal.fire('❌ Failed', 'Could not delete petition', 'error');
          console.error(err);
        }
      }
    };

    const startEditing = (petition) => {
      newPetition.value = { ...petition };
      editingPetition.value = petition.id;
    };

    const resetForm = () => {
      newPetition.value = { title: '', description: '', photo: '' };
      editingPetition.value = null;
    };

    const filteredPetitions = computed(() => {
      return petitions.value.filter(p =>
          p.title.toLowerCase().includes(searchQuery.value.toLowerCase())
      );
    });

    const paginatedPetitions = computed(() => {
      const start = (currentPage.value - 1) * petitionsPerPage;
      const end = start + petitionsPerPage;
      return filteredPetitions.value.slice(start, end);
    });

    const totalPages = computed(() => {
      return Math.ceil(filteredPetitions.value.length / petitionsPerPage);
    });

    const nextPage = () => {
      if (currentPage.value < totalPages.value) {
        currentPage.value++;
      }
    };

    const prevPage = () => {
      if (currentPage.value > 1) {
        currentPage.value--;
      }
    };

    onMounted(fetchPetitions);

    return {
      petitions,
      newPetition,
      savePetition,
      deletePetition,
      startEditing,
      searchQuery,
      currentPage,
      nextPage,
      prevPage,
      paginatedPetitions,
      totalPages,
      editingPetition,
    };
  }
};
</script>

<template>
  <div class="petition-page">
    <h1>📜 Petition Management</h1>

    <div class="create-form">
      <h2>{{ editingPetition ? "Edit Petition" : "Create Petition" }}</h2>
      <input v-model="newPetition.title" type="text" placeholder="Title" />
      <textarea v-model="newPetition.description" placeholder="Description"></textarea>
      <input v-model="newPetition.photo" type="text" placeholder="Photo URL (optional)" />
      <button @click="savePetition">{{ editingPetition ? "Update Petition" : "Create Petition" }}</button>
    </div>

    <div class="search-bar">
      <input v-model="searchQuery" type="text" placeholder="🔎 Search petitions..." />
    </div>

    <div v-if="paginatedPetitions.length > 0" class="petition-list">
      <div v-for="petition in paginatedPetitions" :key="petition.id" class="petition-card">
        <h3>{{ petition.title }}</h3>
        <p>{{ petition.description }}</p>
        <div class="button-actions">
          <button @click="startEditing(petition)" class="edit-btn">✏️ Edit</button>
          <button @click="deletePetition(petition.id)" class="delete-btn">🗑️ Delete</button>
        </div>
      </div>
    </div>

    <div v-else class="no-results">
      No petitions found.
    </div>

    <div class="pagination" v-if="totalPages > 1">
      <button @click="prevPage" :disabled="currentPage === 1">Previous</button>
      <span>Page {{ currentPage }} of {{ totalPages }}</span>
      <button @click="nextPage" :disabled="currentPage === totalPages">Next</button>
    </div>
  </div>
</template>

<style scoped>
.petition-page {
  padding: 30px;
  font-family: 'Poppins', sans-serif;
}

h1 {
  text-align: center;
  font-size: 32px;
  margin-bottom: 20px;
}

.create-form {
  background: #f9f9f9;
  padding: 20px;
  margin-bottom: 30px;
  border-radius: 10px;
}

.create-form input,
.create-form textarea {
  width: 100%;
  padding: 12px;
  margin: 8px 0;
  border-radius: 8px;
  border: 1px solid #ccc;
}

.create-form button {
  background: #4CAF50;
  color: white;
  padding: 12px 20px;
  border: none;
  border-radius: 8px;
  margin-top: 10px;
  cursor: pointer;
}

.create-form button:hover {
  background: #45a049;
}

.search-bar {
  margin: 20px 0;
}

.search-bar input {
  width: 100%;
  padding: 12px;
  border-radius: 8px;
  border: 1px solid #ccc;
}

.petition-list {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 20px;
}

.petition-card {
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  position: relative;
}

.petition-card h3 {
  margin-bottom: 10px;
}

.petition-card p {
  margin-bottom: 10px;
  color: #666;
}

.button-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.edit-btn {
  background: #3498db;
  border: none;
  padding: 8px 12px;
  color: white;
  border-radius: 8px;
  cursor: pointer;
}

.edit-btn:hover {
  background: #2980b9;
}

.delete-btn {
  background: #ff4d4d;
  border: none;
  padding: 8px 12px;
  color: white;
  border-radius: 8px;
  cursor: pointer;
}

.delete-btn:hover {
  background: #e60000;
}

.no-results {
  text-align: center;
  font-size: 18px;
  color: #777;
  margin-top: 20px;
}

.pagination {
  margin-top: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 20px;
}

.pagination button {
  background: #3498db;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
}

.pagination button:disabled {
  background: #ccc;
  cursor: not-allowed;
}
</style>
