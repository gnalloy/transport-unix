//go:build !linux && !darwin && !freebsd && !netbsd && !openbsd && !dragonfly

package unix

import (
	"gnalloy.org/gnalloy/channel"
	"gnalloy.org/gnalloy/transport"
)

func listenUnix(string, socketOptions) (listenSocket, error) {
	return listenSocket{}, ErrUnsupportedUnixSocket
}

func acceptUnix(transport.FDRef) (transport.FDRef, bool, error) {
	return transport.FDRef{FD: -1}, false, ErrUnsupportedUnixSocket
}

func dialUnix(string, socketOptions) (transport.FDRef, error) {
	return transport.FDRef{FD: -1}, ErrUnsupportedUnixSocket
}

func cleanupSocket(Address) error {
	return nil
}

func closeFD(transport.FDRef) error {
	return nil
}

func newNativeReadWriter() channel.FDReadWriter {
	return nil
}
