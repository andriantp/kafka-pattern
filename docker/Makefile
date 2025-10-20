# Nama file: Makefile
COMPOSE=docker compose

# Kafka
kafka-up:
	@echo "🚀 Starting Kafka (broker + connect + akhq)..."
	rm -rf kafka kafka-connect
	mkdir -p kafka 
	chmod 777 kafka 
	$(COMPOSE) -f kafka.docker-compose.yml up --force-recreate -d --build 

kafka-down:
	@echo "🛑 Stopping Kafka..."
	$(COMPOSE) -f kafka.docker-compose.yml down

kafka-ps:
	@echo "📋 Kafka containers status:"
	$(COMPOSE) -f kafka.docker-compose.yml ps -a

# Postgres
postgres-up:
	@echo "🚀 Starting Postgres..."
	rm -rf pgdata
	mkdir -p pgdata
	chown -R 999:999 pgdata
	$(COMPOSE) -f postgres.docker-compose.yml up --force-recreate -d --build 

postgres-down:
	@echo "🛑 Stopping Postgres..."
	$(COMPOSE) -f postgres.docker-compose.yml down

postgres-ps:
	@echo "📋 Postgres container status:"
	$(COMPOSE) -f postgres.docker-compose.yml ps -a

# ElasticSearch + Kibana
elastic-up:
	@echo "🚀 Starting ElasticSearch + Kibana..."
	rm -rf esdata
	mkdir -p esdata
	chmod 777 esdata
	$(COMPOSE) -f elastic.docker-compose.yml up --force-recreate -d --build 

elastic-down:
	@echo "🛑 Stopping ElasticSearch + Kibana..."
	$(COMPOSE) -f elastic.docker-compose.yml down

elastic-ps:
	@echo "📋 ElasticSearch + Kibana containers status:"
	$(COMPOSE) -f elastic.docker-compose.yml ps -a
