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

-   Troquei a caixa S por uma caixa pseudo-aleatória.
    
    - [sbox](./internal/myaes/sbox.go)

-   Fiz a criptografia/descriptografia dos blocos em paralelo.

    - [encrypt](./internal/myaes/encrypt.go)
    - [decrypt](./internal/myaes/decrypt.go)
