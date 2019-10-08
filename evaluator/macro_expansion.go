package evaluator

//)
//
//func DefineMacros(program *ast.Program, env *object.Environment) {
//
//	definitions := []int{}
//	for i, statement := range program.Statements {
//		if isMacroDefinition(statement) {
//			addMacro(statement, env)
//			definitions = append(definitions, i)
//		}
//	}
//	for i := len(definitions) - 1; i >= 0; i = i -1 {
//		definitionIndex := definitions[i]
//		program.Statements = append(
//			program.Statements[:definitionIndex],
//			program.Statements[definitionIndex+1:]...,
//		)
//	}
//}
