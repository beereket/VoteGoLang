<template>
  <div class="petition-comments-page">
    <h1>🗳️ Petition Comments Management</h1>

    <div class="petition-list">
      <h2>Select Petition:</h2>
      <div v-for="petition in petitions" :key="petition.id" class="petition-item">
        <button @click="selectPetition(petition)">
          {{ petition.title }}
        </button>
      </div>
    </div>

    <div v-if="selectedPetition" class="comments-section">
      <h2>Comments for: {{ selectedPetition.title }}</h2>

      <div v-if="comments.length > 0" class="comments-list">
        <div v-for="comment in comments" :key="comment.id" class="comment-card">
          <p><strong>User ID:</strong> {{ comment.user_id }}</p>
          <p><strong>Comment:</strong> {{ comment.comment_text }}</p>
          <p><strong>Upvotes:</strong> {{ comment.upvotes }} | <strong>Downvotes:</strong> {{ comment.downvotes }}</p>
          <button @click="deleteComment(comment.id)" class="delete-button">❌ Delete Comment</button>
        </div>
      </div>

      <div v-else class="no-comments">
        No comments for this petition.
      </div>
    </div>

    <div v-else class="no-selection">
      <p>Please select a petition to view comments 📋</p>
    </div>

  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import api from '@/services/api';

export default {
  name: "PetitionsCommentPage",
  setup() {
    const petitions = ref([]);
    const selectedPetition = ref(null);
    const comments = ref([]);

    const fetchPetitions = async () => {
      try {
        const res = await api.get('/petitions');
        petitions.value = res.data;
      } catch (err) {
        console.error('❌ Failed to fetch petitions', err);
      }
    };

    const fetchComments = async (petitionId) => {
      try {
        const res = await api.get(`/petitions/${petitionId}/comments`);
        comments.value = res.data; // NO FILTER here
      } catch (err) {
        console.error('❌ Failed to fetch comments', err);
      }
    };

    const selectPetition = (petition) => {
      selectedPetition.value = petition;
      comments.value = [];
      fetchComments(petition.id);
    };


    const deleteComment = async (commentId) => {
      if (confirm('⚠️ Are you sure you want to delete this comment?')) {
        try {
          await api.delete(`/admin/petitions/comments/delete/${commentId}`);
          fetchComments(selectedPetition.value.id);
        } catch (err) {
          console.error('❌ Failed to delete comment', err);
        }
      }
    };

    onMounted(fetchPetitions);

    return {
      petitions,
      selectedPetition,
      comments,
      selectPetition,
      deleteComment
    };
  }
};
</script>

<style scoped>
.petition-comments-page {
  padding: 30px;
  font-family: 'Poppins', sans-serif;
  background: #f5f7fa;
  min-height: 100vh;
}

h1 {
  text-align: center;
  margin-bottom: 30px;
  font-size: 36px;
}

.petition-list {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;
  justify-content: center;
  margin-bottom: 30px;
}

.petition-item button {
  background: #4CAF50;
  color: white;
  padding: 12px 18px;
  border: none;
  border-radius: 10px;
  cursor: pointer;
}

.petition-item button:hover {
  background: #45a049;
}

.comments-section {
  margin-top: 30px;
}

.comments-list {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
}

.comment-card {
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

.delete-button {
  margin-top: 10px;
  background: #e74c3c;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 8px;
  cursor: pointer;
}

.delete-button:hover {
  background: #c0392b;
}

.no-selection, .no-comments {
  text-align: center;
  margin-top: 50px;
  color: #777;
}
</style>
