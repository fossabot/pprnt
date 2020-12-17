package pprnt

/*
	Version: 2.0.0
	Author: Alan Maldonado

	== OpenSource Project ==
*/

import (
	"log"
	"reflect"
	"strings"

	"github.com/DrN3MESiS/pprnt/cleaner"
	"github.com/DrN3MESiS/pprnt/helpers"
	"github.com/DrN3MESiS/pprnt/printer"
)

var (
	detailMode   bool = false
	initialDepth int  = 0
)

//DetailMode ...
/*
Code: 1 = Enables Detail Mode
Code: 0 = Disables Detail Mode
*/
func DetailMode(code int) {
	switch code {
	case 1:
		detailMode = true
		break
	case 0:
		detailMode = false
		break
	}
}

//IdentLength ...
/*
Defines the length of each identation level
*/
func IdentLength(length int) {
	printer.IdentString = strings.Repeat(" ", length)
}

//Print ...
/*
Print single object or value
*/
func Print(arg interface{}) {
	errMessage := "[PPRNT]"
	switch reflect.ValueOf(arg).Kind() {
	case reflect.Map, reflect.Struct:
		MTP, err := helpers.ValueToMap(arg)
		if err != nil {
			log.Printf("%+v Couldn't print the provided data > %+v", errMessage, err.Error())
			return
		}

		err = printer.PrintMap(MTP, initialDepth, detailMode)
		if err != nil {
			log.Printf("%+v Couldn't print the provided data > %+v", errMessage, err.Error())
			return
		}
		break
	case reflect.Array, reflect.Slice:
		tempArray := arg.([]interface{})
		err := printer.PrintArray(tempArray, initialDepth, detailMode)
		if err != nil {
			log.Printf("%+v Couldn't print the provided data > %+v", errMessage, err.Error())
			return
		}
		break
	default:
		err := printer.PrintNormal(arg, initialDepth, detailMode)
		if err != nil {
			log.Printf("%+v Couldn't print the provided data > %+v", errMessage, err.Error())
			return
		}
		break
	}

}

// //TrPrint ...
// /**/
// func TrPrint(data interface{}) error {
// 	errMessage := "[pprnt][Print()] > "

// 	if reflect.ValueOf(data).Kind() == reflect.Struct {
// 		mapData, err := helpers.StructToMap(data)
// 		if err != nil {
// 			log.Printf("%+v", errMessage+err.Error())
// 			return errors.New(errMessage + err.Error())
// 		}

// 		depth := 1
// 		printer.PrintData(mapData, &depth)

// 	} else if reflect.ValueOf(data).Type() == reflect.TypeOf(map[string]interface{}{}) {
// 		depth := 1
// 		printer.PrintData(data.(map[string]interface{}), &depth)

// 	} else if reflect.ValueOf(data).Type() == reflect.TypeOf(map[string]string{}) {
// 		depth := 1
// 		toSend := map[string]interface{}{}
// 		for key, value := range data.(map[string]string) {
// 			toSend[key] = value
// 		}
// 		printer.PrintData(toSend, &depth)

// 	} else {
// 		fmt.Printf("%+v", data)
// 	}

// 	return nil
// }

//SuperDepthMapCleaning ...
func SuperDepthMapCleaning(curMap map[string]interface{}) (map[string]interface{}, error) {
	return cleaner.CleanMap(curMap), nil
}
