# tcpproxy
本案例实现tcp代理实现的流程, 包括tcp粘包问题的解决


# 本案例实现的功能 
	1 tcp 粘包问题
		demo.go
	2 tcp 服务搭建 
		server.go
	3 tcp 代理
		proxy.go
	3 tcp 客户端连接
		client.go

# 使用的库
github.com/xutao123777/tcppack


# 流程
   server.go < proxy.go < client.go

go run server.go
go run proxy.go
go run client.go
