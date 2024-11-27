import Link from "next/link";
import { FaShoppingCart, FaUser } from "react-icons/fa";

export default function Navbar() {
  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light">
      <div className="container">
        <Link href="/" className="navbar-brand">
          E-Commerce
        </Link>
        <div className="collapse navbar-collapse">
          <ul className="navbar-nav ms-auto">
            <li className="nav-item">
              <Link href="/cart" className="nav-link">
                <FaShoppingCart /> Cart
              </Link>
            </li>
            <li className="nav-item">
              <Link href="/login" className="nav-link">
                <FaUser /> Login
              </Link>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
}
