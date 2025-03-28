# Readme

These scripts wrap the core functions of [Shamir’s Secret Sharing (SSS)](https://en.wikipedia.org/wiki/Shamir%27s_secret_sharing) into a little CLI. Uses [github.com/SSSaaS/sssa-golang](github.com/SSSaaS/sssa-golang) for implementation of SSS.

WARNING: Handling secrets on your file system may expose them to malware attacks. Exercise extreme caution in handling seed phrases and other critical secrets. The code in this repo is intended for practice and learning, and should not be introduced as a dependency into critical software environments.

Intended use-case: share-splitting blockchain seed phrases (mnemonics) and other critical secrets to avoid single-points of failure for losses and leaks.


# Slip39

with slip 39 the shares are represented as series of english words, via the [slip39 wordslist](https://github.com/satoshilabs/slips/blob/master/slip-0039/wordlist.txt)

## usage

### create shares

will spit shares into ./shares/ 
mkdir this if you need to


```shell
secret="I am a really cool top secret secret" passphrase="asdf" threshold="4" total_shares="6" ./create-slip39-shares/create-slip39-shares
```

### recovery

threshold number of shares must be in shares/
```shell
ls shares
```
```console
3.txt	4.txt	5.txt	6.txt
```
```shell
 cat shares/3.txt
```
```console
inherit typical academic ambition carve talent fawn fridge usual frost fangs shaped sugar surface decent quick auction slim treat idea flame frozen move huge client craft mandate hamster black science either prune lunar freshman elite deal
```

note that recovery will fail if too many or too few shares are present

```shell
passphrase="asdf" ./recover-from-slip39-shares/recover-from-slip39-shares
```
```console
Recovered secret bytes:
[73 32 97 109 32 97 32 114 101 97 108 108 121 32 99 111 111 108 32 116 111 112 32 115 101 99 114 101 116 32 115 101 99 114 101 116]
Recovered secret (hex):
4920616d2061207265616c6c7920636f6f6c20746f702073656372657420736563726574
```
```shell
passphrase="asdf" ./recover-from-slip39-shares/recover-from-slip39-shares  | xxd -r -p
```
```console
�s2��)s!�!2��!2Q�Ac!�I am a really cool top secret secret%
```
