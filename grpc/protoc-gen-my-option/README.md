# ProtoOption

## 简介

### extension

能够对已经存在的消息进行扩展，这种扩展只是动态增加字段，而没有增加接口。  
例子：扩展protobuf的FieldOptions  
```protobuf
message MyFileOptions {
  optional bool ok = 1;
}

// 扩展protobuf的MessageOptions
extend google.protobuf.FileOptions {
  optional MyFileOptions my_file_option = 51234; // 取一个自定义不重复值
}
```

### protobuf options
- google.protobuf.FileOptions
- google.protobuf.ServiceOptions
- google.protobuf.MethodOptions
- google.protobuf.MessageOptions
- google.protobuf.FieldOptions
- google.protobuf.EnumOptions
- google.protobuf.EnumValueOptions
- google.protobuf.OneofOptions
- google.protobuf.ExtensionRangeOptions

## 实现自定义option

### 扩展实现
```shell
mkdir protoc-gen-my-option
cd protoc-gen-my-option
```

实现一个FileOptions
```shell
mkdir proto

# 在proto中实现自己的option,可以将EOF的内容拷贝到proto/my_option.proto中
echo > proto/my_option.proto  <<EOF
syntax = "proto3";

package my_option;

option go_package="protoc-gen-my-option/proto";

import "google/protobuf/descriptor.proto";

message MyFileOptions {
  optional bool ok = 1;
}

// 扩展protobuf的MessageOptions
extend google.protobuf.FileOptions {
  optional MyFileOptions my_file_option = 51234; // 取一个自定义不重复值
}
EOF 
```

生成后proto/my_option.proto中的内容
```protobuf
syntax = "proto3";

package my_option;

option go_package="protoc-gen-my-option/proto";

import "google/protobuf/descriptor.proto";

message MyFileOptions {
  optional bool ok = 1;
}

// 扩展protobuf的MessageOptions
extend google.protobuf.FileOptions {
  optional MyFileOptions my_file_option = 51234; // 取一个自定义不重复值
}
```

生成扩展的pb文件
```shell
protoc --go_out=../ proto/my_option.proto
```


### 解析扩展插件

#### 创建工程文件目录
__扩展插件的开头必须以protoc-gen开头__，在创建的protoc-gen-my-option目录下编写插件，并初始化项目。

```shell
go mod init protoc-gen-my-option
```
#### 解析demo
```go
package main

import (
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
			getFileOptions(f)
		}
		return nil
	})
}
```

安装插件
```shell
go install .
```

#### 使用MyFileOptions
```shell
mkdir protoc-gen-my-option/test_proto

# 使用my-option
echo > test_proto/test_proto.proto  <<EOF
syntax = "proto3";

package my_option;

option go_package="protoc-gen-my-option/test_proto";

import "my_option.proto";

// 使用FiledOptions

option (my_option.my_file_option).ok = true;
EOF 
```

生成的test_proto/test_proto.proto内容如下：
```protobuf
syntax = "proto3";

package my_option;

option go_package="protoc-gen-my-option/test_proto";

import "my_option.proto";

// 使用FiledOptions

option (my_option.my_file_option).ok = true;
```

#### 验证插件解析MyFileOptions
```shell
protoc -I=./proto --go_out=../ --my-option_out=../ --proto_path=test_proto test_option.proto
```
成功后会输出以下内容,能够看到my_options中的ok被成功解析。
```shell
2023/02/18 20:12:16 protogen plugin
2023/02/18 20:12:16 getFileOptions
2023/02/18 20:12:16 my_options: ok:true
```

### 解析扩展插件后生成我们自己的代码
在上面的demo中添加代码生成逻辑
```go
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
```

[以上代码地址](https://github.com/wxxhub/go_study/tree/master/grpc/protoc-gen-my-option)