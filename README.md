# ⚡ Golang Curl Clone

A curl clone built with golang for making HTTP requests with support for benchmarking, custom headers, and more.

---

## 🚀 Usage


```go run main.go [flags]```

All flags are specified in the format:

```--flag=value```

## Help

```go run main.go help```

## 🛠️ Required Flags

- --url / -U | The target URL to send the request to.
- --method / -M | HTTP method (GET, POST, PUT, PATCH, DELETE).

## ⚙️ Available Flags

- --url | -U | Request URL.
- --method | -M | HTTP method.
- --protocol | -P | Protocol to use (http or https).
- --headers | -H | Custom headers. Format: "[header1]:[headerValue1];[header2]:[headerValue2]"
- --skipTLSVerification | -K | Skip TLS certificate validation. ⚠️
- --output | -O | Output response to a file. Specify the file path here.
- --timeout | -T | Request timeout (in milliseconds).
- --data | -D | Request body (for POST, PUT, etc).
- --benchmark | -B | Enable benchmarking mode. Displays average request time in the end.
- --requestsCount | -R | Number of requests for benchmarking. By default 10.
- --verbose | -V | Print detailed request/response logs.

## 🌐 Supported HTTP Methods

- GET

- POST

- PUT

- PATCH

- DELETE

## 🧪 Benchmarking

Enable benchmarking with:

```--benchmark --requestsCount=20```
