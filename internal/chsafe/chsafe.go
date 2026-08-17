package chsafe

func Close(ch chan struct{}) {
	if ch == nil {
		return
	}
	close(ch)
}
