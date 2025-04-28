<template>
  <div class="news-page">
    <h1>📰 News Management</h1>

    <div class="create-form">
      <input v-model="newNews.title" type="text" placeholder="News Title" />
      <input v-model="newNews.photo" type="text" placeholder="Photo URL" />
      <textarea v-model="newNews.paragraph" placeholder="Paragraph"></textarea>
      <button @click="createNews">➕ Create News</button>
    </div>

    <div class="news-list">
      <div v-for="news in newsList" :key="news.id" class="news-card">
        <img v-if="news.photo" :src="news.photo" alt="news" class="news-photo" />
        <div class="news-content">
          <h2>{{ news.title }}</h2>
          <p>{{ news.paragraph }}</p>
          <div class="news-meta">
            <small>Created: {{ formatDate(news.created_at) }}</small>
            <button @click="deleteNews(news.id)">🗑️ Delete</button>
          </div>
        </div>
      </div>
    </div>

    <div v-if="newsList.length === 0" class="no-news">
      No news found. 😢
    </div>
  </div>
</template>

<script>
import api from '@/services/api';
import Swal from 'sweetalert2';

export default {
  name: 'NewsPage',
  data() {
    return {
      newsList: [],
      newNews: {
        title: '',
        paragraph: '',
        photo: ''
      }
    }
  },
  methods: {
    async fetchNews() {
      try {
        const res = await api.get('/news');
        this.newsList = res.data;
      } catch (error) {
        console.error('❌ Failed to fetch news', error);
      }
    },
    async createNews() {
      if (!this.newNews.title || !this.newNews.paragraph) {
        Swal.fire('❗ Error', 'Please fill all fields', 'error');
        return;
      }
      try {
        await api.post('/admin/news/create', this.newNews);
        Swal.fire('✅ Success', 'News created!', 'success');
        this.newNews = { title: '', paragraph: '', photo: '' };
        this.fetchNews();
      } catch (error) {
        console.error('❌ Failed to create news', error);
        Swal.fire('❌ Error', 'Failed to create news', 'error');
      }
    },
    async deleteNews(id) {
      const confirm = await Swal.fire({
        title: 'Are you sure?',
        text: "This news will be deleted!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonText: 'Yes, delete it!',
      });
      if (confirm.isConfirmed) {
        try {
          await api.delete(`/admin/news/delete/${id}`);
          Swal.fire('✅ Deleted!', 'News has been deleted.', 'success');
          this.fetchNews();
        } catch (error) {
          console.error('❌ Failed to delete news', error);
          Swal.fire('❌ Error', 'Failed to delete news', 'error');
        }
      }
    },
    formatDate(dateStr) {
      const options = { year: 'numeric', month: 'short', day: 'numeric' };
      return new Date(dateStr).toLocaleDateString(undefined, options);
    }
  },
  mounted() {
    this.fetchNews();
  }
}
</script>

<style scoped>
.news-page {
  padding: 30px;
  font-family: 'Poppins', sans-serif;
  min-height: 100vh;
  background: #f7f9fb;
}

h1 {
  text-align: center;
  margin-bottom: 30px;
  color: #2c3e50;
}

.create-form {
  display: flex;
  flex-direction: column;
  gap: 12px;
  max-width: 500px;
  margin: 0 auto 50px auto;
}

.create-form input,
.create-form textarea {
  padding: 12px;
  border: 1px solid #ccc;
  border-radius: 8px;
  resize: none;
}

.create-form button {
  background-color: #4CAF50;
  color: white;
  padding: 12px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
}

.create-form button:hover {
  background-color: #45a049;
}

.news-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
  gap: 30px;
  padding: 20px;
}

.news-card {
  background: white;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.news-card:hover {
  transform: translateY(-6px) scale(1.02);
  box-shadow: 0 12px 20px rgba(0,0,0,0.15);
}

.news-photo {
  width: 100%;
  height: 250px;
  object-fit: cover;
}

.news-content {
  padding: 20px;
}

.news-content h2 {
  margin-bottom: 10px;
  font-size: 22px;
  color: #333;
}

.news-content p {
  font-size: 16px;
  color: #555;
}

.news-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 20px;
}

.news-meta button {
  background: #e74c3c;
  border: none;
  color: white;
  padding: 8px 12px;
  border-radius: 6px;
  cursor: pointer;
}

.news-meta button:hover {
  background: #c0392b;
}

.no-news {
  text-align: center;
  margin-top: 50px;
  color: #777;
}
</style>