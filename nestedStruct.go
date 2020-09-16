package main

import "fmt"

type salary struct {
	Basic, HRA, TA float64
}

type employee struct {
	Name, LastName, Email string
	MonthlyIncome         []salary
	Domain                []programming
	Lang                  []language
}

type programming struct {
	JS    string
	Go    string
	React string
}

type language struct {
	Lang []string
}

func main() {
	emp := employee{
		Name:     "Saran",
		LastName: "Raj",
		Email:    "saran@gmail.com",
		MonthlyIncome: []salary{
			salary{
				Basic: 30.00,
				HRA:   20.00,
				TA:    10.00,
			},
		},
		Domain: []programming{
			programming{
				JS: "JS",
			},
		},
		Lang: []language{
			language{
				Lang: []string{
					"js",
				},
			},
		},
	}

	fmt.Println("emp ", emp)
}
