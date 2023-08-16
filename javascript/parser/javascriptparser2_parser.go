// Code generated from JavaScriptParser2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // JavaScriptParser2

import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type JavaScriptParser2 struct {
	*antlr.BaseParser
}

var JavaScriptParser2ParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func javascriptparser2ParserInit() {
  staticData := &JavaScriptParser2ParserStaticData
  staticData.LiteralNames = []string{
    "", "", "", "", "", "'['", "']'", "'('", "')'", "'{'", "", "'}'", "';'", 
    "','", "'='", "'?'", "'?.'", "':'", "'...'", "'.'", "'++'", "'--'", 
    "'+'", "'-'", "'~'", "'!'", "'*'", "'/'", "'%'", "'**'", "'??'", "'#'", 
    "'>>'", "'<<'", "'>>>'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", 
    "'==='", "'!=='", "'&'", "'^'", "'|'", "'&&'", "'||'", "'*='", "'/='", 
    "'%='", "'+='", "'-='", "'<<='", "'>>='", "'>>>='", "'&='", "'^='", 
    "'|='", "'**='", "'=>'", "'null'", "", "", "", "", "", "", "", "", "", 
    "", "'break'", "'do'", "'instanceof'", "'typeof'", "'case'", "'else'", 
    "'new'", "'var'", "'catch'", "'finally'", "'return'", "'void'", "'continue'", 
    "'for'", "'switch'", "'while'", "'debugger'", "'function'", "'this'", 
    "'with'", "'default'", "'if'", "'throw'", "'delete'", "'in'", "'try'", 
    "'as'", "'from'", "'class'", "'enum'", "'extends'", "'super'", "'const'", 
    "'export'", "'import'", "'async'", "'await'", "'yield'", "'implements'", 
    "", "", "'private'", "'public'", "'interface'", "'package'", "'protected'", 
    "'static'", "", "", "", "", "", "", "", "", "'${'",
  }
  staticData.SymbolicNames = []string{
    "", "HashBangLine", "MultiLineComment", "SingleLineComment", "RegularExpressionLiteral", 
    "OpenBracket", "CloseBracket", "OpenParen", "CloseParen", "OpenBrace", 
    "TemplateCloseBrace", "CloseBrace", "SemiColon", "Comma", "Assign", 
    "QuestionMark", "QuestionMarkDot", "Colon", "Ellipsis", "Dot", "PlusPlus", 
    "MinusMinus", "Plus", "Minus", "BitNot", "Not", "Multiply", "Divide", 
    "Modulus", "Power", "NullCoalesce", "Hashtag", "RightShiftArithmetic", 
    "LeftShiftArithmetic", "RightShiftLogical", "LessThan", "MoreThan", 
    "LessThanEquals", "GreaterThanEquals", "Equals_", "NotEquals", "IdentityEquals", 
    "IdentityNotEquals", "BitAnd", "BitXOr", "BitOr", "And", "Or", "MultiplyAssign", 
    "DivideAssign", "ModulusAssign", "PlusAssign", "MinusAssign", "LeftShiftArithmeticAssign", 
    "RightShiftArithmeticAssign", "RightShiftLogicalAssign", "BitAndAssign", 
    "BitXorAssign", "BitOrAssign", "PowerAssign", "ARROW", "NullLiteral", 
    "BooleanLiteral", "DecimalLiteral", "HexIntegerLiteral", "OctalIntegerLiteral", 
    "OctalIntegerLiteral2", "BinaryIntegerLiteral", "BigHexIntegerLiteral", 
    "BigOctalIntegerLiteral", "BigBinaryIntegerLiteral", "BigDecimalIntegerLiteral", 
    "Break", "Do", "Instanceof", "Typeof", "Case", "Else", "New", "Var", 
    "Catch", "Finally", "Return", "Void", "Continue", "For", "Switch", "While", 
    "Debugger", "Function_", "This", "With", "Default", "If", "Throw", "Delete", 
    "In", "Try", "As", "From", "Class", "Enum", "Extends", "Super", "Const", 
    "Export", "Import", "Async", "Await", "Yield", "Implements", "StrictLet", 
    "NonStrictLet", "Private", "Public", "Interface", "Package", "Protected", 
    "Static", "Identifier", "StringLiteral", "BackTick", "WhiteSpaces", 
    "LineTerminator", "HtmlComment", "CDataComment", "UnexpectedCharacter", 
    "TemplateStringStartExpression", "TemplateStringAtom",
  }
  staticData.RuleNames = []string{
    "program", "statement", "functionDeclaration", "functionBody",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 128, 32, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 5, 
	0, 10, 8, 0, 10, 0, 12, 0, 13, 9, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 
	1, 2, 1, 2, 1, 3, 1, 3, 5, 3, 25, 8, 3, 10, 3, 12, 3, 28, 9, 3, 1, 3, 1, 
	3, 1, 3, 1, 26, 0, 4, 0, 2, 4, 6, 0, 0, 29, 0, 11, 1, 0, 0, 0, 2, 14, 1, 
	0, 0, 0, 4, 16, 1, 0, 0, 0, 6, 22, 1, 0, 0, 0, 8, 10, 3, 2, 1, 0, 9, 8, 
	1, 0, 0, 0, 10, 13, 1, 0, 0, 0, 11, 9, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 
	12, 1, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 14, 15, 3, 4, 2, 0, 15, 3, 1, 0, 
	0, 0, 16, 17, 5, 89, 0, 0, 17, 18, 5, 119, 0, 0, 18, 19, 5, 7, 0, 0, 19, 
	20, 5, 8, 0, 0, 20, 21, 3, 6, 3, 0, 21, 5, 1, 0, 0, 0, 22, 26, 5, 9, 0, 
	0, 23, 25, 9, 0, 0, 0, 24, 23, 1, 0, 0, 0, 25, 28, 1, 0, 0, 0, 26, 27, 
	1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 27, 29, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 
	29, 30, 5, 11, 0, 0, 30, 7, 1, 0, 0, 0, 2, 11, 26,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// JavaScriptParser2Init initializes any static state used to implement JavaScriptParser2. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJavaScriptParser2(). You can call this function if you wish to initialize the static state ahead
// of time.
func JavaScriptParser2Init() {
  staticData := &JavaScriptParser2ParserStaticData
  staticData.once.Do(javascriptparser2ParserInit)
}

// NewJavaScriptParser2 produces a new parser instance for the optional input antlr.TokenStream.
func NewJavaScriptParser2(input antlr.TokenStream) *JavaScriptParser2 {
	JavaScriptParser2Init()
	this := new(JavaScriptParser2)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &JavaScriptParser2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "JavaScriptParser2.g4"

	return this
}


// JavaScriptParser2 tokens.
const (
	JavaScriptParser2EOF = antlr.TokenEOF
	JavaScriptParser2HashBangLine = 1
	JavaScriptParser2MultiLineComment = 2
	JavaScriptParser2SingleLineComment = 3
	JavaScriptParser2RegularExpressionLiteral = 4
	JavaScriptParser2OpenBracket = 5
	JavaScriptParser2CloseBracket = 6
	JavaScriptParser2OpenParen = 7
	JavaScriptParser2CloseParen = 8
	JavaScriptParser2OpenBrace = 9
	JavaScriptParser2TemplateCloseBrace = 10
	JavaScriptParser2CloseBrace = 11
	JavaScriptParser2SemiColon = 12
	JavaScriptParser2Comma = 13
	JavaScriptParser2Assign = 14
	JavaScriptParser2QuestionMark = 15
	JavaScriptParser2QuestionMarkDot = 16
	JavaScriptParser2Colon = 17
	JavaScriptParser2Ellipsis = 18
	JavaScriptParser2Dot = 19
	JavaScriptParser2PlusPlus = 20
	JavaScriptParser2MinusMinus = 21
	JavaScriptParser2Plus = 22
	JavaScriptParser2Minus = 23
	JavaScriptParser2BitNot = 24
	JavaScriptParser2Not = 25
	JavaScriptParser2Multiply = 26
	JavaScriptParser2Divide = 27
	JavaScriptParser2Modulus = 28
	JavaScriptParser2Power = 29
	JavaScriptParser2NullCoalesce = 30
	JavaScriptParser2Hashtag = 31
	JavaScriptParser2RightShiftArithmetic = 32
	JavaScriptParser2LeftShiftArithmetic = 33
	JavaScriptParser2RightShiftLogical = 34
	JavaScriptParser2LessThan = 35
	JavaScriptParser2MoreThan = 36
	JavaScriptParser2LessThanEquals = 37
	JavaScriptParser2GreaterThanEquals = 38
	JavaScriptParser2Equals_ = 39
	JavaScriptParser2NotEquals = 40
	JavaScriptParser2IdentityEquals = 41
	JavaScriptParser2IdentityNotEquals = 42
	JavaScriptParser2BitAnd = 43
	JavaScriptParser2BitXOr = 44
	JavaScriptParser2BitOr = 45
	JavaScriptParser2And = 46
	JavaScriptParser2Or = 47
	JavaScriptParser2MultiplyAssign = 48
	JavaScriptParser2DivideAssign = 49
	JavaScriptParser2ModulusAssign = 50
	JavaScriptParser2PlusAssign = 51
	JavaScriptParser2MinusAssign = 52
	JavaScriptParser2LeftShiftArithmeticAssign = 53
	JavaScriptParser2RightShiftArithmeticAssign = 54
	JavaScriptParser2RightShiftLogicalAssign = 55
	JavaScriptParser2BitAndAssign = 56
	JavaScriptParser2BitXorAssign = 57
	JavaScriptParser2BitOrAssign = 58
	JavaScriptParser2PowerAssign = 59
	JavaScriptParser2ARROW = 60
	JavaScriptParser2NullLiteral = 61
	JavaScriptParser2BooleanLiteral = 62
	JavaScriptParser2DecimalLiteral = 63
	JavaScriptParser2HexIntegerLiteral = 64
	JavaScriptParser2OctalIntegerLiteral = 65
	JavaScriptParser2OctalIntegerLiteral2 = 66
	JavaScriptParser2BinaryIntegerLiteral = 67
	JavaScriptParser2BigHexIntegerLiteral = 68
	JavaScriptParser2BigOctalIntegerLiteral = 69
	JavaScriptParser2BigBinaryIntegerLiteral = 70
	JavaScriptParser2BigDecimalIntegerLiteral = 71
	JavaScriptParser2Break = 72
	JavaScriptParser2Do = 73
	JavaScriptParser2Instanceof = 74
	JavaScriptParser2Typeof = 75
	JavaScriptParser2Case = 76
	JavaScriptParser2Else = 77
	JavaScriptParser2New = 78
	JavaScriptParser2Var = 79
	JavaScriptParser2Catch = 80
	JavaScriptParser2Finally = 81
	JavaScriptParser2Return = 82
	JavaScriptParser2Void = 83
	JavaScriptParser2Continue = 84
	JavaScriptParser2For = 85
	JavaScriptParser2Switch = 86
	JavaScriptParser2While = 87
	JavaScriptParser2Debugger = 88
	JavaScriptParser2Function_ = 89
	JavaScriptParser2This = 90
	JavaScriptParser2With = 91
	JavaScriptParser2Default = 92
	JavaScriptParser2If = 93
	JavaScriptParser2Throw = 94
	JavaScriptParser2Delete = 95
	JavaScriptParser2In = 96
	JavaScriptParser2Try = 97
	JavaScriptParser2As = 98
	JavaScriptParser2From = 99
	JavaScriptParser2Class = 100
	JavaScriptParser2Enum = 101
	JavaScriptParser2Extends = 102
	JavaScriptParser2Super = 103
	JavaScriptParser2Const = 104
	JavaScriptParser2Export = 105
	JavaScriptParser2Import = 106
	JavaScriptParser2Async = 107
	JavaScriptParser2Await = 108
	JavaScriptParser2Yield = 109
	JavaScriptParser2Implements = 110
	JavaScriptParser2StrictLet = 111
	JavaScriptParser2NonStrictLet = 112
	JavaScriptParser2Private = 113
	JavaScriptParser2Public = 114
	JavaScriptParser2Interface = 115
	JavaScriptParser2Package = 116
	JavaScriptParser2Protected = 117
	JavaScriptParser2Static = 118
	JavaScriptParser2Identifier = 119
	JavaScriptParser2StringLiteral = 120
	JavaScriptParser2BackTick = 121
	JavaScriptParser2WhiteSpaces = 122
	JavaScriptParser2LineTerminator = 123
	JavaScriptParser2HtmlComment = 124
	JavaScriptParser2CDataComment = 125
	JavaScriptParser2UnexpectedCharacter = 126
	JavaScriptParser2TemplateStringStartExpression = 127
	JavaScriptParser2TemplateStringAtom = 128
)

// JavaScriptParser2 rules.
const (
	JavaScriptParser2RULE_program = 0
	JavaScriptParser2RULE_statement = 1
	JavaScriptParser2RULE_functionDeclaration = 2
	JavaScriptParser2RULE_functionBody = 3
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParser2RULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParser2RULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParser2RULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParser2Listener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParser2Listener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavaScriptParser2Visitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *JavaScriptParser2) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JavaScriptParser2RULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == JavaScriptParser2Function_ {
		{
			p.SetState(8)
			p.Statement()
		}


		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDeclaration() IFunctionDeclarationContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParser2RULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParser2RULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParser2RULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParser2Listener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParser2Listener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavaScriptParser2Visitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *JavaScriptParser2) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JavaScriptParser2RULE_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.FunctionDeclaration()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	OpenParen() antlr.TerminalNode
	CloseParen() antlr.TerminalNode
	FunctionBody() IFunctionBodyContext

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParser2RULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParser2RULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParser2RULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) Function_() antlr.TerminalNode {
	return s.GetToken(JavaScriptParser2Function_, 0)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(JavaScriptParser2Identifier, 0)
}

func (s *FunctionDeclarationContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParser2OpenParen, 0)
}

func (s *FunctionDeclarationContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParser2CloseParen, 0)
}

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionBodyContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParser2Listener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParser2Listener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavaScriptParser2Visitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *JavaScriptParser2) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JavaScriptParser2RULE_functionDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(JavaScriptParser2Function_)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(17)
		p.Match(JavaScriptParser2Identifier)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(18)
		p.Match(JavaScriptParser2OpenParen)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(19)
		p.Match(JavaScriptParser2CloseParen)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(20)
		p.FunctionBody()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpenBrace() antlr.TerminalNode
	CloseBrace() antlr.TerminalNode

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParser2RULE_functionBody
	return p
}

func InitEmptyFunctionBodyContext(p *FunctionBodyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JavaScriptParser2RULE_functionBody
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParser2RULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParser2OpenBrace, 0)
}

func (s *FunctionBodyContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParser2CloseBrace, 0)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParser2Listener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParser2Listener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (s *FunctionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavaScriptParser2Visitor:
		return t.VisitFunctionBody(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *JavaScriptParser2) FunctionBody() (localctx IFunctionBodyContext) {
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JavaScriptParser2RULE_functionBody)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(JavaScriptParser2OpenBrace)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(23)
			p.MatchWildcard()



		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Match(JavaScriptParser2CloseBrace)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


