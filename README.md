# Grafos e Implementações

printf("olá, mundo\n");

Os códigos deste mísero repositório têm por objetivo analisar o desempenho de algorítmos de <em>Grafos</em> em uma linguagem de programação relativamente conhecida: Golang! Você pode acessar o website oficial dessa Linguagem de Programação [aqui](https://go.dev/).

**OBS1:** Se for pesquisar essa linguagem no Google, trate de pesquisar por <em>**GOLANG**</em> e não apenas <em>GO</em>. Reflita sobre essa afirmação!

A versão utilizada para este trabalho é a **1.21.0** e você pode acompanhar todas as versões no próprio website deles, [neste link](https://go.dev/dl/).
Se você utiliza Linux, é recomendável utilizar um versionador de linguagens de programação universal como o [asdf](https://asdf-vm.com/guide/getting-started.html).

## Instruções

### Importando as bibliotecas

Você perceberá que não existem executáveis no repositório, afinal, em outro Sistema Operacional que não aquele em que foi escrito o programa tentar utilizar o mesmo executável, problemas desastrosos podem ocorrer. Portando, você precisa gerar o executável para utilizar o programa.

**OBS2:** Se você for utilizar <em>**Matrizes de Adjacência**</em> FAÇA AS CONTAS!!! O Go não rodará o programa para matrizes muito grandes. Como base de comparação entre minha máquina e os grafos disponibilizados, a partir do grafo 5 em uma máquina com 12GB de memória RAM, a chamada que cria a matriz dá <em>panic</em> (mais sobre isso [aqui](https://go.dev/doc/effective_go#panic)).

Para utilizar as bibliotecas em algum arquivo de teste (.go) você pode importar as bibliotecas que têm o nome das estruturas de dados utilizadas para lidar com grafos: <em>list</em> e <em>matrix</em>. Para importar, você pode fazer:

```{go}
    // Importando como uma matriz de adjacência
    import (
        "github.com/mhscardoso/matrix"
    )

    func main() {
        A, err := matrix.CreateMatrix("filename.txt")
    }
```

ou

```{go}
    // Importando como uma lista de adjacência
    import (
        "github.com/mhscardoso/list"
    )
    
    func main() {
        A, err := list.CreateList("filename.txt")
    }
```

Geralmente, para teste, utilizava o **<em>main.go</em>**, que é um arquivo na raiz do projeto, para testar as funcionalidades e importar funções.

### Compilando um código em Go

Estando com seu arquivo com os testes feitos, suponhamos que seja **<em>main.go</em>**

```
- goraphs
    - list
    - matrix
    - ...
    - main.go
```

e com o PATH do seu sistema operacional preparado para chamar o compilador do terminal, basta fazer:

```[bash]
go build main.c
```

ficando com

```
- goraphs
    - list
    - matrix
    - ...
    - main.go
    - main.exe
```

isso, se você utiliza windows.

### Executando sem ter gerado um executável

Uma forma de rodar o código apenas em um arquivo, é fazer:

```
go run main.go
```

que o código em **<em>main.go</em>** irá rodar.
