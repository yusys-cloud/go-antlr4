// Code generated from JavaScriptParser2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // JavaScriptParser2

import "github.com/antlr4-go/antlr/v4"


// JavaScriptParser2Listener is a complete listener for a parse tree produced by JavaScriptParser2.
type JavaScriptParser2Listener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)
}
