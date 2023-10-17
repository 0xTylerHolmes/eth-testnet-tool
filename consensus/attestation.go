package consensus

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	fuzz "github.com/google/gofuzz"
)

func RandomAttestation() *phase0.Attestation {
	var attestation phase0.Attestation
	f := fuzz.New().NilChance(0)
	for true {
		f.Fuzz(&attestation)
		_, err := attestation.MarshalSSZ()
		if err == nil {
			break
		}
	}
	return &attestation
}
