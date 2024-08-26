"use client"

import { useState } from "react"

import { cn } from "@/lib/utils"

import { ChatPrompts } from "./chat-propmts"
import { Icons } from "./icons"
import { Button } from "./ui/button"

export default function NewId() {
  return (
    <>
    <div className="h-5"></div>
    <div className="h-max">
      <div className="flex h-max">
        <div className="w-1/2 pr-4">
          <h2 className="text-6xl mb-4">
            Lorem, ipsum dolor sit amet consectetur adipisicing elit. Assumenda,
            sunt!
          </h2>
          <p className="">
            Lorem ipsum dolor sit, amet consectetur adipisicing elit. Incidunt
            saepe perspiciatis sunt, dolore ullam itaque quisquam unde nisi qui
            officiis rerum! Dolorum quidem consequatur omnis.
          </p>
        </div>
        <div className="w-3/6">
          <img
            src="data:image/jpeg;base64,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"
            alt=""
            className="h-full"
          />
        </div>
      </div>

      <div></div>
    </div>
    </>
  )
}
