package config

import (
	"fmt"
	"os"
	"strings"
)

var configFile string = "config.yaml"

type configStr map[string]string

// createConfig
func CreateConfig() configStr {
	str := configStr{"blacklist": "none", "source": "none", "chapters" : "none"}
	_writeConfig(str)
	return str
}

// setBlacklist
func SetBlacklist( file string) configStr {
	str := _readConfig()
	str["blacklist"] = file
	err := _writeConfig(str)
	if err != nil {
		panic(err)
	}
	return str
}

// setSource
func SetSource( file string) configStr {
	str := _readConfig()
	str["source"] = file
	err := _writeConfig(str)
	if err != nil {
		panic(err)
	}
	return str
}

// setChapters
func SetChapters( file string) configStr {
	str := _readConfig()
	str["chapters"] = file
	err := _writeConfig(str)
	if err != nil {
		panic(err)
	}
	return str
}
// getBlacklist
func GetBlacklist() string {
	str := _readConfig()
	return str["blacklist"]
}

// getSource
func GetSource() string {
	str := _readConfig()
	return str["source"]
}

// getChapters
func GetChapters() string {
	str := _readConfig()
	return str["chapters"]
}

// _readConfig
func _readConfig() configStr {
	var str = make(configStr)

	c, err := os.ReadFile(configFile)
	if err != nil {
		str = CreateConfig()
		return str
	}
	lns := strings.Split(string(c), "\n")

	for ln := range lns {	
		lnNoW := strings.ReplaceAll(lns[ln], " ", "")
		if len(lnNoW) <= 0 {
			continue
		}
		kvPair := strings.Split(lnNoW, ":")
		str[kvPair[0]] = kvPair[1]
	}
	return str
}

// _writeConfig
func _writeConfig(str configStr) error {
	f, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer f.Close()

	for k, v := range str {
		ln := k + " : " + v + "\n"
		f.WriteString(ln)
	}
	return nil
}

func DisplayConfig() {
	str := _readConfig()
	for key, val := range str {
		fmt.Printf("%s : %s\n", key, val)
	}
}
