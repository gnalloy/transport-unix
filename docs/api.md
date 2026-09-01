# API Reference

[简体中文](api.zh-CN.md) | [Docs Index](README.md)

This inventory is generated from `go doc -short` for the packages in this repository. It is a quick public-surface map; source files and tests remain the authority for exact semantics.

## Packages

### `gnalloy.org/transport-unix`

Package name: `unix`

```text
var ErrUnsupportedUnixSocket = errors.New("gnalloy/transport/unix: unsupported unix domain socket") ...
func ReceiveFD(transport.FDRef) (int, error)
func SendFD(transport.FDRef, int) error
type Address struct{ ... }
    func ParseAddress(address string) (Address, error)
type AllocatorFactory func(loop *transport.EventLoop) (buffer.Allocator, error)
type Config struct{ ... }
    func DefaultConfig() Config
type Credentials struct{ ... }
    func PeerCredentials(transport.FDRef) (Credentials, error)
type DatagramEndpoint struct{ ... }
    func ListenDatagram(address string, cfg Config) (*DatagramEndpoint, error)
type Server struct{ ... }
type Transport struct{ ... }
    func NewTransport(cfg Config) *Transport
```
