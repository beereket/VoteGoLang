<template>
  <div class="petition-admin-page">
    <h1 class="page-title">📝 Manage Petitions</h1>

    <div class="create-form">
      <h2>Create New Petition</h2>
      <form @submit.prevent="createPetition">
        <input v-model="newPetition.title" type="text" placeholder="Title" required />
        <input v-model="newPetition.photo" type="text" placeholder="Photo URL" required />
        <textarea v-model="newPetition.description" placeholder="Description" required></textarea>
        <button type="submit">Create Petition</button>
      </form>
    </div>

    <div class="petition-list">
      <h2>Existing Petitions</h2>
      <div v-for="petition in petitions" :key="petition.id" class="petition-card">
        <div v-if="editingId === petition.id">
          <input v-model="editPetition.title" placeholder="Title" />
          <input v-model="editPetition.photo" placeholder="Photo URL" />
          <textarea v-model="editPetition.description" placeholder="Description"></textarea>
          <button @click="updatePetition(petition.id)">Save</button>
          <button @click="cancelEdit">Cancel</button>
        </div>
        <div v-else>
          <h3>{{ petition.title }}</h3>
          <img src="https://as2.ftcdn.net/v2/jpg/05/39/88/27/1000_F_539882795_JhDKGYvDELRfiPr0fSChlkwMhblvHoQH.jpg" alt="Petition Photo" class="petition-photo" />
          <p>{{ petition.description }}</p>
          <div class="buttons">
            <button @click="startEdit(petition)">Edit</button>
            <button @click="deletePetition(petition.id)">Delete</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import api from '@/services/api';
import { ref, onMounted } from 'vue';

export default {
  name: 'PetitionAdminPage',
  setup() {
    const petitions = ref([]);
    const newPetition = ref({ title: '', photo: '', description: '' });
    const editPetition = ref({ title: '', photo: '', description: '' });
    const editingId = ref(null);

    const fetchPetitions = async () => {
      const res = await api.get('/petitions');
      petitions.value = res.data;
    };

    const createPetition = async () => {
      await api.post('/admin/petitions/create', newPetition.value);
      await fetchPetitions();
      newPetition.value = { title: '', photo: '', description: '' };
    };

    const startEdit = (petition) => {
      editingId.value = petition.id;
      editPetition.value = { ...petition };
    };

    const cancelEdit = () => {
      editingId.value = null;
      editPetition.value = { title: '', photo: '', description: '' };
    };

    const updatePetition = async (id) => {
      await api.put(`/admin/petitions/update/${id}`, editPetition.value);
      editingId.value = null;
      await fetchPetitions();
    };

    const deletePetition = async (id) => {
      if (confirm('Are you sure you want to delete this petition?')) {
        await api.delete(`/admin/petitions/delete/${id}`);
        await fetchPetitions();
      }
    };

    onMounted(fetchPetitions);

    return {
      petitions,
      newPetition,
      editPetition,
      editingId,
      createPetition,
      startEdit,
      cancelEdit,
      updatePetition,
      deletePetition,
    };
  }
};
</script>

<style scoped>
.petition-admin-page {
  padding: 40px;
}
.page-title {
  text-align: center;
  margin-bottom: 30px;
}
.create-form, .petition-list {
  background: white;
  padding: 20px;
  border-radius: 12px;
  margin-bottom: 40px;
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
}
input, textarea {
  width: 100%;
  margin-bottom: 15px;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 8px;
}
button {
  margin-right: 10px;
  padding: 10px 20px;
  border: none;
  background: #3498db;
  color: white;
  border-radius: 8px;
  cursor: pointer;
}
button:hover {
  background: #2980b9;
}
.petition-card {
  margin-bottom: 20px;
  padding: 20px;
  background: #f9f9f9;
  border-radius: 10px;
}
.petition-photo {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 8px;
  margin: 10px 0;
}
.buttons {
  margin-top: 10px;
}
</style>
