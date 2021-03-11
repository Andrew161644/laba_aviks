docker rm $(docker ps -aq) -f
docker volume prune -f
docker container prune -f
docker rmi rest_app
cd ../rest
docker-compose up