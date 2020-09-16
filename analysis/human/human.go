package human

type Human interface {
	Structure() []string
	Performance() string
}

type Man string

func (m Man) Structure() []string {
	var parts = []string{"Brain", "Heart", "Nose"}
	return parts
}

func (m Man) Performance() string {
	return "8 hrs/Day"

}
