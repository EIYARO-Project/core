# Architecture Overview

Protocol module, wallet module, and network module are the three modules for the top view of the project. Protocol module is the core of Eiyaro, it's response for validate blocks, validate transactions, update chain status, and maintain the UTXO storage. The network module is used for manage the connections and synchronize logic between different Eiyaro nodes. Wallet module is an optional module for Eiyaro full node. The purpose of the wallet module is to provide user ability to send or receive transactions.

# Protocol Module

There are four submodules under protocol module

-  Transaction Pool: Transaction pool is designed for tempory store all the unconfirmed transactions in memory. Unconfirmed Transactions are either local user-submitted or synchronize from other Eiyaro node through network module. A transaction will be removed from the transaction pool if it has been packed into a block or become invalid due to time range condition.

-  Data Storage: Data storage module is responsible for persistent store main chain blocks, orphan blocks, chain status, and maintain the newest UTXO view map.

-  Block Index: Block index is a tree structure to manage the main chain block status. Block index will contain all the block's header information include orphan blocks due to the orphan blocks might become the main chain if it satisfies the consensus logic.

- Consensus: Consensus module controls the validation rules for Eiyaro chain include POW work proof, Merkle proof validation, BVM validation, time range check, transaction validation, UTXO check, and gas limit rules check.

# wallet module

There are four submodules under wallet module

-  HSM module: HSM module is created to generate the private key with a mnemonic, sign transaction, and symmetric encryption/decryption private key

-  Account module: Account module implemented the multi-account hierarchy deterministic wallets. A user can use one or multiple private keys to create a unique account, and each account may have infinitely many addresses.

-  Asset module: Asset module helps the user create and issue their own asset. Each asset has its own unique Id and smart contract.

- Transaction module: Transaction module will record all the user related transactions and UTXOs in persistent storage. Transaction module is also responsible for UTXO select in build transaction process.

# Network module

-  Node Discover: Node discover is using UDP connection to find all the alive Eiyaro node on the internet.

-  Transaction Synchronize: After two Eiyaro nodes finished hands shake process, this module will exchange all the transactions in their memory pool.

-  Block Synchronize: Block synchronize is a passive module, it will request the block if the connected node has higher blocks.

-  New Block Fast Broadcast: It's a special module for only broadcast new mined blocks. The main purpose of this module is to reduce the latency of the whole network gets the newest block.