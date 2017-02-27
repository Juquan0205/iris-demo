package service

type DemoService struct {

}

/**
  service层:
  处理业务逻辑
 */
func (*DemoService) DealBusinessLogic(name string) (string, error){
	if name == "" {
		return "hello, world", nil
	}
	return "hello, " + name, nil
}
