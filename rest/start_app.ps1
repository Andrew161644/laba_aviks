docker rm $(docker ps -aq) -f
docker volume prune -f
docker rmi rest_app
docker-compose up