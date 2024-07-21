<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { type Note } from '../models/Note'

const route = useRoute()
const note = ref<Note | null>(null)
const API_URL = import.meta.env.VITE_API_URL

onMounted(() => {
  GetNoteByID()
})

const GetNoteByID = async () => {
  try {
    let apiUrl = API_URL + '/note/' + route.params.uuid
    const response = await fetch(apiUrl, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    const data = await response.json()
    if (response.ok) {
      note.value = data
    } else {
      console.log(data.error)
      alert(data.error)
    }
  } catch (error) {
    console.error('Error: ', error)
  }
}
</script>

<template>
  <div class="viewNote-container" v-if="note">
    <h1>{{ note.title }}</h1>
    <p id="view_viewer_number">viewer number : {{ note.current_views }} / {{ note.max_views }}</p>
    <p id="view_content">{{ note.content }}</p>
  </div>
</template>

<style scoped>
.viewNote-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  margin-top: 20vh;
}

h1 {
  font-size: 24px;
  margin-bottom: 20px;
  text-align: center;
}
</style>
