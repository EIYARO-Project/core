## Download

Download [the latest release](https://github.com/Eiyaro/eiyaro/releases/tag/v0.3.0) according to your system:

![download](http://upload-images.jianshu.io/upload_images/127313-1008ae62a5219782.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## Initialize

Open *terminal* for macOS and Linux users. If using windows, you could use *cmd* or *powershell*.

enter the directory where the latest eiyaro application are downloaded, run the following command:

```bash
$ ./eiyarod init --chain_id testnet
```
![eiyarod init --chain_id testnet](http://upload-images.jianshu.io/upload_images/127313-0a5610cd630d46a4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

`--chain_id testnet` means to connect to the public testnet. If you use `--chain_id mainnet`, it would be standalone mode.

##  Launch the node

```bash
$ ./eiyarod node --mining
```

![eiyarod node --wallet.enable --mining](http://upload-images.jianshu.io/upload_images/127313-97693a07ac82c06c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Then you could access the dashboard by visiting `http://localhost:9888` via your brower.

## New key

Whither creating a new account or asset, you have to create a key first, which is essential for asset and account.

![new key](http://upload-images.jianshu.io/upload_images/127313-4f333e8e172b7be7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Let's say, new a key naming `alice`:

![key alias: alice](http://upload-images.jianshu.io/upload_images/127313-b0af950643a847bd.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![alice](http://upload-images.jianshu.io/upload_images/127313-822c3ecff0a29938.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## New account

*Accounts* -> *New*, choose an account alias, e.g., *alice*:

![account alias: alice](http://upload-images.jianshu.io/upload_images/127313-d3e49eb21b053b81.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Select an avaliable key for account *alice*:

![choose a key for account alice](http://upload-images.jianshu.io/upload_images/127313-a62b4b964c4153d9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![created account alice](http://upload-images.jianshu.io/upload_images/127313-dd925471a390bf93.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## New asset

Let's new an asset naming *gold*:

![asset alias](http://upload-images.jianshu.io/upload_images/127313-2d87c9bf7dd66092.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Choose a key as well.

![choose key for asset gold](http://upload-images.jianshu.io/upload_images/127313-7d14d6eb052fc909.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![created asset gold](http://upload-images.jianshu.io/upload_images/127313-f054c81879f857d9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## Send Transaction

Before sending a transaction, ensure you have enough balances:

- every transaction needs to consume at least 20,000,000 EY gas.

![balances](http://upload-images.jianshu.io/upload_images/127313-e7f0c07680c19a1a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### Transfer EY

This section demonstrates the transaction *alice* sending 100000000 EY to *bob*.

![new transaction](http://upload-images.jianshu.io/upload_images/127313-f3c766ff82f633c5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

There are several types of actions. 

- `Spend from account`: **sender** of transaction.
- `Control with account` or `Control with receiver`: **receiver** of transaction.

Note: The prompter of trading should pay the transaction fee, i.e., at least 20,000,000 EY gas.

![spend from account](http://upload-images.jianshu.io/upload_images/127313-41187e1ad4b18fe5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Sender(*alice*) pays the EY as transaction fee:

![spend from account: EY gas](http://upload-images.jianshu.io/upload_images/127313-b806fad17294d557.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Sender(*alice*) wants transfer 100000000 EY to receiver:

![spend from account: alice, ey, 100000000](http://upload-images.jianshu.io/upload_images/127313-e5b9a766e3dc7551.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Configure receiver(*bob*) accordingly:

![control with account: bob, ey, 100000000](http://upload-images.jianshu.io/upload_images/127313-bf9428b1e2c2ec2d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Now, bob should have 100000000 EY.

![check out the balances](http://upload-images.jianshu.io/upload_images/127313-4993027a878eece0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### Issue and Spend Other Assets

Aside from transferring EY, which is a native asset in the system, you can issue some new assets in Eiyaro, e.g., gold. Whether issuing or spending an asset, you have to build a transaction to complete, which is similar to the above instruction.

 
