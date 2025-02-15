package evaluator

import (
	"fmt"
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			// TODO fix this
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				len := len(arg.Value)
				return &object.Integer{Value: int64(len)}
			case *object.Array:
				len := len(arg.Elements)
				return &object.Integer{Value: int64(len)}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				len := len(arg.Elements)
				if len == 0 {
					return NULL
				}
				return arg.Elements[0]
			default:
				return newError("argument to `first` must be ARRAY, got %s",
					args[0].Type())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				len := len(arg.Elements)
				if len == 0 {
					return NULL
				}
				return arg.Elements[len-1]
			default:
				return newError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			switch arr := args[0].(type) {
			case *object.Array:
				// Arrays are immutable in monkey
				length := len(arr.Elements)
				newElements := make(
					[]object.Object,
					length+1,
				)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]
				return &object.Array{Elements: newElements}

			default:
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			switch arr := args[0].(type) {
			case *object.Array:
				// Arrays are immutable in monkey
				length := len(arr.Elements)
				if length == 0 {
					return NULL
				}
				newElements := make(
					[]object.Object,
					length-1,
				)
				copy(newElements, arr.Elements[1:])
				return &object.Array{Elements: newElements}

			default:
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
		},
	},
}
