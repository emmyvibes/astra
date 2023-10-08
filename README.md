# Astra

Devpost writeup: https://devpost.com/software/astra-pfmt0n

Design notes: [DESIGN-NOTES.md](/DESIGN-NOTES.md)

## Components

- `backend/`: CouchDB database interaction layer
- `libastra/`: Distribute networking and PouchDB (CouchDB implementation) automated startup and configuration
- `wanda/`: Combined API which is consumed by applications using this framework

`libvirt/` contains scripts and tooling for testing yggdrasil network auto-reconfiguration and speed across NATs with and without port-forwarding
