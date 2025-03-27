
## usage

### Split secret into shares

docker run --rm go-shamir "holy cow this thing is pretty great i bet you will never be able to steal my secrets now hahahahaha" 3 7 > shares.txt



docker run --rm -v $(pwd):/data share-combiner /data/shares.txt



