package main

import (
    "flag"
    "streak/cmd"
)

func main() {
    // have support for both --help and -h for UNIX convention
    var callHelp bool
    flag.BoolVar(&callHelp, "help", false, "Show help message")
    flag.BoolVar(&callHelp, "h", false, "Short help message")

    callBrowse := flag.Bool("browse", false, "browse images")
    callConfig := flag.Bool("config", false, "config setup")
    callDownload := flag.Bool("download", false, "download media")

    // todo: add defaults for all other options
    // todo: add file and directory support

    flag.Parse()

    switch {
    case callHelp:
	cmd.RunHelp()

    case *callConfig:
	cmd.RunConfig()

    case *callBrowse:
	cmd.RunBrowse()

    case *callDownload:
	cmd.RunDownload()

    // option to call by default when no other is explicitly stated
    default:
	cmd.RunUpload()
    }
}
