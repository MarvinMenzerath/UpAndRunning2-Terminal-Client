package lib

import "os"

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
		{"--help", "-h", "Prints this help."},
		{"--version", "-v", "Prints the client's version."},
		{"--websites", "-w", "Prints every website's status in a table."},
		{"--status [URL]", "-s [URL]", "Prints the website's status-information."},
		{"--results [URL] {limit} {offset}", "-r [URL] {limit} {offset}", "Prints the website's last results."},
	})
}

func PrintVersion(version string) {
	println("You are currently using v" + version + " of this client.")
	println("Please check https://github.com/MarvinMenzerath/UpAndRunning2-Terminal-Client/releases regularly for updates.")
}
