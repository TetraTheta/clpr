package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/integrii/flaggy"
	"github.com/maruel/natural"
	"golang.design/x/clipboard"
)

const (
	defaultName = "_default"
	version     = "1.0.0"
)

var (
	getCmd  *flaggy.Subcommand
	listCmd *flaggy.Subcommand
	setCmd  *flaggy.Subcommand

	// set default value of variable
	getName    = ""
	setName    = ""
	setContent = readPipe()
)

// This is initializer
func init() {
	flaggy.SetName("clpr")
	flaggy.SetDescription("Universal Clipboard Utility")
	flaggy.SetVersion(version)

	// subcommand: get
	getCmd = flaggy.NewSubcommand("get")
	getCmd.Description = "Get content of clipboard or LCC"
	getCmd.String(&getName, "name", "n", "Name of LCC")

	// subcommand: list
	listCmd = flaggy.NewSubcommand("list")
	listCmd.Description = "List stored LCC"

	// subcommand: set
	setCmd = flaggy.NewSubcommand("set")
	setCmd.Description = "Set content of clipboard or LCC"
	setCmd.String(&setName, "name", "n", "Name of LCC")
	setCmd.AddPositionalValue(&setContent, "content", 1, false, "New content of clipboard or LCC. If not present, piped string value will be used.")

	flaggy.AttachSubcommand(getCmd, 1)
	flaggy.AttachSubcommand(listCmd, 1)
	flaggy.AttachSubcommand(setCmd, 1)

	flaggy.Parse()
}

func main() {
	var exitCode int

	switch {
	case getCmd.Used:
		exitCode = getCommand(getName)
	case listCmd.Used:
		exitCode = listCommand()
	case setCmd.Used:
		if setContent != "" {
			exitCode = setCommand(setName, setContent)
		} else {
			fmt.Fprintln(os.Stderr, "No content is provided for set.")
			exitCode = 2
		}
	default:
		fmt.Fprintln(os.Stderr, "clpr: No valid subcommand found.")
		exitCode = 1
	}

	os.Exit(exitCode)
}

//region Helper function

// Get string from Pipe
func readPipe() string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return ""
	}
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		data, err := io.ReadAll(bufio.NewReader(os.Stdin))
		if err == nil {
			return strings.TrimSpace(string(data))
		}
	}
	return ""
}

//endregion Helper function

//region Command function

// Process 'get' command
func getCommand(name string) int {
	if name == "" {
		if err := clipboard.Init(); err == nil {
			fmt.Println(string(clipboard.Read(clipboard.FmtText)))
			return 0
		} else {
			fmt.Fprintln(os.Stderr, "Clipboard is not available:", err)
			name = defaultName
		}
	}

	content, err := os.ReadFile(filepath.Join(getLCCDir(), name+".txt"))
	if err == nil {
		fmt.Print(string(content))
		return 0
	} else {
		fmt.Fprintf(os.Stderr, "Failed to read LCC: %q\n%v\n", name, err)
		return 2
	}
}

// Process 'list' command
func listCommand() int {
	dir := getLCCDir()
	files, err := filepath.Glob(filepath.Join(dir, "*.txt"))

	if err == nil {
		if len(files) == 0 {
			fmt.Println("No LCC files in", dir)
			return 0
		} else {
			names := make([]string, 0, len(files))
			for _, f := range files {
				base := filepath.Base(f)
				names = append(names, strings.TrimSuffix(base, filepath.Ext(base)))
			}
			sort.Sort(natural.StringSlice(names))
			for _, n := range names {
				fmt.Println("*", n)
			}
			return 0
		}
	} else {
		fmt.Fprintln(os.Stderr, "Error globbing LCC files:", err)
		return 1
	}
}

// Process 'set' command
func setCommand(name, content string) int {
	if name == "" {
		if err := clipboard.Init(); err == nil {
			clipboard.Write(clipboard.FmtText, []byte(content))
			return 0
		} else {
			fmt.Fprintln(os.Stderr, "Clipboard is not available:", err)
			name = defaultName
		}
	}

	dir := getLCCDir()
	if err := os.MkdirAll(dir, 0755); err == nil {
		path := filepath.Join(dir, name+".txt")
		if err2 := os.WriteFile(path, []byte(content), 0644); err2 == nil {
			return 0
		} else {
			fmt.Fprintln(os.Stderr, "Failed to write LCC:", err2)
			return 2
		}
	} else {
		fmt.Fprintf(os.Stderr, "Cannot create LCC directory: %q\n%v\n", dir, err)
		return 2
	}
}

//endregion Command function
