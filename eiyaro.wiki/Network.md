There are mainly the following network types than the original chain network, which have different uses:

- mainnet main network: the currency of circulation is ey with real value.
- testnet test network: The test network for developers to conduct transactions or smart contracts.
- solonet test network: single-node network, not connected to other nodes, equivalent to private chain.
- National Secret Test Network: A network that uses the national secret encryption algorithm.

This article will focus on the methods of compiling and starting each network, and compare the similarities and differences between several networks.

## 1 Compile and start the network

First of all, enter the eiyaro working directory and compile:

```
$ cd $GOPATH/src/github.com/eiyaro
$ make eiyarod
```

Initialize the network:

```
$ cd cmd/eiyarod
$ ./eiyarod init --chain_id <net-type> --home <config_and_data_path>
```

`--chain_id` specifies the type of network to start. There are three network types available:

- mainnet：main network
- testnet：wisdom test network
- solonet：single test network

If you do not specify a network type, the solonet is started by default.

`--home` specifies the path where the configuration and data files are stored. If you do not specify a storage path, the default path for different platforms is:

- MacOS(darwin): ~/Library/Application Support/Eiyaro
- Windows(windows): ~/AppData/Roaming/Eiyaro
- Others(eg.Linux): ~/.eiyaro

For example, to start solonet, specify the configuration and data file storage directory as `$HOME/eiyaro/solonet`:

```
$ ./eiyarod init --chain_id solonet --home $HOME/eiyaro/solonet
```

The corresponding configuration file `config.toml` will appear in the `$HOME/eiyaro/solonet` folder:

```
$ cat $HOME/eiyaro/solonet/config.toml
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml
fast_sync = true
db_backend = "leveldb"
api_addr = "0.0.0.0:9888"
chain_id = "solonet"
[p2p]
laddr = "tcp://0.0.0.0:46658"
seeds = ""
```

After initializing the network, start the network:

```
$ ./eiyarod node --mining --home <config_and_data_path>
```

The `--mining` option specifies that mining is enabled when the node is started. If the network type specified during the initial initialization is mainnet or testnet, there is no need to open mining. If you specify solonet, it is highly recommended to start mining, otherwise it will not be packaged into the block after the transaction test.

Take the startup solonet node as an example:

```
$ ./eiyarod node --mining --home $HOME/eiyaro/solonet
```

## 2 main network

Eiyaro main net.

## 3 wisdom test network

The name of the testnet is the wisdom, and has the ability to communicate with other nodes. When developers develop smart contracts, they need to deploy tests on the wisdom test network before they are officially released to the main network. The function of the wisdom test network is consistent with the function of the main network. The debug results obtained by the developer on the wisdom test network will also be the same as those of the main network.

There are several ways for developers to obtain the test currency of the consumer:

- Turn on the `--mining` option when starting the node. The current test power of the test network is about 1~2KHash/s. It is difficult for a normal developer to use a personal computer to dig up test coins unless a professional than a raw ore is used.
- Contact to obtain test coins from Eiyaro technology operations.
- Test coins can be obtained from the <https://matpool.io> website.

## 4 solonet test network

The solonet test network is essentially a single-node network, does not connect with neighboring nodes, and its initial difficulty value is much lower than the main network and the wisdom test network. At present, the initial speed of the solonet is about 10 seconds, which ensures that developers can get network running results faster. Developers can first open the solonet for single-node mining, test their smart contracts, then go to the same test network as the main network environment for formal testing, and finally deploy the actual smart contract on the Eiyaro main network.

Because the solonet test network is a single-node network, this node can obtain all the mining revenue as the only miner.

## 5 Several network differences

Individual network address prefixes:

network | Address prefix
--------|---------------
mainnet | bm
testnet | tm
solonet | sm