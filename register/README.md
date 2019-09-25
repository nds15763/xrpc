自动负载均衡

Nginx 的动态负载均衡实现流程如下：

以相同的 Consul 标签对 Web Server 进行服务标记和分类，新增或者删除 Web Server 服务器节点；
Registrator 监控到 Web Server 的状态更新，自动在 Consul服务注册中心将它注册或者注销；
Consul-template 订阅了 Consul 服务注册中心的服务消息，接收到 Consul 的消息推送，即 Web Server 服务节点状态发生改变。
Consul-template 自动去修改和替换 Nginx 服务器下的 nginx配置文件中的模板，并重新加载服务达到自动负载均衡的目的。

作者：零壹技术栈
链接：https://juejin.im/post/5b2a6bc351882574cf66a211
来源：掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


现在有两种方式，第一种是我可以从consul中取出ip地址，自己写算法去负载均衡
第二种方式是我直用上面的方式去做这个负载均衡。

第一种实现可以用一致性哈希，但是这个我还没研究
第二种就直接写一些配置就好了。我是觉得第二种比较靠谱