package vehicle

type Vehicle interface {
	Structure() []string
	Speed() string
}

type Car string

func (c Car) Structure() []string {
	var parts = []string{
		"ECU",
		"ENGINE",
		"Air Filters",
		"Wipers",
	}
	return parts
}

func (c Car) Speed() string {
	return "200 km/hrs"

}
