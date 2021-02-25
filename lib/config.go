package lib

import (
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"os"
	//"reflect"
	"strings"

	"github.com/spf13/viper"
)

var Cfg map[string]*ichuntCfg

type ichuntCfg struct {
	*viper.Viper
}


func (v *ichuntCfg) GetStringSL(key string) string {
	fmt.Println(v.Get(key))
	fmt.Println(v.Get(key))
	return cast.ToString(v.Get(key))
}

func (v *ichuntCfg) GetLixinString(key string) string {
	tmp := v.Get(key)
	str := fmt.Sprintf("%v", tmp)
	return str
}


//func GetStringSlicea(key string) []string { return ichuntCfg.GetStringSlice(key) }
//func (c *ichuntCfg) GetStringSlicea(key string) []string {
//	return cast.ToStringSlice(c.Get(key))
//}

func initViperConf(ConfEnvPath string) error {
	f, err := os.Open(ConfEnvPath + "/")
	if err != nil {
		return err
	}
	fileList, err := f.Readdir(1024)
	if err != nil {
		return err
	}
	for _, fo := range fileList {
		if !fo.IsDir() {
			bts, err := ioutil.ReadFile(ConfEnvPath + "/" + fo.Name())
			if err != nil {
				return err
			}
			v := viper.New()
			v.SetConfigType("toml")
			//fmt.Println(bytes.NewBuffer(bts))
			v.ReadConfig(bytes.NewBuffer(bts))
			pathArr := strings.Split(fo.Name(), ".")
			if Cfg == nil {
				Cfg = make(map[string]*ichuntCfg)
			}
			tmp := &ichuntCfg{v}
			Cfg[pathArr[0]] = tmp
		}
	}
	return nil
}

func Instance(name string) *ichuntCfg {
	return Cfg[name]
}

func Init(path string) ( err error){
	err = initViperConf(path)
	return err
}


