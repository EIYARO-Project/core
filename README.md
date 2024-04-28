# EIYARO Core

[![Build](https://github.com/EIYARO-Project/core/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/EIYARO-Project/core/actions)
[![Supports Windows](https://img.shields.io/badge/support-Windows-blue?logo=Windows)](https://github.com/EIYARO-Project/core/releases/latest)
[![Supports Linux](https://img.shields.io/badge/support-Linux-yellow?logo=Linux)](https://github.com/EIYARO-Project/core/releases/latest)
[![License](https://img.shields.io/github/license/EIYARO-Project/core)](https://github.com/EIYARO-Project/core/blob/master/LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/EIYARO-Project/core?label=latest%20release)](https://github.com/EIYARO-Project/core/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/EIYARO-Project/core/total)](https://github.com/EIYARO-Project/core/releases)
[![EIYARO Project Community](https://img.shields.io/discord/1217075571528564736?label=EIYARO%20Project%20Discord&logo=discord)](https://discord.gg/W5eBcfbMjM)

**Official golang implementation of the EIYARO protocol.**

Automated builds are available for stable releases and the unstable master branch. Binary archives are published at https://github.com/EIYARO-Project/core/releases.

## What is EIYARO?

The goal of the EIYARO project is to create a blockchain ecosystem with a wide range of application scenarios, and provide users with safe, efficient and convenient digital asset transaction services. We will reduce transaction costs, increase transaction speed, and ensure network security and stability through continuous technological innovation and optimisation.

In the current state `eiyaro` is able to:

- Manage key, account as well as asset
- Send transactions, i.e., issue, spend and retire asset


## Building from source

### Requirements

- [Go](https://golang.org/doc/install) version 1.18 or higher, with `$GOPATH` set to your preferred directory

### Installation

Ensure Go with the supported version is installed properly:

```console
$ go version
$ go env GOROOT GOPATH
```

- Get the source code

```console
$ git clone https://github.com/EIYARO-Project/core.git $GOPATH/src/eiyaro/core
```

- Build source code

```console
$ cd $GOPATH/src/eiyaro/core
$ go mod tidy
$ make eiyarod    
$ make eiyarocli  
```

When successfully building the project, the `eiyarod` and `eiyarocli` binary should be present in `cmd/eiyarod` and `cmd/eiyarocli` directory, respectively.

### Executables

The EIYARO project comes with several executables found in the `cmd` directory.

| Command      | Description                                                  |
| ------------ | ------------------------------------------------------------ |
| **eiyarod**   | eiyarod command can help to initialize and launch eiyaro domain by custom parameters. `eiyarod --help` for command line options. |
| **eiyarocli** | Our main EIYARO CLI client. It is the entry point into the EIYARO network (main-, test- or private net), capable of running as a full node archive node (retaining all historical state). It can be used by other processes as a gateway into the EIYARO network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `eiyarocli --help` and the [eiyaroccli API page](https://github.com/EIYARO-Project/core/blob/main/API-Reference.md) for command line options. |

## Running eiyaro

Currently, eiyaro is still in active development and a ton of work needs to be done, but we also provide the following content for these eager to do something with `eiyaro`. This section won't cover all the commands of `eiyarod` and `eiyarocli` at length, for more information, please the help of every command, e.g., `eiyarocli help`.

### Initialize

First of all, initialize the node:

```console
$ cd ./cmd/eiyarod
$ ./eiyarod init --chain_id mainnet
```

There are three options for the flag `--chain_id`:

- `mainnet`: connect to the mainnet.
- `testnet`: connect to the testnet wisdom.
- `solonet`: standalone mode.

After that, you'll see `config.toml` generated, then launch the node.

### Proposed Configuration File

We have seen good performance with this configuration file:

```toml
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml
fast_sync = true
db_backend = "leveldb"
api_addr = "0.0.0.0:9888"
node_alias = "MyMostAwesomeNodeAlias"
moniker = "MyMostAwesomeNode"
chain_id = "mainnet"
[p2p]
laddr = "tcp://0.0.0.0:46657"
lan_discoverable = false
seeds = "148.135.99.102:46657,217.194.133.61:46657,154.201.78.187:46657,103.115.46.201:46657,154.44.8.62:46657,24.233.3.133:46657,27.25.156.254:46657"
keep_dial = "148.135.99.102:46657,217.194.133.61:46657,154.201.78.187:46657,103.115.46.201:46657,154.44.8.62:46657,24.233.3.133:46657,27.25.156.254:46657"
max_num_peers = 50
[web]
closed = true
```

The added field `keep_dial` allows for a more stable connection to the `mainnet`.

### Launch

```console
$ nohup ./eiyarod node &
```

available flags for `eiyarod node`:

```
Flags:
      --auth.disable                     Disable rpc access authenticate
      --chain_id string                  Select network type
  -h, --help                             help for node
      --log_file string                  Log output file (default "log")
      --log_level string                 Select log level(debug, info, warn, error or fatal)
      --mining                           Enable mining
      --p2p.dial_timeout int             Set dial timeout (default 3)
      --p2p.handshake_timeout int        Set handshake timeout (default 30)
      --p2p.keep_dial string             Peers addresses try keeping connecting to, separated by ',' (for example "1.1.1.1:46657;2.2.2.2:46658")
      --p2p.laddr string                 Node listen address. (0.0.0.0:0 means any interface, any port) (default "tcp://0.0.0.0:46656")
      --p2p.lan_discoverable             Whether the node can be discovered by nodes in the LAN (default true)
      --p2p.max_num_peers int            Set max num peers (default 50)
      --p2p.node_key string              Node key for p2p communication
      --p2p.proxy_address string         Connect via SOCKS5 proxy (eg. 127.0.0.1:1086)
      --p2p.proxy_password string        Password for proxy server
      --p2p.proxy_username string        Username for proxy server
      --p2p.seeds string                 Comma delimited host:port seed nodes
      --p2p.skip_upnp                    Skip UPNP configuration
      --prof_laddr string                Use http to profile eiyarod programs
      --simd.enable                      Enable SIMD mechan for tensority
      --vault_mode                       Run in the offline enviroment
      --wallet.disable                   Disable wallet
      --wallet.rescan                    Rescan wallet
      --wallet.txindex                   Save global tx index
      --web.closed                       Lanch web browser or not
      --ws.max_num_concurrent_reqs int   Max number of concurrent websocket requests that may be processed concurrently (default 20)
      --ws.max_num_websockets int        Max number of websocket connections (default 25)

Global Flags:
      --home string   root directory for config and data
  -r, --root string   DEPRECATED. Use --home (default "/Users/zcc/Library/Application Support/EIYARO")
      --trace         print out full stack trace on errors
```

Given the `eiyarod` node is running, the general workflow is as follows:

- create key, then you can create account and asset.
- send transaction, i.e., build, sign and submit transaction.
- query all kinds of information, let's say, avaliable key, account, key, balances, transactions, etc.

__simd feature:__

You could enable the _simd_ feature to speed up the _PoW_ verification (e.g., during mining and block verification) by simply:
```console
$ eiyarod node --simd.enable
```

To enable this feature you will need to compile from the source code by yourself, and `make eiyarod-simd`. 

What is more,

+ if you are using _Mac_, please make sure _llvm_ is installed by `brew install llvm`.
+ if you are using _Windows_, please make sure _mingw-w64_ is installed and set up the _PATH_ environment variable accordingly.

For more details about using `eiyarocli` command please refer to [API Reference](https://github.com/EIYARO-Project/core/blob/main/API-Reference.md)

### Launching using `systemd`

#### Under root

```ini
[Unit]
Description=EIYARO Node
After=syslog.target
After=network.target

[Service]
RestartSec=2s
Type=simple
User=USERNAME
Group=GROUP
WorkingDirectory=/home/USERNAME
ExecStart=/home/USERNAME/bin/eiyarod node
Restart=always
PrivateTmp=true
ProtectSystem=full
NoNewPrivileges=true

[Install]
WantedBy=multi-user.target
```

Drop the above file under `/etc/systemd/system/eiyarod.service`

To enable it:
```console
systemctl enable eiyarod.service
```

To start it:
```console
systemctl start eiyarod.service
```

To query it's status:
```console
systemctl status eiyarod.service
```

#### Under user

```ini
[Unit]
Description=EIYARO Node
After=syslog.target
After=network.target

[Service]
RestartSec=2s
Type=simple
WorkingDirectory=/home/USERNAME
ExecStart=/home/USERNAME/bin/eiyarod node
Restart=always

[Install]
WantedBy=default.target
```


Drop the above file under `~/.config/systemd/user/eiyarod.service`

To enable it:
```console
systemctl --user enable eiyarod.service
```

To start it:
```console
systemctl --user start eiyarod.service
```

To query it's status:
```console
systemctl --user status eiyarod.service
```

### Dashboard

Copy and save your tokename, tokename is used to log into your node page Access the dashboard:

```console
$ cd ./cmd/eiyarocli
$ ./eiyarocli create-access-token eiyaro

$ open http://localhost:9888/ OR Login with your IP + 9888 port
```

### In Docker

Ensure your [Docker](https://www.docker.com/) version is 17.05 or higher.

```console
$ docker build -t eiyaro .
```

For the usage please refer to [running-in-docker-wiki](https://github.com/EIYARO-Project/core/wiki/Running-in-Docker).


## Installing with Homebrew

```console
$ brew tap eiyaro/eiyaro && brew install eiyaro
```


## Contributing

Thank you for considering helping out with the source code! Any contributions are highly appreciated, and we are grateful for even the smallest of fixes!

If you run into an issue, feel free to [eiyaro issues](https://github.com/EIYARO-Project/core/issues/) in this repository. We are glad to help!

## License

[AGPL v3](./LICENSE)
