// src/app/components/ProductCard.js

import React from "react";
import Link from "next/link";

const ProductCard = ({ product, imageUrl }) => {
  return (
    <div className="card">
      {imageUrl ? (
        <img
          src={imageUrl}
          className="card-img-top"
          alt={product.title}
          style={{ objectFit: "cover", height: "150px" }} // Added styles for better image display
        />
      ) : (
        <div className="placeholder-image">Image not available</div>
      )}
      <div className="card-body">
        <h5 className="card-title">{product.title}</h5>
        <p className="card-text">Price: ${product.Price}</p>
        <p className="card-text">Stock: {product.Stock}</p>
        <Link href={`/products/${product.id}`}>
          <button className="btn btn-primary w-100">View Details</button>
        </Link>
      </div>
    </div>
  );
};

export default ProductCard;
