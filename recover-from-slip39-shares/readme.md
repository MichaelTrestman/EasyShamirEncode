docker build -t recover-from-slip39-shares-builder .
docker create --name temp-container recover-from-slip39-shares-builder
docker cp temp-container:/app/recover-from-slip39-shares ./recover-from-slip39-shares