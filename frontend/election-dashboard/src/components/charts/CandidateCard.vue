<template>
  <div class="candidate-card">
<!--    <img :src="`/uploads/${candidate.photo}`" alt="Candidate" />-->
    <h3>{{ candidate.name }}</h3>
    <p>{{ candidate.party }}</p>
    <button
        :disabled="hasVoted"
        @click="$emit('vote')"
    >
      {{ hasVoted ? "✅ Voted" : "Vote" }}
    </button>
  </div>
</template>

<script>
export default {
  props: {
    candidate: Object,
    hasVoted: Boolean
  }
};
</script>

<style scoped>
.candidate-card {
  background: white;
  border-radius: 15px;
  padding: 20px;
  width: 200px;
  text-align: center;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  transition: transform 0.3s;
}
.candidate-card:hover {
  transform: translateY(-5px);
}

img {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 15px;
}

h3 {
  font-size: 18px;
  margin-bottom: 5px;
}

p {
  font-size: 14px;
  color: gray;
  margin-bottom: 15px;
}

button {
  padding: 10px 20px;
  background: #3498db;
  border: none;
  border-radius: 8px;
  color: white;
  cursor: pointer;
}
button:hover {
  background: #2980b9;
}
button:disabled {
  background: #ccc;
  cursor: not-allowed;
}
</style>
