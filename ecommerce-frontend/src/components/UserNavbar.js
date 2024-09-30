// components/UserNavbar.js
import Link from "next/link";
import { useRouter } from "next/router";
import Swal from "sweetalert2"; // Import SweetAlert2
import "sweetalert2/dist/sweetalert2.min.css"; // Use this instead of .scss

const UserNavbar = () => {
  const router = useRouter();

  const handleLogout = () => {
    // Show SweetAlert2 confirmation
    Swal.fire({
      title: "Are you sure?",
      text: "Do you really want to logout?",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Yes, logout!",
    }).then((result) => {
      if (result.isConfirmed) {
        // If confirmed, remove the token and redirect to login page
        localStorage.removeItem("token");
        localStorage.removeItem("role");

        Swal.fire(
          "Logged out!",
          "You have been successfully logged out.",
          "success"
        );

        router.push("/login");
      }
    });
  };

  return (
    <nav className="bg-gray-800 p-4">
      <ul className="flex space-x-4">
        <li>
          <Link
            href="/admin/dashboard"
            className="text-white hover:text-gray-300"
          >
            Dashboard
          </Link>
        </li>
        <li>
          <Link
            href="/admin/products"
            className="text-white hover:text-gray-300"
          >
            Products
          </Link>
        </li>
        <li>
          <Link href="/orders" className="text-white hover:text-gray-300">
            Orders
          </Link>
        </li>
        <li>
          <Link href="/admin/users" className="text-white hover:text-gray-300">
            Users
          </Link>
        </li>
        <li>
          <Link
            href="/admin/settings"
            className="text-white hover:text-gray-300"
          >
            Settings
          </Link>
        </li>
        {/* Logout */}
        <li>
          <button
            onClick={handleLogout}
            className="text-white hover:text-gray-300"
          >
            Logout
          </button>
        </li>
      </ul>
    </nav>
  );
};

export default UserNavbar;
