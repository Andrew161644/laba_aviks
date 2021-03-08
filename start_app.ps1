docker rm $(docker ps -aq) -f
docker volume prune -f
docker rmi aics_project_app
docker-compose up