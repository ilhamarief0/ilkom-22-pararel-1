const API_URL = "http://localhost:8083"; // Payment Service URL

export const processPayment = async (paymentData) => {
  const response = await fetch(`${API_URL}/payments`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(paymentData),
  });
  return response.json();
};
