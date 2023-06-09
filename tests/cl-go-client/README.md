# Concentrated Liquidity Go Client

This Go-client allows connecting to an Merlins chain via Ignite CLI and
setting up a concentrated liquidity pool with positions.

## General Setup FAQ

- Update constants at the top of the file accordingly.
   * Make sure keyring is set up.
   * Client home is pointing to the right place.

## LocalMerlins Setup

Make sure that you run `localmerlins` in the background and have keys
added to your keyring with:

```bash
make set-env localmerlins # sets environment to $HOME/.merlinsd-local

make localnet-start

make localnet-keys
```

See `tests/localmerlins` for more info.

## Running

```bash
make localnet-cl-create-positions
```

In the current state, it does the following:
- Queries status of the chain to make sure it's running.
- Queries pool with id 1. If does not exist, creates it
- Sets up one CL position
