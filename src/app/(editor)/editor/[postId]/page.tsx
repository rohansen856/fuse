import { notFound, redirect } from "next/navigation"
import { News, User } from "@prisma/client"
import axios from "axios"

import { authOptions } from "@/lib/auth"
import { db } from "@/lib/db"
import { getCurrentUser } from "@/lib/session"
import { Editor } from "@/components/editor"

async function getPostForUser(postId: News["id"], userId: User["id"]) {
  console.log(postId)
  const response = await axios.get(`http://localhost:8000/news/${postId}`)
  console.log(response)
  return response.data
}

interface EditorPageProps {
  params: { postId: string }
}

export default async function EditorPage({ params }: EditorPageProps) {
  const user = await getCurrentUser()

  if (!user) {
    redirect(authOptions?.pages?.signIn || "/login")
  }
  console.log(params)
  const post = await getPostForUser(params.postId, user.id)

  if (!post) {
    notFound()
  }

  return (
    <Editor
      post={{
        id: post.id,
        title: post.title,
        content: post.content,
        published: post.published,
      }}
    />
  )
}
