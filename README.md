# Go Bootcamp: Intensive Deep Dive into Golang

This repository contains solutions to tasks from an intensive Go course covering key aspects of the Go programming language. The course is a series of practical assignments, each focusing on different features of the language and the standard library.

## Main Topics

### Working with File Systems and CLI:
- Implementation of Unix-like utilities (find, wc, xargs)  
- Processing large files and data streams  
- Parallel task execution  

### Data Structures and Algorithms:
- Binary trees and traversal algorithms  
- Heaps and priority queues  
- Knapsack problem solutions  
- Statistical calculations  

### Data Handling:
- Processing XML/JSON formats  
- Data structure comparison  
- Serialization and deserialization  

### Advanced Concepts:
- Concurrency (goroutines and channels)  
- Reflection and interface usage  
- Interlanguage interaction (CGO)  
- Performance optimization  

## Project Features

- **Practical approach:** Each assignment solves a real-world task  
- **Idiomatic Go:** Code follows Go best practices  
- **Comprehensive testing:** All solutions come with unit tests  
- **Optimization:** Includes profiling and performance tuning  
- **Documentation:** Detailed explanations of solutions and approaches  

## Technology Stack

- **Language:** Go 1.20+  
- **Libraries:**  
  - Go standard library (sync, net/http, encoding, etc.)  
  - Testing (go test, testify)  
  - Profiling (pprof)  
  - Documentation (godoc)  

## Repository Structure

```text
├── Day00_Statistics/        # Basic statistical calculations
├── Day01_DataComparison/    # Comparing XML/JSON and file systems
├── Day02_UnixTools/         # Implementation of Unix-like utilities
├── Day05_DataStructures/    # Trees, heaps, algorithms
├── Day07_Algorithms/        # Coin change tasks, optimization
├── Day08_AdvancedConcepts/  # Pointers, reflection, CGO
└── Day09_Concurrency/       # Concurrency programming patterns
````

Each directory contains:

* Source code solutions
* Unit tests
* Usage examples
* Documentation

## How to Use

Clone the repository:

```bash
git clone https://github.com/your-username/go-bootcamp.git
```

Navigate to the directory for the desired day:

```bash
cd go-bootcamp/DayXX_Topic
```

Build and run the solution:

```bash
go build -o app main.go
./app [parameters]
```

Run the tests:

```bash
go test -v ./...
```
