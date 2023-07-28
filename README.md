# FIBER REST API
> Simple REST API using GO Lang, Fiber and MONGO!

A API with CRUD of Bitcoin's Transactions consuming [Blockchain.com](Blockchain.com)'s API

## Installing / Getting started

To run this project you must have GO installed and a MongoDB running!

```shell
git clone https://github.com/limarodrigoo/fiber-rest-api.git
cd fiber-rest-api
go mod tidy  # install's all packages
```


### Initial Configuration

For this projects you need to create a app.env file that contains your configurations of MONGO:
```shell
    DB_HOST=127.0.0.1
    DB_USER=mongo
    DB_PASSWORD=1234
    DB_NAME=cryptos
    DB_PORT=27017
```

## Running

```shell
go run main.go
```

### End points

### GET /transactions/
Get all transactions `transactions on db only`

**Response**

```
// Array of transactions on db
{
    "data": [
        {
            "hash": "e8fb9e06e8bcc16b695c5b2e8e39f9afafd80beb7dee0b1cb34caa375b1e3d0e",
            "ver": 1,
            "vin_sz": 1,
            "vout_sz": 2,
            "size": 207,
            "weight": 720,
            "fee": 0,
            "relayed_by": "0.0.0.0",
            "lock_time": 0,
            "tx_index": 500981282829641,
            "double_spend": false,
            "time": 1690518134,
            "block_index": 800566,
            "block_height": 800566,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "01200000000000000000000000000000000000000000000000000000000000000000",
                    "script": "0336370c0b2f4d61726120506f6f6c2f1a01000200000003000000181fa5b2050000008f010000",
                    "index": 0,
                    "prev_out": {
                        "addr": "",
                        "n": 4294967295,
                        "script": "",
                        "spending_outpoints": [
                            {
                                "tx_index": 500981282829641,
                                "n": 0
                            }
                        ],
                        "spent": true,
                        "tx_index": 0,
                        "type": 0,
                        "value": 0
                    }
                }
            ],
            "out": [
                {
                    "type": 0,
                    "spent": false,
                    "value": 636721454,
                    "spending_outpoints": [],
                    "n": 0,
                    "tx_index": 500981282829641,
                    "script": "76a9142fc701e2049ee4957b07134b6c1d771dd5a96b2188ac",
                    "addr": "15MdAHnkxt9TMC2Rj595hsg8Hnv693pPBB"
                },
                {
                    "type": 0,
                    "spent": false,
                    "value": 0,
                    "spending_outpoints": [],
                    "n": 1,
                    "tx_index": 500981282829641,
                    "script": "6a24aa21a9edd60c498b28aa8c05acad9ed34ab63a80654c176134274b674a52e70b8c3013bf",
                    "addr": ""
                }
            ]
        },
    ]
}    

```

### /transaction/$hash
Get transactions by hash

**Parameters**

|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `hash` | required | string  | The hash of the transaction                                  |

**Response**

```
// Transaction
{
    "data": {
        "hash": "e8fb9e06e8bcc16b695c5b2e8e39f9afafd80beb7dee0b1cb34caa375b1e3d0e",
        "ver": 1,
        "vin_sz": 1,
        "vout_sz": 2,
        "size": 207,
        "weight": 720,
        "fee": 0,
        "relayed_by": "0.0.0.0",
        "lock_time": 0,
        "tx_index": 500981282829641,
        "double_spend": false,
        "time": 1690518134,
        "block_index": 800566,
        "block_height": 800566,
        "inputs": [
            {
                "sequence": 4294967295,
                "witness": "01200000000000000000000000000000000000000000000000000000000000000000",
                "script": "0336370c0b2f4d61726120506f6f6c2f1a01000200000003000000181fa5b2050000008f010000",
                "index": 0,
                "prev_out": {
                    "addr": "",
                    "n": 4294967295,
                    "script": "",
                    "spending_outpoints": [
                        {
                            "tx_index": 500981282829641,
                            "n": 0
                        }
                    ],
                    "spent": true,
                    "tx_index": 0,
                    "type": 0,
                    "value": 0
                }
            }
        ],
        "out": [
            {
                "type": 0,
                "spent": false,
                "value": 636721454,
                "spending_outpoints": [],
                "n": 0,
                "tx_index": 500981282829641,
                "script": "76a9142fc701e2049ee4957b07134b6c1d771dd5a96b2188ac",
                "addr": "15MdAHnkxt9TMC2Rj595hsg8Hnv693pPBB"
            },
            {
                "type": 0,
                "spent": false,
                "value": 0,
                "spending_outpoints": [],
                "n": 1,
                "tx_index": 500981282829641,
                "script": "6a24aa21a9edd60c498b28aa8c05acad9ed34ab63a80654c176134274b674a52e70b8c3013bf",
                "addr": ""
            }
        ]
    }
}

```