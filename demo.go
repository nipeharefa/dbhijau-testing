package demo

//go:generate mockgen --typed --package=mocks --destination=mocks/finder.go demo.nipeharefa.dev/demo Finder
type (
	Finder interface {
		Say() string
	}
)
