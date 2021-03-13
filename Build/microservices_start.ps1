docker rm $(docker ps -aq) -f
docker volume prune -f
docker container prune -f
docker rmi grpc_service_app -f
cd ../grpc_service
docker-compose up