# Normal Transaction

## Introduction

A blockchain contain many transactions.Each transaction contains one or more inputs, and one or more outputs.more detail please refer to UTXO model

An input either issues new units of an asset, or transfers existing units by naming an output of an earlier transaction as the source for the transfer. 

Outputs take the asset units from the inputs and define how they will be allocated.

Transactions are divided into normal transactions and advanced transactions,we just introduce the normal transactions in this doc.

## Normal transaction summary

you can see a normal transaction summary below

![](images/normal_transaction_summary.png)

- **Spend**: a spend output,spended by the account written in the end(martin)
- **Control**: a input,control by the account written in the end,if the account is same as spend account,it means a change
- **amount**: the number of asset
- **asset**: the type of asset
- **account**: the asset owner

## Normal transaction detail

you can see a normal transaction detail below

![](images/normal_transaction_detail.png)

- **ID**: transaction ID, you can use ID to search the information on blockchain explorer
- **Timestamp**: the time of block generation
- **Block ID**: the unique symbol of block
- **Block Height**：the unique number of block
- **Position**: the serial number of transaction in this block
- **Gas**: transaction fee, pay for miner

## Creating new transaction

you can press "New" button on the right of Transactions page to create a new transaction

![](images/new_transaction.png)

- **From**:the source of transaction, you must specify the account and asset imformation
- **To**: the destination of transaction, you must specify the address and amount
- **Gas**:transaction fee, It is divided into three ways:Stardard,Fast and Customize
	- Standard:the standard fee,calculate by the systom automatically
	- Fast: double standard fee,if you choose Fast,it will be confirmed priority
	- Customize: you can specify transaction fee yourself.and it's very obvious that more gas you paid faster your transaction comfirmed,but upper limit is 0.4 EY,if you give more than 0.4 EY,it's no help for accelerating transaction confirmation
- **Password**:used for protecting private key,only you have your password,you can use your balances to create a transaction

## Transaction record

if the transaction is comfirmed by miner, It will show in your transaction list(you have to synchronize the highest block)

![](images/transaction_record.png)