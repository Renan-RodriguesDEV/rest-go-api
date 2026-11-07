# REST Go API

API REST desenvolvida em Go. Este documento explica o necessário para executar, compilar e entender o projeto sem esconder os conceitos básicos.

## Índice

- [Ferramentas](#ferramentas)
- [Pré-requisitos](#pré-requisitos)
- [Como iniciar](#como-iniciar)
- [Como compilar e gerar o build](#como-compilar-e-gerar-o-build)
- [Como foi estruturado](#como-foi-estruturado)
- [Fluxo de uma requisição](#fluxo-de-uma-requisição)
- [Glossário](#glossário)
- [Fontes](#fontes)

## Ferramentas

**Go (Golang)** é a linguagem usada no backend. Ela gera um programa compilado, possui tipagem estática e inclui ferramentas para formatar, testar e compilar o código.

**HTTP** é o protocolo usado para a comunicação entre cliente e servidor. A API recebe uma requisição, executa uma regra e devolve uma resposta com status, cabeçalhos e, normalmente, JSON.

**REST** é um estilo de construção de APIs que usa recursos identificados por URLs e os métodos HTTP: `GET` consulta, `POST` cria, `PUT` ou `PATCH` altera e `DELETE` remove.

**JSON** é o formato de texto usado para transportar dados. Exemplo: `{"name":"Ana"}`.

**Go Modules** gerencia o nome do módulo e as dependências no arquivo `go.mod`. O arquivo `go.sum` registra hashes para verificar dependências baixadas.

O pacote padrão `net/http` fornece servidor, cliente e tipos HTTP. Se o projeto usar um roteador ou banco externo, essas dependências devem aparecer no `go.mod`; não é necessário instalar uma ferramenta adicional manualmente.

## Pré-requisitos

Instale o Go pela página oficial e confirme no terminal:

```bash
go version
```

A versão exigida pelo projeto está na diretiva `go` do arquivo `go.mod`. Execute os comandos na pasta raiz, onde esse arquivo está localizado.

## Como iniciar

Baixe e organize as dependências:

```bash
go mod download
go mod tidy
```

Inicie a API sem criar um executável permanente:

```bash
go run .
```

O servidor normalmente fica disponível em `http://localhost:8080`. Use a porta definida pelo código ou por uma variável de ambiente. Para testar uma rota:

```bash
curl http://localhost:8080/
```

O terminal que executa `go run .` mantém o servidor ativo. Pressione `Ctrl+C` para encerrá-lo.

## Como compilar e gerar o build

Primeiro valide e formate o código:

```bash
gofmt -w .
go test ./...
go vet ./...
```

Gere um executável para o sistema atual:

```bash
go build -o rest-go-api .
```

No Windows, o arquivo resultante será `rest-go-api.exe` caso o nome não tenha extensão. Execute-o com:

```powershell
.\rest-go-api.exe
```

`go build` compila e preserva o executável. `go run` compila temporariamente e já executa. Assim, o build de produção deve ser feito com `go build` e o executável deve ser distribuído junto com as configurações necessárias.

Para compilar para Linux a partir do PowerShell:

```powershell
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o rest-go-api .
```

## Como foi estruturado

O ponto de entrada de uma aplicação Go é uma função `main` dentro do pacote `main`. Ela prepara as dependências, registra as rotas e inicia o servidor HTTP.

Uma rota associa método e caminho a um handler. O handler recebe `http.ResponseWriter` e `*http.Request`, lê parâmetros ou JSON, chama a regra necessária e escreve o status e a resposta.

O fluxo recomendado é separar responsabilidades: o handler trata HTTP; o serviço contém a regra de negócio; o repositório acessa banco ou outra fonte de dados. Essa separação facilita testes e evita colocar toda a lógica em uma única função. O nome exato das pastas depende da implementação deste projeto.

Entradas externas devem ser validadas antes do uso. Respostas devem informar um status HTTP correto, retornar JSON quando esse for o contrato e não expor senhas, tokens ou detalhes internos de erro.

## Fluxo de uma requisição

1. O cliente envia método, URL, cabeçalhos e, opcionalmente, um corpo JSON.
2. O servidor encontra a rota correspondente.
3. O handler valida a entrada e chama a regra de negócio.
4. A aplicação consulta ou altera dados quando necessário.
5. O servidor devolve status HTTP, cabeçalhos e corpo.

Para descobrir as rotas reais, procure por `HandleFunc`, `Handle`, `GET`, `POST` ou pelo registro do roteador no código. Procure também por `os.Getenv` para saber quais variáveis de ambiente são necessárias.

## Glossário

**API**: contrato que permite a comunicação entre programas.

**Endpoint**: combinação de método HTTP e URL que executa uma operação.

**Handler**: função que atende uma requisição HTTP.

**Rota**: regra que direciona uma URL e método para um handler.

**Status HTTP**: número que indica o resultado, como `200` sucesso, `201` criação, `400` entrada inválida, `404` não encontrado e `500` erro interno.

**Middleware**: código executado antes ou depois do handler, por exemplo para log, autenticação ou CORS.

**Dependência**: pacote externo usado pelo projeto.

**Build**: processo de transformar o código em um executável.

**Compilar**: converter o código-fonte em instruções que o sistema consegue executar.

**Variável de ambiente**: configuração fornecida pelo sistema operacional, sem gravá-la necessariamente no código.

**Idempotência**: propriedade de uma operação que produz o mesmo efeito quando repetida; `GET`, `PUT` e `DELETE` normalmente são tratados dessa forma, conforme o contrato da API.

## Fontes

- Go, documentação oficial: https://go.dev/doc/
- Go Command, incluindo `run`, `build`, `test` e `mod`: https://go.dev/ref/mod e https://pkg.go.dev/cmd/go
- Pacote HTTP padrão: https://pkg.go.dev/net/http
- RFC 9110, semântica HTTP: https://www.rfc-editor.org/rfc/rfc9110
- JSON, RFC 8259: https://www.rfc-editor.org/rfc/rfc8259
