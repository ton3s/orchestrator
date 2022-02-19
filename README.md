# Task Orchestrator in Go

## Setup

```bash
git clone git@github.com:ton3s/orchestrator.git
```

## Run

```bash
❯ go run main.go
task: {d96489fc-5c71-4473-911e-c23743ac79eb Task-1 0 Image-1 0 1024 1 map[] map[]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
task event: {8c621f2e-3e16-4d62-98b6-5a15421a9d83 0 2022-02-18 21:24:13.050269 -0600 CST m=+0.002833144 {d96489fc-5c71-4473-911e-c23743ac79eb Task-1 0 Image-1 0 1024 1 map[] map[]  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}}
worker: { {<nil> <nil> 0} map[] 0}
I will collect stats
I will start or stop a task
I will start a task
I will stop a task
manager: {{<nil> <nil> 0} map[] map[] [] map[] map[]}
I will select an appropriate worker based on the task requirements
I will update tasks
I will send work to workers
node: {Node-1 192.168.1.1 4 1024 0 25 0 worker 0}
create a test container
{"status":"Pulling from library/postgres","id":"13"}
{"status":"Digest: sha256:8b8ff4fcdc9442d8a1d38bd1a77acbdfbc8a04afda9c19df47383cb2d08fc04b"}
{"status":"Status: Image is up to date for postgres:13"}
Container 3c2baef66b05c1b4be6cad84dfb273f44dcf6940cc0db64f3e961f3623be8cc6 is running with config {test-container-1 false false false map[] [] postgres:13 0 0 0 [POSTGRES_USER=cube POSTGRES_PASSWORD=secret] }
stopping container 3c2baef66b05c1b4be6cad84dfb273f44dcf6940cc0db64f3e961f3623be8cc6
2022/02/18 21:24:19 Attempting to stop container 3c2baef66b05c1b4be6cad84dfb273f44dcf6940cc0db64f3e961f3623be8cc6
Container 3c2baef66b05c1b4be6cad84dfb273f44dcf6940cc0db64f3e961f3623be8cc6 has been stopped and removed
```
