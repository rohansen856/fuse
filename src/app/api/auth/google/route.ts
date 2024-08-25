import { NextResponse } from "next/server"

import { absoluteUrl } from "@/lib/utils"

export async function GET() {
  const redirectUri = absoluteUrl("/api/auth/callback") // Your redirect URI
  const clientId = process.env.GOOGLE_CLIENT_ID // Set in .env.local
  const scope = "profile email" // Scopes you need

  const googleAuthUrl = `https://accounts.google.com/o/oauth2/v2/auth?response_type=code&client_id=${clientId}&redirect_uri=${redirectUri}&scope=${scope}&access_type=offline`

  return NextResponse.redirect(googleAuthUrl)
}
