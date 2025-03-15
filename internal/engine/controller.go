package engine

type Controller interface {
	Init(ctx Context)
	Tick(ctx Context)
}
