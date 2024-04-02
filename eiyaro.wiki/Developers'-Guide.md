
This document is the entry point for developers of the Go implementation of Eiyaro. 

## Installation Instructions

If you just want to run eiyaro, use the normal [Installation Instructions](https://github.com/Eiyaro/eiyaro/wiki/Build-and-Install)

## Building and Testing

### Go Environment

We assume that you have [`go` v1.8 or higher installed](https://golang.org/doc/install), and `GOPATH` is set.

**Note**:You must have your working copy under `$GOPATH/src/github.com/eiyaro`.

Since `go` does not use relative path for import, working in any other directory will have no effect, since the import paths will be appended to `$GOPATH/src`, and if the lib does not exist, the version at master HEAD will be downloaded.

Most likely you will be working from your fork of `eiyaro`, let's say from `github.com/nirname/eiyaro`. Clone or move your fork into the right place:

```
git clone git@github.com:nirname/eiyaro.git $GOPATH/src/github.com/eiyaro
```

### Building Executables
Switch to the eiyaro repository root directory.
```bash
$ cd $GOPATH/src/github.com/eiyaro
$ make eiyarod    # build eiyarod
$ make eiyarocli  # build eiyarocli
```
When successfully building the project, the `eiyaro` and `eiyarocli` binary should be present in `cmd/eiyarod` and `cmd/eiyarocli` directory, respectively.

#### Initialize

First of all, initialize the node:

```bash
$ cd ./cmd/eiyarod
$ ./eiyarod init --chain_id testnet
```

There are two options for the flag `--chain_id`:

- `testnet`: connect to the testnet.
- `mainnet`: standalone mode.

After that, you'll see `.eiyarod` generated in current directory, then launch the node.

#### launch

``` bash
$ ./eiyarod node --mining
```

available flags for `eiyarod node`:

```
      --auth.disable                Disable rpc access authenticate
      --mining                      Enable mining
      --p2p.dial_timeout int        Set dial timeout (default 3)
      --p2p.handshake_timeout int   Set handshake timeout (default 30)
      --p2p.laddr string            Node listen address.
      --p2p.max_num_peers int       Set max num peers (default 50)
      --p2p.pex                     Enable Peer-Exchange
      --p2p.seeds string            Comma delimited host:port seed nodes
      --p2p.skip_upnp               Skip UPNP configuration
      --prof_laddr string           Use http to profile eiyarod programs
      --wallet.disable              Disable wallet
      --web.closed                  Lanch web browser or not
```

Given the `eiyarod` node is running, the general workflow is as follows:

- create key, then you can create account and asset.
- send transaction, i.e., build, sign and submit transaction.
- query all kinds of information, let's say, avaliable key, account, key, balances, transactions, etc.

### Testing

Testing one library:

```
go test -v ./account 
```

Using options `-v` (logging even if no error) is recommended.

Testing only some methods:

```
go test -v ./account -run TestCreateAccount
```

**Note**: here all tests with prefix _TestMethod_ will be run, so if you got TestMethod, TestMethod1, then both!

**Running benchmarks**, eg.:
switch to test directory.
```bash
cd $GOPATH/src/github.com/eiyaro/test
go test -v -bench=. -benchtime=3s -run=none
```

Using options `-bench` to specify test directory, and `-benchtime` to specify the test time.

for more see [go test flags](http://golang.org/cmd/go/#hdr-Description_of_testing_flags)



## Contributing

Thank you for considering to help out with the source code! We welcome contributions from
anyone on the internet, and are grateful for even the smallest of fixes!

GitHub is used to track issues and contribute code, suggestions, feature requests or
documentation.

If you'd like to contribute to eiyaro, please fork, fix, commit and send a pull
request (PR) for the maintainers to review and merge into the main code base. 

PRs need to be based on and opened against the `master` branch (unless by explicit
agreement, you contribute to a complex feature branch).

We encourage a PR early approach, meaning you create the PR the earliest even without the
fix/feature. This will let core devs and other volunteers know you picked up an issue.
These early PRs should indicate 'in progress' status.
