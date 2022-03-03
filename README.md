


SQLCMD模块
* [数据库利用工具](http://ryze-t.com/posts/2022/02/16/%E6%95%B0%E6%8D%AE%E5%BA%93%E8%BF%9E%E6%8E%A5%E5%88%A9%E7%94%A8%E5%B7%A5%E5%85%B7-Sylas.html]
* [MDUT](https://github.com/SafeGroceryStore/MDUT)


https://github.com/sairson/Yasso/blob/6a99f1143d78e4c8224e49d00c0cfae39353f893/cmd/tools.go#L100

//https://stackoverflow.com/questions/27803654/explanation-of-checking-if-value-implements-interface
https://stackoverflow.com/questions/59831642/how-to-get-a-list-of-a-structs-methods-in-go
//检查某个方法是否实现了接口：https://go.dev/play/p/tNNDukK4wRi

redis密码一样的时候去重 已完成
redis和Mysql之类的Extra信息，可以增加一个result的字段 已完成


* 完成SQLCMD模块
  -m ls  <dst path>
  -m cat <dst file>
  -m upload <src path> <dst path>
  -m exec <cmd string>


Sqlcmd 传入多个参数：
http://liuqh.icu/2021/11/07/go/package/28-cobra/

设计一下sqlcmd的使用
cube sqlcmd -s 127.0.0.1 -l root -p root -x mssql exec "whoami"
cube sqlcmd -s 127.0.0.1 -l root -p root -x mssql upload  <src> <dst>
cube sqlcmd -s 127.0.0.1 -l root -p root -x mssql ls  <src>
cube sqlcmd -s 127.0.0.1 -l root -p root -x mssql cat  <src> 

* Probe充分测试
* 完善Help信息， Probe和Crack的
* 增加打印version信息
* 添加代码注释
* 完成sqlcmd
* 
* 线程异常超时退出
  https://www.cnblogs.com/bigdataZJ/p/golang-timeout.html
  
httpbasic的时候还检查了端口是否开放
phpmyadmin需要测试一下