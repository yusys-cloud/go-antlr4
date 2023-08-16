// Code generated from JavaScriptParser2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // JavaScriptParser2

import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by JavaScriptParser2.
type JavaScriptParser2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JavaScriptParser2#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by JavaScriptParser2#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by JavaScriptParser2#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by JavaScriptParser2#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

}