"use client"

import * as React from "react"
import Image from "next/image"
import Link from "next/link"
import { useSearchParams } from "next/navigation"
import { zodResolver } from "@hookform/resolvers/zod"
import { signIn } from "next-auth/react"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { cn } from "@/lib/utils"
import { userAuthSchema } from "@/lib/validations/auth"
import { buttonVariants } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { toast } from "@/components/ui/use-toast"
import { Icons } from "@/components/icons"

interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> {}

type FormData = z.infer<typeof userAuthSchema>

export function UserAuthForm({ className, ...props }: UserAuthFormProps) {
  const [isLoading, setIsLoading] = React.useState<boolean>(false)
  const searchParams = useSearchParams()

  return (
    <div className={cn("grid gap-6", className)} {...props}>
      <div className="relative">
        <div className="relative flex justify-center text-xs uppercase">
          <Link
            className={buttonVariants({
              className: "w-[400px] max-w-full gap-4",
            })}
            href={"/api/auth/google"}
          >
            <Image src={"/icons/google.svg"} height={20} width={20} alt="G" />
            google
          </Link>
        </div>
      </div>
    </div>
  )
}
