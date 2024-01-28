# GO

Learning GO (basic)
<img src="/gopher-image.png">

# Sobre GO

> Golang é uma linguagem de programação de código aberto, ou seja, qualquer pessoa pode inspecionar, modificar e aprimorar. ELa é uma linguagem _compilada_ e _estaticamente tipada_

- Compilada: é quando o código-fonte é executado diretamente pelo sistema operacional ou pelo processador. Isso ocorre após ele ser traduzido por um processo de compilação que utiliza um programa chamado compilador

- Estaticamente tipada: define no código os tipos das variáveis de um programa. Assim, eles são conhecidos durante a compilação.

## Bytes[] vs Rune[]

OS 2 SÃO UTILIZADOS PARA MANIPULAR STRINGS

`Byte`

- 8 bits (1 byte) na memória
- Atendem ao padrão ASCII

`Rune`

- 34 bites (4 bytes) na memória
- Atendem alguns caracteres que não existem no padrão ASCII, por exemplo, que fazem parte do UTF-8, mas não do ASCII

> Como escolher byte ou rune?

- ASCII Characters:
  - Se você estiver lidando apenas com caracteres ASCII, pode optar por usar byte, pois cada caractere ASCII é representado por 1 byte.
- UTF-8 (Unicode Transformation Format) and Unicode Characters:
  - Se estiver lidando com caracteres que ultrapassam o conjunto ASCII, como caracteres acentuados, emojis, ou qualquer caractere do Unicode, é mais apropriado usar rune. Isso ocorre porque rune representa caracteres em Unicode e é capaz de lidar com caracteres multibyte, como os codificados em UTF-8.
- Eficiência de Memória:
  - Em termos de eficiência de memória, byte geralmente ocupa menos espaço do que rune, já que byte é apenas 1 byte, enquanto rune é 4 bytes. Se a economia de memória for crucial e você estiver lidando apenas com caracteres ASCII, byte pode ser mais eficiente.
- Velocidade de Iteração:
  - Em relação à velocidade de iteração em loops sobre strings, a escolha entre byte e rune também dependerá do tipo de manipulação que você está realizando. Se estiver manipulando caracteres individualmente e não precisar se preocupar com caracteres multibyte, usar byte pode ser mais eficiente em termos de desempenho. No entanto, se precisar lidar com caracteres Unicode, rune é mais apropriado para garantir que cada caractere seja tratado corretamente.
- Manutenção e Leitura do Código:
  - Além disso, considere a clareza e a manutenção do código. Se você estiver trabalhando com Unicode e caracteres multibyte, usar rune pode tornar o código mais legível, pois expressa a intenção de lidar com caracteres, independentemente do tamanho na memória.

## Ponteiros

Retém o endereço da memória dos dados, mas não os dados em si.
O endereço da memória diz à função onde encontrar os dados, mas não o valor dos dados. Você pode enviar o ponteiro para a função em vez dos dados e a função poderá, então, alterar a variável original em seu lugar original. Isso é chamado de passagem por referência porque o valor da variável não é enviado para a função, apenas a sua localização.

Se colocar um “&” na frente de um nome de variável, você está declarando que quer obter o ENDEREÇO ou um ponteiro para aquela variável

O segundo elemento da sintaxe é o uso asterisco (\*) ou do operador de desreferenciação

Quando se declara uma variável de ponteiro, vc segue o nome da variável para a qual o ponteiro aponta, prefixado com \*, desta forma:

> Var myPointer \*int32 = &someint

É mais comum usar ponteiro:

- Ao definirmos argumentos de funções e valores de retorno
- Definirmos métodos em tipos personalizados
- Quando estiver usando com grandes quantidades de dados

Quando você escreve uma função, pode definir os argumentos a serem passados por valor ou por referência

- _Passar por valor_ > uma cópia desse valor é enviada para a função e quaisquer alterações feitas no argumento - dentro daquela função - afetam somente a variável naquela função e não o local de onde ela foi passada.

- _Passar por referência_ > ou seja, passar um ponteiro para aquela argumento, vc poderá alterar o valor de dentro da função e também alterar o valor da variável original que foi enviado.

```go
package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var ptr *int = &a

	fmt.Printf("Endereço de a: %p\n", &a)
	fmt.Printf("Endereço de ptr: %p\n", ptr)
}
```

Saída

> Endereço de a: 0xc00001c1e0
> Endereço de ptr: 0xc00001c1e0

## Aplicações

Conceitos importantes:

- Serialização é quando um objeto é transformado, em uma cadeia de bytes e desta forma pode ser manipulado de maneira mais fácil, seja através de transporte pela rede ou salvo no disco.
- Decodificar: Geralmente, é usado quando se trata de converter dados de uma representação codificada (como JSON, XML, etc.) para uma representação nativa em uma linguagem de programação específica.
- Desserializar: É comumente usado quando se trata de converter dados serializados (geralmente em um formato binário) de volta para uma forma utilizável em uma linguagem de programação.

`MARSHAL`

A função Marshal no pacote encoding/json em Go é usada para converter (serializar) dados em formato JSON. A serialização é o processo de converter dados de uma representação interna, como uma estrutura de dados em Go, para um formato que pode ser facilmente transmitido ou armazenado, como JSON.

`UNMARSHAL`

A função Unmarshal do pacote encoding/json em Go é usada para decodificar (desserializar) dados JSON para uma estrutura de dados em Go. O processo de decodificação converte dados JSON, que geralmente são transmitidos em formato de texto, para uma estrutura ou tipo de dados na linguagem de programação.

json.Unmarshal O espera que você forneça um slice de bytes ([]byte) que contenha os dados JSON que você deseja decodificar. O tipo de dados []byte é usado para representar uma sequência de bytes, e é uma escolha natural para armazenar dados em formatos de texto, como JSON.

## Interface Writer | os | json

1.  INTERFACE `io.Writer`

        Type Writer interface {
        	Write(p []byte) (n int, err error)		}

    Essa interface é projetada para representar qualquer tipo que pode receber um slice de bytes e grava-los em um fluxo de dados. Isso proporciona uma abstração comum para operações de escritas em go. Ele permite que vc escreva dados em algo que implemente a interface io.Writer

2.  PACOTE `os`

    O pacote `os` em GO fornece funcionalidades para interagir com o sistema operacional. No contexto da aula, estamos interessados na implementação do método `Write` para o tipo `File` no pacote `os`

        package os

        func (f  *File) Write(b []byte) (n int, err error)

    Acima, `os.File` implementa a interface `os.Writer`, significando que você pode usar objetos do tipo `*os.File` em qualquer lugar onde `io.Writer` é esperado

3.  PACOTE `json`

    No pacote `json`, o construtor `NewEncoder` cria um novo codificador JSON que escreve em um `io.Writer`

    func NewEncoder(w io.Writer) \*Encoder

    Isso significa que você pode passar qualquer objeto que implemente `io.Writer` para `NewEncoder`, incluindo objetos do tipo `os.File`.
    Ele pega um io.Writer como entrada e grava a saída codificada json no fluxo de dados subjacente que implementa o io.Writer

4.  Uso do `Println` e `Fprintln`

    A função `Println` em Go é uma função conveniente para escrever no padrão de saída. Ela usa `Fprintln` internamente

        func Println(a …interface{}) (n int, err error) {
        	return Fprintln(os.Stdout, a…)
        }

    Aqui, `Fprintln` é uma função mais genérica que escreve em qualquer `io.Write`

    func Fprintln(w io.Write, a …interface{}) (n int, err error) {}

5.  Padrão de Saída (`os.Stdout`)

    A variável `os.Stdout` representa a saída padrão (stdout), que é o destino padrão para as funções como `Println`

    var Stdout = NewFile(uintptr(sys call.Stdout), ˜/dev/stdout˜)

    A função `NewFile` é usada para criar um objeto do tipo `os.File` a partir do descritor de arquivo associado à saída padrão

    func NewFile(fd unitptr, name string) \* FIle

————————————————————————————

RELACIONANDO OS TÓPICOS

- A interface `io.Writer` é a chave que permite que diferentes pacotes e bibliotecas em GO compartilhem um contrato comum para operações de escrita
- O pacote `os` fornece implementações especificas para operacoes de escritas em arquivos, incluindo a saída padrão (`os.Stdout`)
- O pacote `json` utiliza a interface `io.Writer` em seu design, permitindo que os clientes forneçam qualquer objeto que implemente essa interface ao criar um codificador JSON
- A função `Println` usa a função mais genérica `Fprintln`, que aceita qualquer `io.Writer`. Isso significa que vc pode usa-la para escrever qualquer destino que implemente `io.Writer`, não apenas saída padrão
- A variável `os.Stdout` representa a saída padrão e é um exemplo de um objeto que implementa a interface `io.Writer`

## Pacote sort

A furnção abaixo faz parte do pacote sort e ordena de forma personalizada

```go
func Sort(data Interface)
```

A interface Interface (passada por parâmetro no método Sort) é composta por 3 métodos:

1. método `Len` -> retorna o número de elementos do slice/coleção

   ```go
   func Len() int
   ```

2. método `Less` -> determina se um elemento vem antes do outro ou se ele tem que ser trocado pra ir depois

   ```go
   func (x Float64Slice) Less(i, j int) bool
   ```

- Compara os elementos nas posições i e j do slice e retorna verdadeiro se o elemento na posição i deve vir antes do elemento na posição j na ordenação.

3. método `Swap` -> determina as ordens de um elemento

   ```go
   func (x Float64Slice) Swap(i, j int)
   ```

- Tem um tipo receptor `Float64Slice`.
- É usada para trocar os elementos em um slice
- Recebe dois índices `i` e `j` que representam as posições de dois elementos em um slice.
- Ela troca os valores desses elemetnos, efetivamente trocando suas posições no slice

```go
func (x ordenaPorPotencia) Len() int {
    return len(x)ß
}

func (x ordenaPorPotencia) Less(i, j int) bool {
    return x[i].potencia < x[j].potencia

}

func (x ordenaPorPotencia) Swap(i, j int) {
    x[i], x[j] = x[j], x[i]
}
```

aqui foi criado um tipo ordenaPorPotencia que é um de slice de carro e IMPLEMENTA a Interface interface do pacote sort de GO

Dentro do método Sort do pacote sort, ele vai usar a interface que foi implementada no tipo ordenaPorPotencia pra ordenar usando bubble sorte, porém poderia escolher outro como quick sort, insertion sort, selection sort, merge sort, heap sort, shell sort, comb sort, vai depender do caso.

——————————————————————————————

A função sem ser chamada pelo método Sort e sem ter a interface implementada:

```go
func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}
```

Mas se ela for usada pelo método Sort e o objeto/slice dela tiver implementado a interface (Interface) corretamente

```go
func bubbleSort(arr []int) {
    for i := 0; arr.Len(); i++ {
        for j := 0; arr.Len()1; j++ {
            if arr.Less(j,i) {
               arr.Swap(i,j)
            }
        }
    }
}
```
