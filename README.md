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


### Probe
收集内网信息，可用插件：`oxid,ms17010,zookeeper`

```bash
cube probe -x oxid -i 192.168.2.1/24
cube probe -x ALL -i 192.168.2.1/24
```

#### TODO
- [ ] NTLM SSP信息收集扫描
  https://www.mi1k7ea.com/2021/02/24/%E6%8E%A2%E6%B5%8B%E5%86%85%E7%BD%91%E5%AD%98%E6%B4%BB%E4%B8%BB%E6%9C%BA/


### Crack
- [x] 批量扫描存在Bug，无法显示多个模块的扫描结果
- [x] 批量扫描存在Bug，使用多个模块扫描不准确

爆破弱密码，可用插件：`ssh，mysql，redis，elastic，ftp，httpbasic，mongo，mssql，phpmyadmin，smb，postgres, jenkins`

```
Examples:
cube crack -u root -p root -i 192.168.1.1 -x ssh
cube crack -u root -p root -i 192.168.1.1 -x ssh --port 2222
cube crack -u root,ubuntu -p 123,000111,root -x ssh -i 192.168.1.1
cube crack -u root -p root -i 192.168.1.1/24 -x ssh
cube crack -u root --pass-file pass.txt -i 192.168.1.1/24 -x ssh
cube crack -u root --pass-file pass.txt -i 192.168.1.1/24 -x ssh,mysql

phpmyadmin和httpbasic只能单独使用，不可组合:
cube crack -u root --pass-file pass.txt -i http://192.168.1.1 -x phpmyadmin
cube crack -u root --pass-file pass.txt -i http://192.168.1.1 -x httpbasic
cube crack -u root --pass-file pass.txt -i http://192.168.1.1 -x jenkins
```

sqlserver爆破密码的代码(Event Code): 18456

#### Sqlcmd
执行命令，可用插件： `ssh`,`mssql`,`mssql-wscript`,`mssql-com`,`mssql-clr`
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

### TODO
##### Probe模块：
NTLM信息识别收集：

NTLM SSP信息扫描：https://github.com/EddieIvan01/ntlmssp
https://github.com/FeigongSec/NTLMINFO
https://github.com/RowTeam/SharpDetectionNTLMSSP
https://github.com/checkymander/Sharp-SMBExec/blob/master/SharpInvoke-SMBExec/Program.cs


- [ ] 增加输出CSV
- [x] 增加sqlcmd的mssql命令执行
- [x] 增加请求间隔延迟 --delay，当设定这个选项的时候，线程强制设为1，这个选项大概用不上？
- [ ] 变量名和函数名优化、
- [ ] 增加蜜罐识别：<https://www.secrss.com/articles/28577>
~~- [ ] SMB和OXID输出的中文乱码问题~~
- [ ] **尝试改造为interface实现**

httpx -title --follow-redirects --status-code -tech-detect --title -ports 8000,8080,8888

### 参考
* <https://github.com/shadow1ng/fscan>
* <https://github.com/k8gege/LadonGo>
* <https://github.com/OJ/gobuster>
* <https://github.com/netxfly/x-crack>
* <https://github.com/mabangde/pentesttools/blob/master/golang/sqltool.go>
