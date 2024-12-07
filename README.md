# Trabalho 1 de Criptografia

## Como usar

### Compilação

```sh
mage
```

Caso não tenha mage, você pode fazer `go run mage.go`.

Caso queira instalar mage, você pode usar o target ensureMage:
```sh
go run mage.go ensureMage
```

O binário gerado será salvo em `./bin`.

### Uso

Para detalhes sobre como usar o programa, execute:

```sh
./bin/t1 --help
```

#### Exemplos

-   `./bin/t1 enc example/key example/text > enc`

    Criptografa o texto `example/text` usando a chave `example/key` e a minha
    implementação do AES.

-   `./bin/t1 dec example/key enc > dec`

    Descriptografa o texto criptografado `enc` usando a chave `example/key` e a
    minha implementação do AES.

-   `./bin/t1 enc -s example/key example/text > enc`

    Criptografa o texto `example/text` usando a chave `example/key` e a
    implementação do AES da biblioteca padrão do Go.

-   `./bin/t1 dec -s example/key enc > dec`

    Descriptografa o texto criptografado `enc` usando a chave `example/key` e a
    implementação do AES da biblioteca padrão do Go.

## Detalhes da minha implementação do AES

-   Troquei a caixa S por uma cifra de substituição simples (monográfica).

    [sbox](./internal/myaes/sbox.go)

-   Fiz a criptografia/descriptografia dos blocos em paralelo.

    [myaes](./internal/myaes/myaes.go)

## Benchmark

Essa é a implementação da criptografia de um bloco da biblioteca padrão Go:

```go
func encryptBlockGo(xk []uint32, dst, src []byte) {
	_ = src[15] // early bounds check
	s0 := byteorder.BeUint32(src[0:4])
	s1 := byteorder.BeUint32(src[4:8])
	s2 := byteorder.BeUint32(src[8:12])
	s3 := byteorder.BeUint32(src[12:16])

	// First round just XORs input with key.
	s0 ^= xk[0]
	s1 ^= xk[1]
	s2 ^= xk[2]
	s3 ^= xk[3]

	// Middle rounds shuffle using tables.
	// Number of rounds is set by length of expanded key.
	nr := len(xk)/4 - 2 // - 2: one above, one more below
	k := 4
	var t0, t1, t2, t3 uint32
	for r := 0; r < nr; r++ {
		t0 = xk[k+0] ^ te0[uint8(s0>>24)] ^ te1[uint8(s1>>16)] ^ te2[uint8(s2>>8)] ^ te3[uint8(s3)]
		t1 = xk[k+1] ^ te0[uint8(s1>>24)] ^ te1[uint8(s2>>16)] ^ te2[uint8(s3>>8)] ^ te3[uint8(s0)]
		t2 = xk[k+2] ^ te0[uint8(s2>>24)] ^ te1[uint8(s3>>16)] ^ te2[uint8(s0>>8)] ^ te3[uint8(s1)]
		t3 = xk[k+3] ^ te0[uint8(s3>>24)] ^ te1[uint8(s0>>16)] ^ te2[uint8(s1>>8)] ^ te3[uint8(s2)]
		k += 4
		s0, s1, s2, s3 = t0, t1, t2, t3
	}

	// Last round uses s-box directly and XORs to produce output.
	s0 = uint32(sbox0[t0>>24])<<24 | uint32(sbox0[t1>>16&0xff])<<16 | uint32(sbox0[t2>>8&0xff])<<8 | uint32(sbox0[t3&0xff])
	s1 = uint32(sbox0[t1>>24])<<24 | uint32(sbox0[t2>>16&0xff])<<16 | uint32(sbox0[t3>>8&0xff])<<8 | uint32(sbox0[t0&0xff])
	s2 = uint32(sbox0[t2>>24])<<24 | uint32(sbox0[t3>>16&0xff])<<16 | uint32(sbox0[t0>>8&0xff])<<8 | uint32(sbox0[t1&0xff])
	s3 = uint32(sbox0[t3>>24])<<24 | uint32(sbox0[t0>>16&0xff])<<16 | uint32(sbox0[t1>>8&0xff])<<8 | uint32(sbox0[t2&0xff])

	s0 ^= xk[k+0]
	s1 ^= xk[k+1]
	s2 ^= xk[k+2]
	s3 ^= xk[k+3]

	_ = dst[15] // early bounds check
	byteorder.BePutUint32(dst[0:4], s0)
	byteorder.BePutUint32(dst[4:8], s1)
	byteorder.BePutUint32(dst[8:12], s2)
	byteorder.BePutUint32(dst[12:16], s3)
}
```

Como é possível perceber aqui, a implementação das etapas do AES estão todas
juntas. Dessa forma, não faz sentido comparar o custo de cada etapa do
algoritmo. Portanto, foi comparado apenas o tempo total da
criptografia/descriptografia do mesmo arquivo.

Os testes foram realizados com:
-   image.jps (26K)
-   tgpl.pdf (5,1M)

e em uma máquina com:
-   CPU: 11th Gen Intel i5-1135G7 (8) @ 4.200GHz 
-   GPU: Intel TigerLake-LP GT2 [Iris Xe Graphics] 
-   Memory: 7740MiB 

### Criptografia


| **Arquivo**  | **Minha Implementação** | **Implementação da Biblioteca Padrão** |
|-----------|--------|--------| 
| image.jpg | 0,007s | 0,005s |
| tgpl.pdf  | 0,151s | 0,031s |

### Descriptografia

| **Arquivo**  | **Minha Implementação** | **Implementação da Biblioteca Padrão** |
|-----------|--------|--------|
| image.jpg | 0,011s | 0,005s |
| tgpl.pdf  | 0,417s | 0,036s |


## Conclusão

A biblioteca padrão do Go é mais rápida em praticamente todos os casos. Esse
resultado já era esperado, porque a implementação padrão usa uma caixa-s
constante e várias "tables-lookup" para pré-calcular algumas rodadas do AES
além de usar uma implementação de Assembly quando suportado.  Enquanto isso, a
minha implementação é totalmente feita do zero em Go somente, não usa essas
tables-lookup's e tem que calcular a caixa s toda vez.
