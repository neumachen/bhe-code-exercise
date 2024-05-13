# BHE Software Engineer Coding Exercise

[Original README](README.old.md)

This is my submission for the coding challenge using Golang as the language of
implementation.

## Prime Number Sieve API

This project implements a Prime Number Sieve using the Sieve of Eratosthenes algorithm in Go. It provides an API to fetch the nth prime number, based on a zero-based index, where the 0th prime is 2.

## Components

- [sieve.go](go/pkg/sieve/sieve.go): Implements the Sieve of Eratosthenes algorithm.
- [sieve_test.go](go/pkg/sieve/sieve_test.go) and [sieve_bench_test.go](go/pkg/sieve/sieve_bench_test.go): Contains unit tests and benchmark tests for the sieve algorithm.
- [Dockefile](go/Dockerfile) and [docker-compose.yml](docker-compose.yml): Configuration files for containerizing the application and setting up the environment with Docker.

## Prerequisites

- Go (1.19 or later)
- Docker (for containerization)
- Docker Compose (for running multi-container Docker applications)

## Setup

To set up and run the project, follow these steps:

### Running Locally

1. Clone the repository:
   ```
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Run tests to verify the setup:
   ```
   go test -v ./...
   ```

3. Run benchmark tests:
   ```
   go test -bench=. ./...
   ```

### Using Docker

1. Build the Docker image:
   ```
   docker compose build
   ```

2. Running tests:
   ```
   docker-compose run test
   ```

3. Running benchmarks:
   ```
   docker-compose run benchmark
   ```

## Author’s Reflection: Navigating Misunderstandings and Overcoming Mathematical Challenges

### Initial Misunderstanding and Correction

When I first approached the task of implementing the Sieve of Eratosthenes to find the nth prime number, I ran into a critical misunderstanding. Initially, I misinterpreted the assignment’s requirement to use a zero-based index for prime numbers. This oversight led me to develop and test a solution based on a one-based index, which is common in mathematical discussions but incorrect for the specified requirement.

The error became apparent during my routine check against project requirements—a crucial step in my development process designed to catch such oversights. Realizing the mistake prompted a comprehensive review and correction of both the implementation and the tests, reinforcing the incorrect behavior. This discovery, while initially disheartening, became a valuable lesson on the importance of fully understanding requirements before proceeding with implementation.

### Challenges with Limited Math Background

My journey was not only about correcting a misunderstanding but also about tackling a significant mathematical challenge. With a limited background in advanced mathematics, understanding and implementing an algorithm as mathematically intensive as the Sieve of Eratosthenes was daunting. The task was made even more complex by the need to dynamically calculate the number of primes required based on the input.

To manage this, I adopted a granular approach, breaking down the problem into smaller, more manageable components. This method allowed me to focus on understanding one piece of the puzzle at a time. For instance, after ensuring a basic understanding of the sieve, I moved on to optimize its performance and scalability.

### Integrating the Prime Number Theorem

A turning point in my implementation was deciding to use the Prime Number Theorem to dynamically adjust the prime calculation range. Despite my initial discomfort with the theorem due to my mathematical limitations, I tackled it by breaking down the theorem and applying it step-by-step. This approach not only improved the performance of my implementation but also deepened my understanding of the mathematical concepts involved.

### Reflecting on the Learning Process

This project was as much about developing software as it was about personal growth and learning. It taught me the importance of methodically verifying requirements and the value of resilience when facing daunting challenges. I chose to be transparent about these experiences in my documentation to help others learn from my approach and possibly encourage a culture of continuous improvement and honest reflection in the development community.

By sharing these insights, I hope to inspire those who might feel overwhelmed by similar challenges. It demonstrates that with persistence and a structured approach, it is possible to overcome significant hurdles and achieve one’s objectives, even when starting from a position of limited expertise.
