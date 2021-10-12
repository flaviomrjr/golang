PACKAGES
========

## Modulos

Conjunto de pacotes que você criou, externos e do proprio Go.

Criando modulo:
```shell
$ go mod init modulo
```

Compilando com o modulo
```shell
$ go build
```

Será criado um arquivo binário

Executando o binário
```shell
./modulo
```

Caso queira enviar o modulo para a raiz do Go, executar o comando `go install`

## Funções

Funções que iniciam com letra maiuscula são publicas e com letra minuscula privada, sendo possivel utilizar apenas dentro do pacote.

## Modulo Externo

Instalando:
```shell
$ go get github.com/badoux/checkmail
```