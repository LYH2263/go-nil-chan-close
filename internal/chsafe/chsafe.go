package chsafe

func Close(ch chan struct{}) {
	// 可选回调通道可能未初始化，需对 nil 做安全关闭
	if ch == nil {
		return
	}
	close(ch)
}
