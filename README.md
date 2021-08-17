# Dojo - Go

## Testes em Go

Go já possui uma biblioteca de testes embutida na linguagem chamada de **testing**. Recomendo fortemente, para quem deseja aprender mais sobre Go e TDD, o tutorial [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/).


## Como criar um teste

Primeiro crie um arquivo com o final **_test.go**; todas as funções de testes devem possuir o **Test** no começo, ex: TestFunctionName(t *testing.T). Não esqueça de **import "testing"**.

## Como rodar

```
go test ./...               # Esse comando roda todos os testes
go test filename_test.go    # Roda um arquivo específico de teste
go test filename_test.go -v # Flag para visualização mais verbosa
```