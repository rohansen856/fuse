import { cookies } from "next/headers"
import { NextResponse } from "next/server"

import { absoluteUrl } from "@/lib/utils"

export async function GET(req: Request) {
  try {
    cookies().delete("user")
    return NextResponse.redirect(absoluteUrl("/"))
  } catch (error) {
    console.error("Error during logout:", error)
    return NextResponse.redirect(absoluteUrl("/login"))
  }
}
