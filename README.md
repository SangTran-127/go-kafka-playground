# go-kafka-playground



# Requirement
## 1.Kafka 
   Using Kafka in KRaft mode (Kafka without ZooKeeper) with Docker Compose is straightforward and simplifies the deployment. Here's how you can set it up with the confluentinc/cp-kafka image and configure a Go application to interact with it.

    Here's the docker-compose.yml file:

    ```yaml
        version: '3.8'
        services:
        kafka:
            image: confluentinc/cp-kafka:7.5.0 # Use the latest Kafka image supporting KRaft
            container_name: kafka
            ports:
            - "9092:9092"
            environment:
            KAFKA_BROKER_ID: 1
            KAFKA_PROCESS_ROLES: broker,controller
            KAFKA_NODE_ID: 1
            KAFKA_LISTENERS: INTERNAL://kafka:9093,EXTERNAL://0.0.0.0:9092
            KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: INTERNAL:PLAINTEXT,EXTERNAL:PLAINTEXT
            KAFKA_ADVERTISED_LISTENERS: INTERNAL://kafka:9093,EXTERNAL://localhost:9092
            KAFKA_INTER_BROKER_LISTENER_NAME: INTERNAL
            KAFKA_CONTROLLER_QUORUM_VOTERS: 1@kafka:9093
            KAFKA_LOG_DIRS: /tmp/kraft-combined-logs
            KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
            KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1
            KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1
            volumes:
            - kafka-data:/tmp/kraft-combined-logs

        volumes:
        kafka-data:

    ```