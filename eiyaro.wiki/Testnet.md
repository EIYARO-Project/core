# Run Testnet

## Initialize

First of all, initialize the node:

```
$ cd ./cmd/eiyarod
$ ./eiyarod init --chain_id testnet --root "/Users/sheng/Library/Eiyaro_testnet"
```

There are three options for the flag `--chain_id`:

- `mainnet`: connect to the mainnet.
- `testnet`: connect to the testnet wisdom.
- `solonet`: standalone mode.

After that, you'll see `config.toml` generated, then launch the node.

## Default root path

- MacOS(`darwin`)：`~/Library/Eiyaro`
- Windows(`windows`): `~/AppData/Roaming/Eiyaro`
- Others（eg.Linux）：`~/.eiyaro`

**use `--root` parameter to customize root path.**

## launch

```
$ ./eiyarod node --mining --root "/Users/sheng/Library/Eiyaro_testnet"
```

available flags for `eiyarod node`:

```
      --auth.disable                Disable rpc access authenticate
      --chain_id string             Select network type
  -h, --help                        help for node
      --mining                      Enable mining
      --home		   				root directory for config and data
  -r, --root   						DEPRECATED. Use --home (default "/Users/sheng/Library/Eiyaro")
      --p2p.dial_timeout int        Set dial timeout (default 3)
      --p2p.handshake_timeout int   Set handshake timeout (default 30)
      --p2p.laddr string            Node listen address.
      --p2p.max_num_peers int       Set max num peers (default 50)
      --p2p.pex                     Enable Peer-Exchange  (default true)
      --p2p.seeds string            Comma delimited host:port seed nodes
      --p2p.skip_upnp               Skip UPNP configuration
      --prof_laddr string           Use http to profile eiyarod programs
      --vault_mode                  Run in the offline enviroment
      --wallet.disable              Disable wallet
      --wallet.rescan               Rescan wallet
      --web.closed                  Lanch web browser or not
```

Given the `eiyarod` node is running, the general workflow is as follows:

- create key, then you can create account and asset.
- send transaction, i.e., build, sign and submit transaction.
- query all kinds of information, let's say, avaliable key, account, key, balances, transactions, etc.

## Dashboard

You can see the result in dashboard.

Access the dashboard:

```
$ open http://localhost:9888/
```

