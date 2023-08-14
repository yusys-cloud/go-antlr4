// Code generated from java.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type JavaMethodParserLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var JavaMethodParserLexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func javamethodparserlexerLexerInit() {
  staticData := &JavaMethodParserLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
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
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
    "T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "Identifier", "WS", 
    "LINE_COMMENT", "MULTILINE_COMMENT",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 19, 159, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 1, 1, 1, 
	1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 
	1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 
	1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 
	1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 
	1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 
	13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 
	1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 
	15, 5, 15, 123, 8, 15, 10, 15, 12, 15, 126, 9, 15, 1, 16, 4, 16, 129, 8, 
	16, 11, 16, 12, 16, 130, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 
	139, 8, 17, 10, 17, 12, 17, 142, 9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 
	18, 1, 18, 5, 18, 150, 8, 18, 10, 18, 12, 18, 153, 9, 18, 1, 18, 1, 18, 
	1, 18, 1, 18, 1, 18, 1, 151, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 
	13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 
	16, 33, 17, 35, 18, 37, 19, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97, 122, 4, 
	0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 
	10, 10, 13, 13, 162, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 
	0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 
	0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 
	0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 
	1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 
	37, 1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 41, 1, 0, 0, 0, 5, 43, 1, 0, 0, 0, 
	7, 50, 1, 0, 0, 0, 9, 58, 1, 0, 0, 0, 11, 68, 1, 0, 0, 0, 13, 75, 1, 0, 
	0, 0, 15, 77, 1, 0, 0, 0, 17, 79, 1, 0, 0, 0, 19, 82, 1, 0, 0, 0, 21, 85, 
	1, 0, 0, 0, 23, 87, 1, 0, 0, 0, 25, 91, 1, 0, 0, 0, 27, 96, 1, 0, 0, 0, 
	29, 103, 1, 0, 0, 0, 31, 120, 1, 0, 0, 0, 33, 128, 1, 0, 0, 0, 35, 134, 
	1, 0, 0, 0, 37, 145, 1, 0, 0, 0, 39, 40, 5, 40, 0, 0, 40, 2, 1, 0, 0, 0, 
	41, 42, 5, 41, 0, 0, 42, 4, 1, 0, 0, 0, 43, 44, 5, 112, 0, 0, 44, 45, 5, 
	117, 0, 0, 45, 46, 5, 98, 0, 0, 46, 47, 5, 108, 0, 0, 47, 48, 5, 105, 0, 
	0, 48, 49, 5, 99, 0, 0, 49, 6, 1, 0, 0, 0, 50, 51, 5, 112, 0, 0, 51, 52, 
	5, 114, 0, 0, 52, 53, 5, 105, 0, 0, 53, 54, 5, 118, 0, 0, 54, 55, 5, 97, 
	0, 0, 55, 56, 5, 116, 0, 0, 56, 57, 5, 101, 0, 0, 57, 8, 1, 0, 0, 0, 58, 
	59, 5, 112, 0, 0, 59, 60, 5, 114, 0, 0, 60, 61, 5, 111, 0, 0, 61, 62, 5, 
	116, 0, 0, 62, 63, 5, 101, 0, 0, 63, 64, 5, 99, 0, 0, 64, 65, 5, 116, 0, 
	0, 65, 66, 5, 101, 0, 0, 66, 67, 5, 100, 0, 0, 67, 10, 1, 0, 0, 0, 68, 
	69, 5, 115, 0, 0, 69, 70, 5, 116, 0, 0, 70, 71, 5, 97, 0, 0, 71, 72, 5, 
	116, 0, 0, 72, 73, 5, 105, 0, 0, 73, 74, 5, 99, 0, 0, 74, 12, 1, 0, 0, 
	0, 75, 76, 5, 44, 0, 0, 76, 14, 1, 0, 0, 0, 77, 78, 5, 123, 0, 0, 78, 16, 
	1, 0, 0, 0, 79, 80, 5, 47, 0, 0, 80, 81, 5, 42, 0, 0, 81, 18, 1, 0, 0, 
	0, 82, 83, 5, 42, 0, 0, 83, 84, 5, 47, 0, 0, 84, 20, 1, 0, 0, 0, 85, 86, 
	5, 125, 0, 0, 86, 22, 1, 0, 0, 0, 87, 88, 5, 105, 0, 0, 88, 89, 5, 110, 
	0, 0, 89, 90, 5, 116, 0, 0, 90, 24, 1, 0, 0, 0, 91, 92, 5, 118, 0, 0, 92, 
	93, 5, 111, 0, 0, 93, 94, 5, 105, 0, 0, 94, 95, 5, 100, 0, 0, 95, 26, 1, 
	0, 0, 0, 96, 97, 5, 83, 0, 0, 97, 98, 5, 116, 0, 0, 98, 99, 5, 114, 0, 
	0, 99, 100, 5, 105, 0, 0, 100, 101, 5, 110, 0, 0, 101, 102, 5, 103, 0, 
	0, 102, 28, 1, 0, 0, 0, 103, 104, 5, 121, 0, 0, 104, 105, 5, 111, 0, 0, 
	105, 106, 5, 117, 0, 0, 106, 107, 5, 114, 0, 0, 107, 108, 5, 95, 0, 0, 
	108, 109, 5, 99, 0, 0, 109, 110, 5, 117, 0, 0, 110, 111, 5, 115, 0, 0, 
	111, 112, 5, 116, 0, 0, 112, 113, 5, 111, 0, 0, 113, 114, 5, 109, 0, 0, 
	114, 115, 5, 95, 0, 0, 115, 116, 5, 116, 0, 0, 116, 117, 5, 121, 0, 0, 
	117, 118, 5, 112, 0, 0, 118, 119, 5, 101, 0, 0, 119, 30, 1, 0, 0, 0, 120, 
	124, 7, 0, 0, 0, 121, 123, 7, 1, 0, 0, 122, 121, 1, 0, 0, 0, 123, 126, 
	1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 32, 1, 0, 
	0, 0, 126, 124, 1, 0, 0, 0, 127, 129, 7, 2, 0, 0, 128, 127, 1, 0, 0, 0, 
	129, 130, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 
	132, 1, 0, 0, 0, 132, 133, 6, 16, 0, 0, 133, 34, 1, 0, 0, 0, 134, 135, 
	5, 47, 0, 0, 135, 136, 5, 47, 0, 0, 136, 140, 1, 0, 0, 0, 137, 139, 8, 
	3, 0, 0, 138, 137, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 
	0, 140, 141, 1, 0, 0, 0, 141, 143, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 
	144, 6, 17, 0, 0, 144, 36, 1, 0, 0, 0, 145, 146, 5, 47, 0, 0, 146, 147, 
	5, 42, 0, 0, 147, 151, 1, 0, 0, 0, 148, 150, 9, 0, 0, 0, 149, 148, 1, 0, 
	0, 0, 150, 153, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 
	152, 154, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 155, 5, 42, 0, 0, 155, 
	156, 5, 47, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 6, 18, 0, 0, 158, 38, 
	1, 0, 0, 0, 5, 0, 124, 130, 140, 151, 1, 6, 0, 0,
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

// JavaMethodParserLexerInit initializes any static state used to implement JavaMethodParserLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJavaMethodParserLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JavaMethodParserLexerInit() {
  staticData := &JavaMethodParserLexerLexerStaticData
  staticData.once.Do(javamethodparserlexerLexerInit)
}

// NewJavaMethodParserLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJavaMethodParserLexer(input antlr.CharStream) *JavaMethodParserLexer {
  JavaMethodParserLexerInit()
	l := new(JavaMethodParserLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &JavaMethodParserLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "java.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JavaMethodParserLexer tokens.
const (
	JavaMethodParserLexerT__0 = 1
	JavaMethodParserLexerT__1 = 2
	JavaMethodParserLexerT__2 = 3
	JavaMethodParserLexerT__3 = 4
	JavaMethodParserLexerT__4 = 5
	JavaMethodParserLexerT__5 = 6
	JavaMethodParserLexerT__6 = 7
	JavaMethodParserLexerT__7 = 8
	JavaMethodParserLexerT__8 = 9
	JavaMethodParserLexerT__9 = 10
	JavaMethodParserLexerT__10 = 11
	JavaMethodParserLexerT__11 = 12
	JavaMethodParserLexerT__12 = 13
	JavaMethodParserLexerT__13 = 14
	JavaMethodParserLexerT__14 = 15
	JavaMethodParserLexerIdentifier = 16
	JavaMethodParserLexerWS = 17
	JavaMethodParserLexerLINE_COMMENT = 18
	JavaMethodParserLexerMULTILINE_COMMENT = 19
)

