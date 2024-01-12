package po

var Beans []interface{}

func register(bean interface{}) {
	Beans = append(Beans, bean)
}
