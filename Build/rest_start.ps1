docker rm $(docker ps -aq) -f
docker volume prune -f
docker container prune -f
docker rmi rest_app_rest_service
cd rest_app
docker-compose up