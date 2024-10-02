// utils/auth.js
import Cookies from "js-cookie";
import bcrypt from "bcryptjs";

export async function login(email, password) {
  try {
    // Fetch semua user dari backend
    const response = await fetch("http://localhost:8081/users");
    if (response.ok) {
      const users = await response.json();

      const user = users.find((user) => user.Email === email);
      if (!user) {
        throw new Error("Invalid Email or password");
      }

      const passwordMatch = await bcrypt.compare(password, user.PasswordHash);
      if (!passwordMatch) {
        throw new Error("Invalid username or password");
      }

      // ... (kode untuk generate JWT dan menyimpan data user - opsional)

      return true;
    } else {
      const errorData = await response.json();
      throw new Error(errorData.message || "Login failed");
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}
