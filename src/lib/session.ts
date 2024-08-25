import { cookies } from "next/headers"

export async function getCurrentUser() {
  const user = cookies().get("user")?.value
  const session = (user ? JSON.parse(user) : null) as {
    id: string
    email: string
    name: string
    image: string
  } | null

  return session
}
