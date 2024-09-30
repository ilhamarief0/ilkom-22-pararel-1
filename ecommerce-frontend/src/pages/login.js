import { useState, useEffect } from "react";
import { useRouter } from "next/router";
import Swal from "sweetalert2"; // Import SweetAlert2

export default function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const router = useRouter();

  // Check if the user is already logged in
  useEffect(() => {
    const token = localStorage.getItem("token");
    const userId = localStorage.getItem("token");
    if (token) {
      // Redirect to appropriate page if already logged in
      const role = localStorage.getItem("role");
      if (role === "admin") {
        router.push("/admin/dashboard");
      } else {
        router.push("/products");
      }
    }
  }, [router]);

  const handleLogin = async (e) => {
    e.preventDefault();

    const response = await fetch("http://localhost:8081/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    });

    const data = await response.json();

    if (response.ok) {
      localStorage.setItem("token", data.token);
      localStorage.setItem("role", data.role);

      // Show success message using SweetAlert2
      Swal.fire({
        title: "Success!",
        text: "Login successful!",
        icon: "success",
        confirmButtonText: "OK",
      }).then(() => {
        // Redirect based on user role
        if (data.role === "admin") {
          router.push("/admin/dashboard");
        } else {
          router.push("/products");
        }
      });
    } else {
      setError(data.error || "Login failed");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <form className="bg-white p-6 rounded shadow-md" onSubmit={handleLogin}>
        <h1 className="text-xl font-bold mb-4">Login</h1>
        {error && <p className="text-red-500">{error}</p>}
        <input
          type="text"
          placeholder="Username"
          className="border p-2 mb-4 w-full"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Password"
          className="border p-2 mb-4 w-full"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button type="submit" className="bg-blue-500 text-white p-2 rounded">
          Login
        </button>
      </form>
    </div>
  );
}
