# Grafos e Implementações

printf("olá, mundo\n");

# Atenção, corretor - TP2!

Este aviso foi agora mais acima para avisar que este branch está em "manutenção" devido a segunda parte do trabalho.
É provável que façamos melhoras nas partes anteriores, mas alguns erros podem ser gerados pois nada está pronto
até (23/10/2023).

# Atenção, corretor!

Para fazer os estudos de caso, utilizamos os testes unitários do Go. Para utilizá-los, você pode entrar nas pastas com testes, (alist e amatrix) via terminal fazendo

```
cd src/graphs
```

e rodar todos os testes, fazendo

```
go test
```

ou rodar apenas uma função do estudo de caso, fazendo

```
go test -run <Function Name>
```

o nome das funções é visível nos arquivos de teste.

**OBS0:** Os arquivos de teste são aqueles que terminam com <em>_test.go</em>

# Apresentação

Os códigos deste mísero repositório têm por objetivo analisar o desempenho de algorítmos de <em>Grafos</em> em uma linguagem de programação relativamente conhecida: Golang! Você pode acessar o website oficial dessa Linguagem de Programação [aqui](https://go.dev/).

**OBS1:** Se for pesquisar essa linguagem no Google, trate de pesquisar por <em>**GOLANG**</em> e não apenas <em>GO</em>. Reflita sobre essa afirmação!

A versão utilizada para este trabalho é a **1.21.0** e você pode acompanhar todas as versões no próprio website deles, [neste link](https://go.dev/dl/).
Se você utiliza Linux, é recomendável utilizar um versionador de linguagens de programação universal como o [asdf](https://asdf-vm.com/guide/getting-started.html).

## Instruções

### Importando as bibliotecas

Você perceberá que não existem executáveis no repositório, afinal, se outro Sistema Operacional, que não aquele em que foi escrito o programa, tentar utilizar o mesmo executável, problemas desastrosos podem ocorrer. Portando, você precisa gerar o executável para utilizar o programa ou utilizar a ferramenta do go que "executa um arquivo".

**OBS2:** Se você for utilizar <em>**Matrizes de Adjacência**</em> FAÇA AS CONTAS!!! O Go não rodará o programa para matrizes muito grandes. Como base de comparação entre minha máquina e os grafos disponibilizados, a partir do grafo 5 em uma máquina com 12GB de memória RAM, a chamada que cria a matriz dá <em>panic</em> (mais sobre isso [aqui](https://go.dev/doc/effective_go#panic)).

Para utilizar as bibliotecas em algum arquivo de teste (.go) você pode importar a biblioteca <em>graphs</em> e escolher
um de seus métodos de criação:
- List
- Matrix

Para utilizá-las basta escolher os métodos e ler o arquivo de entrada.

```{go}
// Importando como uma matriz de adjacência
import (
    "github.com/mhscardoso/graphs"
)

func main() {
    A := graphs.Matrix()

    graphs.ReadFile(A, "filename.txt")
}
```

ou

```{go}
// Importando como uma lista de adjacência
import (
    "github.com/mhscardoso/graphs"
)

func main() {
    A := graphs.List()

    graphs.ReadFile(A, "filename.txt")
}
```

Repare que a função para ler o arquivo é a mesma, a qual recebe ou uma lista ou uma matriz. (Pode receber os dois tipos).

Geralmente, para teste, utilizava o **<em>main.go</em>**, que é um arquivo na raiz do projeto, para testar as funcionalidades e importar funções.

### Compilando um código em Go

Estando com seu arquivo com os testes feitos, suponhamos que seja **<em>main.go</em>**

```
- goraphs
    - alist
    - amatrix
    - ...
    - main.go
```

e com o PATH do seu sistema operacional preparado para chamar o compilador do terminal, basta fazer:

```[bash]
go build main.go
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
