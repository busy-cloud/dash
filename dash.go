package dash

import (
	"embed"
	"encoding/json"
	"github.com/busy-cloud/boat/apps"
	"github.com/busy-cloud/boat/store"
	_ "github.com/busy-cloud/dash/internal"
	"log"
)

//go:embed assets
var assets embed.FS

//go:embed pages
var pages embed.FS

//go:embed manifest.json
var manifest []byte

func init() {
	//注册为内部插件
	var a apps.App
	err := json.Unmarshal(manifest, &a)
	if err != nil {
		log.Fatal(err)
	}
	apps.Register(&a)

	//注册资源
	a.AssetsFS = store.PrefixFS(&assets, "assets")
	a.PagesFS = store.PrefixFS(&pages, "pages")
	//a.ProtocolsFS = store.PrefixFS(&protocols, "protocols"))
}
