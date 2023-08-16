// Code generated from JavaScriptParser2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // JavaScriptParser2

import "github.com/antlr4-go/antlr/v4"

// BaseJavaScriptParser2Listener is a complete listener for a parse tree produced by JavaScriptParser2.
type BaseJavaScriptParser2Listener struct{}

var _ JavaScriptParser2Listener = &BaseJavaScriptParser2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJavaScriptParser2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJavaScriptParser2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJavaScriptParser2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJavaScriptParser2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseJavaScriptParser2Listener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseJavaScriptParser2Listener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseJavaScriptParser2Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJavaScriptParser2Listener) ExitStatement(ctx *StatementContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseJavaScriptParser2Listener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseJavaScriptParser2Listener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseJavaScriptParser2Listener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseJavaScriptParser2Listener) ExitFunctionBody(ctx *FunctionBodyContext) {}
