# Go-Kafka-Playground

A simple project to experiment with Kafka using KRaft mode (Kafka without ZooKeeper) and integrating it with a Go application.

---

## Requirements

### 1. Kafka in KRaft Mode

Kafka in KRaft mode simplifies deployment by removing the dependency on ZooKeeper. This project uses `confluentinc/cp-kafka` Docker image to run Kafka in KRaft mode.

### Docker Compose Configuration

Below is the `docker-compose.yml` file to set up Kafka in KRaft mode:

```yaml
version: "3.8"

services:
  kafka:
    image: confluentinc/cp-kafka:7.5.0
    container_name: kafka
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_PROCESS_ROLES: broker,controller
      KAFKA_NODE_ID: 1
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_LISTENERS: CONTROLLER://kafka:9093,EXTERNAL://0.0.0.0:9092
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: CONTROLLER:PLAINTEXT,EXTERNAL:PLAINTEXT
      KAFKA_ADVERTISED_LISTENERS: EXTERNAL://localhost:9092
      KAFKA_INTER_BROKER_LISTENER_NAME: EXTERNAL # Inter-broker communication must use a listener in advertised.listeners
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@kafka:9093
      KAFKA_LOG_DIRS: /tmp/kraft-combined-logs
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1
      CLUSTER_ID: "RGHjASfRSuK24tH6J9AmsA"
    volumes:
      - kafka-data:/tmp/kraft-combined-logs

volumes:
  kafka-data:

