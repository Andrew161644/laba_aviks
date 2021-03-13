docker rm $(docker ps -aq) -f
docker volume prune -f
docker image prune -f -a