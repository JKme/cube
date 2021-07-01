# cube

```
 ._____A_____,
 |`          :\
 | `         : B
 |  `        :  \
 C   +-----A-----+
 |   :       :   :
 |__ : _A____:   :
 `   :        \  :
  `  :         B :
   ` :          \:
    `:_____A_____>
```


### CUBE
学习Go语言

参考gobuster的几种爆破模式，Cube现在有三种模式， `Probe`、`Crack`、`Sqlcmd`
(~~给fscan和Ladon做了个分类~~😀😀，[X-Crack](https://github.com/netxfly/x-crack) 套壳

#### Probe
内网探测信息，比如OXID多网卡探测，Windows版本探测，MS17010扫描
可用插件：
- [x] oxid多网卡扫描
- [ ] ms17010内网扫描
- [ ] 插件多了再做一个大的分类， POC对应ms17010, INFO对应信息收集， ALL表示全部

```
cube probe -x oxid -i 192.168.2.1/24
```

#### Crack
爆破弱密码，可用插件：SSH/Mysql/Redis

```
Examples:
cube crack -u root -p root -i 192.168.1.1 -x ssh
cube crack -u root -p root -i 192.168.1.1 -x ssh --port 2222
cube crack -u root,ubuntu -p 123,000111,root -x ssh -i 192.168.1.1
cube crack -u root -p root -i 192.168.1.1/24 -x ssh
cube crack -u root --pass-file pass.txt -i 192.168.1.1/24 -x ssh
cube crack -u root --pass-file pass.txt -i 192.168.1.1/24 -x ssh,mysql
```

#### Sqlcmd
执行命令，可用插件： ssh

#### TODO
#####Probe模块：
NTLM信息识别收集：https://github.com/FeigongSec/NTLMINFO
https://github.com/RowTeam/SharpDetectionNTLMSSP
https://github.com/checkymander/Sharp-SMBExec/blob/master/SharpInvoke-SMBExec/Program.cs


- [ ] phpmyadmin weblogic tomcat httpBasic // phpmyadmin的爆破存在问题
- [x] REDIS未授权

- [ ] ZOOKEEPER未授权
- [ ] MS17010  ms17010探测：https://github.com/Gh057Pr1nc3/smb_scanner


- [x] 增加内置词典
- [ ] 增加输出CSV
- [x] log的输出带颜色
- [x] 增加请求间隔延迟 --delay，当设定这个选项的时候，线程强制设为1，这个选项大概用不上？
- [ ] 变量名和函数名优化
~~- [ ] SMB和OXID输出的中文乱码问题~~
- [ ] **尝试改造为interface实现**
httpx -title --follow-redirects --status-code -tech-detect --title -ports 8000,8080,8888
###参考
* <https://github.com/shadow1ng/fscan>
* <https://github.com/k8gege/LadonGo>
* <https://github.com/OJ/gobuster>
* <https://github.com/netxfly/x-crack>
