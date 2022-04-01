export COMPOSE_DOCKER_CLI_BUILD=1 && export DOCKER_BUILDKIT=1

docker-compose --profile bpi build

docker-compose --profile infra up -d

docker exec -it kafka /opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server kafka:9092 --replication-factor 1  --partitions 1 --topic orgReg
docker exec -it kafka /opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server kafka:9092 --replication-factor 1  --partitions 1 --topic workgroupReg
docker exec -it kafka /opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server kafka:9092 --replication-factor 1  --partitions 1 --topic workflowSyncFPL
docker exec -it kafka /opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server kafka:9092 --replication-factor 1  --partitions 1 --topic workflowSyncACK


docker-compose --profile bpi up -d

docker-compose --profile frontend build

docker-compose --profile frontend up -d
