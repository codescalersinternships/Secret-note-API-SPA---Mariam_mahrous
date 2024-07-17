<script setup lang="ts">
</script>


<template>
     <h1 id="created_message">Your Note Has been created</h1>
    <div id="note_link">
        <a href="{{ link }}" >{{ link }}</a>
    </div>
</template>

<style scoped>

h1 {
  font-size: 24px;
  margin-bottom: 20px;
  text-align: center;
}

#note_link  {
    text-align: center;
    padding-top: 5vh;
    font-size: 20px;
}




</style>

