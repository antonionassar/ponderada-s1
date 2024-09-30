package starter

// importação de pacotes para formatação de strings e entrada/saida, pacotes de funções matematicas e de requisições http
import (
	"fmt" 
	"math" 
	"net/http" 
)

// SayHello retorna um greeting personalizado
func SayHello(name string) string {
	// Formata e retorna um greeting utilizando o nome fornecido
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

// OddOrEven determina se um número é ímpar ou par
func OddOrEven(num int) string {
	// Calcula o resto da divisão do número por 2
	criteria := math.Mod(float64(num), 2)
	// Verifica se o resto é 1 ou -1 para determinar se o número é ímpar
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", num)
	}
	// Caso contrário, o número é par
	return fmt.Sprintf("%v is an even number", num)
}

// Checkhealth verifica o status de saúde e retorna uma mensagem
func Checkhealth(writer http.ResponseWriter, req *http.Request) {
	// Escreve a mensagem de status de saúde na resposta HTTP
	fmt.Fprintf(writer, "health check passed")
}