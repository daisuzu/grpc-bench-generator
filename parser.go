package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	unaryArgs  = 3
	streamArgs = 2
)

func detectPkgName(path string) string {
	for _, p := range filepath.SplitList(os.Getenv("GOPATH")) {
		if strings.HasPrefix(path, p) {
			// remove the leading slash
			return strings.Replace(path, filepath.Join(p, "src"), "", 1)[1:]
		}
	}
	// relative import if not found in GOPATH
	return "./" + filepath.Base(path)
}

func parseFile(bci *BenchClientInfo, f *ast.File) {
	for _, d := range f.Decls {
		switch decl := d.(type) {
		case *ast.GenDecl:
			parseGenDecl(bci, decl)
		case *ast.FuncDecl:
			parseFuncDecl(bci, decl)
		}
	}
}

// parseGenDecl searches interfaces and its methods of Client API from decl,
// and add it to BenchClientInfo.
func parseGenDecl(bci *BenchClientInfo, decl *ast.GenDecl) {
	if decl.Tok != token.TYPE {
		return
	}

	for _, s := range decl.Specs {
		ts := s.(*ast.TypeSpec)
		iface, ok := ts.Type.(*ast.InterfaceType)
		if !ok {
			continue
		}

		ifaceName := ts.Name.String()
		if !strings.HasSuffix(ifaceName, "Client") {
			continue
		}
		if strings.Index(ifaceName, "Service_") > 0 {
			// TODO: support streaming RPC
			continue
		}
		bci.Services[ifaceName] = &service{RPCs: make(map[string]rpc)}

		for _, m := range iface.Methods.List {
			methodName := m.Names[0].String()
			mt := m.Type.(*ast.FuncType)

			var streaming bool
			for _, r := range mt.Results.List {
				if strings.Index(types.ExprString(r.Type), "Service_") > 0 {
					// TODO: support streaming RPC
					streaming = true
					break
				}
			}

			switch {
			case mt.Params.NumFields() == unaryArgs && !streaming:
				argName := types.ExprString(mt.Params.List[1].Type)[1:]
				bci.Services[ifaceName].RPCs[methodName] = rpc{Unary: true, In: argName}
			default:
				// TODO: support streaming RPC
				bci.Services[ifaceName].RPCs[methodName] = rpc{Stream: true}
			}
		}
	}
}

// parseFuncDecl searches factory functions of Client API Interface from decl,
// and add it to BenchClientInfo.
func parseFuncDecl(bci *BenchClientInfo, decl *ast.FuncDecl) {
	if !strings.HasPrefix(decl.Name.String(), "New") {
		return
	}

	funcName := decl.Name.String()
	for _, r := range decl.Type.Results.List {
		retName := types.ExprString(r.Type)
		if s, ok := bci.Services[retName]; ok {
			// s has already been added in parseGenDecl()
			s.Factory = funcName
		}
	}
}

func Parse(path string) BenchClientInfo {
	bci := BenchClientInfo{
		PkgName:  detectPkgName(path),
		Services: make(map[string]*service),
	}

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range pkgs {
		for _, f := range p.Files {
			parseFile(&bci, f)
		}
	}

	return bci
}
