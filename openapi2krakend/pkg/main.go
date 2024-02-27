package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"github.com/okhuz/openapi2krakend/pkg/converter"
	"github.com/okhuz/openapi2krakend/pkg/utility"
)

func main() {
	swaggerDirectory := flag.String("directory", "./swagger", "Directory of the swagger files")
	outputFilePath := flag.String("output", "output/krakend.json", "Output file path")

	flag.Parse()

	encoding := utility.GetEnv("ENCODING", "json")
	globalTimeout := utility.GetEnv("GLOBAL_TIMEOUT", "3000ms")

	configuration := converter.Convert(*swaggerDirectory, encoding, globalTimeout)

	file, _ := json.MarshalIndent(configuration, "", " ")
	outputDir := filepath.Dir(*outputFilePath)
	_ = os.MkdirAll(outputDir, 0777)
	_ = os.WriteFile(*outputFilePath, file, 0644)
}
