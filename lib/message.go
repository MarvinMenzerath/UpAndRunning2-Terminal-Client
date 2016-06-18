package lib

import (
	"os"
	"runtime"
)

func PrintHead() {
	println("#################################")
	println("# UpAndRunning2 Terminal Client #")
	println("#################################\n")
}

func PrintWrongArguments() {
	println("No / Wrong arguments given.\n")
	PrintHelp()
	os.Exit(1)
}

func PrintHelp() {
	println("HELP:")
	PrintAsTable([]string{"Command", "Short Command", "Description"}, [][]string{
		{"help", "-h", "Prints this help."},
		{"version", "-v", "Prints the client's version."},
		{"websites", "-w", "Prints every website's status in a table."},
		{"status [URL]", "-s [URL]", "Prints the website's current status-information."},
		{"results [URL] {limit} {offset}", "-r [URL] {limit} {offset}", "Prints the website's latest results."},
	})
}

func PrintVersion(version string) {
	println("You are currently using v" + version + " [" + runtime.Version() + "@" + runtime.GOOS + "_" + runtime.GOARCH + "] of this client.")
	println("Please check https://github.com/MarvinMenzerath/UpAndRunning2-Terminal-Client/releases regularly for updates.")
}
