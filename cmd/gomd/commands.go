package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var implementedCommands = []string{
	"copy",
	"cp",
	"move",
	"mv",
	"remove",
	"rm",
	"mkdir",
	"makedir",
}

func (a *app) cmdAutocomplete(currentText string) (entries []string) {
	if len(currentText) <= 1 {
		return entries
	}
	for _, cmd := range a.commands {
		if strings.HasPrefix(strings.ToLower(cmd), strings.ToLower(currentText)) {
			entries = append(entries, cmd)
		}
	}
	return entries
}

func (a *app) executeCommand(command string) {
	if command == "" {
		return
	}
	c := strings.Split(command, " ")
	switch c[0] {
	case "copy", "cp":
		if err := a.cmdCopy(); err != nil {
			fmt.Fprintln(a.appOut, "error: ", err)
		}
	case "move", "mv":
		if err := a.cmdMove(); err != nil {
			fmt.Fprintln(a.appOut, "error: ", err)
		}
		a.left.Folder.Update()
		a.left.makeTableView()
		a.right.Folder.Update()
		a.right.makeTableView()
	}
	a.cmd.SetText("")
}

func (a *app) cmdCopy() error {

	return nil
}

func (a *app) cmdMove() error {
	oldpath := filepath.Join(
		a.left.Folder.Path,
		a.left.Folder.SelectedFile().Name())
	newpath := filepath.Join(
		a.right.Folder.Path,
		a.left.Folder.SelectedFile().Name())
	if oldpath == newpath {
		a.Println("nothing to move here")
		return nil
	}
	return os.Rename(oldpath, newpath)
}
