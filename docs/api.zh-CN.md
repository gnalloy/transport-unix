# API 参考

[English](api.md) | [文档索引](README.zh-CN.md)

本清单由本仓库 package 的 `go doc -short` 生成，用于快速查看公共面。精确语义以源码和测试为准。

## 包

### `gnalloy.org/transport-unix`

包名：`unix`

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
