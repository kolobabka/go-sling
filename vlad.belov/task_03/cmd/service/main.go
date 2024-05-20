package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

type ValuteInfo struct {
	NumCode  int64   `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

type ValCurs struct {
	Valute []ValuteInfo `xml:"Valute"`
}

type ValuteInfoDecoded struct {
	NumCode  int64  `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func (valute *ValuteInfo) UnmarshalXML(decoder *xml.Decoder, startElem xml.StartElement) error {
	valuteDecoded := ValuteInfoDecoded{}
	errorCode := decoder.DecodeElement(&valuteDecoded, &startElem)
	if errorCode != nil {
		return errorCode
	}

	valute.NumCode = valuteDecoded.NumCode
	valute.CharCode = valuteDecoded.CharCode

	valuteDecoded.Value = strings.Replace(valuteDecoded.Value, ",", ".", 1)
	valute.Value, errorCode = strconv.ParseFloat(valuteDecoded.Value, 64)
	if errorCode != nil {
		return errorCode
	}
	return nil
}

var wordPtr string

func init() {
	flag.StringVar(&wordPtr, "config", "config.yaml", "yaml file")
}

func main() {
	flag.Parse()
	yamlFile, errorCode := os.ReadFile(wordPtr)
	if errorCode != nil {
		fmt.Printf("Unable to read a file: %s\n", wordPtr)
		panic(errorCode)
	}

	var config Config
	errorCode = yaml.Unmarshal(yamlFile, &config)
	if errorCode != nil {
		fmt.Printf("Unable to parse a YAML file: %s\n", wordPtr)
		panic(errorCode)
	}

	xmlFile, errorCode := os.Open(config.Input)

	if errorCode != nil {
		fmt.Printf("Unable to open a file: %s\n", config.Input)
		panic(errorCode)
	}

	defer xmlFile.Close()

	xmlDecoder := xml.NewDecoder(xmlFile)

	xmlDecoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}

	var vals ValCurs
	errorCode = xmlDecoder.Decode(&vals)
	if errorCode != nil {
		fmt.Printf("Unable to decode a file: %s\n", config.Input)
		panic(errorCode)
	}

	sort.SliceStable(vals.Valute[:], func(i, j int) bool {
		return vals.Valute[i].Value > vals.Valute[j].Value
	})

	jsonFile, errorCode := json.MarshalIndent(vals.Valute, "", " ")
	if errorCode != nil {
		fmt.Println("Unable to convert to json format")
		panic(errorCode)
	}

	dir := filepath.Dir(config.Output)
	errorCode = os.MkdirAll(dir, 0777)
	if errorCode != nil && !os.IsExist(errorCode) {
		fmt.Printf("Unable to create a directory %s\n", dir)
		panic(errorCode)
	}

	file, errorCode := os.Create(config.Output)

	if errorCode != nil && !os.IsExist(errorCode) {
		fmt.Printf("Unable to create a file %s\n", config.Output)
		panic(errorCode)
	}

	errorCode = os.WriteFile(config.Output, jsonFile, 0777)
	if errorCode != nil {
		fmt.Printf("Unable to write to json file: %s", config.Output)
		panic(errorCode)
	}

	file.Close()
}
