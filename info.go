package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

type rpc struct {
	In     string
	Unary  bool
	Stream bool
}

type service struct {
	Factory string
	RPCs    map[string]rpc
}

func (s service) DefaultRPC() string {
	return keys(s.RPCs)[0]
}

type BenchClientInfo struct {
	PkgName  string
	Services map[string]*service
}

func (bci BenchClientInfo) NumServices() int {
	return len(bci.Services)
}

func (bci BenchClientInfo) DefaultService() string {
	return keys(bci.Services)[0]
}

func (bci BenchClientInfo) DefaultRPC() string {
	return bci.Services[bci.DefaultService()].DefaultRPC()
}

func (bci BenchClientInfo) UsageService() string {
	b := bytes.NewBufferString("`")
	fmt.Fprintln(b, "Specifies service to call:")
	for _, svcName := range keys(bci.Services) {
		fmt.Fprintf(b, "\t\t   %s\n", svcName)
	}
	b.WriteByte('`')
	return b.String()
}

func (bci BenchClientInfo) UsageRPC() string {
	b := bytes.NewBufferString("`")
	fmt.Fprintln(b, "Specifies rpc to call:")
	for _, svcName := range keys(bci.Services) {
		fmt.Fprintf(b, "\t   Valid rpc of %s:\n", svcName)
		for _, rpcName := range keys(bci.Services[svcName].RPCs) {
			fmt.Fprintf(b, "\t\t   %s\n", rpcName)
		}
	}
	b.WriteByte('`')
	return b.String()
}

func keys(v interface{}) []string {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		return []string{}
	}

	ret := make([]string, 0, len(rv.MapKeys()))
	for _, k := range rv.MapKeys() {
		ret = append(ret, k.String())
	}
	sort.Strings(ret)
	return ret
}
