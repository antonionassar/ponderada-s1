# Ponderada de Programação (Semana 1) - TDD

## Introdução ao TDD (Test-Driven Development)

Test-Driven Development (TDD) é uma metodologia de desenvolvimento de software que consiste em criar testes antes de escrever o código de uma funcionalidade. TDD segue três etapas principais: Red, Green e Refactor.

## Ciclo TDD

1. **Red (Falha)**

   - **Escrever um teste que falha**: Começamos criando um teste para a funcionalidade que desejamos. Como a funcionalidade ainda não existe, o teste deve falhar nesse estágio.
   - **Exemplo**:
     - Criamos os testes `TestSayHello`, `TestOddOrEven` e `TestCheckhealth`. Inicialmente, eles falham, pois as funções `SayHello`, `OddOrEven` e `Checkhealth` ainda não foram implementadas.

2. **Green (Sucesso)**

   - **Escrever o código para passar no teste**: Desenvolvemos o código necessário para que o teste passe, sem focar em aperfeiçoamentos. A meta é apenas que o teste tenha sucesso.
   - **Exemplo**:
     - Escrevemos as funções `SayHello`, `OddOrEven` e `Checkhealth` de forma que os testes correspondentes passem.

3. **Refactor (Refatoração)**
   - **Melhorar o código**: Após o teste passar, aprimoramos o código para torná-lo mais limpo e eficiente, garantindo que todos os testes continuem passando.
   - **Exemplo**:
     - Refatoramos as funções `SayHello`, `OddOrEven` e `Checkhealth`, mantendo a funcionalidade e os testes intactos.

## Vantagens do TDD

- **Melhor Design do Código**: TDD incentiva a criação de um código fácil de testar, o que melhora a manutenção.
- **Feedback Rápido**: Testes automatizados fornecem feedback imediato sobre alterações no código.
- **Confiança no Código**: Como os testes são escritos para cobrir casos esperados, há menos chances de erros.

## Documentação dos Testes

### IMGS

- Na pasta "imgs", voce irá encontrar os resultados para os testes "sayHello", "oddOrEven" e "checkHealth", respectivamente.
