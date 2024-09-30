const ProductTable = ({ products, fetchProducts, setEditProduct }) => {
  const handleDelete = async (id) => {
    await fetch(`http://localhost:8082/products/${id}`, {
      method: "DELETE",
    });
    fetchProducts();
  };

  return (
    <table className="min-w-full bg-white border border-gray-300 rounded-lg">
      <thead>
        <tr className="bg-gray-200">
          <th className="p-4 border-b">ID</th>
          <th className="p-4 border-b">Name</th>
          <th className="p-4 border-b">Description</th>
          <th className="p-4 border-b">Price</th>
          <th className="p-4 border-b">Stock</th>
          <th className="p-4 border-b">Actions</th>
        </tr>
      </thead>
      <tbody>
        {products.map((product) => (
          <tr key={product.id}>
            <td className="p-4 border-b">{product.id}</td>
            <td className="p-4 border-b">{product.name}</td>
            <td className="p-4 border-b">{product.description}</td>
            <td className="p-4 border-b">{product.price}</td>
            <td className="p-4 border-b">{product.stock}</td>
            <td className="p-4 border-b">
              <button
                className="text-blue-500 hover:underline"
                onClick={() => setEditProduct(product)}
              >
                Edit
              </button>
              <button
                className="text-red-500 hover:underline ml-2"
                onClick={() => handleDelete(product.id)}
              >
                Delete
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default ProductTable;
