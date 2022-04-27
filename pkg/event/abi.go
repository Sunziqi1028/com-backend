package event

type Contract struct {
	Abi      string
	Event    string
	EventHex string
}

var (
	StartupContract = Contract{
		Event:    "created",
		EventHex: "0x822c16987e5c88fd1ec8ce2935c0b5daf646231496d234745d143a9b62673973",
		Abi: `[
	{
		"inputs": [],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "name",
				"type": "string"
			},
			{
				"components": [
					{
						"internalType": "string",
						"name": "name",
						"type": "string"
					},
					{
						"internalType": "enum Startup.Mode",
						"name": "mode",
						"type": "uint8"
					},
					{
						"internalType": "string",
						"name": "logo",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "mission",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "overview",
						"type": "string"
					},
					{
						"internalType": "bool",
						"name": "isValidate",
						"type": "bool"
					}
				],
				"indexed": false,
				"internalType": "struct Startup.Profile",
				"name": "startUp",
				"type": "tuple"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "msg",
				"type": "address"
			}
		],
		"name": "created",
		"type": "event"
	},
	{
		"stateMutability": "payable",
		"type": "fallback"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "string",
						"name": "name",
						"type": "string"
					},
					{
						"internalType": "enum Startup.Mode",
						"name": "mode",
						"type": "uint8"
					},
					{
						"internalType": "string",
						"name": "logo",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "mission",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "overview",
						"type": "string"
					},
					{
						"internalType": "bool",
						"name": "isValidate",
						"type": "bool"
					}
				],
				"internalType": "struct Startup.Profile",
				"name": "p",
				"type": "tuple"
			}
		],
		"name": "newStartup",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"name": "startups",
		"outputs": [
			{
				"internalType": "string",
				"name": "name",
				"type": "string"
			},
			{
				"internalType": "enum Startup.Mode",
				"name": "mode",
				"type": "uint8"
			},
			{
				"internalType": "string",
				"name": "logo",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "mission",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "overview",
				"type": "string"
			},
			{
				"internalType": "bool",
				"name": "isValidate",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address payable",
				"name": "receiver",
				"type": "address"
			}
		],
		"name": "suicide0",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"stateMutability": "payable",
		"type": "receive"
	}
]`,
	}
)
