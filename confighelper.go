package confighelper

import (
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// BindEnvs binds the struct fields with the config, it look-ups 'mapstructure' tag
// for keys
func BindEnvs(iface interface{}, parts ...string) error {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			err := viper.BindEnv(strings.Join(append(parts, tv), "."))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
