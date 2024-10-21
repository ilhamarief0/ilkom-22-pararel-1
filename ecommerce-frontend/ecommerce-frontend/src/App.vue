<template>
  <div id="app" class="min-h-screen flex flex-col items-center justify-center bg-gray-100">
    <!-- Navigasi -->
    <nav class="bg-white shadow-lg p-4 w-full flex justify-center space-x-4">
      <router-link to="/" class="text-gray-800 hover:text-blue-600">Home</router-link>
      <router-link v-if="token" to="/dashboard" class="text-gray-800 hover:text-blue-600">Dashboard</router-link>
      <router-link v-if="!token" to="/login" class="text-gray-800 hover:text-blue-600">Login</router-link>
      <!-- <router-link to="/payment" class="text-gray-800 hover:text-blue-600">Payment</router-link> -->
    </nav>

    <!-- Konten -->
    <div class="container mx-auto my-10 p-8 bg-white shadow-md rounded-lg">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
      token: null, // Buat token sebagai bagian dari state
    };
  },
  created() {
    // Ambil token dari localStorage ketika komponen diinisialisasi
    this.token = localStorage.getItem("jwt_token");
  },
};
</script>

<style scoped>
/* Kamu bisa menambahkan styling tambahan di sini jika diperlukan */
</style>
