package main

type Subscriber struct {
	name string
}

func New(name string, connector Connector) *Subscriber {
	r := &Subscriber{name}

	connector.Subscribe(r.print)

	return r
}

func (s *Subscriber) print(state string) {
	println(s.name, "receive new state:", state)
}
