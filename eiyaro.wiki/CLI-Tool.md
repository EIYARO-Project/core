# Introduction to the command line tool 
`Eiyaro` provides `CLI` client (`eiyarocli`), it is the entry point into the `Eiyaro` network, capable of running as a full node archive node. 

It can be used by other processes as a gateway into the Eiyaro network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports.

The full `CLI` option can be viewed using `eiyarocli --help`, or the [eiyarocli Wiki](https://github.com/Eiyaro/eiyaro/wiki/Command-Line-Options) page.

# Specific command line options
Including accounts、keys、assets、token、transactions、blocks and other aspects.

- accounts

| Options | Meaning | Parameters | Examples |
| :------| :------ | :------ |:------ |
| create-account | create an new account | `<alias> <xpub(s)> [flags]` |   |
| list-accounts | list the existing accounts | `[flags]` |  |
| delete-account | delete the existing account | `<<accountID/alias> [flags]` |  |
| create-account-receiver | Create an account receiver | `<accountAlias> [accountID] [flags]` |  |
| list-addresses | list the account addresses| `[flags]` |  |
| validate-address | validate the account addresses | `<address> [flags]` |  |
| list-pubkeys | list the account pubkeys | `<accountInfo> [publicKey] [flags` |  |
| list-balances | list the accounts balances | `[flags]` |  |
| list-unspent-outputs | list the accounts unspent outputs | `[flags]` |  |

- keys

| Options | Meaning | Parameters | Examples |
| :------| :------ | :------ |:------ |
| create-key | create a key | `<alias> <password> [flags]` |   |
| list-keys | list the existing keys | `[flags]` |  |
| delete-key | delete a key | `<xpub> <password> [flags]` |  |
| reset-key-password | reset key password | `<xpub> <old-password> <new-password> [flags]` |  |
| check-key-password | check key password | `<xpub> <password> [flags]` |  |
| sign-message | sign message to generate signature | `<address> <message> <password> [flags]` |  |
| verify-message | verify signature for specified message | `<address> <xpub> <message> <signature> [flags]` |  |

- assets

| Options | Meaning | Parameters | Examples |
| :------| :------ | :------ |:------ |
| create-asset | create an asset | `<alias> <xpub(s)> [flags]` |   |
| get-asset | get asset by assetID | `<assetID> [flags]` |  |
| list-assets | list the existing assets | `[flags]` |  |
| update-asset-alias | update the asset alias | `<assetID> <newAlias> [flags]` |  |

- token

| Options | Meaning | Parameters | Examples |
| :------| :------ | :------ |:------ |
| create-access-token | create a new access token | `<tokenID> [flags]` |   |
| list-access-tokens | list the existing access tokens| `[flags]` |  |
| delete-access-token | delete an access token | `<tokenID> [flags]` |  |
| check-access-token | check an access token | `<tokenID> <secret> [flags]` |  |


- transactions

| Options | Meaning | Parameters | Examples |
| :------| :------ | :------ |:------ |
| build-transaction | build one transaction template,default use account id and asset id | `<accountID/alias> <assetID/alias> <amount>[outputID] [flags]` |   |
| sign-transaction | sign transaction templates with account password | `<json templates> [flags]` |  |
| submit-transaction | submit signed transaction | `<json templates> [flags]` |  |
| create-transaction-feed| create a transaction feed filter | `<alias> <filter> [flags]` |   |
| list-transaction-feeds| list all of transaction feeds | `[flags]` |   |
| delete-transaction-feed| delete a transaction feed filter | `<alias> [flags]` |   |
| get-transaction-feed| get a transaction feed by alias | `<alias> [flags]` |   |
| update-transaction-feed| update transaction feed | `<alias> <fiter> [flags]` |   |

- blocks

| Options | Meaning | Parameters | Examples |
| :------| :------ | :------ |:------ |
| get-block-hash | get the hash of most recent block | `[flags]` |   |
| get-block-count | get the number of most recent block | `[flags]` |  |
| get-block | get a whole block matching the given hash or height | `<hash> / <height> [flags]` |  |
| get-block-header | get the header of a block matching the given hash or height | `<hash> / <height> [flags]` |  |
| get-difficulty | get the difficulty of most recent block | `[flags]` |  |
| get-hash-rate | get the nonce of most recent block | `[flags]` |  |

- others

| Options | Meaning | Parameters | Examples |
| :------| :------ | :------ |:------ |
| is-mining | if client is actively mining new blocks | `[flags]` |   |
| set-mining | start or stop mining | `<true or false> [flags]` |  |
| net-info | print the summary of network | `[flags]` |  |
| decode-program | decode program to instruction and data | `<program> [flags]` |  |
| version | print the version number of Eiyarocli | `[flags]` |  |
| wallet-info | print the information of wallet | `[flags]` |  |
| rescan-wallet | trigger to rescan block information into related wallet | `[flags]` |  |


