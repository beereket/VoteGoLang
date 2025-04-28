<template>
  <div class="user-management-page">
    <h1>👥 User Management</h1>

    <!-- Search input -->
    <div class="search-bar">
      <input
          v-model="searchTerm"
          type="text"
          placeholder="Search users..."
          class="search-input"
      />
    </div>

    <!-- Users Table -->
    <table class="users-table">
      <thead>
      <tr>
        <th>Username</th>
        <th>Full Name</th>
        <th>Birth Date</th>
        <th>Address</th>
        <th>Status</th>
        <th>Actions</th>
      </tr>
      </thead>

      <tbody>
      <tr v-for="user in paginatedUsers" :key="user.id">
        <td>{{ user.username }}</td>
        <td>{{ user.user_full_name }}</td>
        <td>{{ user.birth_date }}</td>
        <td>{{ user.address }}</td>
        <td>
          <span v-if="user.deleted_at" class="banned">Banned</span>
          <span v-else class="active">Active</span>
        </td>
        <td>
          <button
              v-if="!user.deleted_at"
              @click="banUser(user.id)"
              class="ban-btn"
          >
            Ban
          </button>
          <button
              v-else
              @click="unbanUser(user.id)"
              class="unban-btn"
          >
            Unban
          </button>
        </td>
      </tr>
      </tbody>
    </table>

    <!-- Pagination -->
    <div class="pagination">
      <button @click="prevPage" :disabled="currentPage === 1">Prev</button>
      <span>Page {{ currentPage }} / {{ totalPages }}</span>
      <button @click="nextPage" :disabled="currentPage === totalPages">Next</button>
    </div>
  </div>
</template>

<script>
import api from "@/services/api";

export default {
  name: "UserManagementPage",
  data() {
    return {
      users: [],
      searchTerm: "",
      currentPage: 1,
      pageSize: 8,
    };
  },
  computed: {
    filteredUsers() {
      const term = this.searchTerm.toLowerCase();
      return this.users.filter(
          (u) =>
              u.username.toLowerCase().includes(term) ||
              u.user_full_name.toLowerCase().includes(term)
      );
    },
    paginatedUsers() {
      const start = (this.currentPage - 1) * this.pageSize;
      const end = start + this.pageSize;
      return this.filteredUsers.slice(start, end);
    },
    totalPages() {
      return Math.ceil(this.filteredUsers.length / this.pageSize);
    },
  },
  methods: {
    async fetchUsers() {
      try {
        const res = await api.get("/admin/users");
        this.users = res.data;
      } catch (err) {
        console.error("❌ Failed to fetch users", err);
      }
    },
    async banUser(id) {
      if (confirm("❗ Are you sure you want to ban this user?")) {
        try {
          await api.delete(`/admin/users/ban/${id}`);
          this.fetchUsers();
        } catch (err) {
          console.error("❌ Failed to ban user", err);
        }
      }
    },
    async unbanUser(id) {
      if (confirm("❗ Are you sure you want to unban this user?")) {
        try {
          await api.put(`/admin/users/unban/${id}`);
          this.fetchUsers();
        } catch (err) {
          console.error("❌ Failed to unban user", err);
        }
      }
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
      }
    },
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
      }
    },
  },
  mounted() {
    this.fetchUsers();
  },
};
</script>

<style scoped>
.user-management-page {
  padding: 40px;
  font-family: 'Poppins', sans-serif;
}

h1 {
  text-align: center;
  margin-bottom: 30px;
}

.search-bar {
  text-align: center;
  margin-bottom: 20px;
}

.search-input {
  padding: 10px 20px;
  width: 300px;
  border-radius: 8px;
  border: 1px solid #ccc;
}

.users-table {
  width: 100%;
  border-collapse: collapse;
  background: #fff;
}

.users-table th,
.users-table td {
  padding: 12px 15px;
  border-bottom: 1px solid #eee;
  text-align: left;
}

.banned {
  color: red;
  font-weight: bold;
}

.active {
  color: green;
  font-weight: bold;
}

.ban-btn, .unban-btn {
  padding: 8px 16px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  color: white;
  font-weight: 600;
}

.ban-btn {
  background: #e74c3c;
}

.unban-btn {
  background: #2ecc71;
}

.pagination {
  margin-top: 20px;
  text-align: center;
}

.pagination button {
  margin: 0 10px;
  padding: 8px 16px;
  background: #3498db;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
}

.pagination span {
  margin: 0 10px;
}
</style>
