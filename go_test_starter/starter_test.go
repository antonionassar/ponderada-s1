package starter_test

// importando pacotes para operações de entrada/saida, simular requisições http e pacotes e funções de teste
import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
)
//função de teste SayHello
func TestSayHello(t *testing.T) {
	// Chama a função SayHello com o nome "William"
	greeting := starter.SayHello("William")
	// Verifica se o greeting gerado está correto
	assert.Equal(t, "Hello William. Welcome!", greeting)
	// Chama a função SayHello com outra string
	another_greeting := starter.SayHello("asdf ghjkl")
	// Verifica se o greeting gerado está correto p/ a segunda entrada
	assert.Equal(t, "Hello asdf ghjkl. Welcome!", another_greeting)
}
// função de teste para OddOrEven
func TestOddOrEven(t *testing.T) {
	// testes para números não negativos agrupados
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		// Verifica se 45 é identificado como ímpar
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		// Verifica se 42 é identificado como par
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		// Verifica se 0 é identificado como par
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})

	// testes para números negativos agrupados
	t.Run("Check Negative Numbers", func(t *testing.T) {
		// Verifica se -45 é identificado como ímpar
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		// Verifica se -42 é identificado como par
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
}

// funcão de teste para Checkhealth
func TestCheckhealth(t *testing.T) {
	// verificar o status de saúde agrupado
	t.Run("Check health status", func(t *testing.T) {
		// Cria uma nova requisição HTTP 
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		// Cria um gravador para capturar a resposta
		writer := httptest.NewRecorder()
		// Chama a função Checkhealth, passando o gravador e a requisição
		starter.Checkhealth(writer, req)
		// Obtém a resposta
		response := writer.Result()
		// Lê o body da resposta
		body, err := io.ReadAll(response.Body)

		// Verifica se o body da resposta está correto
		assert.Equal(t, "health check passed", string(body))
		// Verifica se o status da resposta é 200 OK
		assert.Equal(t, 200, response.StatusCode)
		// Verifica se o cabeçalho Content-Type está correto
		assert.Equal(t,
			"text/plain; charset=utf-8",
			response.Header.Get("Content-Type"))

		// Verifica se não houve erro ao ler o corpo da resposta
		assert.Equal(t, nil, err)
	})
}