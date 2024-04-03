# Set Network

first you must install and run Eiyaro

the detail please ref to [build and install](https://github.com/Eiyaro/eiyaro/wiki/Build-and-Install)

## Choose a Network

when you run Eiyaro_Wallet.exe，you may see the user interface like below

![](images/setting-network.png)

there two networks, one is mainnet,the other is Testnet

**mainnet:** the formal network if Eiyaro network publish , before this it just a single node network,you can test mining,but you can not transfer because you can not connect any other node

**testnet:** the test network,it will be used for test feature and help to find bug.the testnet include many nodes,you can connect to other node and transfer to other account

we choose testnet

## Block Synchronization

When you choose the network and starts to run the node, it will synchronize all the blocks first
Before the synchronization of last block you cannot trade ( even is someone send you coins, you can’t see them) either the mining process cannot be made.

you will see the information below

![](images/network_syc_status.png)

after the label "Synchronizing",there are two numbers:the first one means current synchronized block,the next one means the highest block.
If the numbers are not the same, the node is downloading the blockchain and synchronizing,or your network is synchronized.

after synchronized,you can use all features, have fun!
