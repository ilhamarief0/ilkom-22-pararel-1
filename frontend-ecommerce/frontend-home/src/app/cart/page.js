"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { API_BASE_URL } from "../../../config";
import { useAuth } from "../../hooks/useAuth";

export default function Cart() {
  const [cartItems, setCartItems] = useState([]);
  const router = useRouter();
  useAuth();

  useEffect(() => {
    const fetchCartItems = async () => {
      const token = localStorage.getItem("jwtToken");

      if (!token) {
        router.push("/login");
        return;
      }

      try {
        const res = await fetch(`${API_BASE_URL}/api/cart`, {
          method: "GET",
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        });

        if (!res.ok) {
          throw new Error("Failed to fetch cart items");
        }

        const data = await res.json();
        setCartItems(data.items); // Assuming the structure contains an 'items' array
      } catch (error) {
        console.error(error);
      }
    };

    fetchCartItems();
  }, [router]);

  return (
    <div className="container py-5">
      <h1 className="text-center">Shopping Cart</h1>
      {cartItems.length === 0 ? (
        <p className="text-center">Your cart is empty.</p>
      ) : (
        <ul className="list-group">
          {cartItems.map((item) => (
            <li key={item.id} className="list-group-item">
              {item.title} - ${item.price}
              <span className="badge bg-secondary float-end">
                {item.quantity}
              </span>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
