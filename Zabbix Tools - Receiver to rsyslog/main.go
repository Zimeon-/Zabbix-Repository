/*
** Copyright (C) 2001-2025 Zabbix SIA
**
** Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
** documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
** rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
** permit persons to whom the Software is furnished to do so, subject to the following conditions:
**
** The above copyright notice and this permission notice shall be included in all copies or substantial portions
** of the Software.
**
** THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
** WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
** COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
** TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
** SOFTWARE.
**/

package main

import (
	"Example_interface/server"
	"flag"
	"fmt"
	"os"
)

const (
	confDescription = "Path to the configuration file"
	helpDescription = "Display help message and application information"
)

var (
	confFlag string
	helpFlag bool
)

func main() {
	loadFlags()

	if helpFlag {
		flag.Usage()
		printInfo()
		os.Exit(0)
	}

	conf, err := loadConfiguration(confFlag)
	if err != nil {
		panic(err)
	}

	initLogger(conf.LogPath)
	defer closeLogFile()

	err = server.Run(conf.Port, conf.CertFile, conf.KeyFile, conf.DataPath, conf.EnableTls)
	if err != nil {
		panic(err)
	}
}

func loadFlags() {
	flag.StringVar(&confFlag, "config", defaultConfigPath, confDescription)
	flag.StringVar(&confFlag, "c", defaultConfigPath, confDescription+" (shorthand)")
	flag.BoolVar(&helpFlag, "help", false, helpDescription)
	flag.BoolVar(&helpFlag, "h", false, helpDescription+" (shorthand)")

	flag.Parse()
}

func printInfo() {
	fmt.Println()
	fmt.Println(applicationInfo())
	fmt.Println()
	fmt.Println(copyrightMessage())
}

func applicationInfo() string {
	return "Interface to receive data from Zabbix over http. For usage information please check the README.md"
}

func copyrightMessage() string {
	return "Copyright (C) 2001-2025 Zabbix SIA\n" +
		"Permission is hereby granted, free of charge, to any person obtaining a copy of\n" +
		"this software and associated documentation files (the \"Software\"), to deal in\n" +
		"the Software without restriction, including without limitation the rights to\n" +
		"use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies\n" +
		"of the Software, and to permit persons to whom the Software is furnished to do\n" +
		"so, subject to the following conditions:\n\n" +

		"The above copyright notice and this permission notice shall be included in all\n" +
		"copies or substantial portions of the Software.\n\n" +

		"THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n" +
		"IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n" +
		"FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n" +
		"AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n" +
		"LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n" +
		"OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE\n" +
		"SOFTWARE."
}
