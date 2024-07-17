<script setup lang="ts">

const GetNotes = async () => {
      try {
        let apiUrl = 'http://localhost:8080/note'
        const response = await fetch(apiUrl, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ email: email.value, password: password.value })
        })

        const data = await response.json()
        if (response.ok) {
          console.log(data)
          router.replace('/note')
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

