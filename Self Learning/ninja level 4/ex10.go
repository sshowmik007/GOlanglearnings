// /Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["fleming_ian"] = []string{`01`, `02`, `03`}

	delete(m, `no_dr`)

	for k, v := range m {
		fmt.Println("this is the record fot ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
