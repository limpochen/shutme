# SHUTME

在私有部署的环境下，检测到断电断网后自动执行关机操作的一个终端命令，用来代替自动关机脚本，运行更加稳定。  
初学GO语言的一个小程序。

## 应用环境
* 运行 Windows / Linux 操作系统的计算机或服务器
* 计算机或服务器无人值守长时间运行
* 有后备电源但不能持久供电，并且后备电源不具备通讯条件

## 使用方法
```shell
shutme -help
```
侦测到异常到处理异常操作的延迟时间 ≈ 侦测间隔 * 重复次数 （单位：秒）

*正式部署前需要测试运行结果*

### 在交互模式下运行
```shell
shutme -t <address>
```

#### 执行非内置关机命令  
SHUTME 内置关机命令  
* Windows: `shutdown -s -t 0`  
* Linux: `shutdown -t 0 -h`  
如果内置命令无效，可通过 -c 选项指定需要的命令。  
```shell
shutme -t <address> -c <cmdline>
```

### 以服务模式后台运行  

#### 安装并启动服务：  
```shell
shutme -s install -t <address>
```

#### 停止并卸载服务：  
```shell
shutme -s uninstall
```

#### 其他可用的服务操作：  
启动服务：`shutme -s start`  
停止服务：`shutme -s stop`  
重启服务：`shutme -s restart`  
