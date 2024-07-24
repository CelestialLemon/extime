# Extime

A simple CLI tool that measures how long a CLI command takes to execute.

### Installation

Download the windows binary from the releases section. Add the path where the binary is located to PATH environment variable.

### Usage

```bash
extime <command>
```

### Sample Usage

```bash
>extime ping www.youtube.com

Pinging youtube-ui.l.google.com [142.250.182.238] with 32 bytes of data:
Reply from 142.250.182.238: bytes=32 time=5ms TTL=117
Reply from 142.250.182.238: bytes=32 time=5ms TTL=117
Reply from 142.250.182.238: bytes=32 time=5ms TTL=117
Reply from 142.250.182.238: bytes=32 time=5ms TTL=117

Ping statistics for 142.250.182.238:
    Packets: Sent = 4, Received = 4, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
    Minimum = 5ms, Maximum = 5ms, Average = 5ms

extime:  3095 ms
```
