// Code generated from java.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // java

import "github.com/antlr4-go/antlr/v4"

// BasejavaListener is a complete listener for a parse tree produced by javaParser.
type BasejavaListener struct{}

var _ javaListener = &BasejavaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasejavaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasejavaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasejavaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasejavaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BasejavaListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {}

// ExitMethodDeclaration is called when production methodDeclaration is exited.
func (s *BasejavaListener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {}

// EnterModifiers is called when production modifiers is entered.
func (s *BasejavaListener) EnterModifiers(ctx *ModifiersContext) {}

// ExitModifiers is called when production modifiers is exited.
func (s *BasejavaListener) ExitModifiers(ctx *ModifiersContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BasejavaListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BasejavaListener) ExitModifier(ctx *ModifierContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BasejavaListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BasejavaListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BasejavaListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BasejavaListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BasejavaListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BasejavaListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterType is called when production type is entered.
func (s *BasejavaListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasejavaListener) ExitType(ctx *TypeContext) {}
