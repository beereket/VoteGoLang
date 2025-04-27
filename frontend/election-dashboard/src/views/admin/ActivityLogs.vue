<template>
  <div class="logs-page">
    <h1>📜 Activity Logs</h1>

    <div class="logs-table">
      <table>
        <thead>
        <tr>
          <th>#</th>
          <th>User</th>
          <th>Action</th>
          <th>Description</th>
          <th>Time</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(log, index) in logs" :key="log.id">
          <td>{{ index + 1 }}</td>
          <td>{{ log.username || 'Unknown' }}</td>
          <td>{{ log.action }}</td>
          <td>{{ log.description }}</td>
          <td>{{ formatTime(log.created_at) }}</td>
        </tr>
        </tbody>
      </table>

      <div v-if="logs.length === 0" class="no-logs">
        No activity logs found 📭
      </div>
    </div>
  </div>
</template>

<script>
import api from '@/services/api';

export default {
  name: 'ActivityLogs',
  data() {
    return {
      logs: [],
    };
  },
  methods: {
    async fetchLogs() {
      try {
        const res = await api.get('/admin/logs');
        this.logs = res.data;
      } catch (err) {
        console.error('❌ Failed to fetch activity logs', err);
      }
    },
    formatTime(datetime) {
      const options = { year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' };
      return new Date(datetime).toLocaleString(undefined, options);
    }
  },
  mounted() {
    this.fetchLogs();
  }
};
</script>

<style scoped>
.logs-page {
  padding: 30px;
  font-family: 'Poppins', sans-serif;
}

h1 {
  text-align: center;
  font-size: 36px;
  margin-bottom: 30px;
}

.logs-table {
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
}

thead {
  background: #f0f0f0;
}

th, td {
  padding: 14px 20px;
  text-align: left;
}

tbody tr:nth-child(odd) {
  background: #fafafa;
}

.no-logs {
  text-align: center;
  padding: 30px;
  color: #999;
}
</style>
