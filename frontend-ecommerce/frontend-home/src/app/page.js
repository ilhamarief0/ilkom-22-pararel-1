"use client";

import { useEffect, useState } from "react";
import { API_BASE_URL } from "../../config";
import ProductCard from "../components/ProductCard";
import { useAuth } from "../hooks/useAuth";

export default function ProductList() {
  const [products, setProducts] = useState([]);
  const [error, setError] = useState("");
  const [imageUrls, setImageUrls] = useState({});
  useAuth(); // Check authentication

  useEffect(() => {
    const fetchProducts = async () => {
      const token = localStorage.getItem("jwtToken");
      if (!token) {
        setError("Please log in to view products.");
        return;
      }

      try {
        const res = await fetch(`${API_BASE_URL}/api/product`, {
          method: "GET",
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        });

        if (!res.ok) {
          throw new Error("Failed to fetch products");
        }

        const data = await res.json();

        if (data.success && Array.isArray(data.data)) {
          setProducts(data.data);
        }

        const urls = data.data.reduce((acc, product) => {
          acc[product.id] = `${API_BASE_URL}/api/gambarproduk/${product.image}`;
          return acc;
        }, {});
        setImageUrls(urls);
      } catch (err) {
        console.error(err);
        setError("Error fetching products: " + err.message);
      }
    };

    fetchProducts();
  }, []);

  return (
    <div className="container py-5">
      <h1 className="text-center mb-4">Our Products</h1>
      {error && <div className="alert alert-danger">{error}</div>}
      <div className="product-list">
        {" "}
        {/* Add the flex container class */}
        {products.map((product) => (
          <ProductCard
            key={product.id}
            product={product}
            imageUrl={imageUrls[product.id]}
          />
        ))}
      </div>
    </div>
  );
}
