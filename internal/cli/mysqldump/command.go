package mysqldump

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/

import (
	"go-mysqldump/pkg/mysql"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dump",
		Short: "MySQL dump database",
		RunE: func(cmd *cobra.Command, args []string) error {
			mysql.Dump()
			return nil
		},
	}

	cmd.Flags().Int("mysql.port", 3306, "Mysql port")
	cmd.Flags().String("compress.type", "bzip2", "Method compression to use")
	cmd.Flags().String("storage.directory", "", "Diretório de armazenamento")

	return cmd
}
