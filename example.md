# Code Extractor Example

This is a sample Markdown file to demonstrate the code extraction tool. It contains code blocks in different languages.

## Shell Commands

Here are some shell commands with comments explaining their purpose:

```shell
# Install required packages
sudo apt update
sudo apt install -y curl wget

# Check system information
uname -a
```

```shell
# Create a new directory and navigate to it
mkdir -p ~/projects/myapp
cd ~/projects/myapp

# Initialize a git repository
git init
```

## Python Code

Some Python examples with comments:

```python
# Calculate Fibonacci numbers
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

# Print first 10 Fibonacci numbers
for i in range(10):
    print(f"Fibonacci({i}) = {fibonacci(i)}")
```

```python
# Simple HTTP server
from http.server import HTTPServer, SimpleHTTPRequestHandler

def run(server_class=HTTPServer, handler_class=SimpleHTTPRequestHandler):
    server_address = ('', 8000)
    httpd = server_class(server_address, handler_class)
    print("Starting server on port 8000...")
    httpd.serve_forever()

if __name__ == '__main__':
    run()
```

## Go Code

Some Go examples with comments:

```go
// Simple HTTP server
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on port 8080...")
    http.ListenAndServe(":8080", nil)
}
```

```go
// Concurrent prime number calculation
package main

import (
    "fmt"
    "sync"
)

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var wg sync.WaitGroup
    numbers := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

    for _, n := range numbers {
        wg.Add(1)
        go func(num int) {
            defer wg.Done()
            if isPrime(num) {
                fmt.Printf("%d is prime\n", num)
            }
        }(n)
    }

    wg.Wait()
}
```

## Mixed Commands

Here's a sequence of commands that combines different languages:

```shell
# Create a new project directory
mkdir -p ~/projects/mixed
cd ~/projects/mixed

# Initialize a Go module
go mod init example.com/mixed
```

```python
# Generate some test data
import json

data = {
    "name": "Test Project",
    "version": "1.0.0",
    "dependencies": ["go", "python"]
}

with open("config.json", "w") as f:
    json.dump(data, f, indent=2)
```

```shell
# Build and run the Go application
go build
./mixed
```