package iface

import (
	"context"

	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
)

type FileInfo interface {
	GetHash() string
	GetName() string
	GetPeers() []string
}

type ScanAPI interface {
	// Start scan network
	Startup(privKey ic.PrivKey, cfg map[string]interface{}) error

	// Publish file to scan
	PublishFile(context.Context, FileInfo, host.Host) error

	GetFilePeers(ctx context.Context, hash string) (FileInfo, error)
}
