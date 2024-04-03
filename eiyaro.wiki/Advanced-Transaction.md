# Advanced Transaction

## Introduction

Advanced transaction is also a type of transaction,so it's consist of  one or more input and one or more output and need transaction fee.

Advanced transaction support a flexible combination of input and output, you can take multi-assets as input in one transaction , you can send the assets to many people and you can retire some assets. It can do many things depends on your business.

## Advanced transaction summary

Advanced transaction contain four type of actions

![](images/advanced_transaction_summary.png)

| Action     | Explain     |Input      |
| ---- | ---- | ---- |
| Issue     |issue some assets(the assets must be defined in your wallet)      | assets,amount     |
| Spend from account     | spend some assets from your account,regard as a input     | account,assets,amount     |
| Control with address     | specify the owner of the assets,regard as an output     |address,assets,amount      |
| Retire     |destroy the assets      |assets,amount      |

notice that four type of actions are Independent and relevant,we will introduce some usage below.

## Assets registration

if you want to register your assets to Eiyaro blockchain,advanced transaction will help you to achieve this goal.

### 1. Create an assets type

first of all, you must create a new assets type.

go to the asset page in Eiyaro wallet,press the "new" button.

![](images/advanced_transaction_new_assets.png)

fill in the alias,definition and key, submit the assets,and then you will find a new assets in your assets page.

but now it has no balance, you must issue the assets.

### 2. Issue asset

go to the transaction page and create a new transaction,choose advanced transaction.  

**add an "issue" action**(if you want to issue more,add more "issue" action)

like the picture below,I issue 1000 "MYassets"

![](images/advanced_transaction_issue_new.png)

**add a "Control with address" action**

fill in the control address(the owner of the assets),the assets and amount,the amount of the assets must equal to the amount when you issue the assets(if more than one address control the assets,the sum must equal to issue assets amount).

![](images/advanced_transaction_control_with_address.png)  

**add a "Spend by account" action**

this action is used for paying transaction fee,advanced transaction is known as a special transaction.the EY fee is relevant to the volume of the input,more input more EY.wallet version 1.0.2 don't support automatic calculation of transaction fee,so you may estimate the fee, the suggested price is 0.1 ey, if you fail you have to raise it .

![](images/advanced_transaction_spend_from_account.png)

now we complete the action and submit the transaction,wait for miner confirm.

after the transaction confirmed,you can see your assets in your balance page.

![](images/advanced_transaction_new_assets_balance.png)

## Batch transaction

you can send several assetss to several addresses in one transaction by using advanced transaction.for example, I will send 0.1 EY and 500 MYassets to more than one addresses.

new transaction,choose advanced transaction.

**create two "Spend from account" actions**:

![](images/advanced_transaction_batch_transaction.png)

**create two "Control with address" actions**
notice that the sum of output action must equal to the input amount(the sum must 0.1EY and 500 MYassets)

![](images/advanced_transaction_batch_address.png)

**Don't forget transaction fee**

create a "Spend by account" action,fill in the EY amount

submit the transaction and wait for confirming

## Destrory assets

if the assets has finished its misson(mainnet launch and destrory ERC20 token),Eiyaro wallet has support destrorying assets.

new transaction,choose advanced transaction

create a "Spend from account" action, fill in the assets you want to destrory

![](images/advanced_transaction_retire_assets.png)

create a "retire" action,choose what you want to destrory. the amount must equal to the input amount 

![](images/advanced_transaction_retire.png)

also don't forget to pay the fee

## Summary

we have introduced four actions of advanced transaction and three use scenarios, It is convenient and flexible.

but notice that it is just an ability, the assetss you regiestered are not certified by a professional institution,also have no legal effect, and no one will receive it if he can't trust you.so we must wait for professional application in Eiyaro blockchain.


