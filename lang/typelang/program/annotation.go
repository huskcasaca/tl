package program

import (
	"errors"
	"strings"
)

const (
	AnnotationConstructorTag = "const"
	AnnotationParamTag       = "param"
	AnnotationReturnTag      = "return"
)

// Annotation represents an annotation comment, like //@name value.
type Annotation struct {
	// Tag
	//
	// it can be one of:
	//	"const" if Value is description for predict or function
	//	"param" if Value is description for parameters
	//	"return" if Value is description for type or function return type
	Tag string `json:"Case"`
	// Value of annotation.
	Value string `json:"value"`
}

func (a Annotation) String() string {
	return "//@" + a.Tag + " " + a.Value
}

func (a Annotation) SplitParam() (param, comment string) {

	const commentParts = 2
	parts := strings.SplitN(a.Value, " ", commentParts)
	comment = ""
	param = parts[0]
	if len(parts) >= commentParts {
		comment = strings.TrimSpace(parts[1])
	}
	return
}

func ParseAnnotation(annotation string) (*Annotation, error) {
	comment := strings.TrimSpace(strings.TrimPrefix(annotation, "//"))
	var commentTag string
	if strings.HasPrefix(comment, "@") {
		const commentParts = 2

		comment = strings.TrimPrefix(comment, "@")
		parts := strings.SplitN(comment, " ", commentParts)

		commentTag = parts[0]
		if len(parts) >= commentParts {
			comment = strings.TrimSpace(parts[1])
		} else {
			return nil, nil
		}
	} else {
		return nil, nil
	}

	switch tag := commentTag; tag {
	case AnnotationConstructorTag, AnnotationParamTag, AnnotationReturnTag:
		return &Annotation{Tag: tag, Value: comment}, nil
	default:
		return nil, errors.New("@" + commentTag + ": invalid comment tag")
	}
}
