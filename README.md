# Concurrent URL Health Checker

A production-style concurrent URL health monitoring CLI written in Go.

This project begins with a simple sequential URL checker and progressively evolves into a concurrent worker-pool based monitoring system using Go concurrency primitives such as goroutines, channels, and WaitGroups.

---

# Objectives

* Learn Go concurrency deeply
* Understand worker-pool architectures
* Practice HTTP/network programming
* Build production-style CLI systems
* Compare sequential vs concurrent execution

---

# Features

## Core Features

* URL health checking
* HTTP status reporting
* Latency measurement
* Timeout handling
* Error reporting

## Concurrent Features

* Goroutines
* WaitGroups
* Channels
* Worker pools
* Configurable concurrency

## Advanced Features

* Retry support
* Rate limiting
* Periodic monitoring
* JSON/CSV output
* Graceful shutdown

---

# Project Evolution

The project is intentionally developed in stages.

```text
Sequential Checker
        ↓
Structured Architecture
        ↓
Timeout Handling
        ↓
Goroutines
        ↓
WaitGroups
        ↓
Channels
        ↓
Worker Pool
        ↓
Retries + Rate Limiting
        ↓
Periodic Monitoring
```

---

# Final Architecture

```text
                +-------------------+
                |   URL Input File  |
                +---------+---------+
                          |
                          v
                 +--------+--------+
                 |   URL Producer  |
                 +--------+--------+
                          |
                     jobs channel
                          |
      +-------------------+-------------------+
      |                   |                   |
      v                   v                   v
+-----------+      +-----------+      +-----------+
| Worker 1  |      | Worker 2  | ...  | Worker N  |
+-----------+      +-----------+      +-----------+
      |                   |                   |
      +-------------------+-------------------+
                          |
                    results channel
                          |
                          v
                 +--------+--------+
                 | Result Collector|
                 +--------+--------+
                          |
                          v
                  Console / JSON / CSV
```

---

# Sequential Algorithm

```text
for each URL:
    start timer
    make HTTP request
    measure latency
    collect result
    print output
```

---

# Concurrent Worker-Pool Algorithm

## Producer

```text
for URL in URLs:
    jobs <- URL

close(jobs)
```

## Worker

```text
for job := range jobs:
    result := check(job)
    results <- result
```

## Collector

```text
for result := range results:
    process(result)
```

---

# Recommended Project Structure

```text
healthchecker/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── checker/
│   ├── reader/
│   └── output/
│
├── urls.txt
├── go.mod
└── README.md
```

---

# Concepts Practiced

* Goroutines
* Channels
* WaitGroups
* Worker pools
* Synchronization
* HTTP clients
* Timeouts
* File I/O
* CLI flags
* Context cancellation
* Rate limiting
* Error handling

---

# Development Roadmap

## Phase 1

* Sequential checker
* Hardcoded URLs
* Basic HTTP requests

## Phase 2

* Result structs
* Timeout support
* File input

## Phase 3

* Goroutines
* WaitGroups
* Channels

## Phase 4

* Worker-pool architecture
* Configurable workers

## Phase 5

* Retries
* Rate limiting
* Monitoring mode

## Phase 6

* JSON/CSV output
* Graceful shutdown
* Metrics/logging

---

# Example Usage

```bash
go run . -workers=10 -input=urls.txt
```

Monitoring mode:

```bash
go run . -workers=10 -watch=30s
```

---

# Stretch Goals

* TLS certificate checks
* DNS timing analysis
* Prometheus metrics
* Web dashboard
* Slack/email alerts
* Distributed workers

---

# Why This Project?

This project demonstrates real-world Go backend engineering patterns beyond basic CRUD applications:

* Concurrent systems design
* Network programming
* Resource management
* Synchronization patterns
* Production-grade architecture

Ideal for:

* Go backend interviews
* Systems engineering practice
* Concurrency mastery
* Portfolio projects

