// pages/products.js

import { useEffect, useState } from "react";

export default function ProductsPage() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const fetchProducts = async () => {
      const token = localStorage.getItem("token");
      const response = await fetch("http://localhost:8082/products", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      const data = await response.json();
      setProducts(data);
    };

    fetchProducts();
  }, []);

  return (
    <div className="p-6">
      <h1 className="text-xl font-bold mb-4">Products</h1>
      <ul>
        {products.map((product) => (
          <li key={product.id} className="border p-2 mb-2">
            {product.name} - ${product.price}
          </li>
        ))}
      </ul>
    </div>
  );
}
