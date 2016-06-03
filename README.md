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
    "0.60",
    "1.15",
    "1.29"
  ],
  "runningprocesses": "2",
  "totalprocesses": "1466"
}
```
