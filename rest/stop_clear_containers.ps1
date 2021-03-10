docker rm $(docker ps -aq) -f
docker volume prune -f -a