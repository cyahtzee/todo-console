package app

type EditTab struct {
	Tab
}

func (t *EditTab) HandleInput(input string) (string, string) {
	return "", ""
}
