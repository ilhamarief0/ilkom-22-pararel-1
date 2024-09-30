// pages/admin/dashboard.js
import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import ProductForm from "../../components/ProductForm";
import ProductTable from "../../components/ProductTable";
import AdminNavbar from "../../components/AdminNavbar"; // Import the navbar component

const Dashboard = () => {
  const router = useRouter();
  const [products, setProducts] = useState([]);
  const [editProduct, setEditProduct] = useState(null);

  // Check user role from localStorage
  useEffect(() => {
    const role = localStorage.getItem("role");

    // If user is not admin, redirect
    if (role !== "admin") {
      router.push("/products"); // Redirect to products or another page
    }
  }, [router]);

  const fetchProducts = async () => {
    const response = await fetch("http://localhost:8082/products");
    if (!response.ok) {
      console.error("Failed to fetch products");
      return;
    }
    const data = await response.json();
    setProducts(data);
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  return (
    <div className="min-h-screen bg-gray-100 p-8">
      <AdminNavbar /> {/* Include the navbar here */}
      <h1 className="text-3xl font-bold mb-6">Admin Dashboard</h1>
      <ProductForm fetchProducts={fetchProducts} editProduct={editProduct} />
      <ProductTable
        products={products}
        fetchProducts={fetchProducts}
        setEditProduct={setEditProduct}
      />
    </div>
  );
};

export default Dashboard;
