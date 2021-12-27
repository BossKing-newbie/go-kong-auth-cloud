docker stop kong-auth-center
docker rm kong-auth-center
docker rmi kong-auth-image
docker build -t kong-auth-image .
docker run -itd --network=kong-net --ip 172.19.0.3 --name kong-auth-center -p 8081:8081 kong-auth-image:latest