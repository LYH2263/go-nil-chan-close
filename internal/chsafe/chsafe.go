package chsafe

func Close(ch chan struct{}) {
	// BUG: no nil check
	close(ch)
}
