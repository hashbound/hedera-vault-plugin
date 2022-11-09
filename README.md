# Hedera Vault Plugin

Hedera Vault Plugin is a backend service to be integrated with Hashicorp Vault in order to bring following features to vault:

- Key Management
  - supports ECDSA/secp256k1 and ED25519 Algorithms
  - supports Import, Generate, Retreive, List and Delete operations
  - supports signing in the box without exposing private keys
- Hedera Account Management
  - supports Import, Create, Retreive, List and Delete operations
  - relates Accounts to its keypair stored in Vault
- Hedera Token Service interactions
  - supports Creating new Tokens
  - supports FT and NFT tokens
  - supports Mint, Burn, transfer/transferFrom, wipe, delete, approve, Freeze/Unfreeze, grant/revoke KYC, pause/unpause, associate/dissociate, etc transactions
- Hedera Topic and Message Management
  - supports creating topics
  - supports creating messages
