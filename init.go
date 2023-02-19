// author: wsfuyibing <websearch@163.com>
// date: 2023-02-18

// Package log
// publish with async mode, allow kafka, file, terminal configuration.
//
//   // Import V8 Engine.
//   import "github.com/fuyibing/log/v8"
//
// Initialize.
//   func init(){
//       // Follow configurations are optional. All of follows can be
//       // configured in `config/log.yaml`. Filled with default if not
//       // configured.
//       log.Config.Set(
//           conf.SetTimeFormat("2006-01-02 15:04:05.999999"),
//           conf.SetLevel(conf.Error),
//           conf.SetPrefix("Prefix"),
//
//           conf.SetServiceAddr("172.16.0.110"),
//           conf.SetServicePort(8080),
//           conf.SetServiceEnvironment("production"),
//           conf.SetServiceName("myapp"),
//           conf.SetServiceVersion("1.2.3"),
//       )
//
//       // If adapter changed by code, You must call log.Client.Reset()
//       // to apply it.
//       log.Config.Set(conf.SetAdapter(adapters.AdapterKafka))
//       log.Client.Reset()
//   }
//
// // Main process.
//   func main(){
//       // Wait for a while
//       // until all logs publish completed.
//       //
//       // If the Close method `log.Client.Close()` is not set, Some logs
//       // end of the application may be lost.
//       defer log.Client.Close()
//
//       ...
//       log.Debug("debug info")
//       log.Infof("info message: adapter=%s, level=%v", conf.Adapter, conf.Level)
//       ...
//   }
package log

import (
	"github.com/fuyibing/log/v8/conf"
	"github.com/fuyibing/log/v8/core"
	"sync"
)

var (
	Config conf.Configuration
	Client core.Client
)

func init() {
	new(sync.Once).Do(func() {
		Config = conf.Config
		Client = core.NewClient()
	})
}
