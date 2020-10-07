# Mars DB
A distributed,persistant and fault tolerant  peer to peer key-value store built on top of levelDB.
The purpose of building the key value store is to enhance my understanding of various storage techinqies
and implement various distributed systems protocol.

## RoadMap
__________________________

### Core and Storage
- [ ] Benchmark the reads and writes latency and ops/sec against different workloads
- [ ] Employee caching technique to decrease write and append latency
- [ ] Change the architecture of storage to improve the performance of writes and append operations.


### Distributed 
- [x ] Consistent hashing
- [  ] Sharding
- [  ] Replication across Nodes
- [  ] Fault Tolerance
- [  ] Range Paritioning

