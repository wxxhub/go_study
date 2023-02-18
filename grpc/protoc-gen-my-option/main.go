package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"log"
	my_proto "protoc-gen-my-option/proto"
)

func getFileOptions(f *protogen.File) *my_proto.MyFileOptions {
	log.Println("getFileOptions")
	// 解析my_options
	if myOptions, ok := proto.GetExtension(f.Proto.Options, my_proto.E_MyFileOption).(*my_proto.MyFileOptions); ok {
		log.Println("my_options:", myOptions)
		return myOptions
	} else {
		log.Println("get my_options failed.")
	}

	return nil
}

func genCode(myOptions *my_proto.MyFileOptions) string {
	if myOptions.Ok != nil && myOptions.GetOk() {
		return `
			func HelloMyOptions() {
				fmt.Println("Hello MyOptions, MyOptions is ok!")
			}
	`
	}
	return `
		func HelloMyOptions() {
			fmt.Println("Hello MyOptions, MyOptions is not ok...")
		}
	`
}

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		log.Println("protogen plugin")
		// protc时会把proto文件传入，解析所有文件
		for _, f := range plugin.Files {
			// 只处理我们处理的文件
			if !f.Generate {
				continue
			}
			// 处理getFiledOptions
			myOptions := getFileOptions(f)

			head := fmt.Sprintf(`
				package %s
				import "fmt"
			`, f.GoPackageName) // go文件, package名是必须的，import生成后一般IDE会自动import,可以不用写，除非某些场景需要提前加好import

			outFile := plugin.NewGeneratedFile(fmt.Sprintf("%s.my.go", f.GeneratedFilenamePrefix), f.GoImportPath)
			outFile.Write([]byte(head))
			outFile.Write([]byte(genCode(myOptions)))
		}
		return nil
	})
}
