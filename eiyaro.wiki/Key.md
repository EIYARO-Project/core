# Mnemonic

Eiyaro already implemented BIP39, for more information please visit: https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki

# Generate Root Private Key

- Input: 64 byte of seed, which can be random generate or calculate from mnemonic.
- Output: 64 byte of private key
```
$ xprv = hmac(sha512, seed, []byte{'R', 'o', 'o', 't'})
$ xprv[0] &= 248
$ xprv[31] &= 31
$ xprv[31] |= 64
```

# Private Key Calculate Public Key
- Input: 64 byte of private key
- Output: 64 byte of public key
```
$ xpub[:32] = edwards25519.GeScalarMultBase(xprv[:32])
$ xpub[32:] = xprv[32:]
```

# Derives the Child Private Key
- Input: 64 byte of private key, derive path byte array
- Output: 64 byte of child private key
```
$ xpub = xprv.XPub()
$ child_xprv = hmac(sha512, []byte{'N'} + xpub[:32] + derive_path, xpub[32:])
$ child_xprv[0] &= 248
$ child_xprv[29] &= 1
$ child_xprv[30] = 0
$ child_xprv[31] = 0
$
$ carry = 0
$ for i := 0; i < 32; i++ {
$     sum := int(xprv[i]) + int(child_xprv[i]) + carry
$     child_xprv[i] = byte(sum & 0xff)
$     carry = (sum >> 8)
$ }
```

# Derives the Child Public Key
- Input: 64 byte of public key, derive path byte array
- Output: 64 byte of child public key
```
$ child_xpub = hmac(sha512, []byte{'N'} + xpub[:32] + derive_path, xpub[32:])
$ child_xpub[0] &= 248
$ child_xpub[29] &= 1
$ child_xpub[30] = 0
$ child_xpub[31] = 0
$ f = edwards25519.GeScalarMultBase(child_xpub[:32])
$ p = edwards25519.Point(xpub[:32])
$ child_xpub[:32] = edwards25519.AddPoint(p, f)
```

# Expand Private Key for Sign
- Input: 64 byte of private key
- Output: 64 byte of expanded private key
```
$ expanded = hmac(sha512, xprv, []byte{'E', 'x', 'p', 'a', 'n', 'd'})
$ expanded[:32] = xprv[:32]
```