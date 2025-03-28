# build
docker build -t recover-from-slip39-shares-builder .
docker create --name temp-container recover-from-slip39-shares-builder
docker cp temp-container:/app/recover-from-slip39-shares ./recover-from-slip39-shares

# usage
# requires the shares to be in shares/
# each share must be on a single line in a .txt file

./recover-from-slip39-shares