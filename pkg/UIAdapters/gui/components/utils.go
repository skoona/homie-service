package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/skoona/homie-service/pkg/deviceCore"
)

var CurrentVariant fyne.ThemeVariant

// SynTax --> v, ok := t.(SomeConcreteType)
// switch v := item.(type) {
//	case map[string]string:
//  case map[string]deviceCore.DeviceAttribute:
//  case map[string]deviceCore.DeviceNode:
//	...}

// MapKeys returns a list of keys
func MapKeys(m map[string]interface{}) (keys []string) {
	keysLen := len(m)
	keys = make([]string, keysLen)
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
// MapValues returns a list of values
func MapValues(m map[string]interface{}) []string {
	valuesLen := len(m)
	values := make([]string, valuesLen)
	var item string
	
	for _, v := range m {
		switch t := v.(type) {
		case deviceCore.Device: item = v.(deviceCore.Device).Name
		case deviceCore.DeviceAttribute: item = v.(deviceCore.DeviceAttribute).Value
		case deviceCore.DeviceAttributeProperty: item = v.(deviceCore.DeviceAttributeProperty).Value
		case deviceCore.DeviceAttributePropertyProperty: item = v.(deviceCore.DeviceAttributePropertyProperty).Value
		case deviceCore.DeviceNode: item = v.(deviceCore.DeviceNode).Name
		case deviceCore.DeviceNodeAttribute: item = v.(deviceCore.DeviceNodeAttribute).Value
		case deviceCore.DeviceNodeProperty: item = v.(deviceCore.DeviceNodeProperty).Value
		case deviceCore.DeviceNodePropertyAttribute: item = v.(deviceCore.DeviceNodePropertyAttribute).Value
		default:
			fmt.Printf("\nUNKNOWN TYPE=(%T)\n", t)
			item = v.(string)
		}

		values = append(values, item)
	}
	return values
}

