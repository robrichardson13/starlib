package starlib

import (
	"fmt"

	"github.com/robrichardson13/starlib/bsoup"
	"github.com/robrichardson13/starlib/encoding/base64"
	"github.com/robrichardson13/starlib/encoding/csv"
	"github.com/robrichardson13/starlib/encoding/json"
	"github.com/robrichardson13/starlib/encoding/yaml"
	"github.com/robrichardson13/starlib/geo"
	"github.com/robrichardson13/starlib/hash"
	"github.com/robrichardson13/starlib/html"
	"github.com/robrichardson13/starlib/http"
	"github.com/robrichardson13/starlib/math"
	"github.com/robrichardson13/starlib/re"
	"github.com/robrichardson13/starlib/time"
	"github.com/robrichardson13/starlib/xlsx"
	"github.com/robrichardson13/starlib/zipfile"
	"go.starlark.net/starlark"
)

// Version is the current semver for the entire starlib library
const Version = "0.4.1"

// Loader presents the starlib library as a loader
func Loader(thread *starlark.Thread, module string) (dict starlark.StringDict, err error) {
	switch module {
	case time.ModuleName:
		return time.LoadModule()
	case http.ModuleName:
		return http.LoadModule()
	case xlsx.ModuleName:
		return xlsx.LoadModule()
	case html.ModuleName:
		return html.LoadModule()
	case bsoup.ModuleName:
		return bsoup.LoadModule()
	case zipfile.ModuleName:
		return zipfile.LoadModule()
	case re.ModuleName:
		return re.LoadModule()
	case base64.ModuleName:
		return base64.LoadModule()
	case csv.ModuleName:
		return csv.LoadModule()
	case json.ModuleName:
		return json.LoadModule()
	case yaml.ModuleName:
		return yaml.LoadModule()
	case geo.ModuleName:
		return geo.LoadModule()
	case math.ModuleName:
		return math.LoadModule()
	case hash.ModuleName:
		return hash.LoadModule()
	}

	return nil, fmt.Errorf("invalid module '%s'", module)
}
