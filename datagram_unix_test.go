//go:build linux || darwin || freebsd || netbsd || openbsd || dragonfly

package unix

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDatagramEndpointSendReceive(t *testing.T) {
	dir := shortSocketDir(t)
	left, err := ListenDatagram(filepath.Join(dir, "l.sock"), DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer left.Close()
	right, err := ListenDatagram(filepath.Join(dir, "r.sock"), DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer right.Close()

	if err := left.SendTo([]byte("ping"), right.Addr().String()); err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, 16)
	n, from, err := right.ReceiveFrom(buf)
	if err != nil {
		t.Fatal(err)
	}
	if string(buf[:n]) != "ping" || from.String() != left.Addr().String() {
		t.Fatalf("n=%d payload=%q from=%s", n, buf[:n], from.String())
	}
}

func shortSocketDir(t *testing.T) string {
	t.Helper()
	base := "/tmp"
	if _, err := os.Stat(base); err != nil {
		base = os.TempDir()
	}
	dir, err := os.MkdirTemp(base, "ga-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	return dir
}
