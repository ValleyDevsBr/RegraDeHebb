package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const size = 10

// Inicializa a matriz 10x10 e os pesos
func inicializarMatriz() [size][size]int {
	var matriz [size][size]int
	return matriz
}

// Inicializa a matriz 10x10 e define a letra A
func inicializarLetraA() [size][size]int {
	var matriz [size][size]int
	// Adiciona os dados da letra A
	matriz[0] = [size]int{-1, -1, -1, 1, 1, 1, 1, -1, -1, -1}
	matriz[1] = [size]int{-1, -1, -1, 1, 1, 1, 1, -1, -1, -1}
	matriz[2] = [size]int{-1, -1, -1, 1, -1, -1, 1, -1, -1, -1}
	matriz[3] = [size]int{-1, -1, -1, 1, -1, -1, 1, -1, -1, -1}
	matriz[4] = [size]int{-1, -1, -1, 1, -1, -1, 1, -1, -1, -1}
	matriz[5] = [size]int{-1, -1, -1, 1, 1, 1, 1, -1, -1, -1}
	matriz[6] = [size]int{-1, -1, 1, -1, -1, -1, -1, 1, -1, -1}
	matriz[7] = [size]int{-1, 1, -1, -1, -1, -1, -1, -1, 1, -1}
	matriz[8] = [size]int{1, -1, -1, -1, -1, -1, -1, -1, -1, 1}
	matriz[9] = [size]int{1, -1, -1, -1, -1, -1, -1, -1, -1, 1}
	return matriz
}

// Inicializa a matriz 10x10 e define a letra B
func inicializarLetraB() [size][size]int {
	var matriz [size][size]int
	// Adiciona os dados da letra B
	matriz[0] = [size]int{1, 1, 1, 1, 1, 1, 1, 1, 1, -1}
	matriz[1] = [size]int{1, 1, -1, -1, -1, -1, -1, -1, -1, 1}
	matriz[2] = [size]int{1, 1, -1, -1, -1, -1, -1, -1, -1, 1}
	matriz[3] = [size]int{1, 1, -1, -1, -1, -1, -1, -1, 1, -1}
	matriz[4] = [size]int{1, 1, 1, 1, 1, 1, 1, 1, -1, -1}
	matriz[5] = [size]int{1, 1, 1, 1, 1, 1, 1, 1, -1, -1}
	matriz[6] = [size]int{1, 1, -1, -1, -1, -1, -1, -1, 1, -1}
	matriz[7] = [size]int{1, 1, -1, -1, -1, -1, -1, -1, -1, 1}
	matriz[8] = [size]int{1, 1, -1, -1, -1, -1, -1, -1, -1, 1}
	matriz[9] = [size]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	return matriz
}

// Treina o modelo usando a Regra de Hebb
func treinarHebb(entrada [size][size]int, y int, w *[size][size]int, b *int) {
	print("Matriz de entrada: ")
	for _, linha := range entrada {
		fmt.Println(linha)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Println("w[", i, "][", j, "] = ", (*w)[i][j], " + ", entrada[i][j], " * ", y)
			w[i][j] += entrada[i][j] * y
		}
	}
	print("Matriz de pesos: ")
	for _, linha := range *w {
		fmt.Println(linha)
	}

	*b += y
}

// Testa a matriz desenhada pelo usuário
func testarHebb(entrada [size][size]int, w [size][size]int, b int) int {

	
	deltaTeste := 0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			deltaTeste += w[i][j] * entrada[i][j]

		}
	}
	
	deltaTeste += b

	if deltaTeste >= 0 {

		return 1 // Letra A

	} else {

		return -1 // Letra B

	}
}

func main() {
	// vai conter toda a informação do treinamento
	b := 0

	w := inicializarMatriz()

	entradaUser := inicializarMatriz()

	fmt.Println("Matriz do usuario: ")
	for _, linha := range entradaUser {
		fmt.Println(linha)
	}

	// Desenha as letras (A e B) nas matrizes de entrada
	letraA := inicializarLetraA()
	letraB := inicializarLetraB()

	// Treinamento
	y1, y2 := 1, -1 // Saída esperada para as letras A e B
	treinarHebb(letraA, y1, &w, &b)
	treinarHebb(letraB, y2, &w, &b)


	myApp := app.New()
	myWindow := myApp.NewWindow("Regra de Hebb - Reconhecimento de Letras")

	// Inicializa a matriz de entrada 10x10
	entrada := inicializarMatriz()
	buttons := make([][]*widget.Button, size)
	for i := range buttons {
		buttons[i] = make([]*widget.Button, size)
	}

	// Função para alternar o estado de um botão
	toggleButton := func(i, j int) {
		if entrada[i][j] == 1 {
			entrada[i][j] = -1
			buttons[i][j].SetText("0")
		} else {
			entrada[i][j] = 1
			buttons[i][j].SetText("1")
		}
	}

	// Função auxiliar para criar um botão com valores capturados
	createButton := func(i, j int) *widget.Button {
		return widget.NewButton("0", func() {
			toggleButton(i, j)
		})
	}

	// Cria os botões da matriz
	grid := container.New(layout.NewGridLayout(size))
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			buttons[i][j] = createButton(i, j)
			grid.Add(buttons[i][j])
		}
	}


	// Rótulo para exibir o resultado
	resultLabel := widget.NewLabel("Clique em 'Executar' para ver o resultado")

	// Função que executa o algoritmo da Regra de Hebb
	runHebb := func() {

		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
			// Converte o texto para um valor inteiro
			valor, err := strconv.Atoi(buttons[i][j].Text)
			if err != nil {
				valor = 0
			}
			entradaUser[i][j] = valor
				
			}
		}
		
		fmt.Println("Matriz do usuario: ")
		for _, linha := range entradaUser {
			fmt.Println(linha)
		}

		// Testa a matriz teste desenhada 
		teste := testarHebb(entrada, w, b)

		// Exibe o resultado na interface
		if(teste == 1){
			resultLabel.SetText("O que você escreveu é uma letra A")
		}else{
			resultLabel.SetText("O que você escreveu é uma letra B")
		}

	}

	resetHebb := func() {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ { 
				entrada[i][j] = 0
				buttons[i][j].SetText("0")
			}
		}
		resultLabel.SetText("Clique em 'Executar' para ver o resultado")
	}

	testHebb := func() {

		entrada := inicializarLetraB()
		
		// Testa a matriz teste desenhada 
		teste := testarHebb(entrada, w, b)
		
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if entrada[i][j] == 1 {
					buttons[i][j].SetText("1")
				} else {
					buttons[i][j].SetText("0")
				}
			}
		}

		if(teste == 1){
			resultLabel.SetText("O que você escreveu é uma letra A")
		}else{
			resultLabel.SetText("O que você escreveu é uma letra B")
		}

	}

	// Botão para executar o algoritmo
	runButton := widget.NewButton("Executar", runHebb)
	resetButton := widget.NewButton("Limpar", resetHebb)
	testButton := widget.NewButton("Testar", testHebb)

	// Organiza os widgets na janela
	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Programa Regra de Hebb para reconhecimento de A e B"),
		grid,
		runButton,
		resetButton,
		testButton,
		resultLabel,
	))

	// Inicia a aplicação
	myWindow.ShowAndRun()
}
