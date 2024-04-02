# 1. Mining by Command Line Configuration

## Initialize

First of all, initialize the node:

```
$ cd ./cmd/eiyarod
$ ./eiyarod init --chain_id mainnet
```

There are three options for the flag `--chain_id`:

- `mainnet`: connect to the mainnet.
- `testnet`: connect to the testnet wisdom.
- `solonet`: standalone mode.

After that, you'll see `config.toml` generated, then launch the node.

## launch and mining

```
$ ./eiyarod node --mining
```

# 2. Minning by Api Configuration

## API Endpoint

Default JSON-RPC endpoints:

| Client | URL                                             |
| ------ | ----------------------------------------------- |
| Go     | [http://localhost:9888](http://localhost:9888/) |

A complete request example via `curl`:

## `set-mining`

Start up node mining.

#### Parameters

`Object`:

- `Boolean` - *is_mining*, whether the node is mining.

#### Example

```
// Request
curl -X POST set-mining -d '{"is_mining": true}'

// Result
```

### Dashboard

You can see the mining result in dashboard.

Access the dashboard:

```
$ open http://localhost:9888/
```
