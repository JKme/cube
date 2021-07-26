# cube

### CUBE

学习Go语言，适用内网渗透测试。 ~~F-Scrack的翻版，给fscan和Ladon做了个分类，[X-Crack](https://github.com/netxfly/x-crack) 套壳~~
，包括三个模块信息收集(probe)、弱密码爆破(crack)、命令执行(sqlcmd)，此处参考了gobuster的爆破模式

```bash
Usage:
  cube [command]

Available Commands:
  crack       crack service password, avaliable plugin: smb,mongo,postgres,ssh,mysql,ftp,redis,elastic,mssql
  help        Help about any command
  probe       collect pentest environment information
  sqlcmd      exec sql query or cmd

Flags:
      --delay int     delay for request 
  -h, --help          help for cube
  -n, --threads int   Number of concurrent threads (default 30)
      --timeout int   Timeout each thread waits (default 5)
  -v, --verbose       Verbose output (errors)
```

### Flags
下面都是全局参数，适用任何模块
#### --delay 每次的请求之间的时间延迟，设定参数之后，多线程数量强制设为1，用于在流量监控特别敏感的内网(~~感觉没有什么卵用~~)
#### --threads  设定多线程数量，默认30
#### --verbose  Debug模式输出


### Probe 内网信息收集
内网探测信息，已实现的有三个插件：OXID多网卡探测，MS17010、zookeeper、smbghost扫描
可用插件：`oxid,ms17010,zookeeper,smbghost`

```bash
cube probe -x oxid -i 192.168.2.1/24
cube probe -x ALL -i 192.168.2.1/24
```

### Crack 弱密码爆破
```bash
cube crack -h
crack service password, avaliable plugin: redis,postgres,mssql,elastic,ssh,mysql,ftp,smb,mongo

Usage:
  cube crack [flags]

Flags:
  -h, --help               help for crack
  -i, --ip string          ip (e.g. 192.168.2.1
      --ip-file string     ip file
  -p, --pass string        login password
      --pass-file string   login password file
  -x, --plugin string      avaliable plugin: redis,postgres,mssql,elastic,ssh,mysql,ftp,smb,mongo
      --port string        if the service is on a different default port, define it here
  -u, --user string        login user
      --user-file string   login user file

Global Flags:
      --delay int     delay for request
  -n, --threads int   Number of concurrent threads (default 30)
      --timeout int   Timeout each thread waits (default 5)
  -v, --verbose       Verbose output (errors)
```
用户名(`-u/--user-file`)和密码(`-p/--pass-file`)成对出现，可以任意组合， 可用插件：`ssh，mysql，redis，elastic，ftp，httpbasic，mongo，mssql，phpmyadmin，smb，postgres, jenkins`

```
Examples:
cube crack -u root -p root -i 192.168.1.1 -x ssh
cube crack -u root -p root -i 192.168.1.1 -x ssh --port 2222
cube crack -u root,ubuntu -p 123,000111,root -x ssh -i 192.168.1.1
cube crack -u root -p root -i 192.168.1.1/24 -x ssh
cube crack -u root --pass-file pass.txt -i 192.168.1.1/24 -x ssh
cube crack -u root --pass-file pass.txt -i 192.168.1.1/24 -x ssh,mysql

phpmyadmin、httpbasic、jenkins插件只能单独使用，不可组合:
cube crack -u root --pass-file pass.txt -i http://192.168.1.1 -x phpmyadmin
cube crack -u root --pass-file pass.txt -i http://192.168.1.1 -x httpbasic
cube crack -u root --pass-file pass.txt -i http://192.168.1.1 -x jenkins
```

sqlserver爆破密码的代码(Event Code): 18456

#### Sqlcmd 命令执行
执行命令，可用插件： `ssh`,`mssql`(开启xp_cmdshell),`mssql-wscript`,`mssql-com`,`mssql-clr`
```
Examples:
cube sqlcmd -x ssh://172.16.157.163:2222 -usa -p123456 -e "whoami"


cube sqlcmd -x mssql://172.16.157.163 -usa -p123456 -e "whoami"
cube sqlcmd -x mssql://172.16.157.163 -usa -p123456 -e "close" //close xp_cmdshell

cube sqlcmd -x mssql-wscript://172.16.157.163 -usa -p123456 -e "whoami"
cube sqlcmd -x mssql-wscript://172.16.157.163 -usa -p123456 -e "close" //close sp_oacreate


cube sqlcmd -x mssql-com://172.16.157.163 -usa -p123456 -e "whoami"
cube sqlcmd -x mssql-com://172.16.157.163 -usa -p123456 -e "close" //close sp_oacreate


cube sqlcmd -x mssql-clr://172.16.157.163 -usa -p123456 -e "whoami"
cube sqlcmd -x mssql-clr://172.16.157.163 -usa -p123456 -e "close" //close CLR
```

#### ELK SIEM Detections Rule
```
1. mssql execute cmd
process where event.type in ("start", "process_started") and
process.name : "cmd.exe" and process.parent.name : "sqlservr.exe"

```

### TODO
NTLM信息识别收集：

https://github.com/FeigongSec/NTLMINFO

https://github.com/RowTeam/SharpDetectionNTLMSSP

https://github.com/checkymander/Sharp-SMBExec/blob/master/SharpInvoke-SMBExec/Program.cs

- [ ] NTLM SSP信息收集扫描
- [ ] 增加输出CSV
- [x] 增加sqlcmd的mssql命令执行
- [x] 增加请求间隔延迟 --delay，当设定这个选项的时候，线程强制设为1，这个选项大概用不上？
- [ ] 变量名和函数名优化
~~- [ ] SMB和OXID输出的中文乱码问题~~
- [ ] **尝试改造为interface实现**

httpx -title --follow-redirects --status-code -tech-detect --title -ports 8000,8080,8888

### 参考
* <https://github.com/shadow1ng/fscan>
* <https://github.com/k8gege/LadonGo>
* <https://github.com/OJ/gobuster>
* <https://github.com/netxfly/x-crack>
* <https://github.com/mabangde/pentesttools/blob/master/golang/sqltool.go>

### 声明
>特别声明：此工具仅限于安全研究，禁止使用该项目进行违法操作，否则自行承担相关责任
