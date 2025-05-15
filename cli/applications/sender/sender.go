package sender

type SimpleSender struct{}

func (s *SimpleSender) Run(str string) {
	println("false send was done")
}
