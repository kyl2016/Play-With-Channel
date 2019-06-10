package main

type Connector interface {
	Subscribe(StateChanged func(state string))
}
