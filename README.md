# Readme

These scripts wrap the core functions of [Shamir’s Secret Sharing (SSS)](https://en.wikipedia.org/wiki/Shamir%27s_secret_sharing) into a little CLI. Uses [github.com/SSSaaS/sssa-golang](github.com/SSSaaS/sssa-golang) for implementation of SSS.

WARNING: Handling secrets on your file system may expose them to malware attacks. Exercise extreme caution in handling seed phrases and other critical secrets. The code in this repo is intended for practice and learning, and should not be introduced as a dependency into critical software environments.

Intended use-case: share-splitting blockchain seed phrases (mnemonics) and other critical secrests to prevent single-point of failure for losses and leaks.

## build docker images

In ./create_shares, run:
```shell
docker build -t create-shares-builder .
docker create --name temp-container create-shares-builder
docker cp temp-container:/app/create-shares ./create-shares
docker rm temp-container
```

In ./combine_shares, run:

```shell
docker build -t combine-shares-builder .
docker create --name temp-container combine-shares-builder
docker cp temp-container:/app/combine-shares ./combine-shares
docker rm temp-container
```

## usage

### Secret to shares

Split a secret into N hashed chunks, any M of which can be used to recover the secret.

Input (the secret) comes from `secret.txt`, output goes to `shares.txt`.

```shell
./secret_to_shares $SECRET  $M $N
```

E.g., split a secret from a file `secret.txt` into three chunks, any 2 of which can be used for recovery:

```
./create-shares "$(<secret.txt)"  2 3
```

```console
1VGL8lfsV80a941rBXMsusmrWpcHwLXwxjnAUb9iYd8=9yT6O4M6EtrVx-uL58yIbVtAFjXPghiMPDHtJz8YiQk=Pr-MJJcdGxaKczap_AETlGV70c7W2HkTZBg5Egv_S9U=-vnSA9gRZel101GgeEHTw_EtTF0tZigEpOo29r2a6ZE=0AR-pd4adDDVkVMUQGs4d-Ku5w_4gbR73w3E-ViaPqg=YEPdKJd-ca5DFz_yayub3NiBnqQozYv1tDbvZMUhhB8=
wbRiYAchSvtTMG74vH3WqkGbIMIODyPOzaFDWTU_zXA=cYGjH5rJt-hzMYa9rTi4SiXsvTTcQ8e6uwaP9OCb_5c=JTiaNuB62f8a7C5g0KzPTp9Drq6aobhLYNa1B1-iLbU=CCNmqGjQLJfbNpOMInePYHrcGlLxrMRqIqflk1a1F3U=Q14sZLw3jU85UP8iwVoHPv16IJ9EtSV67oME1CegIpo=_OpdTwvdx0a4BeuXR7cSYX8j3z7sFJsdilU6F4Os4ac=
h90J9KyqWJqCoA7wH6bjfjL1rnmu__X0d-y8fPM7zII=LG_OpmZcf3De_KRVL73J5xXmZf-Ki2IUutv1ODSOE6Y=KQVOEvSZ0xJVDlJUrUVis4-_aU7FomIlsfj8dqmVJd8=nY-z8LzNorDofCKK3_vwM3WMYr-86CPDtu4ZgDTkNEI=LHDtUocPGIlE36J-TYytL9GqzGVkwiq_Zxn5kEu-0yc=CXClNnuu4qNatzm14w9R778BB-FNPW259T7hbmttpCU=

```
## Shares to secret

Regenerate the secret from a set of shares. Will fail unless M valid chunks are on subsequent lines of `shares.txt`. Secret will be output to stdout. 


```shell
./combine-shares shares.txt
```
```console
Recovered secret: congratulations! you decoded the message. you have mastered easyshamirencodeCLI!
```
