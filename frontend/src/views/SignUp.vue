<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
let signUp = ref(true)
let email = ref('')
let password = ref('')
const API_URL = import.meta.env.VITE_API_URL


const handleClick = () => {
  signUp.value = !signUp.value
}

const handleSubmit = async () => {
  try {
    let endpoint = signUp.value ? '/signup' : '/login'
    let apiUrl = API_URL + endpoint
    console.log(apiUrl)
    const response = await fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ email: email.value, password: password.value })
    })

    const data = await response.json()
    if (response.ok) {
      localStorage.setItem('token', data.token)
      router.replace('/home')
    } else {
      alert(data.error)
    }
  } catch (error) {
    console.error('Error: ', error)
  }
}
</script>

<template>
  <div>
    <h1>{{ signUp ? 'Sign up' : 'Log in' }}</h1>
    <form @submit.prevent="handleSubmit">
      <label for="Email">Email:</label> <br />
      <br />
      <input type="email" placeholder="Enter Email" v-model="email" />

      <label for="Password">Password:</label> <br /><br />
      <input type="password" placeholder="Enter Password" minlength="8" v-model="password" />
      <p>
        {{ signUp ? 'Already have an account?' : 'Create an account' }}
        <span @click="handleClick">{{ signUp ? 'Log in' : 'Sign up' }}</span>
      </p>
      <button type="submit">{{ signUp ? 'Sign Up' : 'Log in' }}</button>
    </form>
  </div>
</template>

<style scoped>
div {
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

span {
  color: #4caf50;
}

p,
span {
  font-size: medium;
}
</style>
