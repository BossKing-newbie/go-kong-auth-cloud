docker rmi kong-auth-image
docker build -t kong-auth-image .
docker run -itd --network=kong-net --network-alias kongAuthNet --name kong-auth-center -p 8081:8081 kong-auth-image:latest