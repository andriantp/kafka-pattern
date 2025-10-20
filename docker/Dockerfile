FROM confluentinc/cp-kafka-connect:7.7.0

# Install Debezium Postgres Connector
RUN curl -fsSL https://repo1.maven.org/maven2/io/debezium/debezium-connector-postgres/2.7.3.Final/debezium-connector-postgres-2.7.3.Final-plugin.tar.gz \
    | tar -xvz -C /usr/share/java/

# Install Confluent Hub Client
RUN confluent-hub install --no-prompt confluentinc/kafka-connect-elasticsearch:14.0.14

# Plugins folder
ENV CONNECT_PLUGIN_PATH="/usr/share/java,/usr/share/confluent-hub-components"
