export const ContractAddress = "0xa7670DAa354D729C5f19efefE3d87985096f5026"

export const ContractABI = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "string",
        name: "id",
        type: "string",
      },
      {
        indexed: false,
        internalType: "string",
        name: "title",
        type: "string",
      },
      {
        indexed: false,
        internalType: "string",
        name: "content",
        type: "string",
      },
      {
        indexed: false,
        internalType: "string",
        name: "createdAt",
        type: "string",
      },
      {
        indexed: false,
        internalType: "string",
        name: "updatedAt",
        type: "string",
      },
      {
        indexed: false,
        internalType: "string",
        name: "publisher",
        type: "string",
      },
    ],
    name: "NewsPublish",
    type: "event",
  },
  {
    inputs: [],
    name: "checkRunning",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    name: "news",
    outputs: [
      {
        internalType: "string",
        name: "id",
        type: "string",
      },
      {
        internalType: "string",
        name: "title",
        type: "string",
      },
      {
        internalType: "string",
        name: "content",
        type: "string",
      },
      {
        internalType: "string",
        name: "createdAt",
        type: "string",
      },
      {
        internalType: "string",
        name: "updatedAt",
        type: "string",
      },
      {
        internalType: "string",
        name: "publisher",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "string",
        name: "id",
        type: "string",
      },
      {
        internalType: "string",
        name: "title",
        type: "string",
      },
      {
        internalType: "string",
        name: "content",
        type: "string",
      },
      {
        internalType: "string",
        name: "createdAt",
        type: "string",
      },
      {
        internalType: "string",
        name: "updatedAt",
        type: "string",
      },
      {
        internalType: "string",
        name: "publisher",
        type: "string",
      },
    ],
    name: "publish",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
]
