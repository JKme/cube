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

利用Go的继承和重写分类模块： https://mp.weixin.qq.com/s/3atC9Bt6SInM56kCysMEhw

套餐待加入：

NTLM信息识别收集：https://github.com/FeigongSec/NTLMINFO
https://github.com/RowTeam/SharpDetectionNTLMSSP

ms17010探测：https://github.com/Gh057Pr1nc3/smb_scanner

Go语言写的漏洞扫描，学习一下框架结构：https://github.com/jweny/pocassist

这部分的代码和x-crack有点共通：https://github.com/jweny/pocassist/blob/5f46cf8625e27e8786443342af74da43fdbedaf5/scripts/scripts.go#L10

https://github.com/checkymander/Sharp-SMBExec/blob/master/SharpInvoke-SMBExec/Program.cs

### 命令行参数：
使用Cobra来创建命令行：https://github.com/spf13/cobra

例子：
https://www.cnblogs.com/sparkdev/p/10856077.html

解析，组合，运行

Crack: phpmyadmin weblogic tomcat httpBasic

1. Crack模块开发
2. SMB和OXID输出的中文乱码问题

https://o-my-chenjian.com/2017/09/20/Using-Cobra-With-Golang/
https://monkeywie.cn/2019/10/10/go-cross-compile/

## Context
http://beangogo.cn/2021/03/08/golang-context%E5%BA%94%E7%94%A8/
https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/

- [ ] FTP弱口令
- [ ] SMB弱口令
- [ ] REDIS未授权

ZOOKEEPER未授权
MS17010

- [x] 增加内置词典
- [ ] 增加输出到文件
- [x] log的输出带颜色
- [x] 增加请求间隔延迟 --delay 
- [ ] probe的时候做一个大的分类， POC对应ms17010, INFO对应信息收集， ALL表示全部
- [ ] 变量名和函数名优化
- [ ] **尝试改造为interface实现**

delay 好像并没有什么卵用
https://github.com/ysrc/xunfeng/tree/40d40ecf55910019b8b904ef70ae1eebb6b6d26f/vulscan/vuldb