# nodes

A Eiyaro P2P network has different kinds of nodes with different requirements for connectivity to one another.
This document describes what kind of nodes should enable and how they should work.

## Seeds

Seeds are the first point of contact for a new node.
The new node discovers other nodes in the network through seeds.

## New Full Node

The full node saves the complete data on the chain.
A new node needs a few things to connect to the network:
- a list of seeds, which can be provided to Eiyaro via config file or command flags.
- a `ChainID`, also called `Network` at the p2p layer.
- a recent block height, H, and hash, HASH for the blockchain.

With the above, the node then queries some peers from the discovery service, dials those peers, and runs the net    
 sync protocols with those it successfully connects to.

## Vault Mode Node

Node runs in an offline environment. Can be used to sign a transaction, but can't broadcast a transaction.
