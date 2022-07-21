package core

type Chain interface {
	SetRouter(*Router)
	Start() error
	ID() uint8
	Name() string
	Stop()
}

type ChainConfig struct {
	Name           string            // Human-readable chain name
	ID             uint8             // Human-readable chain id
	Endpoint       string            // Url for client endpoint
	From           string            // External account address and Manage its private key
	KeystorePath   string            // Location of key files
	BlockstorePath string            // Location of blockstore
	FreshStart     bool              // If true, blockstore is ignored at start.
	LatestBlock    bool              // If true, overrides blockstore or latest block in config and starts from current block
	Opts           map[string]string // Per chain options
}
