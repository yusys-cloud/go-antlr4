parser grammar JavaScriptParser2;

options {
    tokenVocab=JavaScriptLexer; // 使用 JavaScriptLexer 中的词法规则
}
program: statement*;

statement: functionDeclaration;

functionDeclaration: 'function' Identifier '(' ')' functionBody;

functionBody: '{' .*? '}';
