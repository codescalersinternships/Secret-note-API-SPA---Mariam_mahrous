<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const naviagteToNotes = () => {
  router.push('/note')
}
const naviagteToCreateNotes = () => {
  router.push('/note/create')
}
</script>

<template>
  <div>
    <h1>Homepage</h1>
    <button type="submit" @click="naviagteToNotes" id="view-Notes">View Notes</button>
    <button type="submit" @click="naviagteToCreateNotes" id="new-note">Create a new note</button>
  </div>
</template>

<style scoped>
button {
  background-color: #4caf50;
  color: white;
  border: none;
  cursor: pointer;
  width: fit-Content;
  margin: 5px;
  padding: 12px;
  border-radius: 5px;
}

h1 {
  font-size: 24px;
  margin-bottom: 20px;
  text-align: center;
}

div {
  width: fit-content;
  margin: 0 auto;
  padding-right: 4%;
  padding-left: 4%;
  padding-top: 2%;
  padding-bottom: 2%;
  border: 1px solid #ccc;
  border-radius: 8px;
  margin-top: 20vh;
  justify-content: space-between;
}
</style>
