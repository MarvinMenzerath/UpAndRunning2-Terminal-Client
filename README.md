# UpAndRunning2 Terminal Client
This is a simple Terminal Client for [UpAndRunning2](https://github.com/MarvinMenzerath/UpAndRunning2).

## Documentation

### Installation
* Download and extract all files into a directory.
* Copy `config/default.json` to `config/local.json` and change this file to your needs.
* Optional: rename `UpAndRunning2-Terminal-Client` to something shorter.
* Done!

### Commands
Just run `./UpAndRunning2-Terminal-Client help` to see all commands available or see this list:

Command                        | Short Command             | Description
-------------------------------|---------------------------|-------------------------------------------------
help                           | -h                        | Prints this help.
version                        | -v                        | Prints the client's version.
websites                       | -w                        | Prints every website's status in a table.
status [URL]                   | -s [URL]                  | Prints the website's current status-information.
results [URL] {limit} {offset} | -r [URL] {limit} {offset} | Prints the website's latest results.
