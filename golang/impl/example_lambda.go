package impl

type ExampleLambda struct {
	isKurtosisCoreDevMode bool
}

func NewExampleLambda(isKurtosisCoreDevMode bool) *ExampleLambda {
	return &ExampleLambda{isKurtosisCoreDevMode: isKurtosisCoreDevMode}
}
