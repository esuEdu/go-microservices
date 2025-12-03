# Microservices Communication Patterns -- Examples

This repository contains hands-on examples of different communication
patterns used in microservice architectures.\
Each branch focuses on a single communication style and includes
minimal, focused examples.

## Purpose

This repo provides practical references for: - How microservices
communicate - Differences between synchronous and asynchronous
communication - Real-world production patterns - Example code &
Docker/Kubernetes setups

## Branch Structure

### 1. `grpc` -- High-Performance RPC (Synchronous)

Includes: - .proto definitions - Generated clients/servers - Unary &
streaming examples - Docker/Kubernetes setup

### 2. `rabbitmq` -- Message Queue (Asynchronous)

Includes: - Producers & consumers - Exchanges & routing - Durable queues
and DLQs - Docker Compose environment

### 3. `kafka` -- Event Streaming (Asynchronous)

Includes: - Producers & consumers - Partitions & consumer groups -
Stream processing samples - Docker/Kubernetes Kafka cluster

### 4. `aws-sns-sqs` -- Cloud Pub/Sub + Queueing (Asynchronous)

Includes: - SNS → SQS fan‑out pattern - Message filtering - FIFO
queues - LocalStack setup

## Architecture Overview

### Synchronous

-   Request/response
-   Example: gRPC

### Asynchronous

-   Event-driven messages
-   Examples: RabbitMQ, Kafka, AWS SNS/SQS

## Usage

``` bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
git checkout grpc
```

Each branch contains its own detailed README.

## Roadmap

-   Add tracing (OpenTelemetry)
-   Add service mesh examples
-   Add CI/CD integrations
