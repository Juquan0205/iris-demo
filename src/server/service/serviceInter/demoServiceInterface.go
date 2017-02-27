package serviceInter

type DemoServiceInterface interface {
	DealBusinessLogic(name string) (string, error)
}