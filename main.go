package main

type conf struct {
	locationArea LocationArea
}

func main() {
	config := &conf{
		locationArea: LocationArea{},
	}
	startRepl(config)
}
