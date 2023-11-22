# Grafos e Implementações

printf("olá, mundo\n");

# Atenção, corretor - TP3!

Grandes corretores! Eis a parte 3 do trabalho de grafos da minha dupla para o período
2023/2. Com grande prazer que anuncio que os testes já podem ser executados em suas
respectivas máquinas e o trabalho foi completamente implementado até o dia anterior
da data de entrega. Um grande alívio para nós - e talvez apenas para nós.

De nossa parte, restam os benchmarks.

**OBS0:** Para este trabalho, foi implementado apenas a lista de adjacências.

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

Para este trabalho prático, a única função implementada foi

```{go}
func TestFord(t *testing.T)
```

nela, você poderá alterar o delta, a fonte e o sumidouro. Além disso, poderá adaptar os nomes
dos arquivos na variável <em>filename</em> que é um array de string.

Para testar, basta fazer:

```{bash}
cd src/graphs
go test -run TestFord
```

é recomendável que coloque um limite
maior para o <em>timeout</em> para que
suporte por mais tempo, fazendo:

```{bash}
go test -run TestFord -timeout=9999s
```

**OBS1:** Os arquivos de teste são aqueles que terminam com <em>_test.go</em>

# Apresentação

Os códigos deste mísero repositório têm por objetivo analisar o desempenho de algorítmos de <em>Grafos</em> em uma linguagem de programação relativamente conhecida: Golang! Você pode acessar o website oficial dessa Linguagem de Programação [aqui](https://go.dev/).

**OBS2:** Se for pesquisar essa linguagem no Google, trate de pesquisar por <em>**GOLANG**</em> e não apenas <em>GO</em>. Reflita sobre essa afirmação!

A versão utilizada para este trabalho é a **1.21.0** e você pode acompanhar todas as versões no próprio website deles, [neste link](https://go.dev/dl/).
Se você utiliza Linux, é recomendável utilizar um versionador de linguagens de programação universal como o [asdf](https://asdf-vm.com/guide/getting-started.html).

## Instruções

### Importando as bibliotecas

Você perceberá que não existem executáveis no repositório, afinal, se outro Sistema Operacional, que não aquele em que foi escrito o programa, tentar utilizar o mesmo executável, problemas desastrosos podem ocorrer. Portando, você precisa gerar o executável para utilizar o programa ou utilizar a ferramenta do go que "executa um arquivo".

Para utilizar as bibliotecas em algum arquivo de teste (.go) você pode importar a biblioteca <em>graphs</em> e escolher
um de seus métodos de criação:
- FList

Para utilizá-las basta escolher os métodos e ler o arquivo de entrada.

```{go}
// Importando como uma lista de adjacência
import (
    "github.com/mhscardoso/graphs"
)

func main() {
    A := graphs.FList()

    graphs.ReadFile(A, "filename.txt")
}
```

Repare que a função para ler o arquivo é a mesma dos outros trabalhos.

Geralmente, para teste, utilizava o **<em>main.go</em>**, que é um arquivo na raiz do projeto, para testar as funcionalidades e importar funções.

### Compilando um código em Go

Estando com seu arquivo com os testes feitos, suponhamos que seja **<em>main.go</em>**

```
- goraphs
    - src
        - alist
        - amatrix
        - flist
        - ...
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
    - src
        - alist
        - amatrix
        - flist
        - ...
    - ...
    - main.go
    - main.exe
```

isso, se você utiliza windows. Se você utiliza linux, provavelmente nem precisa
deste mero tutorial!

### Executando sem ter gerado um executável

Uma forma de rodar o código apenas em um arquivo, é fazer:

```
go run main.go
```

que o código em **<em>main.go</em>** irá rodar.
