# build
docker build -t create-slip-39-shares-builder .
docker create --name temp-container create-slip-39-shares-builder
docker cp temp-container:/app/create-slip39-shares ./create-slip39-shares
docker rm temp-container

# usage
# dumps the shares into ./shares/
# mkdir shares if you need to

secret="Blahblahblah bleep blooop fleep floop bleep blooop fleep floop bleep blooop fleep floop." passphrase="asdf" threshold="4" total_shares="6" ./create-slip39-shares/create-slip39-shares