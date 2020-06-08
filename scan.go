package iface

import (
	"context"

	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
)

type ScanAPI interface {
	// Start scan network
	Startup(privKey ic.PrivKey, cfg map[string]interface{}) error

	// Publish file to scan
	PublishFile(context.Context, string, host.Host) error

	GetFilePeers(ctx context.Context, hash string) ([]string, error)
}
