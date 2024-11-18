## 1. Deploy ERC4337 Web3 Pay Contract
```sh
git clone https://github.com/fumeng00mike/Pay-Contract.git
yarn
yarn run deploy
```

## 2. Run loadtest
```sh
polycli loadtest erc4337 --account-factory $ACCOUNT_FACTORY --config $CONFIG --entry-point $ENTRYPOINT --helper $HELPER --payable-account $PAYABLE_ACCOUNT --token-receiver $TOKEN_RECEIVER --validator $WEBAUTHN_VALIDATOR --rpc-url "$ETH_RPC_URL" --legacy --private-key "$private_key" -v 700 --requests 1000 --concurrency 10
```