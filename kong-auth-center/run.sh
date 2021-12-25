docker build -t kong-auth-image .
docker run -d --network=kong-net --name kong-auth-center -p 8081:8081 kong-auth-image:latest