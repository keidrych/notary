package signer

import (
	pb "github.com/docker/notary/proto"
	"github.com/endophage/gotuf/data"
	"github.com/endophage/gotuf/signed"
)

// SigningService is the interface to implement a key management and signing service
type SigningService interface {
	KeyManager

	// Signer returns a Signer for a given keyID
	Signer(keyID *pb.KeyID) (Signer, error)
}

// CryptoServiceIndex represents a mapping between a service algorithm string
// and a CryptoService
type CryptoServiceIndex map[data.KeyAlgorithm]signed.CryptoService

// KeyManager is the interface to implement key management (possibly a key database)
type KeyManager interface {
	// CreateKey creates a new key and returns it's Information
	CreateKey() (*pb.PublicKey, error)

	// DeleteKey removes a key
	DeleteKey(keyID *pb.KeyID) (*pb.Void, error)

	// KeyInfo returns the public key of a particular key
	KeyInfo(keyID *pb.KeyID) (*pb.PublicKey, error)
}

// Signer is the interface that allows the signing service to return signatures
type Signer interface {
	Sign(request *pb.SignatureRequest) (*pb.Signature, error)
}