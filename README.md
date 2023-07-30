# API Project Readme

This project is an API developed using the Go language, Fiber framework, and MongoDB to manage transactions data. The API provides several endpoints to interact with transaction information stored in the database. Below are the details of each endpoint and the expected responses.

## Endpoints

### Get All Transactions

- **Endpoint:** `GET /transactions`
- **Description:** Retrieves a list of all transactions stored in the database.
- **Example Response:**

```
"data": [
        {
            "hash": "96145f1d41235a760b5e44697303c457689ad312770c9664b66b6e1932437f04",
            "ver": 1,
            "vin_sz": 1,
            "vout_sz": 2,
            "size": 382,
            "weight": 766,
            "fee": 90000,
            "relayed_by": "0.0.0.0",
            "lock_time": 0,
            "tx_index": 158228310863309,
            "double_spend": false,
            "time": 1690517773,
            "block_index": 800566,
            "block_height": 800566,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "0400473044022012134fd89bb8da04a7c19022fcdadf7fdf1c51d193359894bf484c4b9207780f022037a66e1f3e17ce206aea73c0255ba099c5c5af59056e52c339958a6f370fec8b014730440220388d55450fef3ee9d26c70aa60bc969b4a877879b0ae0330e796f6dc03ca6cdf02204d8b2bb3bdcbf2f6b693d1e88d75b845d39e04a26be88d90f834c207a0f9d6b4016952210375e00eb72e29da82b89367947f29ef34afb75e8654f6ea368e0acdfd92976b7c2103a1b26313f430c4b15bb1fdce663207659d8cac749a0e53d70eff01874496feff2103c96d495bfdd5ba4145e3e046fee45e84a8a48ad05bd8dbb395c011a32cf9f88053ae",
                    "script": "",
                    "index": 0,
                    "prev_out": {
                        "addr": "bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej",
                        "n": 4,
                        "script": "0020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d",
                        "spending_outpoints": [
                            {
                                "tx_index": 158228310863309,
                                "n": 0
                            }
                        ],
                        "spent": true,
                        "tx_index": 3369726711104028,
                        "type": 0,
                        "value": 210155170
                    }
                }
            ],
            "out": [
                {
                    "type": 0,
                    "spent": false,
                    "value": 24542995,
                    "spending_outpoints": [],
                    "n": 0,
                    "tx_index": 158228310863309,
                    "script": "76a91490c5c51fc30b04eafb80c91103a7c1a2f3d38ac588ac",
                    "addr": "1ECVDvfKw74jyD9HYxTFZYSyqVWGB5x7cp"
                },
                {
                    "type": 0,
                    "spent": true,
                    "value": 185522175,
                    "spending_outpoints": [
                        {
                            "tx_index": 4162255667144669,
                            "n": 0
                        }
                    ],
                    "n": 1,
                    "tx_index": 158228310863309,
                    "script": "0020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d",
                    "addr": "bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej"
                }
            ]
        },
]
```


### Get Transaction by Hash

- **Endpoint:** `GET /transaction/:hash`
- **Description:** Retrieves a specific transaction using its hash value.
- **Example Response:**

```
"data": 
        {
            "hash": "96145f1d41235a760b5e44697303c457689ad312770c9664b66b6e1932437f04",
            "ver": 1,
            "vin_sz": 1,
            "vout_sz": 2,
            "size": 382,
            "weight": 766,
            "fee": 90000,
            "relayed_by": "0.0.0.0",
            "lock_time": 0,
            "tx_index": 158228310863309,
            "double_spend": false,
            "time": 1690517773,
            "block_index": 800566,
            "block_height": 800566,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "0400473044022012134fd89bb8da04a7c19022fcdadf7fdf1c51d193359894bf484c4b9207780f022037a66e1f3e17ce206aea73c0255ba099c5c5af59056e52c339958a6f370fec8b014730440220388d55450fef3ee9d26c70aa60bc969b4a877879b0ae0330e796f6dc03ca6cdf02204d8b2bb3bdcbf2f6b693d1e88d75b845d39e04a26be88d90f834c207a0f9d6b4016952210375e00eb72e29da82b89367947f29ef34afb75e8654f6ea368e0acdfd92976b7c2103a1b26313f430c4b15bb1fdce663207659d8cac749a0e53d70eff01874496feff2103c96d495bfdd5ba4145e3e046fee45e84a8a48ad05bd8dbb395c011a32cf9f88053ae",
                    "script": "",
                    "index": 0,
                    "prev_out": {
                        "addr": "bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej",
                        "n": 4,
                        "script": "0020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d",
                        "spending_outpoints": [
                            {
                                "tx_index": 158228310863309,
                                "n": 0
                            }
                        ],
                        "spent": true,
                        "tx_index": 3369726711104028,
                        "type": 0,
                        "value": 210155170
                    }
                }
            ],
            "out": [
                {
                    "type": 0,
                    "spent": false,
                    "value": 24542995,
                    "spending_outpoints": [],
                    "n": 0,
                    "tx_index": 158228310863309,
                    "script": "76a91490c5c51fc30b04eafb80c91103a7c1a2f3d38ac588ac",
                    "addr": "1ECVDvfKw74jyD9HYxTFZYSyqVWGB5x7cp"
                },
                {
                    "type": 0,
                    "spent": true,
                    "value": 185522175,
                    "spending_outpoints": [
                        {
                            "tx_index": 4162255667144669,
                            "n": 0
                        }
                    ],
                    "n": 1,
                    "tx_index": 158228310863309,
                    "script": "0020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d",
                    "addr": "bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej"
                }
            ]
        }
```

### Get Transactions by Block Hash

- **Endpoint:** `GET /transactions/by-block/:hash`
- **Description:** Retrieves a list of transactions associated with a specific block hash.
- **Example Response:**

```
"data": [
        {
            "hash": "96145f1d41235a760b5e44697303c457689ad312770c9664b66b6e1932437f04",
            "ver": 1,
            "vin_sz": 1,
            "vout_sz": 2,
            "size": 382,
            "weight": 766,
            "fee": 90000,
            "relayed_by": "0.0.0.0",
            "lock_time": 0,
            "tx_index": 158228310863309,
            "double_spend": false,
            "time": 1690517773,
            "block_index": 800566,
            "block_height": 800566,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "0400473044022012134fd89bb8da04a7c19022fcdadf7fdf1c51d193359894bf484c4b9207780f022037a66e1f3e17ce206aea73c0255ba099c5c5af59056e52c339958a6f370fec8b014730440220388d55450fef3ee9d26c70aa60bc969b4a877879b0ae0330e796f6dc03ca6cdf02204d8b2bb3bdcbf2f6b693d1e88d75b845d39e04a26be88d90f834c207a0f9d6b4016952210375e00eb72e29da82b89367947f29ef34afb75e8654f6ea368e0acdfd92976b7c2103a1b26313f430c4b15bb1fdce663207659d8cac749a0e53d70eff01874496feff2103c96d495bfdd5ba4145e3e046fee45e84a8a48ad05bd8dbb395c011a32cf9f88053ae",
                    "script": "",
                    "index": 0,
                    "prev_out": {
                        "addr": "bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej",
                        "n": 4,
                        "script": "0020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d",
                        "spending_outpoints": [
                            {
                                "tx_index": 158228310863309,
                                "n": 0
                            }
                        ],
                        "spent": true,
                        "tx_index": 3369726711104028,
                        "type": 0,
                        "value": 210155170
                    }
                }
            ],
            "out": [
                {
                    "type": 0,
                    "spent": false,
                    "value": 24542995,
                    "spending_outpoints": [],
                    "n": 0,
                    "tx_index": 158228310863309,
                    "script": "76a91490c5c51fc30b04eafb80c91103a7c1a2f3d38ac588ac",
                    "addr": "1ECVDvfKw74jyD9HYxTFZYSyqVWGB5x7cp"
                },
                {
                    "type": 0,
                    "spent": true,
                    "value": 185522175,
                    "spending_outpoints": [
                        {
                            "tx_index": 4162255667144669,
                            "n": 0
                        }
                    ],
                    "n": 1,
                    "tx_index": 158228310863309,
                    "script": "0020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d",
                    "addr": "bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej"
                }
            ]
        },
]
```

### Get Transactions by Address

- **Endpoint:** `GET /transactions/by-address/:address`
- **Description:** Retrieves a list of transactions associated with a specific address.
- **Example Response:**

```
"data": [
        {
            "hash": "96145f1d41235a760b5e44697303c457689ad312770c9664b66b6e1932437f04",
            "ver": 1,
            "vin_sz": 1,
            "vout_sz": 2,
            "size": 382,
            "weight": 766,
            "fee": 90000,
            "relayed_by": "0.0.0.0",
            "lock_time": 0,
            "tx_index": 158228310863309,
            "double_spend": false,
            "time": 1690517773,
            "block_index": 800566,
            "block_height": 800566,
            "inputs": [
                {
                    "sequence": 4294967295,
                    "witness": "0400473044022012134fd89bb8da04a7c19022fcdadf7fdf1c51d193359894bf484c4b9207780f022037a66e1f3e17ce206aea73c0255ba099c5c5af59056e52c339958a6f370fec8b014730440220388d55450fef3ee9d26c70aa60bc969b4a877879b0ae0330e796f6dc03ca6cdf02204d8b2bb3bdcbf2f6b693d1e88d75b845d39e04a26be88d90f834c207a0f9d6b4016952210375e00eb72e29da82b89367947f29ef34afb75e8654f6ea368e0acdfd92976b7c2103a1b26313f430c4b15bb1fdce663207659d8cac749a0e53d70eff01874496feff2103c96d495bfdd5ba4145e3e046fee45e84a8a48ad05bd8dbb395c011a32cf9f88053ae",
                    "script": "",
                    "index": 0,
                    "prev_out": {
                        "addr": "bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej",
                        "n": 4,
                        "script": "0020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d",
                        "spending_outpoints": [
                            {
                                "tx_index": 158228310863309,
                                "n": 0
                            }
                        ],
                        "spent": true,
                        "tx_index": 3369726711104028,
                        "type": 0,
                        "value": 210155170
                    }
                }
            ],
            "out": [
                {
                    "type": 0,
                    "spent": false,
                    "value": 24542995,
                    "spending_outpoints": [],
                    "n": 0,
                    "tx_index": 158228310863309,
                    "script": "76a91490c5c51fc30b04eafb80c91103a7c1a2f3d38ac588ac",
                    "addr": "1ECVDvfKw74jyD9HYxTFZYSyqVWGB5x7cp"
                },
                {
                    "type": 0,
                    "spent": true,
                    "value": 185522175,
                    "spending_outpoints": [
                        {
                            "tx_index": 4162255667144669,
                            "n": 0
                        }
                    ],
                    "n": 1,
                    "tx_index": 158228310863309,
                    "script": "0020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d",
                    "addr": "bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej"
                }
            ]
        },
]
```

### Delete Transaction by Hash

- **Endpoint:** `DELETE /transaction/:hash`
- **Description:** Deletes a specific transaction using its hash value.
- **Example Response:**

```
{
"deleted": true
}
```


## Dependencies

To run this API project, make sure you have the following dependencies installed:

- Go (Programming Language)
- Fiber (Web Framework)
- MongoDB (Database)

## How to Run

1. Clone this repository to your local machine.
2. Set up the MongoDB database and connection details.
3. Navigate to the project's root directory.
4. Run the following command to install dependencies:

```
go mod download
```

5. Build and run the project:

```
go run main.go
```


The API should now be running and accessible at `http://localhost:3000`.

