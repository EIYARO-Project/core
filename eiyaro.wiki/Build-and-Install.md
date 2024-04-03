## Install from release

## Install from source

### Requirements

- [Go](https://golang.org/doc/install) version 1.8 or higher, with `$GOPATH` set to your preferred directory

### Installation

Ensure Go with the supported version is installed properly:

```
$ go version
$ go env GOROOT GOPATH
```

#### Get the source code

```
$ git clone https://github.com/Eiyaro/ey.git $GOPATH/src/github.com/eiyaro
```

#### Build

```
$ cd $GOPATH/src/eiyaro/ey
$ go mod tidy
$ make eiyarocd    
$ make eiyaroccli  
```

When successfully building the project, the `eiyaro` and `eiyarocli` binary should be present in `cmd/eiyarod` and `cmd/eiyarocli` directory, respectively.

## Running in Docker

### Build the image
```bash
docker build -t eiyaro .
```

### Init eiyaro
```bash
docker run -v <Eiyaro/data/directory/on/host/machine>:/root/.eiyaro eiyaro:latest eiyarod init --chain_id <chainId>
```

The default Eiyaro data directory (on the host) is:

+ Mac: `~/Library/Eiyaro`
+ Linux: `~/.eiyaro` 
+ Windows: `%APPDATA%\Eiyaro`

### Enter the iterative mode
```bash
docker run -it -p 9888:9888 -v <Eiyaro/data/directory/on/host/machine>:/root/.eiyaro eiyaro:latest
```

Then you can use eiyarod and eiyarocli following [Readme](https://github.com/Eiyaro/eiyaro/blob/master/README.md)

Use `exit` to exit Docker's iterative mode

### Daemon mode
For example,
```bash
docker run -d -p 9888:9888 -v <Eiyaro/data/directory/on/host/machine>:/root/.eiyaro eiyaro:latest eiyarod node --web.closed --auth.disable
```

__To list the running containners and check their container id, image, corresponding command, created time, status, name and ports being used:__
```bash
docker container ls
```
or 
```bash
docker ps
```

__To execute a command inside a containner, for example:__
```bash
docker exec -it <containerId> eiyarocli create-access-token <tokenId>
```

__To stop a running containner:__
```bash
docker stop <containerId>
```

__To remove a containner:__
```bash
docker rm <containerId>
```
