const API_URL = "http://localhost:8082"; // Product Service URL

export const fetchProducts = async () => {
  const response = await fetch(`${API_URL}/products`);
  return response.json();
};

export const createProduct = async (productData) => {
  const response = await fetch(`${API_URL}/products`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(productData),
  });
  return response.json();
};
