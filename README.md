# go-dash (alpha)

A low-overhead monitoring web dashboard for a GNU/Linux machine. Simply 
drop-in the app and go!


## Documentation

### Distribution-specific information

`/issue`
```json
{
  "distro": "Ubuntu 14.04.4 LTS",
  "kernel": "3.13.0-85-generic"
}
```

### File system disk space usage

`/df`
```json
{
  "filesystems": [
    {
      "filesystem": "udev",
      "size": "3.9G",
      "available": "3.9G",
      "mountpoint": "/dev"
    },
    {
      "filesystem": "tmpfs",
      "size": "787M",
      "available": "785M",
      "mountpoint": "/run"
    }
  ]
}
```

### CPU and IO Load Averages

`/load`
```json
{
  "load": [
    0.60,
    1.15,
    1.29
  ],
  "runningprocesses": 2,
  "totalprocesses": 1466
}
```

### Swap space and utilization

`/swap`
```json
[
  {
    "filename": "/dev/sda5",
    "type": "partition",
    "size": 3905532,
    "used": 572,
    "priority": -1
  }
]
```

### Network connections and routing tables

`/netstat`
```json
[
  {
    "protocol": "tcp",
    "local_address": "192.168.1.111:42549",
    "foreign_address": "192.30.253.125:443",
    "state": "ESTABLISHED"
  },
  {
    "protocol": "tcp",
    "local_address": "192.168.1.117:47739",
    "foreign_address": "54.192.39.12:443",
    "state": "ESTABLISHED"
  }
]
```

### CPU Information
`/cpuinfo`
```json
{
  "architecture": "x86_64",
  "cpus": 4,
  "sockets": 1,
  "cpu_family": 6,
  "bogomips": 5182.87
}
```
