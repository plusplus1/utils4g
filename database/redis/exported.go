package redis

func NewClient(path string) *WrapClient {
	return mgr.Connect(path)
}

func NewQueue(queueName string, path string) Queue {
	if c := NewClient(path); c == nil {
		return nil
	} else {
		return &queueImpl{name: queueName, conf: c.conf.Path(), client: c}
	}
}
