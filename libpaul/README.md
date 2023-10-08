# libpaul

## init

- starts pouchdb subprocess
- starts yggdrasil subprocess
- connects pouchdb to other pouchdb

## api functions

- **Get all entries** - get a list of all entries, regardless of if they're stored on this device
- **Create** entry (bytes, metadata)
  - creates a File record, with its devices field indicating it is saved on this devicec
  - returns Entry
- **Delete** entry (hash)
  - removes File record, which effectively untracks this file on all devices
  - deletes local copy
- **Save** entry (hash)
  - fetch a copy of this file and save it on this device
  - updates File record's devices field to say that it is saved on this device
  - returns Entry
- **Unsave** entry (hash)
  - deletes local copy
  - updates File record's devices field to say that it is no longer saved on this device
- **Get list of devices**
  - returns list of devices
- _Add peer (peer ip)_
- _Get my IP_
  - _returns my IP_
