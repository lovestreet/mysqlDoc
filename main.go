package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/lovestreet/mysqlDoc/database"
	"github.com/lovestreet/mysqlDoc/docx"

	"github.com/lovestreet/mysqlDoc/config"
)

var configFile = flag.String("cfg", "config.json", "config file name")

func main() {
	flag.Parse()

	//load config
	var cfg, err = loadConfig(*configFile)
	if err != nil {
		fmt.Println("load configuration error : ", err)
		return
	}

	//write table and column
	if err = database.Init(cfg); err != nil {
		fmt.Println("init database error : ", err)
		return
	}
	fmt.Println("init database success")

	for _, schema := range cfg.Schema {
		fmt.Println()
		fmt.Println("[[ begin to process schema : ", schema)
		schema = strings.TrimSpace(schema)
		// get table
		var tables = database.GetTables([]string{schema})
		fmt.Printf("load table defines success")
		// get table columns
		var columns = database.GetColumns([]string{schema})
		fmt.Printf("load table column defines success")

		fmt.Printf("begin to write to file")
		var fileName = fmt.Sprintf("%v.docx", schema)
		docx.WriteTables(fileName, tables, columns)
		fmt.Printf("write to file finished")
		fmt.Println("]] process schema : ", schema, "finished")
	}

	fmt.Println("done")

}

func loadConfig(file string) (cfg config.Configuration, err error) {
	if _, err := os.Stat(file); err != nil {
		fmt.Println(err)
		return cfg, err
	}

	cfg, err = config.LoadConfig(file)
	return cfg, err
}
