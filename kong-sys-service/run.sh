docker rmi kong-sys-image
docker build -t kong-sys-image .
docker run -itd --network=kong-net --network-alias kongSysNet --name kong-sys-center -p 8087:8087 kong-sys-image:latest