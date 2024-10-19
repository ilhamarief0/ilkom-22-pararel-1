<template>
    <div>
      <h2>Pembayaran</h2>
      <div>
        <label for="order_id">Order ID:</label>
        <input type="text" v-model="orderID" id="order_id" />
      </div>
      <div>
        <label for="amount">Jumlah (IDR):</label>
        <input type="number" v-model="amount" id="amount" />
      </div>
      <button @click="payNow">Bayar Sekarang</button>
    </div>
  </template>
  
  <script>
  export default {
    name: "PaymentForm",
    data() {
      return {
        orderID: "ORDER-123",
        amount: 100000,
      };
    },
    methods: {
      async payNow() {
        try {
          // Pastikan Snap.js sudah tersedia
          if (typeof window.snap === 'undefined') {
            alert("Midtrans Snap.js belum dimuat, mohon tunggu beberapa saat dan coba lagi.");
            return;
          }
  
          const response = await fetch('http://localhost:8080/payment', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              order_id: this.orderID,
              gross_amount: this.amount,
            }),
          });
  
          const data = await response.json();
  
          if (data.redirect_url) {
            console.log("Redirect URL dari backend:", data.redirect_url);
  
            // Extract the token from the redirect_url
            const token = data.redirect_url.split('/').pop();
  
            // Gunakan Snap.js untuk membuka popup pembayaran Midtrans dengan token
            window.snap.pay(token, {
              onSuccess: function(result) {
                alert("Pembayaran Berhasil: " + JSON.stringify(result));
              },
              onPending: function(result) {
                alert("Menunggu Pembayaran: " + JSON.stringify(result));
              },
              onError: function(result) {
                alert("Pembayaran Gagal: " + JSON.stringify(result));
              },
              onClose: function() {
                alert("Pembayaran dibatalkan oleh pengguna.");
              },
            });
          } else {
            console.error("Invalid redirect URL:", data);
          }
  
        } catch (error) {
          console.error("Error saat menghubungi server:", error);
        }
      },
    },
    mounted() {
      const midtransScriptUrl = "https://app.sandbox.midtrans.com/snap/snap.js";
      const midtransClientKey = "SB-Mid-client-etaFzjV97U45y5La"; // Pastikan Client Key benar
  
      let scriptTag = document.createElement('script');
      scriptTag.src = midtransScriptUrl;
      scriptTag.setAttribute('data-client-key', midtransClientKey);
  
      document.body.appendChild(scriptTag);
    },
  };
  </script>
  
  <style scoped>
  input {
    display: block;
    margin-bottom: 10px;
  }
  
  button {
    padding: 10px 20px;
    background-color: #28a745;
    color: white;
    border: none;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #218838;
  }
  </style>