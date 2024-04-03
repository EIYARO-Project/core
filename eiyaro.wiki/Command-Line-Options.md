# Command Line Options

[ [English]() | [中文](https://github.com/boy-good/rpc/wiki/3.ey%E7%9A%84cli%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E) ]

* [create-key](#create-key)
* [list-keys](#list-keys)
* [delete-key](#delete-key)
* [create-account](#create-account)
* [list-accounts](#list-accounts)
* [delete-account](#delete-account)
* [create-account-receiver](#create-account-receiver)
* [list-addresses](#list-addresses)
* [validate-address](#validate-address)
* [create-asset](#create-asset)
* [get-asset](#get-asset)
* [list-assets](#list-assets)
* [list-balances](#list-balances)
* [create-access-token](#create-access-token)
* [list-access-tokens](#list-access-tokens)
* [check-access-token](#check-access-token)
* [get-block-count](#get-block-count)
* [get-block-hash](#get-block-hash)
* [get-block](#get-block)
* [get-difficulty](#get-difficulty)
* [get-block-header](#get-block-header)
* [is-mining](#is-mining)
* [gas-rate](#gas-rate)
* [set-mining](#set-mining)
* [build-transaction](#build-transaction)
* [sign-transaction](#sign-transaction)
* [submit-transaction](#submit-transaction)


more API detail:  [API Reference](https://github.com/eiyaro/eiyaro/wiki/API-Reference)

## create-key 

|             Function              |          Description          |
| :---------------------------: | :--------------------: |
| create-key（alias，password） | create key  |
|           **Parameters**            |                        |
|             alias             |        key alias        |
|           password            |        key password        |
|           **Returns**            |                        |
|             alias             |        key alias        |
|             xpub              |     public key     |
|             file              |   key file path   |

**Example**
`./eiyarocli create-key wen 123456`

## list-keys

|   Function    |         Description         |
| :-------: | :------------------: |
| list-keys | acquire all the keys in node |
| **Parameters**  |                      |
|    无     |                      |
| **Returns**  |   return the object array   |
|   alias   |       key alias       |
|   xpub    |    public key    |
|   file    |  key file path  |

**Example**
`./eiyarocli list-keys`

## delete-key

|                Function           |           Description      |
| :-----------------------------------------: | :----------------------: |
|        delete-key（xpub，password）         |         delete key       |
|              **Parameters**                |                          |
|                    xpub                     |      public key      |
|                  password                   |   password   |
|                  **Returns**                |                        |
|           Successfully delete key          |      delete success      |
|       key not found or wrong password     | public error or not found |
| could not decrypt key with given passphrase |     wrong password     |

**Example** `./eiyarocli delete-key 7b328361cb360a2dbd289b4fcc94f38b6dcf4b1f1469a96489950146e8d7feb2d8d5254c3d2f6bd0c6cd6ea7be08d5f7672e8bc49c3a9c3b67ad3f8190d1be79 123456`

## create-account

|      Function      |           Description           |
| :------------: | :----------------------: |
| create-account |         create account         |
|    **Parameters**    |                          |
|     alias      |          account alias          |
|    xpub(s)     |           public key           |
|    **Returns**    |                          |
|     alias      |          account alias          |
|       id       |          account id          |
|   key_index    | the index of account in node |
|     quorum     |          number of signatures       |
|     xpubs      |       public key        |

**Example** `./eiyarocli create-account test 81b6150ee90b7936a8b21f7f3e5028d9060aef1cc633c13cbbb21183bbf66f3b484bf876468b4278f1421f9b66d55d6a4d71542e20faf908590a39d4e96cc0ed`

## list-accounts

|     Function      |              Description              |
| :-----------: | :------------------------------: |
| list-accounts |        get all accounts in your node       |
|   **Parameters**    |                                  |
|      无       |                                  |
|   **Returns**    |         return obj array below         |
|      id       |              帐户id              |
|     alias     |              账户名              |
|   key_index   |     帐户在钱包节点中的索引值     |
|    quorom     | 签名数，即发送交易时所需的密钥数 |
|     xpubs     |             公钥数组             |

**Example** `./eiyarocli list-accounts`

## delete-account

|              Function              | Description                  |
| :----------------------------------: | ------------------------------ |
| delete-account（alias \| accountID） | delete account                 |
|               **Parameters**              |                          |
|          alias or accountID          | account alias or account id   |
|               **Returns**               |                             |
|     Successfully delete account      | the success prompt of delete  |
|         fail to find account         | the error of delete |

**Example** `./eiyarocli delete-account test  or  ./eiyarocli delete-account 0E96GEA9G0A04`

## create-account-receiver

| Function                         | Description                   |
| -------------------------------- | --------------------- |
| create-account-receiver（alias \| accountID） | create new address for account             |
| **Parameters**                                      |                                |
| alias                     | mandatory:account alias               |
| accountID                 | optional， account id                 |
| **Returns**                      |                                |
| address                         | address                          |
| control_program                 | used for transaction |

**Example** `./eiyarocli create-account-receiver zhang or ./eiyarocli create-account-receiver zhang 0DV0J74K00A02`

## list-addresses

|       Function             |          Description                   |
| :---------------------------: | :----------------------------------------: |
| list-addresses（alias \| id） |    get all addresses in a account     |
|           **Parameters**            |                               |
|             alias \ id            |           account alias or id    |
|           **Returns**            |  return the object array below      |
|         account_alias         |                account alias         |
|          account_id           |             account id                |
|            address            |             address                  |
|            change             | if addresschange when the account change，default true |

**Example** `./eiyarocli list-addresses --alias zhang or ./eiyarocli list-addresses --id 0DV0J74K00A02`

## validate-address

|        Function             |         Description         |
| :-------------------------: | :------------------: |
| validate-address（address） |   whether the checkout address is legitimate   |
|          **Parameters**           |                      |
|           address           |         address         |
|          **Returns**           |                      |
|            valid            | legitimate or not，true is legitimate |
|          is_local           |   whether the node is local   |

**Example** `./eiyarocli validate-address sm1qrztgvhxgfy2njgewdhk524uhzhdw03g9l63u48`

## create-asset

|         Function          |         Description         |
| :-------------------: | :------------------: |
|     create-asset      |       create asset       |
|       **Parameters**        |                      |
|         alias         |       asset alias      |
|        xpub(s)        |        public key group        |
|       **Returns**        |                      |
|         alias         |       asset alias       |
|      definition       |       asset definition       |
|          id           |       asset id        |
|   issuance_program    |  control program of issue asset  |
|        quorum         |     number of  signatures        |
|         keys          | JSONObject，parameter below |
| asset_derivation_path |                      |
|     asset_pubkey      |                      |
|       root_xpub       |                      |

**Example** `./eiyarocli create-asset GODS 81b6150ee90b7936a8b21f7f3e5028d9060aef1cc633c13cbbb21183bbf66f3b484bf876468b4278f1421f9b66d55d6a4d71542e20faf908590a39d4e96cc0ed`

## get-asset

|        Function         |         Description        |
| :-----------------: | :----------------------: |
|      get-asset      |     get asset information     |
|      **Parameters**       |                          |
|       assetID       |          asset id          |
|      **Returns**       |                          |
|        alias        |         asset alias         |
|         id          |         asset id          |
|    issue_program    |    control program of issue asset    |
|      key_index      |    index of public key        |
|       quorum        |    number of signatures          |
| raw_definition_byte |      byte of asset definition      |
|        type         |           asset type           |
|     vm_version      | the version of VM |
|        xpubs        |        public key group          |
|     definition      | definition of asset, JSONObject below |
|      decimals       |         precise digits         |
|     description     |         description of asset         |
|        name         |         asset alias          |
|       symbol        |         asset symbol         |

**Example** `./eiyarocli get-asset ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff`

## list-assets

| Function            | Description                         |
| ------------------- | ---------------------------------------- |
| list-assets         | get all asset information                |
| **Parameters**            |                                          |
| none                  |                                          |
| **Returns**            | return asset array                       |
| alias               | asset alias                                 |
| id                  | asset id                         |
| issue_program       | the control program of issue asset           |
| key_index           | index of public key                          |
| quorum              | number of signature                          |
| raw_definition_byte | byte of asset definition                     |
| type                | asset type                                |
| vm_version          | version of VM                              |
| xpubs               | public key group                            |
| definition          | definition of asset, JSONObject below |
| decimals            | precise digits                            |
| description         | description of asset                            |
| name                | asset alias                            |
| symbol              | asset symbol                            |

**Example** `./eiyarocli list-assets`

## list-balances

|       Function       |          Description                  |
| :--------------: | :-------------------------------------: |
|  list-balances   |        get balance in wallet         |
|     **Parameters**     |                                         |
|        none        |                                         |
|     **Returns**     |                                         |
|  account_alias   |               account alias                 |
|    account_id    |               account id                  |
|      amount      |               balance                   |
|   asset_alias    |               asset alias                 |
|     asset_id     |               asset id                  |
| asset_definition |       definition of asset，JSONObject           |

**Example** `./eiyarocli list-balances`

## create-access-token

|        Function         |          Description               |
| :-----------------: | :-----------------------------: |
| create-access-token |            create token            |
|      **Parameters**       |                                 |
|       tokenID       |            token id            |
|       **Returns**         |                                 |
|     created_at      |          create time             |
|         id          |               id                |
|        token        |  token(including password) |

**Example** `./eiyarocli create-access-token test`

## list-access-tokens

|        Function        |      Description         |
| :----------------: | :-------------------: |
| list-access-tokens |     get all token     |
|      **Parameters**      |                       |
|         none         |                       |
|      **Returns**      |                       |
|     created_at     |       create time        |
|         id         |       token id        |
|       token        | token(including password) |

**Example** `./eiyarocli  list-access-tokens`

## check-access-token

|         Function             |             Description              |
| :-------------------------------: | :-------------------------------: |
|        check-access-token         |       check token               |
|             **Parameters**        |                                  |
|              tokenID              |                                   |
|              secret               | token password                  |
|             **Returns**              |                              |
|        Valid access token         |   prompt of right token          |
|        ERROR invalid token    |       prompt of wrong token          |
| ERROR nonexisting access token ID |   prompt of tokenId doesn't exist |

**Example** `./eiyarocli check-access-token test 1061bacc735ae574f84001d617ed17a8302f7c6206481b9f45c7235853e6ab19`

## get-block-count

|      Function       |      Description      |
| :-------------: | :------------: |
| get-block-count | get current block height number |
|    **Parameters**     |                |
|       无        |                |
|    **Returns**     |                |
|   block_count   |      block height number      |

**Example** `./eiyarocli get-block-count`

## get-block-hash

|      Function      |        Description        |
| :------------: | :----------------: |
| get-block-hash | the newest block hash |
|    **Parameters**    |                    |
|       none       |                    |
|    **Returns**    |                    |
|   block_hash   |       block hash       |

**Example** `./eiyarocli get-block-hash`

## get-block

|               Function               |        Description         |
| :---------------------------------------: | :------------------------: |
| get-block（ block_height \| block_hash ） | get block infomation by block height or block hash |
|                 **Parameters**                  |                    |
|               block_height                |        block height       |
|                block_hash                 |       block hash       |
|                 **Retruns**               |                          |
|                   hash                    |       block hash          |
|                   size                    |       block size          |
|                  version                  |       block version       |
|                  height                   |       block height        |
|            previous_block_hash            |       previous block hash |
|                 timestamp                 |     timetamp of block     |
|          transaction_merkle_root          |    merkle root value    |
|          transaction_status_hash          |      merkle status      |
|                   bits                    |       difficulty bit      |
|                   nonce                   |           nonce            |
|                difficulty                 |       difficulty          | |               transactions                |    JSONArray，below     |
|                    id                     |     transaction hash      |
|                  inputs                   |   JSONArray，see below    |
|                  outputs                  |   JSONArray，see below    |
|                   size                    |         transaction size |
|                status_fail                |      request status      |
|                time_range                 |      timetamp of response|
|                  version                  |      transaction version  |

|       Parameters       |            Description            |
| :--------------: | :------------------------------: |
|      inputs      |                                  |
|      amout       |           asset amount            |
|    arbitrary     | exist when coinbase transaction |
| asset_definition |       definition of asset, JSONObject        |
|     asset_id     |              asset id              |
|       type       |              asset type               |

|       Parameters       |        Description         |
| :--------------: | :-----------------: |
|     outputs      |                     |
|     address      |      output address       |
|      amount      |        amount         |
| asset_definition | definition of asset, JSONObject |
|     asset_id     |       asset id        |
| control_program  |    control program of asset     |
|        id        | utxo id  |
|     position     |      output position       |
|       type       |     output type      |

**Example** `./eiyarocli get-block 20a31511b7869fc90b849c4d011ccc0b51ee61e742d20034438309236afd78af or ./eiyarocli get-block 8031`

## get-difficulty

|      Function      |         Description       |
| :------------: | :----------------------: |
| get-difficulty | get current or specify block difficulty |
|    **Parameters**    |                          |
|       none       | if none get current difficulty |
|      hash      |    block hash(optional)    |
|     height     |    block height(optional)    |
|    **Returns**    |                          |
|      bits      |      difficulty bits      |
|   difficulty   |         difficulty        |
|      hash      |          block hash          |
|     height     |          block height      |

**Example** `./eiyarocli get-difficulty or ./eiyarocli get-difficulty --height 8251 or ./eiyarocli get-difficulty  --hash 13d8d75dd33fcd277cf9a2d86e71daeab9aac6f8eb06a7505b2766d564720028`

##get-block-header

|       Function       |       Description             |
| :--------------: | :------------------------------: |
| get-block-header | get block header detail info by block height or block hash |
|     **Parameters**     |                            |
|   block_height   |          block height,optional     |
|    block_hash    |          block hash,optional      |
|     **Returns**     |                                  |
|   block_header   |             block header          |
|      reward      |             block reward           |

**Example** `./eiyarocli get-block-header 12 or ./eiyarocli get-block-header e3b7cbc56b355cce8f0c827edaa6a154298d5d42ec398b4a47fa4af2f14b0a36`

## is-mining

|   Function    |           Description                 |
| :-------: | :-----------------------------------: |
| is-mining |         check mining or not          |
| **Parameters**  |                                   |
|    none     |                                       |
| **Returns**  |                                       |
| is_mining | true mining，false not mining |

**Example** `./eiyarocli is-mining`

## gas-rate

|   Function   |     Description     |
| :------: | :---------------: |
| gas-rate | get gas account when you build this transaction |
| **Parameters** |                   |
|    none    |                   |
| **Returns** |                   |
| gas_rate |      gas account      |

**Example** `./eiyarocli gas-rate`

## set-mining

|    Function    |    Description    |
| :--------: | :----------: |
| set-mining | switch mining status |
|   **Parameters**  |              |
|    true    | start mining |
|   false    | stop mining |
|   **Returns** |              |
|     success message       |    mining status          |

**Example** `./eiyarocli set-mining true or ./eiyarocli set-mining false`

## build-transaction

|    Function    |    Description    |
| :---------------: | :-------------------------------: |
| build-transaction |             build transaction              |
|   **Parameters**   |                                   |
|       alias       |             account alias               |
|       asset       |           asset alias            |
|      amount       |  transaction amount 1EY = 10^8NEU  |
|      address      |     transaction target address       |
|       type        |            transaction type              |
|     receiver      | control program id，when type = spend |
|   **Returns**   |                                   |
|       json        |    return raw-transaction    |

**Example** `./eiyarocli build-transaction --alias zhang EY 1000000000 --type spend --receiver 00145bb44b681e970cf1be37513fe9e39785f56bd72c 
	or ./eiyarocli build-transaction --alias zhang EY 1000000000 --address sm1qtw6yk6q7jux0r03h2yl7ncuhsh6kh4ev92z7pk --type address`

## sign-transaction

|    Function    |    Description    |
| :--------------: | :------------------: |
| sign-transaction |    sign for transaction    |
|     **Parameters**     |                      |
| raw-transaction  | json data from build transaction  |
|     password     |    key password     |
|     **Returns**     |                      |
|     json     | signed json data |

**Example** `【notice:remember add '' in json,or it will show too many parameters】./eiyarocli sign-transaction '{"allow_additional_actions":false,"raw_transaction":"070100010161015f65f1a7c13ebd9d687c853dd28fd125702a7d173439969d857cb01b0b98bd4372ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8099c4d5990100011600142412b41a0f48a6dfad726f7b502e557aa39df36b010002013effffffffffffffff   ffffffffffffffffffffffffffffffffffffffffffffffff80ab94ef950101160014a81d7084fcafaa32798e7355284639000aea0d5600013dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8094ebdc03011600145bb44b681e970cf1be37513fe9e39785f56bd72c00","signing_instructions":[{"position":0,"witness_components":[{"keys":[{"derivation_path":["010100000000000000","0100000000000000"],"xpub":"145b584d3de371405e7b6ceba11d69a3ae6f0b22d2ec879f17977f3c51f5f2d14390c20ead68eb2ab4e210ca0958873b28a9635b8c4a29ba74934aa47137c80c"}],"quorum":1,"signatures":null,"type":"raw_tx_signature"},{"type":"data","value":"9817f7c9a9d8729374cfff24a4b2e400fdfd2cfb4ab09ee339925a9033ab7eb7"}]}]}' --password 12345`

## submit-transaction

|    Function    |    Description    |
| :----------------: | :------------------------: |
| submit-transaction |      broadcast transaction    |
|      **Parameters**      |                            |
|      json     | the json return by sign-transaction |
|      **Returns**      |                            |
|        txid        |        transaction hash         |

**Example** ` ./eiyarocli submit-transaction '{"raw_transaction":"070100010161015fedb1e3ea7fd526a6f5402f95e8b2fd870fa9a766c89b3555b016602aa45197e4ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8099c4d5990100011600142412b41a0f48a6dfad726f7b502e557aa39df36b63024070eb4675106805bd6e72e63d45ee9ff30802106fd11be2d4deac192bc3cb75cbfe5180c9aa1285baae09a93997724d0c8a50cde99c1e619012af7fbce645f809209817f7c9a9d8729374cfff24a4b2e400fdfd2cfb4ab09ee339925a9033ab7eb702013effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80ab94ef9501011600140478426d17f952005f10a38e6779ecd2dd6bca6500013dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8094ebdc03011600145bb44b681e970cf1be37513fe9e39785f56bd72c00"}'`