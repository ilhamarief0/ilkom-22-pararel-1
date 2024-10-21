<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-100">
      <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
        <h2 class="text-2xl font-bold text-center mb-6">Login</h2>
        <form @submit.prevent="login" class="space-y-4">
          <div>
            <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
            <input
              v-model="username"
              type="text"
              id="username"
              required
              class="mt-1 p-2 w-full border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <input
              v-model="password"
              type="password"
              id="password"
              required
              class="mt-1 p-2 w-full border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <button
            type="submit"
            class="w-full bg-blue-500 text-white font-bold py-2 px-4 rounded-lg hover:bg-blue-600 transition duration-200"
          >
            Login
          </button>
        </form>
        <p v-if="errorMessage" class="text-red-500 mt-4">{{ errorMessage }}</p>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    name: "UserLogin",
    data() {
      return {
        username: '',
        password: '',
        errorMessage: '',
      };
    },
    methods: {
      async login() {
        try {
          // Mengirimkan request ke /login di backend
          const response = await axios.post('http://localhost:8000/login', {
            username: this.username,
            password: this.password,
          });
  
          // Jika berhasil, simpan token JWT di localStorage
          const token = response.data.token;
          localStorage.setItem('jwt_token', token);
  
          // Redirect ke halaman setelah login (misalnya dashboard)
          this.$router.push({ name: 'UserDashboard' });
        } catch (error) {
          this.errorMessage = 'Invalid username or password.';
        }
      },
    },
  };
  </script>
  
  <style scoped>
  /* Styling tambahan jika diperlukan */
  </style>
  