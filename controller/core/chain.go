package core

type Chain interface {
	SetRouter(*Router)
	Start() error // Start chain
	ID() uint8
	Name() string
	Stop()
}

type ChainConfig struct {
	Name           string            // Human-readable chain name
	ID             uint8             // ChainID
	Endpoint       string            // url for rpc endpoint
	From           string            // address of key to use
	KeystorePath   string            // Location of key files
	Insecure       bool              // Indicated whether the test keyring should be used
	BlockstorePath string            // Location of blockstore
	FreshStart     bool              // If true, blockstore is ignored at start.
	LatestBlock    bool              // If true, overrides blockstore or latest block in config and starts from current block
	Opts           map[string]string // Per chain options
}
