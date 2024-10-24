// src/router/index.js
import { createRouter, createWebHistory } from "vue-router";
import Home from "@/components/Home.vue";
import Dashboard from "@/components/Dashboard.vue";
import Login from "@/components/Login.vue";

// Definisikan rute-rute yang tersedia
const routes = [
  { path: "/", name: "HomePage", component: Home },
  { path: "/login", name: "UserLogin", component: Login },
  {
    path: "/dashboard",
    name: "UserDashboard",
    component: Dashboard,
    meta: { requiresAuth: true },
  },
];

// Membuat router menggunakan Vue Router versi 3 (untuk Vue 3)
const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Middleware: Cek token sebelum mengakses route
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("jwt_token");

  // Jika user sudah login dan mencoba akses halaman login
  if (token && to.name === "UserLogin") {
    // Redirect ke halaman dashboard jika sudah login
    next({ name: "UserDashboard" });
  } else if (to.matched.some((record) => record.meta.requiresAuth)) {
    // Jika route membutuhkan otentikasi
    if (token) {
      // Jika ada token, lanjutkan ke route yang diminta
      next();
    } else {
      // Jika tidak ada token, redirect ke halaman login
      next({ name: "UserLogin" });
    }
  } else {
    // Jika route tidak membutuhkan otentikasi, lanjutkan
    next();
  }
});

export default router;
