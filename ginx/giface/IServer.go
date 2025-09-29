package giface

/**
 * server-服务器接口层
 * 定义：服务器启动、监听、停止等方法
 */
type IServer interface {
	//Serve server执行方法
	Serve()
	// Start 开启服务器
	Start()
	// ListenAndServe 监听server 端口
	ListenAndServe()
	// Stop 停止运行
	Stop()
}
