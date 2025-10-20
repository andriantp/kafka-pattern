# kafka-pattern

Example toolkit for implementing integration patterns with Apache Kafka, focusing on Change Data Capture (CDC) using Go.
This repository supports the article "Golang x Kafka: Change Data Capture (CDC) & Integration Patterns" by [Andrian Triputra](https://andriantriputra.medium.com/golang-x-kafka-6-change-data-capture-cdc-integration-patterns-391f434655f9).


## 🧭 Overview

CDC (Change Data Capture) is a key pattern in data and event-driven architectures.
It captures changes in a source system (insert, update, delete) and streams them to downstream systems via Kafka.

This repository contains Go-based examples that demonstrate how to:
- Capture data changes
- Stream events through Kafka topics
- Integrate between microservices, databases, and downstream consumers
- Implement various CDC and integration patterns in practice


## 📘 Reference Article

Read the full article on Medium:
Golang x Kafka: [Change Data Capture (CDC) & Integration Patterns](https://andriantriputra.medium.com/golang-x-kafka-6-change-data-capture-cdc-integration-patterns-391f434655f9)

The article explains several integration and CDC patterns conceptually, while this repository provides the practical code examples that support it.

## 📁 Repository Structure
```bash
/docker       - Docker and Docker Compose configuration for Kafka, ZooKeeper, and related services  
/golang       - Go code examples of Kafka and CDC integration patterns  
.gitignore  
LICENSE       - Apache 2.0 License  
README.md  
```

## 🚀 Getting Started
Prerequisites
- Docker and Docker Compose
- Go (version 1.18 or later recommended)
- A running Kafka broker

(Optional) A database that supports CDC or a simulated data source

Run the example
- Clone the repository:
```bash
$ git clone https://github.com/andriantp/kafka-pattern.git
$ cd kafka-pattern
```
- Start Kafka and supporting services:
```bash
$ cd docker
$ docker-compose up -d
```
- Run the Go example:
```bash
cd ../golang
go run main.go
```
- Test the CDC flow by simulating data changes (insert, update, delete) in the source system and verifying that the events are published to the Kafka topic and consumed downstream.

## 📌 Integration Patterns Covered

This repository and the accompanying article demonstrate several integration patterns:

- Change Data Capture (CDC): Capture changes in a database and publish them to Kafka.
- Event-driven architecture: Services communicate through Kafka topics instead of direct calls.
- Outbox pattern / Log-based replication: Separate application data changes from public event models.
- Denormalization / materialized view: Produce enriched events for easier downstream consumption.
- Microservices decoupling: Use Kafka to reduce direct dependencies between services.

## ✅ Benefits

- Real-world examples of Go and Kafka integration for CDC.
- Understand modern event-driven and integration architectures.
- Easily adaptable for your own use cases or experiments.
- Accelerates prototyping of Kafka-based systems with Go.

## 🛠 Notes and Tips

- Ensure Kafka and ZooKeeper are configured properly (ports, topics, replication factors).
- If using a real database for CDC, make sure it supports log-based change capture (for example MySQL binlog or PostgreSQL WAL).
- For quick testing, you can simulate data changes manually and use a simple Kafka consumer to verify results.
- In production, consider additional aspects such as schema evolution, error handling, retries, poison message handling, monitoring, and security (ACLs, encryption).
- Be mindful of schema changes that may break downstream consumers. Using data contracts or denormalized events can help manage schema evolution.

## 📄 License

This project is licensed under the Apache License 2.0.

See the LICENSE file for details.