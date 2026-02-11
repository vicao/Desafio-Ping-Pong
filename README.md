# Desafio Ping Pong

Programa em Go que demonstra concorrência com goroutines e canais: imprime alternadamente
as palavras "ping" e "pong" até o usuário encerrar a execução.

## Recursos
- Alternância garantida entre `ping` e `pong` usando canais de sinalização.
- Pausa configurável entre impressões (atualmente 1 segundo).
- Código simples e didático para aprendizado de goroutines e comunicação por canais.

## Arquitetura (visão rápida)
- `ping` e `pong`: duas goroutines responsáveis por enviar sinais ao canal de mensagens.
- `pingCh` / `pongCh`: canais do tipo `struct{}` usados apenas como sinalizadores para alternância.
- `msg`: canal de `string` usado para transportar as mensagens que serão impressas.
- `Imprimir`: goroutine que lê `msg`, imprime e faz `time.Sleep(1 * time.Second)`.

O fluxo de execução é: `pingCh` sinaliza → goroutine `Ping` envia "ping" para `msg` → `pongCh` é sinalizado →
goroutine `Pong` envia "pong" para `msg` → `pingCh` é sinalizado → repetir.

## Requisitos
- Go (1.20+ recomendado). Instale a partir de https://go.dev/dl/ se necessário.

## Como executar (PowerShell)

Executar sem compilar (modo desenvolvimento):

```powershell
go run .
```

Executar especificando arquivo:

```powershell
go run pingpong.go
```

Compilar e executar o binário (Windows):

```powershell
go build -o pingpong.exe
.\pingpong.exe
```

Parar a execução:
- Pressione `Ctrl+C` no terminal para interromper imediatamente.
- O programa também aguarda entrada (pressionar Enter) na `main` — isso pode ser usado para finalizar de forma controlada.

## Exemplos de uso
- Apenas execute `go run .` e observe a saída alternada:

```
pong
ping
pong
ping
...
```

## Possíveis melhorias
- Adicionar argumento de linha de comando para executar por N iterações e sair automaticamente.
- Permitir configurar o intervalo de tempo via flag `-interval`.
- Adicionar testes automatizados simples.

## Contribuição
1. Faça um fork do repositório.
2. Crie uma branch com a sua feature: `git checkout -b feat/minha-melhora`.
3. Faça commits claros e envie um Pull Request.

## Troubleshooting
- Erro `go: command not found`: instale o Go e garanta que a pasta `go/bin` esteja no `PATH`.
- Não use `go run .\pingpong.exe` — esse comando tenta interpretar um binário como fonte Go.

## Licença
Projeto de exemplo — sinta-se livre para usar como material de estudo.
