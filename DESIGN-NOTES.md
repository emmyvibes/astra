## Overlay Network

Overlay network requirements:

- Needs to be portable to all platforms (Go works for this)
- Needs to be usable with both our own custom code (in Go) + we also need to be able to get the CouchDB implementation to use the overlay network
- Needs to dynamically improve routing when a direct peering is set up, e.g.
  1.  Device A is behind a NAT and connects to a couple public static peers
  2.  Device B is behind a NAT and connects to a couple public static peers
  3.  Device A and B can exchange data but slowly
  4.  NAT traversal (upnp, hole punching, etc) is used to establish a direct peering between A and B, alongside the existing static peers
  5.  **Test case\*\***:\*\* after the direct peering is established, data exchange between the two devices should use the direct peering
      1.  Pinecone - not yet tested
      2.  Yggdrasil - **PASSED**

yggdrasil on iOS and Android implementation examples:

[https://github.com/yggdrasil-network/yggdrasil-android](https://github.com/yggdrasil-network/yggdrasil-android)

- Key file: [https://github.com/yggdrasil-network/yggdrasil-android/blob/main/app/src/main/java/eu/neilalexander/yggdrasil/PacketTunnelProvider.kt](https://github.com/yggdrasil-network/yggdrasil-android/blob/main/app/src/main/java/eu/neilalexander/yggdrasil/PacketTunnelProvider.kt)
- Looks like yggdrasil is started as a subprocess, an android VPN service is established [https://developer.android.com/reference/android/net/VpnService.Builder#establish()](<https://developer.android.com/reference/android/net/VpnService.Builder#establish()>) then the file descriptor from establish() is used to pass packets to/from yggdrasil

[https://github.com/yggdrasil-network/yggdrasil-ios](https://github.com/yggdrasil-network/yggdrasil-ios)

- Key file: [https://github.com/yggdrasil-network/yggdrasil-ios/blob/main/Yggdrasil%20Network%20Extension/PacketTunnelProvider.swift](https://github.com/yggdrasil-network/yggdrasil-ios/blob/main/Yggdrasil%20Network%20Extension/PacketTunnelProvider.swift)
- Similar to android, looks like yggdrasil is started as a subprocess, a VPN is configured with iOS, and then packets are passed to/from yggdrasil

pinecone on iOS and Android implementation examples:

[https://github.com/vector-im/element-android-p2p](https://github.com/vector-im/element-android-p2p)

[https://github.com/vector-im/element-ios-p2p](https://github.com/vector-im/element-ios-p2p)

^ both of these above don't work with pinecone directly, instead they subprocess Dendrite which is bound to pinecone in Go: [https://github.com/matrix-org/dendrite/tree/2259e71c0cc7d154eefd85a6466b08e7afb7f69a/cmd/dendrite-demo-pinecone](https://github.com/matrix-org/dendrite/tree/2259e71c0cc7d154eefd85a6466b08e7afb7f69a/cmd/dendrite-demo-pinecone)

### Summarized research on Yggdrasil and Pinecone

- Both Pinecone and Yggdrasil implement a core go library with packet reader/writer interfaces, that you can directly network with
  - Pinecone: [https://pkg.go.dev/github.com/matrix-org/pinecone@v0.11.0/router#Router](https://pkg.go.dev/github.com/matrix-org/pinecone@v0.11.0/router#Router)
  - Yggdrasil: [https://pkg.go.dev/github.com/yggdrasil-network/yggdrasil-go@v0.4.7/src/core#Core](https://pkg.go.dev/github.com/yggdrasil-network/yggdrasil-go@v0.4.7/src/core#Core)
- Yggdrasil and Pinecone, in addition to the above core library, implements a linux tunnel (VPN) interface
  - Pinecone tunnel: [https://github.com/matrix-org/pinecone/tree/main/cmd/pineconeip](https://github.com/matrix-org/pinecone/tree/main/cmd/pineconeip)
  - Yggdrasil tunnel: [https://github.com/yggdrasil-network/yggdrasil-go/tree/develop/src/tun](https://github.com/yggdrasil-network/yggdrasil-go/tree/develop/src/tun)
- Yggdrasil in addition to the above, also implements a multicast auto-local-peering system
- For development purposes today, I think we should use Yggdrasil due to its simple use as a tun/vpn device on Linux systems
- For future (true) cross platform deployment, our core go library (libpaul) should interface directly with one of the above core APIs.
  - This is how all of the above mobile implementations (ygg android, ios, and dendrite-pinecone) interface with the overlay network
  - This allows the app to use the overlay network without touching the OS's network stack

## CouchDB server implementation selection

**CouchDB implementation we select has to allow accepting connections**

Main two options are PouchDB and CouchDB itself

PouchDB

PouchDB-server supports all of the core features we need, namely accepting connections and being able to participate in master-master replication

CouchDB

CouchDB is built on Erlang, which is hard to port to multiple platforms

## File Layer DB Models

Device - a list of devices

- IPv6 Address (in the overlay network)
- Public key
- Device name

File - a list of files

- File hash (PK)
- Metadata
  - File name
  - (extra metadata, size?)
- List of devices it's saved on
  - device id

## Image App DB Models

Image

- Image date
- Rererence to File

## API

"Entry"

- File model (from above)
- Local path (if it exists on this device, where a copy of this file is accessible locally)

API

**Bold - Exposed API and internal API**

_Italic - internal API only_

strikethrough - not right now

- **Get all entries** - get a list of all entries, regardless of if they're stored on this device
- **Create** entry (bytes, metadata)
  - creates a File record, with its devices field indicating it is saved on this devicec
  - returns Entry
- **Delete** entry (hash)
  - removes File record, which effectively untracks this file on all devices
  - deletes local copy
- **Fetch** entry (hash)
  - fetch a copy of this file for temporary use
  - returns byte array
- **Save** entry (hash)
  - fetch a copy of this file and save it on this device
  - updates File record's devices field to say that it is saved on this device
  - returns Entry
- **Unsave** entry (hash)
  - deletes local copy
  - updates File record's devices field to say that it is no longer saved on this device
- Generate pair code
  - returns pair code
- Use pair code (pair code)
  - returns true/false
- **Get list of devices**
  - returns list of devices
- _Add peer (peer ip)_
- _Get my IP_
  - _returns my IP_
