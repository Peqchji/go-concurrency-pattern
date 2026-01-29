# Go Concurrency Patterns: AI-Driven Learning

This repository is a collection of **Go (Golang) concurrency patterns** I have practiced, developed through a "learning by doing" partnership between a software engineer and AI. 

The focus is on building scalable, race-condition-free systems while preparing for high-performance engineering.

---

## 🚀 Learning Methodology

Unlike traditional tutorials, this repo follows an iterative **feedback-loop**:
1.  **Scenario Synthesis:** AI proposes a high-concurrency challenge (e.g., rate-limited data ingestion).
2.  **Implementation:** Writing idiomatic Go using channels, goroutines, and synchronization primitives.
3.  **Stress Testing:** Identifying bottlenecks and data races using `go test -race`.
4.  **AI Refinement:** Refactoring for "Mechanical Sympathy"—aligning code with the Go runtime scheduler for maximum efficiency.

---

## 🛠️ Implemented Patterns

This repo contains hands-on implementations for the following concurrent structures:

| Pattern | Directory | Key Focus |
| :--- | :--- | :--- |
| **Worker Pool** | `/workerpool` | Resource limiting and managed task distribution. |
| **Fan-In** | `/fanin` | Multiplexing multiple data streams into a single channel. |
| **Fan-Out** | `/fanout` | Distributing a single workload across multiple workers. |
| **Pipeline** | `/pipeline` | Multi-stage stream transformation. |
| **Rate Limiter** | `/ratelimit` | Controlling execution frequency (Token Bucket/Leaky Bucket). |
| **Semaphore** | `/semaphore` | Controlling access to a finite pool of resources. |
| **Single Flight** | `/single-flight` | Suppressing duplicate function calls. |
| **Parallel Search** | `/parellel-search` | Concurrent searching/fetching across datasets. |
| **Atomic Ops** | `/atomic` | Lock-free state management. |
| **Lazy Loading** | `/lazy-load` | Thread-safe, on-demand resource initialization. |

---

## 🏗️ Technical Stack

* **Language:** Go (Golang) 100%
* **Primitives:** Channels, `sync.Mutex`, `sync.WaitGroup`, `sync.Once`, `context.Context`.
* **Testing:** Race Detector (`-race`), Benchmarking, and AI-guided code reviews.
* **Target Domain:** Distributed Systems and Data Engineering pipelines.

---