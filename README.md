# Dojo - Go

## Testes em Go

Go já possui uma biblioteca de testes embutida na linguagem chamada de **testing**. Recomendo fortemente, para quem deseja aprender mais sobre Go e TDD, o tutorial [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/).

## Como criar um teste

Primeiro crie um arquivo com o final **\_test.go**; todas as funções de testes devem possuir o **Test** no começo, ex: TestFunctionName(t \*testing.T). Não esqueça de **import "testing"**.

## Como rodar

```
go test ./...               # Esse comando roda todos os testes
go test filename_test.go    # Roda um arquivo específico de teste
go test filename_test.go -v # Flag para visualização mais verbosa
```

### Requisito 1

Escreva o método greet(name) que recebe um nome e retorna um simples cumprimento. Por exemplo, quando a entrada for "Bob", o método retornará "Hello, Bob."

### Requisito 2

Deve tratar valores nulos (nil, None, null etc) com um cumprimento genérico. Por exemplo, quando a entrada for null, o método retornará "Hello, my friend."

### Requisito 3

Precisa tratar quando se estiver berrando. Caso o nome esteja grafado todo em caixa alta, o método deverá gritar de volta. Por exemplo, quando a entrada for "JERRY", o método retornará "HELLO JERRY!"

### Requisito 4

Necessita tratar dois nomes na entrada. Quando o nome for um array com dois nomes (ou, caso a linguagem suporte, varargs ou splat),então ambos os nomes deverão ser impressos. Por exemplo, quando a entrada for ["Jill", "Jane"], o método retornará a mensagem "Hello, Jill and Jane."

### Requisito 5

Deve tratar uma quantidade arbitrária de nomes na entrada. Caso a lista contenha mais de dois elementos o último item deverá terminar com ", and". Por exemplo se a entrada for ["Amy", "Brian", "Charlotte"], o método deverá retornar "Hello, Amy, Brian, and Charlotte."

### Requisito 6

Precisa permitir a utilização de nomes nomes normais e berrados separando a resposta em dois cumprimentos. Por exemplo, quando a entrada for [“Amy”, “BRIAN”, “Charlotte”], o método deverá retornar a mensagem “Hello, Amy and Charlotte. AND HELLO BRIAN!”

### Requisito 7

Se algum dos nomes usados na entrada possuir uma vírgula, este deverão ser quebrados em entradas individuais. Por exemplo, caso a entrada seja [“Bob”, “Charlie, Dianne”], o método retornará a mensagem “Hello, Bob, Charlie, and Dianne.”

### Requisito 8

Permitir que as vírgulas introduzidas pelo Requisito 7 sejam intencionalmente ignoradas. Para tal deve-se usar do mesmo modo que nos arquivos CSV, com o uso de aspas delimitando a entrada. Por exemplo, caso a entrada seja [“Bob”, “\”Charlie, Dianne\””], o método retornará a mensagem "Hello, Bob and Charlie, Dianne."
