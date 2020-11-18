package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Marcelo",
		"sobrenome": "Santos",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Alice",
			"segundo":  "Santos",
		},
		"test": {"teste": "teste"},
	}

	fmt.Println(usuario2)
	delete(usuario2, "test")
	fmt.Println(usuario2)

	usuario2["tempo"] = map[string]string{
		"hoje": "sol",
	}

	fmt.Println(usuario2)
}
