## ADS-B to NATS

A small application that reads ADS-B messages and sends them to NATS.

## Deployment
Deploy as a simple deployment:
```shell
ko apply -f deploy.yaml
```

## Validate
Check that message are written from a nats-box:
```shell
nats-top -s nats
```