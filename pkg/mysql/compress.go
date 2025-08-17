package mysql

import (
	"log/slog"
	"os"
	"os/exec"
)

func Compress(filename, compressType string) {

	slog.Info("Iniciando compressão do arquivo", "filename", filename, "compress_type", compressType)

	args := createArgs(filename, compressType)

	cmd := exec.Command(compressType, args...)
	cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	slog.Debug("Comando a ser executado", "exec", cmd.String())
	err := cmd.Run()
	if err != nil {
		slog.Error("Erro ao executar a compactação do arquivo", "filename", filename, "error", err)
	}

	slog.Info("Arquivo compactado com sucesso", "filename", filename)
}

func createArgs(filename, compressType string) []string {
	var args []string

	switch compressType {
	case "zstd":
		args = []string{"-z", "-14", "--rm", filename}
	case "bzip2":
		args = []string{filename}
	default:
		args = []string{filename}
	}

	return args
}
