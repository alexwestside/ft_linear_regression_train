package model

type Commander interface {
	NewCommand() error
}

func (l *Model) NewCommand() error {

	return nil
}
