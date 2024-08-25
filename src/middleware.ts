import { cookies } from "next/headers"
import { NextRequest, NextResponse } from "next/server"
import { getToken } from "next-auth/jwt"
import { withAuth } from "next-auth/middleware"

import { absoluteUrl } from "./lib/utils"

export default async function middleware(req: NextRequest) {
  const token = cookies().get("user")
  const isAuth = !!token
  const isAuthPage = req.nextUrl.pathname.startsWith("/login")
  if (isAuthPage) {
    if (isAuth) {
      return NextResponse.redirect(new URL("/dashboard", req.url))
    }

    return null
  }

  if (!isAuth) {
    let from = req.nextUrl.pathname
    if (req.nextUrl.search) {
      from += req.nextUrl.search
    }

    return NextResponse.redirect(
      new URL(`/login?from=${encodeURIComponent(from)}`, req.url)
    )
  }

  fetch(absoluteUrl("/api/admin/views"), {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      route: req.nextUrl.pathname,
      email: (JSON.parse(token.value) as { email: string }).email,
      device: req.headers.get("user-agent") || "unknown",
    }),
  }).catch((e) => console.log(e))
}

export const config = {
  matcher: ["/dashboard/:path*", "/admin/:path*"],
}
