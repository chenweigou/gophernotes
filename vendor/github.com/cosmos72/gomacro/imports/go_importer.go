// this file was generated by gomacro command: import _b "go/importer"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"go/importer"
)

// reflection: allow interpreted code to import "go/importer"
func init() {
	Packages["go/importer"] = Package{
	Binds: map[string]Value{
		"Default":	ValueOf(importer.Default),
		"For":	ValueOf(importer.For),
	},Types: map[string]Type{
		"Lookup":	TypeOf((*importer.Lookup)(nil)).Elem(),
	},
	}
}
