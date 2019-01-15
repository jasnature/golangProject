# [Some golang Projects]  [一些 Go lang 语言开发的项目]  

## [Folder Description]  [项目文件介绍]  

### [ConnProxy 正/反代理工具]  
This is used go language to develope program of proxy server with support http and https.   
这个是使用GO 语言开发的一个程序，它是可以代理访问其他服务器资源的工具，支持HTTP和HTTPS。  

===Propertys and functions===   一些功能介绍  
* 1.can custom config to manager the proxy server(timeout,conn num,wait num,filter ip,etc...).  
* 2.can set direct proxy.  
* 3.can set reverse proxy list.  
* 4.can auto switch reverse proxy server or custom function to handle switch plot.  
* 5.can support multiple connection accept for high request.

* 1.可以自定义配置进行代理工具的设置，目前包含 超时，最大并发连接数，最大等待数，IP过滤，等等  
* 2.支持直接代理访问  
* 3.支持多反向代理服务器设置，类似于nginx代理切换功能。  
* 4.自动切换反向配置的代理服务器，也可以自己实现切换策略  
* 5.支持大并发连接。

### [GOBLog]  by golang program
This is log writer for program runing.
