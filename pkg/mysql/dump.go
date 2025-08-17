package mysql

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"time"

	"go-mysqldump/pkg/config"
)

func Dump() {

	cfg := config.GetConfig()

	for dbItem := range cfg.Databases {
		databaseName := cfg.Databases[dbItem].Name
		charset := cfg.Databases[dbItem].Charset

		// Monta nome do arquivo
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		outputFile := fmt.Sprintf("%s/%s_%s.sql", cfg.Storage.Directory, databaseName, timestamp)

		// Verifica se o diretório de backup existe, caso contrário, cria
		if _, err := os.Stat(cfg.Storage.Directory); os.IsNotExist(err) {
			err := os.Mkdir(cfg.Storage.Directory, 0755)
			if err != nil {
				log.Fatalf("Erro ao criar o diretório de backup: %v", err)
			}
		}

		// Monta comando mysqldump

		slog.Info("Iniciando backup database", "database", databaseName)

		args := []string{
			"compose",
			"exec",
			"mysql56",
			"mysqldump",
			"--routines",
			"--single-transaction",
			"--skip-lock-tables",
			fmt.Sprintf("--default-character-set=%s", charset),
			"-h", "localhost",
			fmt.Sprintf("-P %d", cfg.Mysql.Port),
			"-u", "root",
			fmt.Sprintf("-p%s", "root"),
			databaseName,
		}

		outFile, err := os.Create(outputFile)
		if err != nil {
			slog.Debug("Erro ao criar arquivo de saída:", "error", err)
			return
		}
		defer outFile.Close()

		// Executa mysqldump
		cmd := exec.Command("docker", args...)
		slog.Debug("Comando a ser executado", "exec", cmd.String(), "output_file", outputFile)
		cmd.Stdout = outFile
		// cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			slog.Error("Erro ao executar mysqldump", "error", err)
		}

		slog.Info("Dumping database com sucesso", "database", databaseName)
		Compress(outputFile, cfg.Compress.Type)
	}
}
