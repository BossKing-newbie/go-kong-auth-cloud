docker stop kong-sys-center
docker rm kong-sys-center
docker rmi kong-sys-image
docker build -t kong-sys-image .
docker run -itd --network=kong-net --ip 172.19.0.4 --name kong-sys-center -p 8087:8087 kong-sys-image:latest