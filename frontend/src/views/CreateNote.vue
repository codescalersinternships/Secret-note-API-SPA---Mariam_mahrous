<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { NoteData } from '../models/Note'

const router = useRouter()
let title = ref('')
let content = ref('')
let date = ref('')
let views = ref('')

const createNote = async () => {
  try {
    let apiUrl = 'http://localhost:8080/note/create'
    let data = new NoteData(title.value, content.value)
    if (date.value!='')
      data.expiration_date=date.value
    if (views.value!='')
      data.max_views=views.value
    const response = await fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('token')
      },
      body: JSON.stringify(data)
    })

    const resp = await response.json()
    if (response.ok) {
      console.log(resp)
      router.push('/note/' + resp.unique_url)
    } else {
      console.log(resp.error)
      alert(resp.error)
    }
  } catch (error) {
    console.log('Error: ', error)
  }
}
</script>

<template>
  <div class="createNote-container">
    <h1>Create Note</h1>
    <form @submit.prevent="createNote">
      <label for="title">Title:</label> <br />
      <br />
      <input type="text" placeholder="Enter a Tilte" v-model="title" required /><br /><br />

      <label for="content">Content:</label> <br /><br />
      <textarea id="content" v-model="content" required></textarea><br /><br />

      <label for="expiry_date">Expiry Date:</label> <br />
      <br />
      <input type="date" id="expiry_date" name="expiry_date" v-model="date" /><br /><br />

      <label for="max_views">Max No of View:</label> <br />
      <br />
      <input type="number" id="max_views" name="max_views" min="0" v-model="views" /><br /><br />

      <button type="submit">Create Note</button>
    </form>
  </div>
</template>

<style scoped>
.createNote-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  margin-top: 5vh;
}

h1 {
  font-size: 24px;
  margin-bottom: 20px;
  text-align: center;
}

form input {
  width: 100%;
  padding: 12px;
  margin-bottom: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  box-sizing: border-box;
}

form button {
  background-color: #4caf50;
  color: white;
  border: none;
  cursor: pointer;
  width: 100%;
  padding: 12px;
  border-radius: 5px;
}

button:hover {
  background-color: #45a049;
}

#content {
  height: 20vh;
  width: 100%;
  padding: 12px;
  margin-bottom: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  box-sizing: border-box;
  margin: auto;
  font-family: sans-serif;
}
</style>
