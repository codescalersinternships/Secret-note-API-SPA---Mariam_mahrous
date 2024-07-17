<script setup lang="ts">
    import { useRouter } from 'vue-router'

    const router = useRouter()

    const naviagteToNotes = () => {
        router.push('/note')
    };
    const naviagteToCreateNotes = () => {
        router.push('/note/create')
    };
</script>


<template>
    <button type="submit" @click="naviagteToNotes">View Notes</button>
    <button type="submit" @click="naviagteToCreateNotes">Create a new note</button>

</template>

<style scoped>
button {

  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
  width: fit-Content;
  margin: 5px;
  padding: 12px;
  border-radius: 5px;

}
</style>