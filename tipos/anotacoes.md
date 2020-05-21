* Os tipos dentro do struct não são separados por ","
* Structs possuem receivers que são os métodos
* Busque torna o uso de struct legível, passando os valores com os nomes deles
* Para alterar alguma coisa dentro da struct é necessário o uso de ponteiros (func (p *pessoas) x() {})
* Não existe heranças em GO e sim uma pseudo herança (campos anônimos)
* Ao criar o seu tipo de dados, é possível customizar receivers
* Quando um struct implementa todos os tipos de uma interface, o GO entende que ela é daquele tipo
* Ao utilizar inteface e for modificar um valor dela, é necessário atribuir para var com o &
    * Geralmente as interface são contruidas para a leitura de dados
* O tipo interface permite a criação de tipos mais genericos para uma variável
* O Go tem a convenção de identificadores começando com a letra maiúscula é público
* Para criar um json é preciso mapear o struct e usar json.Marshal(), que retornar o json e o erro