package program

import "github.com/xelaj/tl"

// Program represents TL formal described in https://core.telegram.org/mtproto/TL-formal#syntax
// Syntactically, a TL program consists of a stream of tokens (separated by spaces, which are ignored at this stage). General program structure:
//
// TL-program ::= constr-declarations { --- functions --- fun-declarations | --- types --- constr-declarations }
type Program struct {
	Entries []ProgramEntry `parser:"@@*"`
}

// ProgramEntry represents TL formal described in https://core.telegram.org/mtproto/TL-formal#syntax
//
// StartTypeDecl is used to indicate if the following declarations are type declarations
// constr-declarations ::= { declaration }
//
// StartFuncDecl is used to indicate if the following declarations are function declarations
// Newline is used to represent if the line is empty
// Declaration is used to represent functions and types
// declaration ::= combinator-decl | partial-app-decl | final-decl
// constr-declarations ::= { declaration }
// fun-declarations ::= { declaration }
//
// Comment is used to represent comments with `//`
type ProgramEntry struct {
	TypeDecl    bool         `parser:"@(constraints_decl newline) |"`
	FuncDecl    bool         `parser:"@(functions_decl newline) |"`
	Newline     bool         `parser:"@newline |"`
	Declaration *Declaration `parser:"@@ ws? semicolon ( newline | EOF ) |"`
	Comment     *string      `parser:"@comment ( newline | EOF ) "`
}

func ParseProgram(program *Program) (*tl.Schema, error) {
	var (
		decls    = []tl.Declaration{}
		comments = []Annotation{}
		funcMode = false
	)
	for _, e := range program.Entries {
		switch {
		case e.TypeDecl:
			funcMode = false
		case e.FuncDecl:
			funcMode = true
		case e.Newline:
			comments = []Annotation{}

		case e.Comment != nil:
			a, err := ParseAnnotation(*e.Comment)
			if err != nil {
				return nil, err
			}
			if a != nil {
				comments = append(comments, *a)
			}

		case e.Declaration != nil:
			d, err := ParseDeclaration(e.Declaration, comments)
			if err != nil {
				return nil, err
			}
			if funcMode {
				d.Category = tl.CategoryFunction
			} else {
				d.Category = tl.CategoryPredict
			}
			// todo check crc
			//if decl.getCRC() != d.CRC {
			//	return nil, nil, nil, errors.New(decl.Name.String() + ": CRC mismatch! Described: " + fmt.Sprintf("0x%08x", decl.CRC) + " Calculated: " + fmt.Sprintf("0x%08x", decl.getCRC()))
			//}
			decls = append(decls, *d)
			comments = []Annotation{}
		}
	}
	return &tl.Schema{
		Declarations: decls,
	}, nil
}
