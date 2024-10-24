<template>
  <div class="min-h-screen flex items-center justify-center bg-orange-100">
    <div class="w-full max-w-xs p-8">
      <h2 class="text-4xl font-bold text-center mb-6 text-orange-500">Login</h2>
      <p class="text-center mb-4 text-gray-600">Welcome back! Please login to your account.</p>
      <form @submit.prevent="login" class="space-y-4">
        <div>
          <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
          <input
            v-model="username"
            type="text"
            id="username"
            required
            class="mt-1 p-2 w-full border border-gray-300 focus:outline-none focus:ring-2 focus:ring-orange-500"
          />
        </div>
        <div>
          <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
          <input
            v-model="password"
            type="password"
            id="password"
            required
            class="mt-1 p-2 w-full border border-gray-300 focus:outline-none focus:ring-2 focus:ring-orange-500"
          />
        </div>
        <button
          type="submit"
          class="w-full bg-orange-500 text-white font-bold py-2 px-4 hover:bg-orange-600 transition duration-200"
        >
          Login
        </button>
      </form>
      <p v-if="errorMessage" class="text-red-500 mt-4">{{ errorMessage }}</p>
      <div class="flex justify-between text-sm mt-4">
        <router-link to="/forgot-password" class="text-orange-500 hover:underline">Forgot Password?</router-link>
        <router-link to="/register" class="text-orange-500 hover:underline">Create an Account</router-link>
      </div>
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
        // Send a request to the backend /login endpoint
        const response = await axios.post('http://localhost:3000/api/auth/login', {
          username: this.username,
          password: this.password,
        });

        // On success, store JWT token in localStorage
        const token = response.data.token;
        localStorage.setItem('jwt_token', token);

        // Redirect to the dashboard page after login
        this.$router.push({ name: 'UserDashboard' });
      } catch (error) {
        this.errorMessage = 'Invalid username or password.';
      }
    },
  },
};
</script>

<style scoped>
/* Optional additional styling (if needed) */
body {
  margin: 0; /* Reset default margins */
}
</style>
