package derivator

import "SuperCalculator/src/parser"

var Derivatives = map[string]func(x parser.Node) parser.Node{
	"sqrt": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("1"),
			parser.MakeOperationNode(
				parser.MUL,
				parser.MakeValueNode("2"),
				parser.MakeFunctionNode("sqrt", x),
			),
		)
	},
	"exp": func(x parser.Node) parser.Node {
		return parser.MakeFunctionNode("exp", x)
	},
	"ln": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("1"),
			x,
		)
	},
	"log": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("1"),
			x,
		)
	},
	"sin": func(x parser.Node) parser.Node {
		return parser.MakeFunctionNode("cos", x)
	},
	"cos": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.MUL,
			parser.MakeValueNode("-1"),
			parser.MakeFunctionNode("sin", x),
		)
	},
	"tan": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.ADD,
			parser.MakeValueNode("1"),
			parser.MakeOperationNode(
				parser.POW,
				parser.MakeFunctionNode("tan", x),
				parser.MakeValueNode("2"),
			),
		)
	},
	"asin": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("1"),
			parser.MakeFunctionNode(
				"sqrt",
				parser.MakeOperationNode(
					parser.SUB,
					parser.MakeValueNode("1"),
					parser.MakeOperationNode(
						parser.POW,
						x,
						parser.MakeValueNode("2"),
					),
				),
			),
		)
	},
	"acos": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("-1"),
			parser.MakeFunctionNode(
				"sqrt",
				parser.MakeOperationNode(
					parser.SUB,
					parser.MakeValueNode("1"),
					parser.MakeOperationNode(
						parser.POW,
						x,
						parser.MakeValueNode("2"),
					),
				),
			),
		)
	},
	"atan": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("1"),
			parser.MakeOperationNode(
				parser.ADD,
				parser.MakeValueNode("1"),
				parser.MakeOperationNode(
					parser.POW,
					x,
					parser.MakeValueNode("2"),
				),
			),
		)
	},
	"sinh": func(x parser.Node) parser.Node {
		return parser.MakeFunctionNode("cosh", x)
	},
	"cosh": func(x parser.Node) parser.Node {
		return parser.MakeFunctionNode("sinh", x)
	},
	"tanh": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.SUB,
			parser.MakeValueNode("1"),
			parser.MakeOperationNode(
				parser.POW,
				parser.MakeFunctionNode("tanh", x),
				parser.MakeValueNode("2"),
			),
		)
	},
	"asinh": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("1"),
			parser.MakeFunctionNode(
				"sqrt",
				parser.MakeOperationNode(
					parser.ADD,
					parser.MakeOperationNode(
						parser.POW,
						x,
						parser.MakeValueNode("2"),
					),
					parser.MakeValueNode("1"),
				),
			),
		)
	},
	"acosh": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("1"),
			parser.MakeFunctionNode(
				"sqrt",
				parser.MakeOperationNode(
					parser.SUB,
					parser.MakeOperationNode(
						parser.POW,
						x,
						parser.MakeValueNode("2"),
					),
					parser.MakeValueNode("1"),
				),
			),
		)
	},
	"atanh": func(x parser.Node) parser.Node {
		return parser.MakeOperationNode(
			parser.DIV,
			parser.MakeValueNode("1"),
			parser.MakeOperationNode(
				parser.SUB,
				parser.MakeValueNode("1"),
				parser.MakeOperationNode(
					parser.POW,
					x,
					parser.MakeValueNode("2"),
				),
			),
		)
	},
}
