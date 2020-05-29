* Em GO não se usa tanto os mocks
* Os arquivos de teste devem terminar com "_teste.go"
* go test -> executa os testes na pasta que você está  
* go test ./... -> executa os testes de todas as pastas
* t.Logf -> usado para os logs dos testes
    * go test -v
* t.Parallel() -> executa os testes em paralelos
* go test std -> executa os testes da biblioteca padrão do go
* O GO oferece relatórios de cobertura de testes
    * go test --cover
    * go test --coverprofile=resultado.out
    * go tool cover --func=resultado.out
        * go tool cover --html=resultado.out