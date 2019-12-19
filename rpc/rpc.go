package rpcdemo

import "errors"

// Service.Method 来调用
type DemoService struct {
}

type Args struct {
	A, B int
}

/**
  两个参数
  args是处理的参数
  result是处理结果接收的指针
返回值是发生的错误
*/
func (DemoService) Div(args *Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
