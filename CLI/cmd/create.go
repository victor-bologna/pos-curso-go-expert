/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"

	"github.com/spf13/cobra"
	"github.com/victor-bologna/pos-curso-go-expert-cli/internal/database"
	_ "modernc.org/sqlite"
)

func newCreateCmd(categoryDB *database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples`,
		RunE:  createCategory(categoryDB),
	}
}

func createCategory(category *database.Category) RunECreate {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := category.Create(name, description)
		return err
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDB()))
	categoryCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Name of the Category")
	createCmd.Flags().StringP("description", "d", "", "Description of the Category")
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite", "./category.db")
	if err != nil {
		panic(err)
	}
	return db
}

func GetCategoryDB(db *sql.DB) *database.Category {
	return database.NewCategory(db)
}
