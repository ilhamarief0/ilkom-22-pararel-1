// src/pages/login.js

import { useState, useEffect } from "react";
import { useRouter } from "next/router"; // Pastikan ini diimpor
import Cookies from "js-cookie";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter(); // Inisialisasi router

  useEffect(() => {
    // Cek jika token sudah ada
    const token = Cookies.get("token");
    if (token) {
      // Jika sudah login, redirect ke halaman yang diinginkan
      router.push("/protected"); // Ubah sesuai kebutuhan
    }
  }, [router]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const res = await fetch("http://localhost:8082/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    const data = await res.json();

    if (res.ok) {
      // Simpan token di cookie
      Cookies.set("token", data.token);
      router.push("/protected"); // Redirect ke halaman protected setelah login
    } else {
      // Tangani kesalahan
      alert(data.message);
    }
  };

  return (
    <div className="flex items-center justify-center h-screen">
      <form onSubmit={handleSubmit} className="bg-white p-6 rounded shadow-md">
        <h2 className="mb-4 text-lg font-bold">Login</h2>
        <div className="mb-4">
          <label className="block text-sm font-medium mb-2">Email</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full px-3 py-2 border rounded"
            required
          />
        </div>
        <div className="mb-4">
          <label className="block text-sm font-medium mb-2">Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full px-3 py-2 border rounded"
            required
          />
        </div>
        <button
          type="submit"
          className="w-full bg-blue-500 text-white py-2 rounded"
        >
          Login
        </button>
      </form>
    </div>
  );
}
