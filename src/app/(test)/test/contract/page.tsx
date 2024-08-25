"use client"

import { useEffect, useState } from "react"
import { ethers, type Contract, type Eip1193Provider } from "ethers"

import { ContractABI, ContractAddress } from "@/lib/contract"
import { toast } from "@/components/ui/use-toast"

export default function ContractTest() {
  const [account, setAccount] = useState<string | null>(null)
  const [contract, setContract] = useState<Contract | null>(null)
  // const [votesStatus, setVotesStatus] = useState<any[]>([])

  async function connect() {
    if (!window.ethereum) return

    const accounts = await window.ethereum.request?.({
      method: "eth_requestAccounts",
    })

    if (!accounts) return

    const provider = new ethers.BrowserProvider(
      window.ethereum as Eip1193Provider
    )
    const signer = await provider.getSigner()
    const contractApple = new ethers.Contract(
      ContractAddress,
      ContractABI,
      signer
    )
    window.localStorage.setItem("address", JSON.stringify(accounts[0]))
    setContract(contractApple)
  }

  // async function getVotes() {
  //   if (!contract) return
  //   const data = [
  //     await contract?.getCandidate(0),
  //     await contract?.getCandidate(1),
  //     await contract?.getCandidate(2),
  //   ]
  //   console.log(data)
  // }

  useEffect(() => {
    contract?.on("NewsPublish", (...props) => {
      console.log(props)
    })
    // getVotes()
  }, [contract])

  async function publish(index: number) {
    try {
      await contract?.publish(index)
      toast({
        title: "Voted Successfully",
      })
    } catch (error) {
      toast({
        title: "Error while voting",
        description: (
          <pre>
            <code>{JSON.stringify(error, null, 2)}</code>
          </pre>
        ),
      })
    }
  }

  return <div></div>
}
