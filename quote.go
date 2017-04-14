package main

type Quote map[string]string

func (q *Quote) Speaker() string {
	return (*q)["speaker"]
}

func (q *Quote) Content() string {
	return (*q)["quote"]
}
