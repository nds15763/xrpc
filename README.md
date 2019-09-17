# xrpc
基于TCP的RPC框架

RPC框架包含几大部分

一、通讯

1、协议：TCP和HTTP2的优劣
TCP包较小，有更好的传输速率，但易用性没有HTTP2好
HTTP2虽然会稍微占一些流量，但是由于其具有一些链路复用，头部压缩等新特性，速度应该并不比TCP差不多少。
（我没试过）
但是这里为了学习还是使用TCP

2、加密
我因为不太了解这部分，只简单选用对称加密了。

3、yamux
yamux是一个链接复用库，可以使一个物理的TCP链接虚拟出多个逻辑链接，从而提高传输效率
源码阅读我有提交简书，不过这两天发不了文。

原理是利用一个可配置大小的buffer数组（可以理解为申请的内存块），通过随时判断数组占用大小，
使用channel去通知各个逻辑流的读写进度，多读少写。

4、链接池
虽然使用yamux可以虚拟逻辑链接，但每一个TCP链接的发送速率是有限的，这里还是需要管理一个TCP的连接池。


二、远程调用

1、反射、注册
本来想自己写的，但是看了一下rpc包的源码，感觉可以直接用RPC


服务注册：consul 和eureka 都可以
我这里选用consul的理由很简单，是因为他支持Windows环境，我电脑起虚拟机太累了