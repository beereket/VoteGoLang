<template>
  <div class="p-8">
    <h1 class="text-3xl font-bold text-gray-800 mb-6">Admin Users</h1>

    <div v-if="admins.length > 0">
      <table class="w-full bg-white rounded-lg shadow-md">
        <thead>
        <tr class="text-left border-b">
          <th class="p-4">Username</th>
          <th class="p-4">Full Name</th>
          <th class="p-4">Birth Date</th>
          <th class="p-4">Address</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="admin in admins" :key="admin.id" class="border-b hover:bg-gray-100">
          <td class="p-4">{{ admin.username }}</td>
          <td class="p-4">{{ admin.user_full_name }}</td>
          <td class="p-4">{{ admin.birth_date }}</td>
          <td class="p-4">{{ admin.address }}</td>
        </tr>
        </tbody>
      </table>
    </div>

    <div v-else class="text-gray-600">
      No admin users found.
    </div>

  </div>
</template>

<script>
import api from '@/services/api';

export default {
  name: 'AdminUsersPage',
  data() {
    return {
      admins: [],
    };
  },
  methods: {
    async fetchAdmins() {
      try {
        const response = await api.get('/admin/users/admins');
        this.admins = response.data;
      } catch (error) {
        console.error('❌ Failed to fetch admin users', error);
      }
    }
  },
  mounted() {
    this.fetchAdmins();
  }
};
</script>

<style scoped>
</style>
