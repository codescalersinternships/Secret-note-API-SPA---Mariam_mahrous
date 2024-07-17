<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
const route = useRoute()
const title = ref('')
const content = ref('')
const maxViews = ref('')
const views = ref('')

onMounted(() => {
  GetNoteByID()
})

const GetNoteByID = async () => {
  try {
    let apiUrl = 'http://localhost:8080/note/' + route.params.uuid
    const response = await fetch(apiUrl, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })

    const data = await response.json()
    if (response.ok) {
      console.log(route.params.uuid)
      title.value = data.title
      content.value = data.content
      maxViews.value = data.max_views
      views.value = data.current_views
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
  <div class="viewNote-container">
    <h1>{{ title }}</h1>
    <p id="view_viewer_number">viewer number : {{ views }} / {{ maxViews }}</p>
    <p id="view_content">{{ content }}</p>
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
