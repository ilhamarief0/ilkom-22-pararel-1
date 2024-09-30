import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import UserNavbar from "../components/UserNavbar";
import Swal from "sweetalert2";

export default function ProductsPage() {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true); // For showing a loader while checking login status
  const router = useRouter();

  useEffect(() => {
    const fetchProducts = async () => {
      const token = localStorage.getItem("token");

      // If no token is found, redirect to the login page
      if (!token) {
        router.push("/login");
        return;
      }

      try {
        const response = await fetch("http://localhost:8082/products", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        // If token is invalid, redirect to login
        if (response.status === 401) {
          router.push("/login");
          return;
        }

        const data = await response.json();
        setProducts(data);
      } catch (error) {
        console.error("Error fetching products:", error);
      }

      setLoading(false);
    };

    fetchProducts();
  }, [router]);

  // Function to handle the buy action
  const handleBuyNow = (product) => {
    let quantity = 1; // Initialize with a default quantity of 1

    Swal.fire({
      title: `Order ${product.name}`,
      html: `
        <label for="quantity">Quantity:</label>
        <input type="number" id="quantity" min="1" value="1" class="swal2-input">
        <p>Total Price: Rp. <span id="total-price">${product.price}</span></p>
      `,
      focusConfirm: false,
      showCancelButton: true,
      confirmButtonText: "Order",
      didOpen: () => {
        const quantityInput = Swal.getPopup().querySelector("#quantity");
        const totalPriceElement = Swal.getPopup().querySelector("#total-price");

        // Update the total price when quantity changes
        quantityInput.addEventListener("input", (e) => {
          quantity = parseInt(e.target.value, 10) || 1; // Convert to number using parseInt
          const totalPrice = quantity * product.price;
          totalPriceElement.textContent = totalPrice;
        });
      },
      preConfirm: () => {
        const quantityInput = Swal.getPopup().querySelector("#quantity").value;
        if (!quantityInput || quantityInput <= 0) {
          Swal.showValidationMessage(`Please enter a valid quantity`);
        }
        return { quantity: parseInt(quantityInput, 10) }; // Convert to number
      },
    }).then(async (result) => {
      if (result.isConfirmed) {
        const token = localStorage.getItem("token");
        const userId = 1; // Replace this with actual user ID from localStorage or any other source

        const orderData = {
          productId: product.id,
          productName: product.name,
          userId: userId, // Ensure this field is present
          quantity: result.value.quantity, // Now it will be a number
          totalPrice: product.price * result.value.quantity, // Ensure this field is a number
        };

        console.log("Sending orderData:", orderData); // Log the orderData for debugging

        try {
          const response = await fetch("http://localhost:8085/payments", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify(orderData),
          });

          if (response.ok) {
            const paymentResponse = await response.json();
            Swal.fire(`Ordered ${result.value.quantity} of ${product.name}!`);
          } else {
            const errorData = await response.json();
            Swal.fire(`Error: ${errorData.message || "Payment failed!"}`);
          }
        } catch (error) {
          console.error("Error processing payment:", error);
          Swal.fire("Error processing payment. Please try again.");
        }
      }
    });
  };

  if (loading) {
    // Optionally, show a loading screen while checking token
    return <div>Loading...</div>;
  }

  return (
    <div className="p-6">
      <UserNavbar />
      <h1 className="text-xl font-bold mb-4">Products</h1>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        {products.map((product) => (
          <div
            key={product.id}
            className="border p-4 shadow-md rounded-lg bg-white"
          >
            <h2 className="text-lg font-bold mb-2">{product.name}</h2>
            <p className="mb-2 text-gray-700">Harga: Rp. {product.price}</p>
            <button
              onClick={() => handleBuyNow(product)}
              className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded"
            >
              Beli Produk
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}
