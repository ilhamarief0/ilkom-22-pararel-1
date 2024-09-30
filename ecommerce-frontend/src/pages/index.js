// src/pages/index.js
import Navbar from "../components/Navbar";

export default function Home() {
  return (
    <div>
      <Navbar />
      <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
        <h1 className="text-4xl font-bold text-blue-600">
          Welcome to E-Commerce Platform
        </h1>
        <p className="mt-4 text-gray-700">
          Your one-stop solution for online shopping.
        </p>
      </div>
    </div>
  );
}
