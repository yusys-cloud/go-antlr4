// Code generated from java.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // java

import "github.com/antlr4-go/antlr/v4"


// javaListener is a complete listener for a parse tree produced by javaParser.
type javaListener interface {
	antlr.ParseTreeListener

	// EnterMethodDeclaration is called when entering the methodDeclaration production.
	EnterMethodDeclaration(c *MethodDeclarationContext)

	// EnterModifiers is called when entering the modifiers production.
	EnterModifiers(c *ModifiersContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterMethodBody is called when entering the methodBody production.
	EnterMethodBody(c *MethodBodyContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// ExitMethodDeclaration is called when exiting the methodDeclaration production.
	ExitMethodDeclaration(c *MethodDeclarationContext)

	// ExitModifiers is called when exiting the modifiers production.
	ExitModifiers(c *ModifiersContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitMethodBody is called when exiting the methodBody production.
	ExitMethodBody(c *MethodBodyContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)
}
