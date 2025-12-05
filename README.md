# Agones Mock SDK Example

This project demonstrates how to abstract the [Agones SDK](https://agones.dev/site/docs/guides/client-sdks/) using a Go interface. This allows developers to write game server logic that can run locally without a Kubernetes cluster (using a mock) and seamlessly switch to the real SDK when deployed.

## Project Structure

- **`interface.go`**: Defines the `AgonesSDK` interface, covering Lifecycle, Health, Metadata, and Alpha features.
- **`mock.go`**: A mock implementation of the SDK that logs operations to standard output instead of communicating with the Agones sidecar.
- **`real.go`**: A wrapper around the official Agones Go SDK that satisfies the `AgonesSDK` interface.
- **`example.go`**: The main entry point demonstrating how to inject the appropriate SDK implementation based on an environment variable.

## Usage

### Local Development (Mock Mode)

To run the game server locally using the mock SDK, set the `AGONES_ENV` environment variable to `local`. This will print SDK operations to the console.

```bash
AGONES_ENV=local go run .
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
go run .
```

*Note: This will fail locally if the Agones sidecar is not present, which is expected behavior for the real SDK.*

## Implementation Details

The `GameServerLogic` struct depends on the `AgonesSDK` interface rather than the concrete SDK struct.

```go
type GameServerLogic struct {
    agones AgonesSDK
}
```

This dependency injection pattern makes unit testing easier and decouples your game logic from the infrastructure.