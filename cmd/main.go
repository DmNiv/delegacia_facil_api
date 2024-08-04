package main

import (
	"test/adapter/http"
	"test/container"
)

func main() {
	// Inicializa o container de dependências
	cont := container.NewContainer()
	
	// Configura o router com os handlers do container
	r := http.SetupRouter(cont)

	// Inicia o servidor
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
