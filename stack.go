package rogu

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Stack struct {
	stack string
	funcs []StackInformation
}

func (stack Stack) GetInfo() []StackInformation {
	return stack.funcs
}

func (stack Stack) HowDeep() int {
	return len(stack.funcs)
}

type StackInformation struct {
	funcName       string
	moduleName     string
	funcParameters []string
	fileName       string
	lineNumber     string
}

func (info StackInformation) GetFuncName() string {
	return info.funcName
}

func (info StackInformation) GetModuleName() string {
	return info.moduleName
}

func (info StackInformation) GetFuncParameters() []string {
	return info.funcParameters
}

func (info StackInformation) GetFileName() string {
	return info.fileName
}

func (info StackInformation) GetLineNumber() string {
	return info.lineNumber
}

func (info StackInformation) WriteInfo(writer io.Writer) {
	param := ""
	for index, par := range info.GetFuncParameters() {
		if index == 0 {
			param += par
		}
		param += ", "
		param += par
	}
	writer.Write([]byte(
		fmt.Sprintf("\n%s %s:%s\n      - %s(%s)\n", info.GetModuleName(), info.GetFileName(), info.GetLineNumber(), info.GetFuncName(), param),
	))
}

func NewStack(stackData []byte) Stack {
	stack := Stack{stack: string(stackData)}
	stackString := string(stackData)

	info := StackInformation{}
	for index, line := range strings.Split(stackString, "\n") {
		if index == 0 || line == "" {
			continue
		}
		if index%2 == 1 {
			nLine := strings.ReplaceAll(line, "...", "0")
			lineArg := strings.Split(nLine, ".")
			funcCallText := lineArg[len(lineArg)-1]
			if len(lineArg) > 1 {
				info.moduleName = lineArg[0]
			}
			expression := regexp.MustCompile(`(.*)\((.*)\)$`)
			stringParts := expression.FindStringSubmatch(funcCallText)
			funcName := ""
			// ToDo handle some cases that are not getting the func name probably
			if len(stringParts) > 0 {
				funcName = stringParts[1]
			}
			info.funcName = funcName
			if len(stringParts) >= 2 && stringParts[2] != "" {
				for _, parameter := range strings.Split(stringParts[2], ",") {
					info.funcParameters = append(info.funcParameters, parameter)
				}
			}
		} else {
			parts := strings.Split(line, "/")
			part := parts[len(parts)-1:][0]
			parts2 := strings.Split(part, ":")
			info.fileName = parts2[0]
			info.lineNumber = strings.Split(parts2[1], " ")[0]

			stack.funcs = append(stack.funcs, info)
			info = StackInformation{}
		}

	}
	return stack
}