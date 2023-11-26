# SHUTME
[中文介绍](README_zh.md)

In a private deployment environment, a terminal command that automatically executes a shutdown operation after detecting a power outage, used to replace the automatic shutdown script, more stable operation.  
A small program written by beginners of Go language.  

## Application Environment
* Computer or server running Windows / Linux operating system.
* Computer or server running unattended for long periods of time.
* A backup power supply that doesn't last long, and the backup power supply does not have communication conditions.

## How to use
```shell
shutme -help
```
Delay for Shutdown ≈ interval * Cycle (In second)

*You need to test the running results before the official deployment*

### Running in interactive mode
```shell
shutme -t <address>
```

#### Execute a non-built-in shutdown command
SHUTME built-in shutdown command  
* Windows: `shutdown -s -t 0`  
* Linux: `shutdown -t 0 -h`  
If the builtin command doesn't work, you can specify the required command through the -c flag.  
```shell
shutme -t <address> -c <cmdline>
```

### Run as service in a daemon

#### Install and start the service:  
```shell
shutme -s install -t <address>
```

#### Stop and uninstall the service:  
```shell
shutme -s uninstall
```

#### Other available service operations:    
Start service: `shutme -s start`  
Stop service: `shutme -s stop`  
Restart service: `shutme -s restart`  
