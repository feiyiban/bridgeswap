package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"

	secp256k1 "github.com/ethereum/go-ethereum/crypto"
)

func ECDSASign2Hex(hash []byte, prv *ecdsa.PrivateKey) (string, error) {
	sig, err := ECDSASign(hash, prv)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sig), nil
}

func ECDSASign(digestHash []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
	return secp256k1.Sign(digestHash, prv)
}

func ECDSASignNoRecover(digestHash []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
	signed, err := ECDSASign(digestHash, prv)
	if err != nil {
		return nil, err
	}
	return signed[:len(signed)-1], nil
}

func CompressPubkey(pubkey *ecdsa.PublicKey) []byte {
	return secp256k1.CompressPubkey(pubkey)
}
func ECDSAVerifySignature(pubkey ecdsa.PublicKey, digestHash, signature []byte) bool {

	if len(signature) > 64 {
		signature = signature[:64]
	}
	return secp256k1.VerifySignature(CompressPubkey(&pubkey), digestHash, signature)
}

func zeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}
