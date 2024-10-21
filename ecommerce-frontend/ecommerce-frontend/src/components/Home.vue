<template>
    <div class="container mx-auto my-8">
      <h1 class="text-2xl font-bold mb-4">Products</h1>
  
      <!-- Cek jika ada produk yang diambil dari backend -->
      <div v-if="products.length > 0" class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        <div
          v-for="product in products"
          :key="product.id"
          class="border p-4 rounded-lg shadow-md bg-white"
        >
          <img :src="'http://localhost:3000/gambarproduk/' + product.image" alt="Product Image" class="w-full h-48 object-cover mb-4 rounded-lg" />
          <h2 class="text-xl font-bold">{{ product.title }}</h2>
          <p class="text-gray-600">{{ product.content }}</p>
        </div>
      </div>
  
      <!-- Tampilkan jika tidak ada produk -->
      <p v-else class="text-gray-500">No products available.</p>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    name: "HomePage",
    data() {
      return {
        products: [], // State untuk menyimpan daftar produk
        errorMessage: '', // Jika ada error
      };
    },
    created() {
      // Ambil data produk dari backend Go saat komponen diinisialisasi
      this.fetchProducts();
    },
    methods: {
      async fetchProducts() {
        const token = localStorage.getItem("jwt_token"); // Ambil token dari localStorage
        if (token) {
          try {
            // Atur header Authorization dengan token JWT
            axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  
            const response = await axios.get('http://localhost:3000/api/product');
            if (response.data.success) {
              this.products = response.data.data; // Menyimpan daftar produk yang diterima dari backend
            } else {
              this.errorMessage = 'No products found';
            }
          } catch (error) {
            this.errorMessage = 'Failed to fetch products';
          }
        } else {
          this.errorMessage = 'No token found. Please login.';
        }
      },
    },
  };
  </script>
  
  <style scoped>
  /* Tambahkan styling sesuai kebutuhan */
  .container {
    max-width: 1200px;
  }
  </style>
  