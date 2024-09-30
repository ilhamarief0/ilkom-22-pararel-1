const API_URL = "http://localhost:8084"; // Shipping Service URL

export const createShipment = async (shipmentData) => {
  const response = await fetch(`${API_URL}/shipments`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(shipmentData),
  });
  return response.json();
};
