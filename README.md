
# Readme

These scripts wrap the core functions of [Shamir's Secret Sharing (SSS)]() into a little cli.

Uses github.com/SSSaaS/sssa-golang for implement of SSS algo.

Intended use-case: share-splitting seed phrases (mnemonics) for Bittensor coldkeys.

## usage

## Secret to shares

```shell
./secret_to_shares $M $N
```
Split a secret into N hashed chunks, any M of which can be used to recover the secret.

Input (the secret) comes from `secret.txt`, output goes to `shares.txt`.

E.g., split the secret into three chunks, any 2 of which can be used for recovery:

```
./secret_to_shares 2 3
```

## Shares to secret

Regenerate the secret from a set of shares. Will fail unless M valid chunks are on subsequent lines of `shares.txt`. Secret will be output to stdout. 






