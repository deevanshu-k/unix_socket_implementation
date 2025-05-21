## Sample Unix Socket Implementation
Simple unix socket implementation to create ipc connection between processes

### Usage
- Server: `go run . server socket.sock`
- Client: `go run . client socket.sock`

> Here `socket.sock` is a sock file name