# go-papi

An API that knows more about the host system than yourself.

## Capabilities

Recommended to be run as a privileged user.

## Client

Client is initialized with the main server IP.


```bash
papi client 10.10.0.2:9001
```

## Server

The aggregator which collects all information and ask clients info on demand.

```bash
papi server
```