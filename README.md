# go-mysqldump

Uma aplicação CLI para realizar dumps do MySQL.

> Em desenvolvimento...

## Instalação

Clone o repositório e instale as dependências Go:

```sh
git clone <repo-url>
cd go-mysqldump
go build -o go-mysqldump
```

## Uso

Execute o comando principal:

```sh
./go-mysqldump --help
```

## Configuração

A configuração pode ser feita de três formas:

1. **Arquivo de configuração**: Por padrão, o app procura por um arquivo `config.yaml` no diretório atual. Você pode especificar outro arquivo usando a flag `--config`:
   ```sh
   ./go-mysqldump --config=new-config.yaml
   ```
2. **Variáveis de ambiente**: Todas as configurações podem ser definidas via variáveis de ambiente prefixadas por `GO_DUMP_`. Por exemplo, para definir `log.level` via env:
   ```sh
   export GO_DUMP_LOG_LEVEL=debug
   ```
3. **Flags de linha de comando**: Flags sempre têm prioridade sobre variáveis de ambiente e arquivos de configuração.

## Ordem de Precedência das Configurações

O objetivo é garantir uma precedência clara e previsível para as configurações, seguindo a hierarquia abaixo (da mais alta para a mais baixa):

1. **Flag de linha de comando** (ex: `--log.level=debug`)
2. **Variável de ambiente** (ex: `GO_DUMP_LOG_LEVEL=debug`)
3. **Arquivo de configuração** (ex: `log.level: debug` em `config.yaml`)
4. **Valor padrão** (ex: `info`)

Essa hierarquia garante que configurações efêmeras e específicas (como uma flag) sempre prevaleçam, enquanto configurações persistentes e gerais (como um arquivo de configuração) fornecem uma base conveniente.

## Exemplo de arquivo `config.yaml`

```yaml
aws:
  bucket: my-backup-bucket

log:
  path: /var/logs/go-mysqldump
  level: info

storage:
  directory: /var/backups/mysql

compress:
  type: bzip2

databases:
  - name: database1
    charset: latin1
  - name: database2
    charset: utf8
    ignore_tables:
      - table1
      - table2
      - table3
```

## Licença

MIT
