// Code generated from java.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // java

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


type javaParser struct {
	*antlr.BaseParser
}

var JavaParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func javaParserInit() {
  staticData := &JavaParserStaticData
  staticData.LiteralNames = []string{
    "", "'('", "')'", "'public'", "'private'", "'protected'", "'static'", 
    "','", "'{'", "'/*'", "'*/'", "'}'", "'int'", "'void'", "'String'", 
    "'your_custom_type'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Identifier", 
    "WS", "LINE_COMMENT", "MULTILINE_COMMENT",
  }
  staticData.RuleNames = []string{
    "methodDeclaration", "modifiers", "modifier", "formalParameters", "formalParameter", 
    "methodBody", "type",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 19, 58, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 3, 0, 16, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 
	3, 0, 22, 8, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 28, 8, 1, 11, 1, 12, 1, 29, 
	1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 37, 8, 3, 10, 3, 12, 3, 40, 9, 3, 1, 
	4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 48, 8, 5, 10, 5, 12, 5, 51, 9, 5, 
	1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 49, 0, 7, 0, 2, 4, 6, 8, 10, 12, 
	0, 2, 1, 0, 3, 6, 1, 0, 12, 15, 55, 0, 15, 1, 0, 0, 0, 2, 27, 1, 0, 0, 
	0, 4, 31, 1, 0, 0, 0, 6, 33, 1, 0, 0, 0, 8, 41, 1, 0, 0, 0, 10, 44, 1, 
	0, 0, 0, 12, 55, 1, 0, 0, 0, 14, 16, 3, 2, 1, 0, 15, 14, 1, 0, 0, 0, 15, 
	16, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 18, 3, 12, 6, 0, 18, 19, 5, 16, 
	0, 0, 19, 21, 5, 1, 0, 0, 20, 22, 3, 6, 3, 0, 21, 20, 1, 0, 0, 0, 21, 22, 
	1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 24, 5, 2, 0, 0, 24, 25, 3, 10, 5, 0, 
	25, 1, 1, 0, 0, 0, 26, 28, 3, 4, 2, 0, 27, 26, 1, 0, 0, 0, 28, 29, 1, 0, 
	0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 3, 1, 0, 0, 0, 31, 32, 
	7, 0, 0, 0, 32, 5, 1, 0, 0, 0, 33, 38, 3, 8, 4, 0, 34, 35, 5, 7, 0, 0, 
	35, 37, 3, 8, 4, 0, 36, 34, 1, 0, 0, 0, 37, 40, 1, 0, 0, 0, 38, 36, 1, 
	0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 7, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 41, 
	42, 3, 12, 6, 0, 42, 43, 5, 16, 0, 0, 43, 9, 1, 0, 0, 0, 44, 45, 5, 8, 
	0, 0, 45, 49, 5, 9, 0, 0, 46, 48, 9, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48, 51, 
	1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 
	51, 49, 1, 0, 0, 0, 52, 53, 5, 10, 0, 0, 53, 54, 5, 11, 0, 0, 54, 11, 1, 
	0, 0, 0, 55, 56, 7, 1, 0, 0, 56, 13, 1, 0, 0, 0, 5, 15, 21, 29, 38, 49,
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

// javaParserInit initializes any static state used to implement javaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewjavaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JavaParserInit() {
  staticData := &JavaParserStaticData
  staticData.once.Do(javaParserInit)
}

// NewjavaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewjavaParser(input antlr.TokenStream) *javaParser {
	JavaParserInit()
	this := new(javaParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &JavaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "java.g4"

	return this
}


// javaParser tokens.
const (
	javaParserEOF = antlr.TokenEOF
	javaParserT__0 = 1
	javaParserT__1 = 2
	javaParserT__2 = 3
	javaParserT__3 = 4
	javaParserT__4 = 5
	javaParserT__5 = 6
	javaParserT__6 = 7
	javaParserT__7 = 8
	javaParserT__8 = 9
	javaParserT__9 = 10
	javaParserT__10 = 11
	javaParserT__11 = 12
	javaParserT__12 = 13
	javaParserT__13 = 14
	javaParserT__14 = 15
	javaParserIdentifier = 16
	javaParserWS = 17
	javaParserLINE_COMMENT = 18
	javaParserMULTILINE_COMMENT = 19
)

// javaParser rules.
const (
	javaParserRULE_methodDeclaration = 0
	javaParserRULE_modifiers = 1
	javaParserRULE_modifier = 2
	javaParserRULE_formalParameters = 3
	javaParserRULE_formalParameter = 4
	javaParserRULE_methodBody = 5
	javaParserRULE_type = 6
)

// IMethodDeclarationContext is an interface to support dynamic dispatch.
type IMethodDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Identifier() antlr.TerminalNode
	MethodBody() IMethodBodyContext
	Modifiers() IModifiersContext
	FormalParameters() IFormalParametersContext

	// IsMethodDeclarationContext differentiates from other interfaces.
	IsMethodDeclarationContext()
}

type MethodDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodDeclarationContext() *MethodDeclarationContext {
	var p = new(MethodDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_methodDeclaration
	return p
}

func InitEmptyMethodDeclarationContext(p *MethodDeclarationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_methodDeclaration
}

func (*MethodDeclarationContext) IsMethodDeclarationContext() {}

func NewMethodDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodDeclarationContext {
	var p = new(MethodDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = javaParserRULE_methodDeclaration

	return p
}

func (s *MethodDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MethodDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(javaParserIdentifier, 0)
}

func (s *MethodDeclarationContext) MethodBody() IMethodBodyContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodBodyContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodBodyContext)
}

func (s *MethodDeclarationContext) Modifiers() IModifiersContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifiersContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModifiersContext)
}

func (s *MethodDeclarationContext) FormalParameters() IFormalParametersContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParametersContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *MethodDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MethodDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.EnterMethodDeclaration(s)
	}
}

func (s *MethodDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.ExitMethodDeclaration(s)
	}
}




func (p *javaParser) MethodDeclaration() (localctx IMethodDeclarationContext) {
	localctx = NewMethodDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, javaParserRULE_methodDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 120) != 0) {
		{
			p.SetState(14)
			p.Modifiers()
		}

	}
	{
		p.SetState(17)
		p.Type_()
	}
	{
		p.SetState(18)
		p.Match(javaParserIdentifier)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(19)
		p.Match(javaParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 61440) != 0) {
		{
			p.SetState(20)
			p.FormalParameters()
		}

	}
	{
		p.SetState(23)
		p.Match(javaParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(24)
		p.MethodBody()
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


// IModifiersContext is an interface to support dynamic dispatch.
type IModifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllModifier() []IModifierContext
	Modifier(i int) IModifierContext

	// IsModifiersContext differentiates from other interfaces.
	IsModifiersContext()
}

type ModifiersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifiersContext() *ModifiersContext {
	var p = new(ModifiersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_modifiers
	return p
}

func InitEmptyModifiersContext(p *ModifiersContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_modifiers
}

func (*ModifiersContext) IsModifiersContext() {}

func NewModifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifiersContext {
	var p = new(ModifiersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = javaParserRULE_modifiers

	return p
}

func (s *ModifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifiersContext) AllModifier() []IModifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModifierContext); ok {
			len++
		}
	}

	tst := make([]IModifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModifierContext); ok {
			tst[i] = t.(IModifierContext)
			i++
		}
	}

	return tst
}

func (s *ModifiersContext) Modifier(i int) IModifierContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifierContext); ok {
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

	return t.(IModifierContext)
}

func (s *ModifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ModifiersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.EnterModifiers(s)
	}
}

func (s *ModifiersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.ExitModifiers(s)
	}
}




func (p *javaParser) Modifiers() (localctx IModifiersContext) {
	localctx = NewModifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, javaParserRULE_modifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 120) != 0) {
		{
			p.SetState(26)
			p.Modifier()
		}


		p.SetState(29)
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


// IModifierContext is an interface to support dynamic dispatch.
type IModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsModifierContext differentiates from other interfaces.
	IsModifierContext()
}

type ModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierContext() *ModifierContext {
	var p = new(ModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_modifier
	return p
}

func InitEmptyModifierContext(p *ModifierContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_modifier
}

func (*ModifierContext) IsModifierContext() {}

func NewModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierContext {
	var p = new(ModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = javaParserRULE_modifier

	return p
}

func (s *ModifierContext) GetParser() antlr.Parser { return s.parser }
func (s *ModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.EnterModifier(s)
	}
}

func (s *ModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.ExitModifier(s)
	}
}




func (p *javaParser) Modifier() (localctx IModifierContext) {
	localctx = NewModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, javaParserRULE_modifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 120) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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


// IFormalParametersContext is an interface to support dynamic dispatch.
type IFormalParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFormalParameter() []IFormalParameterContext
	FormalParameter(i int) IFormalParameterContext

	// IsFormalParametersContext differentiates from other interfaces.
	IsFormalParametersContext()
}

type FormalParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParametersContext() *FormalParametersContext {
	var p = new(FormalParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_formalParameters
	return p
}

func InitEmptyFormalParametersContext(p *FormalParametersContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_formalParameters
}

func (*FormalParametersContext) IsFormalParametersContext() {}

func NewFormalParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParametersContext {
	var p = new(FormalParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = javaParserRULE_formalParameters

	return p
}

func (s *FormalParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParametersContext) AllFormalParameter() []IFormalParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormalParameterContext); ok {
			len++
		}
	}

	tst := make([]IFormalParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormalParameterContext); ok {
			tst[i] = t.(IFormalParameterContext)
			i++
		}
	}

	return tst
}

func (s *FormalParametersContext) FormalParameter(i int) IFormalParameterContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParameterContext); ok {
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

	return t.(IFormalParameterContext)
}

func (s *FormalParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FormalParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.EnterFormalParameters(s)
	}
}

func (s *FormalParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.ExitFormalParameters(s)
	}
}




func (p *javaParser) FormalParameters() (localctx IFormalParametersContext) {
	localctx = NewFormalParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, javaParserRULE_formalParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.FormalParameter()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == javaParserT__6 {
		{
			p.SetState(34)
			p.Match(javaParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(35)
			p.FormalParameter()
		}


		p.SetState(40)
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


// IFormalParameterContext is an interface to support dynamic dispatch.
type IFormalParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Identifier() antlr.TerminalNode

	// IsFormalParameterContext differentiates from other interfaces.
	IsFormalParameterContext()
}

type FormalParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterContext() *FormalParameterContext {
	var p = new(FormalParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_formalParameter
	return p
}

func InitEmptyFormalParameterContext(p *FormalParameterContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_formalParameter
}

func (*FormalParameterContext) IsFormalParameterContext() {}

func NewFormalParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterContext {
	var p = new(FormalParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = javaParserRULE_formalParameter

	return p
}

func (s *FormalParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FormalParameterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(javaParserIdentifier, 0)
}

func (s *FormalParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FormalParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.EnterFormalParameter(s)
	}
}

func (s *FormalParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.ExitFormalParameter(s)
	}
}




func (p *javaParser) FormalParameter() (localctx IFormalParameterContext) {
	localctx = NewFormalParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, javaParserRULE_formalParameter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Type_()
	}
	{
		p.SetState(42)
		p.Match(javaParserIdentifier)
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


// IMethodBodyContext is an interface to support dynamic dispatch.
type IMethodBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMethodBodyContext differentiates from other interfaces.
	IsMethodBodyContext()
}

type MethodBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodBodyContext() *MethodBodyContext {
	var p = new(MethodBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_methodBody
	return p
}

func InitEmptyMethodBodyContext(p *MethodBodyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_methodBody
}

func (*MethodBodyContext) IsMethodBodyContext() {}

func NewMethodBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodBodyContext {
	var p = new(MethodBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = javaParserRULE_methodBody

	return p
}

func (s *MethodBodyContext) GetParser() antlr.Parser { return s.parser }
func (s *MethodBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MethodBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.EnterMethodBody(s)
	}
}

func (s *MethodBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.ExitMethodBody(s)
	}
}




func (p *javaParser) MethodBody() (localctx IMethodBodyContext) {
	localctx = NewMethodBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, javaParserRULE_methodBody)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(javaParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(javaParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(46)
			p.MatchWildcard()



		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(javaParserT__9)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(javaParserT__10)
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


// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = javaParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = javaParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(javaListener); ok {
		listenerT.ExitType(s)
	}
}




func (p *javaParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, javaParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 61440) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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


