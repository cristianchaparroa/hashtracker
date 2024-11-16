# HashTracker


## Local environment

### Backend
1. Install Golang
2. Run the project

```go
go run cmd/* 
```

### Endpoints

```bash
curl localhost:8080/transactions/0x8611C17eA68caE77762AdF6446a00f5a71dd7784?level=1
```
The following is the expected result

```
{
	"transactions": [
		{
			"hash": "0x6637a2ab8da7405f1b72dfc325ae80abea0f31149d8bcabf259e402d6a1a2e3e",
			"from": "0x8611C17eA68caE77762AdF6446a00f5a71dd7784",
			"to": "0x7cCD87331574Fe39140697962555848236C75DFA",
			"value": "10000000000000000",
			"timestamp": "2024-11-14T06:10:23Z"
		},
		{
			"hash": "0x396c3a2b3a63efa6ff81b2b268068a405bc4b159bd0c5946ca07880589e5058a",
			"from": "0x28C6c06298d514Db089934071355E5743bf21d60",
			"to": "0x8611C17eA68caE77762AdF6446a00f5a71dd7784",
			"value": "13009150000000000",
			"timestamp": "2024-11-14T06:08:23Z"
		},
	],
	"graph": {
		"nodes": [
			{
				"name": "0x8611C17eA68caE77762AdF6446a00f5a71dd7784"
			},
			{
				"name": "0x7cCD87331574Fe39140697962555848236C75DFA"
			},..
		],
		"links": [
			{
				"source": 0,
				"target": 1,
				"value": 10000000000000000
			},
			{
				"source": 2,
				"target": 0,
				"value": 13009150000000000
			},
			{
				"source": 0,
				"target": 3,
				"value": 3198986360403884
			},..
		]
	},
	"mermaid": "Z3JhcGggTFIKICAgIDB4ODYxMS4uLjc3ODQtLT58My4yMCBtRVRIfDB4MjUzNS4uLjMwM2IKICAgIDB4ODYxMS4uLjc3ODQtLT58MjAuMDAgbUVUSHwweDAwMDAuLi5iRTU5CiAgICAweERGZDUuLi45NjNkLS0+fDIuMzggbUVUSHwweDg2MTEuLi43Nzg0CiAgICAweDg2MTEuLi43Nzg0LS0+fDEwLjAwIG1FVEh8MHg3Y0NELi4uNURGQQogICAgMHgyOEM2Li4uMWQ2MC0tPnw0Mi43MyBtRVRIfDB4ODYxMS4uLjc3ODQ="
}
```

- Sankey Graph support
- Mermaid support


## Production Endpoints

```
GET https://hashtracker-em3r7.ondigitalocean.app/transactions/<WALLET_ADDRESS>?level=1 
GET https://hashtracker-em3r7.ondigitalocean.app/transactions/<ENS>?level=1 
```

## Services
The following are the list of services used.

### Blockscout endpoints
[/addresses/{address_hash}/transactions](https://eth.blockscout.com/api-docs)

### Subgraphs
ENS Resolution: https://api.thegraph.com/subgraphs/name/ensdomains/ens
