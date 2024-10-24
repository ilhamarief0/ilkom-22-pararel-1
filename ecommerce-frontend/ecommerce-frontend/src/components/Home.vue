<template>
  <div class="container mx-auto my-8">
    <h1 class="text-2xl font-bold mb-4">Products</h1>

    <!-- Cek jika ada produk yang diambil dari backend -->
    <div v-if="products.length > 0" class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4">
      <div
        v-for="product in products"
        :key="product.id"
        class="border rounded-lg shadow-lg bg-white overflow-hidden p-2"
      >
        <img 
          :src="`http://localhost:3010/api/gambarproduk/${product.image}`" 
          alt="Product Image" 
          class="w-full h-48 object-cover mb-2"
        />
        <div class="p-2">
          <h2 class="text-lg font-semibold text-gray-800">{{ product.title }}</h2>
          <p class="text-gray-600 text-sm mb-2">{{ product.content }}</p>
          <h2 class="text-lg font-semibold text-gray-800">Rp. {{ product.Price }}</h2>
          <button class="bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700 transition">
            Add to Cart
          </button>
        </div>
      </div>
    </div>

    <!-- Tampilkan jika tidak ada produk -->
    <p v-else class="text-gray-500 text-center">No products available.</p>
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

h1 {
  text-align: center;
}

.border {
  border-color: #e2e8f0;  /* Warna border yang lebih lembut */
}
</style>
