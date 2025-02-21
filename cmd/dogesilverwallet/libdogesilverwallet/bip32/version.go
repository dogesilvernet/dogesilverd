package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// DogesilverMainnetPrivate is the version that is used for
// dogesilver mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var DogesilverMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// DogesilverMainnetPublic is the version that is used for
// dogesilver mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var DogesilverMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// DogesilverTestnetPrivate is the version that is used for
// dogesilver testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var DogesilverTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// DogesilverTestnetPublic is the version that is used for
// dogesilver testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var DogesilverTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// DogesilverDevnetPrivate is the version that is used for
// dogesilver devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var DogesilverDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// DogesilverDevnetPublic is the version that is used for
// dogesilver devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var DogesilverDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// DogesilverSimnetPrivate is the version that is used for
// dogesilver simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var DogesilverSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// DogesilverSimnetPublic is the version that is used for
// dogesilver simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var DogesilverSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case DogesilverMainnetPrivate:
		return DogesilverMainnetPublic, nil
	case DogesilverTestnetPrivate:
		return DogesilverTestnetPublic, nil
	case DogesilverDevnetPrivate:
		return DogesilverDevnetPublic, nil
	case DogesilverSimnetPrivate:
		return DogesilverSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case DogesilverMainnetPrivate:
		return true
	case DogesilverTestnetPrivate:
		return true
	case DogesilverDevnetPrivate:
		return true
	case DogesilverSimnetPrivate:
		return true
	}

	return false
}
