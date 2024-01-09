package githosting

type Gitlab struct {
	Config Config
}

func (g *Gitlab) VerifyToken() error {
	return nil
}
