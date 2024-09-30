import { NextResponse } from "next/server";

export async function middleware(req) {
  const { user } = req.cookies; // Assuming you store user data in cookies

  // Parse the user cookie to get user information
  let parsedUser = null;
  try {
    parsedUser = user ? JSON.parse(user) : null;
  } catch (err) {
    console.error("Failed to parse user cookie", err);
  }

  const isAdmin = parsedUser?.role === "admin"; // Check if the user is admin

  // Define protected paths
  const adminPaths = ["/admin/dashboard"];
  const userPaths = ["/products"];

  const pathname = req.nextUrl.pathname;

  // Check if the request is for an admin path
  if (adminPaths.includes(pathname)) {
    if (!isAdmin) {
      // Redirect to a forbidden page if the user is not an admin
      return NextResponse.redirect(new URL("/403", req.url)); // Redirect to 403 page
    }
  }

  // Check if the request is for a user path
  if (userPaths.includes(pathname)) {
    if (isAdmin) {
      // If the user is an admin, you can redirect them or allow access as desired
      return NextResponse.redirect(new URL("/admin/dashboard", req.url)); // Redirect admin users to admin dashboard
    }
  }

  // Allow the request to proceed for other routes
  return NextResponse.next();
}

// Specify the paths the middleware should apply to
export const config = {
  matcher: ["/admin/:path*", "/products"], // Apply to admin routes and products
};
