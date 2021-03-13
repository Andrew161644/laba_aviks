docker rm $(docker ps -aq) -f
docker volume prune -f
cd local_scripts
docker-compose up