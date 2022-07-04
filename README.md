# confighelper
confighelper provides additional feature to viper to support AutomativEnv without having any config file.

## How to use

```
import "github.com/netbookai/confighelper"

//in you config handling code, call BindEnvs function and pass your struct object
configObj := Config{}
err := confighelper.BindEnvs(configObj)

//now this will work perfectly fine without having any empty config file
viper.AutomaticEnv()
```


#### reference: this [comment](https://github.com/spf13/viper/issues/188#issuecomment-399884438) by [@krak3n](https://www.github.com/krak3n)


