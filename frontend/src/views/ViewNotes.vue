<script setup lang="ts">
import { ref, onMounted } from 'vue'

const notes = ref('')

const GetNotes = async () => {
  try {
    const apiUrl = 'http://localhost:8080/note'
    const response = await fetch(apiUrl, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('token')
      }
    })

    const data = await response.json()
    if (response.ok) {
      notes.value = data
    } else {
      console.log(data.error)
      alert(data.error)
    }
  } catch (error) {
    console.error('Error: ', error)
  }
}

onMounted(() => {
  GetNotes()
})
</script>

<template>
  <h1 id="header">My notes</h1>
  <div class="viewNotes" v-if="notes != ''">
    <div v-for="note in notes" :key="note.id" class="viewNote-container">
      <h1>{{ note.title }}</h1>
      <p id="view_viewer_number">viewer number: {{ note.current_views }} / {{ note.max_views }}</p>
      <p id="view_content">{{ note.content }}</p>
    </div>
  </div>
</template>

<style scoped>
#header {
  font-size: 30;
  margin: 30px;
  color: #4caf50;
}
.viewNote-container {
  max-width: 400px;
  min-width: 30%;
  padding: 1%;
  border: 1px solid #ccc;
  border-radius: 8px;
}

.viewNotes {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  margin: auto;
}

h1 {
  font-size: 24px;
  margin-bottom: 20px;
  text-align: center;
}
</style>
