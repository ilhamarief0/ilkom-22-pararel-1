import { createApp } from "vue";
import App from "./App.vue";
import router from "./router"; // Import router
import "./assets/tailwind.css";

const app = createApp(App);

app.use(router); // Pastikan router digunakan dengan benar
app.mount("#app");
