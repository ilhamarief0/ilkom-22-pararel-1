"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { API_BASE_URL } from "../../../config";

export default function ProductList() {
  const [products, setProducts] = useState([]);
  const [error, setError] = useState("");
  const [imageUrls, setImageUrls] = useState({});

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
        } else if (data.success) {
          setProducts([data.data]);
        } else {
          throw new Error("Unexpected response format");
        }

        fetchImages(data.data, token);
      } catch (err) {
        console.error(err);
        setError("Error fetching products: " + err.message);
      }
    };


    const fetchImages = async (products, token) => {
      const urls = {};
      for (let product of products) {
        try {
          const imageRes = await fetch(
            `${API_BASE_URL}/api/gambarproduk/${product.image}`,
            {
              method: "GET",
              headers: {
                Authorization: `Bearer ${token}`,
                "Content-Type": "application/json",
              },
            }
          );

          if (imageRes.ok) {
            urls[
              product.id
            ] = `${API_BASE_URL}/api/gambarproduk/${product.image}`;
          }
        } catch (err) {
          console.error(err);
        }
      }
      setImageUrls(urls);
    };

    fetchProducts();
  }, []);

  return (
    <div className="container py-5">
      <h1 className="text-center mb-4">Our Products</h1>
      {error && <div className="alert alert-danger">{error}</div>}
      <div className="row">
        {products.map((product) => (
          <div key={product.id} className="col-md-3 mb-4">
            <div className="card h-100">
              {imageUrls[product.id] ? (
                <img
                  src={imageUrls[product.id]}
                  className="card-img-top"
                  alt={product.title}
                />
              ) : (
                <div className="placeholder-image">Image not available</div>
              )}

              <div className="card-body">
                <h5 className="card-title">{product.title}</h5>
                <p className="card-text">Price: ${product.Price}</p>
                <p className="card-text">Stock: {product.Stock}</p>
                <Link href={`/products/${product.id}`}>
                  <button className="btn btn-primary w-100">
                    View Details
                  </button>
                </Link>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
