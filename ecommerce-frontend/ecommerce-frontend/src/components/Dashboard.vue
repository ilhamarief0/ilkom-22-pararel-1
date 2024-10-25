<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-100">
      <div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-lg text-center">
        <h2 class="text-2xl font-bold text-gray-800 mb-4">Welcome to the Dashboard</h2>
        <p class="text-gray-600">You are logged in!</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "UserDashboard",
    created() {
      // Kamu bisa menggunakan token di localStorage untuk mendapatkan data dari backend
    //   const token = localStorage.getItem("jwt_token");
  
      // Misalnya, gunakan token untuk mengambil data pengguna
      // axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
      // axios.get('http://localhost:3000/api/user')
    },
  };
  </script>
  
  <style scoped>
  /* Tambahkan styling tambahan jika diperlukan */
  </style>
  