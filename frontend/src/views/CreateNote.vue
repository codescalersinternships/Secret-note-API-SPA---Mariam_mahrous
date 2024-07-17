<script setup lang="ts">
</script>


<template>
  <div class="createNote-container">

    <h1>Create Note</h1>
    <form @submit.prevent="createNote">
      <label for="title">Title:</label> <br> <br>
        <input type="text" id="title" name="title" required><br><br>

        <label for="content">Content:</label> <br><br>
        <textarea id="content" name="content" required></textarea><br><br>

        <label for="expiry_date">Expiry Date:</label> <br> <br>
        <input type="date" id="expiry_date" name="expiry_date"><br><br>

        <label for="max_views">Max No of View:</label> <br> <br>
        <input type="number" id="max_views" name="max_views" min="0"><br><br>

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

  background-color: #4CAF50;
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
    height:20vh;
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

