docker rm $(docker ps -aq) -f
docker volume prune -f
docker container prune -f
docker rmi full_app_rest_service
docker rmi full_app_neiro_service
docker rmi full_app_currency_service
cd full_app
docker-compose up