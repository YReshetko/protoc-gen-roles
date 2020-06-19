package main

import (
	"bytes"
	"fmt"
	"github.com/YReshetko/protoc-gen-roles/proto"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"sort"
	"strings"
)

const methodsPattern = "/%s.%s/%s" // "/<package>.<service>/<rpc-method>"
var imps = []string{
	//"proto \"github.com/golang/protobuf/proto\"",
}

type rpcMethodWithRoles struct {
	method string
	roles []string
}
type arrMapping []rpcMethodWithRoles

func (a arrMapping) Len() int           { return len(a) }
func (a arrMapping) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a arrMapping) Less(i, j int) bool { return a[i].method < a[j].method }


type b struct {
	bytes.Buffer
	prefix string
}

func processRequest(rq *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	rs := new(plugin.CodeGeneratorResponse)
	//log.Println("Files to generate: ", rq.FileToGenerate)

	for _, protoFile := range rq.ProtoFile {
		if !contains(rq.FileToGenerate, *protoFile.Name) {
			continue
		}
		rolesMap := make(map[string][]string)
		buff := new(b)

		packageName := "default"
		if len(protoFile.Options.GetGoPackage()) != 0 {
			packageInd := strings.Index(protoFile.Options.GetGoPackage(), ";")
			packageName = protoFile.Options.GetGoPackage()[packageInd + 1:]
			packageName = strings.Replace(packageName, ".", "_", -1)
			packageName = strings.Replace(packageName, "/", "_", -1)
		} else {
			packageName = strings.Replace(*protoFile.Package, ".", "_", -1)
		}

		buff.header(packageName)

		servicePackage := strings.Replace(*protoFile.Package, ".", "_", -1)
		for _, serv := range protoFile.GetService() {
			serviceName := serv.Name
			//log.Println("Working with service: ", serviceName)
			for _, method := range serv.Method {
				methodName := method.Name
				//log.Println("Working with method: ", methodName)
				rolesLine := strings.TrimSpace(extractRoles(method.Options))
				if len(rolesLine) == 0 {
					continue
				}
				roles := strings.Split(rolesLine, ",")
				key := fmt.Sprintf(methodsPattern, servicePackage, *serviceName, *methodName)
				rolesMap[key] = roles
				//log.Println("Get roles option:", roles)
			}
		}
		if len(rolesMap) == 0{
			continue
		}
		buff.finalize(rolesMap)
		rs.File = append(rs.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(generatedFileName(protoFile)),
			Content: proto.String(buff.String()),
		})
	}
	//log.Println(rs.File)
	return rs
}

func contains(arr []string, str string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}

func extractRoles(options *descriptor.MethodOptions) string {
	//log.Printf("Option %+v\n", options)
	roles, err := proto.GetExtension(options, gen_roles.E_Roles)
	if err != nil {
		//panic(err)
		return ""
	}
	if roles == nil {
		return ""
	}
	var out string
	switch t := roles.(type) {
	case string:
		out = t
	case *string:
		out = *t
	default:
		panic("roles have unknown type")
	}
	return out
}

func (buff *b) header(pack string) {
	buff.write("// Code generated by protoc-gen-example. DO NOT EDIT.")
	buff.newLine()
	buff.write(fmt.Sprintf("package %s", pack))
	buff.newLine()
	buff.imports()
	buff.newLine()
	//buff.write("const _ = proto.ProtoPackageIsVersion3")
	buff.newLine()
}

func (buff *b) imports() {
	if len(imps) == 0 {
		return
	}
	buff.write("import (")
	buff.newLine()
	buff.tab()
	for _, imp := range imps {
		buff.write(imp)
		buff.newLine()
	}
	buff.shiftab()
	buff.write(')')
	buff.newLine()
}
func (buff *b)newLine() {
	buff.write('\n')
}

func (buff *b)finalize(rolesMap map[string][]string) {
	buff.write("var rolesMap = map[string][]string{")
	buff.tab()
	buff.newLine()
	rolesArr := getSortedRolesArrByMap(rolesMap)
	for _, v := range rolesArr {
		buff.write("\"", v.method, "\": {")
		size := len(v.roles)
		for i, role := range v.roles {
			appendAfter := ""
			if i < size - 1 {
				appendAfter = ", "
			}
			buff.appned("\"", role, "\"", appendAfter)
		}
		buff.appned("},")
		buff.newLine()
	}
	buff.shiftab()
	buff.write("}")

	buff.newLine()
	buff.newLine()
	buff.write("func CheckRole(method, role string)bool{")
	buff.newLine()
	buff.tab()

	buff.write("v, ok := rolesMap[method]")
	buff.newLine()
	buff.write("if !ok {")
	buff.newLine()
	buff.tab()
	buff.write("// If there is no method in the map then it's not protected")
	buff.newLine()
	buff.write("return true")
	buff.shiftab()
	buff.newLine()
	buff.write("}")
	buff.newLine()
	buff.write("for _, roleToCheck := range v {")
	buff.newLine()
	buff.tab()
	buff.write("if role == roleToCheck {")
	buff.newLine()
	buff.tab()
	buff.write("return true")
	buff.shiftab()
	buff.newLine()
	buff.write("}")
	buff.newLine()
	buff.shiftab()
	buff.write("}")
	buff.newLine()
	buff.write("return false")

	buff.shiftab()
	buff.newLine()
	buff.write("}")
}

func (buff *b) tab() {
	buff.prefix = buff.prefix + "\t"
}
func (buff *b) shiftab() {
	buff.prefix = buff.prefix[1:]
}

func (buff *b) write(value ...interface{}) {
	buff.WriteString(buff.prefix)
	buff.appned(value...)
}

func (buff *b)appned(value ...interface{}){
	for _, v := range value {
		switch t := v.(type) {
		case string:
			buff.WriteString(t)
		case *string:
			buff.WriteString(*t)
		case byte:
			buff.WriteByte(t)
		case rune:
			buff.WriteRune(t)
		default:
			panic(fmt.Sprintf("Unknown type: %+v", t))
		}
	}
}

func generatedFileName(protoFile *descriptor.FileDescriptorProto) string {
	fileName := protoFile.GetName()
	if strings.HasSuffix(fileName, ".proto") {
		fileName = fileName[:len(fileName)-len(".proto")]
	}
	goPackage := strings.TrimSpace(protoFile.Options.GetGoPackage())
	if len(goPackage) > 0 {
		index := strings.Index(goPackage, ";")
		srcIndex := strings.LastIndex(fileName, "/")
		if srcIndex == -1 {
			fileName = goPackage[index + 1:] + "/" + fileName
		} else {
			fileName = fileName[:srcIndex + 1] + goPackage[index + 1:] + "/" + fileName[srcIndex + 1:]
		}

	}

	outputFile := fileName + ".roles" + ".go"

	return outputFile
}

func getSortedRolesArrByMap(roles map[string][]string) arrMapping {
	rolesArr := []rpcMethodWithRoles{}
	for k, v := range roles {
		rolesArr = append(rolesArr, rpcMethodWithRoles{k, v})
	}

	out := arrMapping(rolesArr)
	sort.Sort(out)
	return out
}