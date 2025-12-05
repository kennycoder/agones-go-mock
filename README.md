# Agones Mock SDK

This package provides an abstraction layer for the [Agones SDK](https://agones.dev/site/docs/guides/client-sdks/) using a Go interface. This allows developers to write game server logic that can run locally without a Kubernetes cluster (using a mock) and seamlessly switch to the real SDK when deployed.

## Installation

```bash
go get github.com/kennycoder/agones-go-mock
```

## Usage

Import the package in your game server code:

```go
import "github.com/kennycoder/agones-go-mock"
```

### Example

See `examples/simple/main.go` for a complete working example.

### Local Development (Mock Mode)

To run the game server locally using the mock SDK, set the `AGONES_ENV` environment variable to `local`. This will print SDK operations to the console.

```bash
AGONES_ENV=local go run examples/simple/main.go
```

Output example:
```text
2023/10/27 10:00:00 Initializing MOCK SDK...
[2023-10-27T10:00:00Z] Agones Mock: Ready called
[2023-10-27T10:00:00Z] Agones Mock [Alpha]: SetPlayerCapacity called with args: [100]
```

### Production (Real Mode)

When deploying to an Agones cluster, run without the environment variable (or set it to anything else). The application will attempt to connect to the Agones sidecar.

```bash
go run examples/simple/main.go
```

*Note: This will fail locally if the Agones sidecar is not present, which is expected behavior for the real SDK.*

## Project Structure

- **`interface.go`**: Defines the `AgonesSDK` interface, covering Lifecycle, Health, Metadata, and Alpha features.
- **`mock.go`**: A mock implementation of the SDK that logs operations to standard output instead of communicating with the Agones sidecar.
- **`real.go`**: A wrapper around the official Agones Go SDK that satisfies the `AgonesSDK` interface.
- **`examples/`**: Contains example usage of the package.

## Implementation Details

Your game logic should depend on the `agonesmock.AgonesSDK` interface rather than the concrete SDK struct.

```go
type GameServerLogic struct {
    agones agonesmock.AgonesSDK
}
```

This dependency injection pattern makes unit testing easier and decouples your game logic from the infrastructure.