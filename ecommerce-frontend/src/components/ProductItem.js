import React from "react";
import axios from "axios";

const ProductItem = ({ product, onDelete, onEdit }) => {
  const handleDelete = async () => {
    await onDelete(product.id);
  };

  return (
    <li>
      <strong>{product.name}</strong> - {product.description} - ${product.price}{" "}
      - Stock: {product.stock}
      <button onClick={() => onEdit(product)}>Edit</button>
      <button onClick={handleDelete}>Delete</button>
    </li>
  );
};

export default ProductItem;
