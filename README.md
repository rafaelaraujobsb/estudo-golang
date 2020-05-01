# estudo-golang
Estudos na linguagem Go para usar Hyperledger Fabric

Apresentada no ano de 2009
    Robert Griesemer, Rob Pike e Ken Thompson
Livre e de código aberto

Pontos de partida: os criadores trabalhavam muito com C++ e por ser uma linguagem complexa
não gostavam de trabalhar com ela, com isso, começaram a perceber que o google precisava
de uma linguagem melhor para os seus produtos. Ela deveria ser limpa, pequena, compilada
e moderna

Princípios:
Simples, segura (memória) e legibilidade
Mínima: uma forma de escrever um determinado tipo de código
    Operadores unários
    Apenas uma forma de fazer um laço
Consolidação de várias ideias e não inovação

Não é OO, trouxe apenas métodos e interfaces

Concorrência (NewSqueak - Rob Pike)

* Compilada
* Multicore
* Concorrência está dentro da linguagem
* Fortemente tipada

Instalação: https://golang.org/doc/install?download=go1.14.2.linux-amd64.tar.gz

Comandos iniciais:
go version
go env
go run teste.go
go help comando
godoc --http=:8080
go vet comandos.go
go build comandos.go
go get -u github.com/go-sql-driver/mysql

https://github.com/cod3rcursos/curso-go
