package output

type Output interface {
	Run()
	Stop()
	Wait()
}
