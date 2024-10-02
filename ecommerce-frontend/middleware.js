// middleware.js

import { NextResponse } from "next/server";
import jwt from "jsonwebtoken";

const jwtSecret = "pass1234"; // Ganti dengan kunci rahasia Anda

export function middleware(req) {
  const { pathname } = req.nextUrl;
  const token = req.cookies.get("token")?.value; // Ambil token dari cookie

  if (pathname.startsWith("/protected")) {
    // Jika mengakses halaman yang dilindungi
    if (!token) {
      // Jika tidak ada token, redirect ke login
      return NextResponse.redirect(new URL("/login", req.url));
    }

    try {
      // Verifikasi token
      jwt.verify(token, jwtSecret);
    } catch (err) {
      // Jika token tidak valid, redirect ke login
      return NextResponse.redirect(new URL("/login", req.url));
    }
  }

  if (pathname.startsWith("/login") && token) {
    // Jika sudah login dan mencoba mengakses halaman login
    return NextResponse.redirect(new URL("/", req.url)); // Redirect ke homepage
  }

  return NextResponse.next(); // Lanjutkan ke halaman berikutnya
}

// Tentukan routes yang menggunakan middleware ini
export const config = {
  matcher: ["/protected/:path*", "/login"],
};
