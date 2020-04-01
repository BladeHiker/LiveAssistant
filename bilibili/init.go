package bilibili

func init() {
	// 初始化一些全局维护的变量
	P = NewPool()
	UserClient = NewClient()
}
