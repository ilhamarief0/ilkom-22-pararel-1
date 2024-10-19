<template>
  <button @click="payNow">Bayar Sekarang</button>
</template>

<script>
export default {
  methods: {
    async payNow() {
      const response = await fetch('http://localhost:8080/payment', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          order_id: 'ORDER-123',
          gross_amount: 100000
        }),
      });
      const data = await response.json();
      if (data.redirect_url) {
        window.snap.pay(data.redirect_url);
      }
    },
  },
};
</script>

<script src="https://app.sandbox.midtrans.com/snap/snap.js" data-client-key="SB-Mid-client-etaFzjV97U45y5La"></script>
