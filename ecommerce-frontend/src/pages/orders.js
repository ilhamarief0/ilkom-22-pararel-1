import { useEffect, useState } from "react";
import UserNavbar from "../components/UserNavbar"; // Import UserNavbar
import Swal from "sweetalert2"; // Import SweetAlert2

const OrdersPage = () => {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    const fetchOrders = async () => {
      try {
        const response = await fetch("http://localhost:8083/orders", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("token")}`, // Ensure correct template string syntax
          },
        });

        if (!response.ok) {
          throw new Error("Failed to fetch orders");
        }

        const data = await response.json();
        setOrders(data); // Ensure data includes product names
      } catch (error) {
        Swal.fire("Error!", error.message, "error");
      }
    };

    fetchOrders();
  }, []);

  return (
    <div>
      <UserNavbar />
      <div className="container mx-auto p-4">
        <h1 className="text-2xl font-bold mb-4">Orders</h1>
        <table className="min-w-full bg-white border border-gray-300">
          <thead>
            <tr className="bg-gray-200">
              <th className="border border-gray-300 p-2">Order ID</th>
              <th className="border border-gray-300 p-2">User ID</th>
              <th className="border border-gray-300 p-2">Product Name</th>{" "}
              <th className="border border-gray-300 p-2">Quantity</th>
              <th className="border border-gray-300 p-2">Total Price</th>
              <th className="border border-gray-300 p-2">Order Date</th>
              <th className="border border-gray-300 p-2">Status</th>
            </tr>
          </thead>
          <tbody>
            {orders.length > 0 ? (
              orders.map((order) => (
                <tr key={order.id}>
                  <td className="border border-gray-300 p-2">{order.id}</td>
                  <td className="border border-gray-300 p-2">
                    {order.user_id}
                  </td>
                  <td className="border border-gray-300 p-2">
                    {order.product_id} {/* Accessing product name */}
                  </td>
                  <td className="border border-gray-300 p-2">
                    {order.quantity}
                  </td>
                  <td className="border border-gray-300 p-2">
                    {order.total_price}
                  </td>
                  <td className="border border-gray-300 p-2">
                    {new Date(order.order_date).toLocaleString()}
                  </td>
                  <td className="border border-gray-300 p-2">{order.status}</td>
                </tr>
              ))
            ) : (
              <tr>
                <td colSpan="7" className="text-center p-2">
                  No orders found
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default OrdersPage;
