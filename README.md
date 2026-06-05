# ChronoKV
What it is? --- a networked database where you can store, retrieve and delete key-value pairs across multiple machines, and also query what any value was at any point in the past(Dr. Strange kinda shit, PS: remind me to remove this after we complete the project )


## Features

- SET / GET / DELETE key-value pairs over a network
- Time travel queries — GET key AT timestamp
- Data persists across restarts via append-only log
- Data distributed across multiple nodes using a simplified version of RAFT
- Fault tolerant — survives one node going down
- Live dashboard showing cluster health

## Replication — Simplified Raft

ChronoKV uses a simplified implementation of the Raft consensus algorithm
for replication and fault tolerance.

### What we implement

Node States
Every node is always in one of three states:
- Follower — default state, listens to leader
- Candidate — requesting votes to become leader
- Leader — handles all writes, sends heartbeats to followers

Leader Election
- On startup nodes vote for a leader
- Node with majority votes becomes leader
- Leader sends periodic heartbeats to followers
- If heartbeats stop, a new election begins
- Every election has a term number — messages from older terms are ignored

Log Replication
- All writes go to the leader only
- Leader appends to its log and sends to all followers
- Write is confirmed only when majority of nodes acknowledge
- Guarantees no data loss even if a node dies mid-write

PS: thanks to claude for explaining components

### What we leave for later (towards full Raft)
--- lagbhag sab jo RAFT ke bare main complicated hai


## Architecture
[ CLI Client ]
↓
[ Any Node ]
↓
[ Consistent Hashing Ring ] → routes to correct node
↓
[ Node: HashMap + Log File ]
↓
[ Replication → Backup Node ]


## Project Structure
chronokv/
├── store/         → in-memory hashmap + log file
├── server/        → TCP server, handles connections
├── hashing/       → consistent hashing ring
├── replication/   → copies data to backup nodes
├── client/        → CLI to interact with the cluster
├── dashboard/     → web dashboard for cluster health
├── main.go        → entry point
└── go.mod         → Go module definition



## Commands

| Command                | Description                  |
|------------------------|------------------------------|
| `SET key value`        | Store a value                |
| `GET key`              | Get current value            |
| `DELETE key`           | Remove a key                 |
| `GET key AT timestamp` | Get value at a point in time |