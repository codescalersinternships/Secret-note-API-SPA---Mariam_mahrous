<script setup lang="ts">
    import { ref } from 'vue';
    let signUp = ref(true);
    let email = ref('');
    let password = ref('');

    const handleClick = () => {
        signUp.value=!signUp.value
    };

</script>


<template>
  <button type="submit" id="button" @click="handleClick">{{signUp ? "Log in" :"Sign up"}}</button>
  <div class="signup-container">

    <h1>{{signUp ? "Sign up":"Log in"}}</h1>
    <form @submit.prevent="signUp">
        <label for="Email">Email:</label> <br> <br>
        <input type="email" placeholder="Enter Email" v-model="email">

        <label for="Password">Password:</label> <br><br>
        <input type="password" placeholder="Enter Password" minlength="8" v-model="password"> 
        <button type="submit">{{signUp ? "Sign Up":"Log in"}}</button>
      </form>
    </div>
</template>

<style scoped>
.signup-container {
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

#button {
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
  width: fit-content;
  padding: 10px;
  border-radius: 5px;
  justify-content: flex-end;
}


</style>

