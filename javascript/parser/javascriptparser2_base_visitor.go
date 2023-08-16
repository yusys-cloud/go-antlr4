// Code generated from JavaScriptParser2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // JavaScriptParser2

import "github.com/antlr4-go/antlr/v4"


type BaseJavaScriptParser2Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJavaScriptParser2Visitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParser2Visitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParser2Visitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavaScriptParser2Visitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}
