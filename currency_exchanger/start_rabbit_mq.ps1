docker run -d -p 5672:5672 rabbitmq:3
docker run -d --hostname localhost --name some-rabbit -p 8085:15672 rabbitmq:3-management