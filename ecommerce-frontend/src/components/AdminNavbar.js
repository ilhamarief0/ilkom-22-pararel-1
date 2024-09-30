// components/AdminNavbar.js
import Link from "next/link";

const AdminNavbar = () => {
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
          <Link href="/admin/orders" className="text-white hover:text-gray-300">
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
      </ul>
    </nav>
  );
};

export default AdminNavbar;
