package environment

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"

	config "github.com/JhonX2011/GFAWBP/pkg/infrastructure/configuration/config_constants"
)

type OsEnvLookup struct{}

type IEnvLookup interface {
	LookupEnv(key string) (string, bool)
}

func NewEnvLookup() IEnvLookup {
	return OsEnvLookup{}
}

func (o OsEnvLookup) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

type OsSetEnv struct{}

type ISetEnv interface {
	Setenv(string, string) error
}

func NewSetEnv() ISetEnv {
	return OsSetEnv{}
}

func (o OsSetEnv) Setenv(key string, value string) error {
	return os.Setenv(key, value)
}

func checkOrAssignEnv(key, fallback string, osLookupEnv IEnvLookup, osSetEnv ISetEnv) {
	if _, ok := osLookupEnv.LookupEnv(key); !ok {
		if err := osSetEnv.Setenv(key, fallback); err != nil {
			log.Panic(err)
		}
	}

	if os.Getenv(key) == "" {
		log.Panicf("[%s] env not initialized", key)
	}
}

func assignEnv(key, fallback string, validate int, osSetEnv ISetEnv) { //nolint:unused
	if err := osSetEnv.Setenv(key, fallback); err != nil {
		log.Panic(err)
	}

	if validate == 1 {
		if os.Getenv(key) == "" {
			log.Panicf("[%s] env not initialized", key)
		}
	}
}

func IsFlagEnabled(key string) bool {
	if v := os.Getenv(key); v == "" || v == "0" {
		return false
	}
	return true
}

func GetIntFeature(key string) int {
	v, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Panic(err)
	}
	return v
}

func GetStringFeature(key string) string {
	if v := os.Getenv(key); len(v) > 0 {
		return v
	}
	return ""
}

func GetEncodeStringFeature(key string) ([]byte, error) {
	val := GetStringFeature(key)
	b, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

func LoadEnvironment() {
	checkOrAssignEnv("APP_NAME", config.AppName, NewEnvLookup(), NewSetEnv())
	checkOrAssignEnv("ENVIRONMENT", config.Environment, NewEnvLookup(), NewSetEnv())
	checkOrAssignEnv("STACK", config.Stack, NewEnvLookup(), NewSetEnv())
	if os.Getenv("ENVIRONMENT") == "development" {
		checkOrAssignEnv("PORT", config.Port, NewEnvLookup(), NewSetEnv())
		checkOrAssignEnv("CONFIG_DIR", config.ConfigDir, NewEnvLookup(), NewSetEnv())
	}
}

func PrintEnv() {
	fmt.Println("***********************************************************************")
	fmt.Println("***********************************************************************")
	fmt.Printf("APP_NAME    ==> [%v]\n", os.Getenv("APP_NAME"))
	fmt.Printf("ENVIRONMENT ==> [%v]\n", os.Getenv("ENVIRONMENT"))
	fmt.Printf("PORT        ==> [%v]\n", os.Getenv("PORT"))
	fmt.Printf("STACK       ==> [%v]\n", os.Getenv("STACK"))
	fmt.Printf("CONFIG_DIR  ==> [%v]\n", os.Getenv("CONFIG_DIR"))
	fmt.Println("***********************************************************************")
	fmt.Println("***********************************************************************")
}
