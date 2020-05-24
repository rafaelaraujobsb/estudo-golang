### Concorrência vs Paralelismo
* Concorrência intercala processos em um unico processador
    * Exige uma estruturação no código
* Paralelismo são os processos sendo executados em processadores diferentes

**Go é uma linguagem concorrente**

* A concorrência pode ser melhor que o paralelismo devido o overhad para fazer o paralelismo, que exige
mais do hardware e SO

* O comando "go" cria uma goroutines para executa a função informada de forma independente
    * A execução da goroutines só fica ativa enquanto o seu programa estiver no ar
* O channel (chan) é um tipo da linguagem que é usado para comunicar goroutines
    * Serve como uma forma de sincronismo para as goroutines
    * ch := make(chan int, 1) -> o segundo parâmetro é o buffer
        * O buffer ajuda a enviar mais itens para o canal
    * ch <- 1  enviando dados para o canal
    * <-ch recebendo dados do canal
    * close(ch) fecha o canal, ideal quando se tem laços de repetição
* Podemos ter um canal que é o multiplexador, ou seja, reune a saída de n canais em um único canal
* Select é uma estrtutra de controle para trabalhar com concorrência, estilo o switch