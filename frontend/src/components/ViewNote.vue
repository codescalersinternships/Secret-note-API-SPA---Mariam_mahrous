<script setup lang="ts">
</script>


<template>
  <div class="viewNote-container">
     <h1> {{title}}</h1>
    <p id="view_viewer_number"> viewer number : {{views}} / {{max_views}}</p>
    <p id="view_content"> {{text}}</p>
    </div>
</template>

<style scoped>
.viewNote-container {
  max-width: 400px;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
 }

h1 {
  font-size: 24px;
  margin-bottom: 20px;
  text-align: center;
}




</style>

