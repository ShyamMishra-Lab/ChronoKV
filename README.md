# ChronoKV
What it is? --- a networked database where you can store, retrieve and delete key-value pairs across multiple machines, and also query what any value was at any point in the past(Dr. Strange kinda shit, PS: remind me to remove this after we complete the project )


## Features

- SET / GET / DELETE key-value pairs over a network
- Time travel queries — GET key AT timestamp
- Data persists across restarts via append-only log
- Data distributed across multiple nodes
- Fault tolerant — survives one node going down
- Live dashboard showing cluster health


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