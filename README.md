# go-random

Socket-based psuedo-rng to imitate /dev/random

## Quickstart

Clone and build the project:
```bash
git clone git@github.com:bcicen/go-random.git && \
cd go-random && \
go build
```

Run without any args for the most basic implementation, using Go math standard library for rng:
```bash
./go-random
```
or run with any number of character device paths as arguments to contribute to the entropy pool:
```bash
./go-random /dev/input/mice /dev/input/event4
```

Connecting to the socket can be tested via socat:
```bash
socat unix-connect:/tmp/rand.sock stdio
```
