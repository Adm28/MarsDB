# Mars DB
A distributed,persistant key-value store built on top of levelDB in Go.
This key-value store is suitable to build OLTP database on top of it.

## Requirements
|Framework/Dependencies | Version|
| :--- | :--- |
| Go | 1.10+ |

## RoadMap
__________________________

### Distributed
- [] Implement Hash Partioning and Range Partioning
- [] Implement Raft Consensus

### Performance
- [] Decide the caching policy and implement caching
- [] Decide the compaction policy

