package cmd

import (
	. "github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"io"
	"os"
	"path/filepath"
)

func GenDocument(c *Command, dir string) error {
	identity := func(s string) string { return "" }
	emptyStr := func(s string) string { return "" }
	return GenDocumentCustom(c, dir, identity, emptyStr)
}
func GenDocumentCustom(c *Command, dir string, prepender, linkHandler func(string) string) error {
	f, err := os.Create(filepath.Join(dir, "README.md"))
	if err != nil {
		return err
	}
	defer f.Close()
	return genDocument(c, f, prepender, linkHandler)
}
func genDocument(c *Command, f io.Writer, prepender, linkHandler func(string) string) error {
	for _, c := range c.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := genDocument(c, f, prepender, linkHandler); err != nil {
			return err
		}
	}

	if _, err := io.WriteString(f, prepender(c.Name())); err != nil {
		return err
	} else if err := doc.GenMarkdownCustom(c, f, linkHandler); err != nil {
		return err
	}
	return nil
}
