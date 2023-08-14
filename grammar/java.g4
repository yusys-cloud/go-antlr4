// JavaMethodParser.g4
grammar java;

// 导入必要的模块

// 定义语法规则
methodDeclaration: modifiers? type Identifier '(' formalParameters? ')' methodBody;

modifiers: modifier+;
modifier: 'public' | 'private' | 'protected' | 'static';

formalParameters: formalParameter (',' formalParameter)*;
formalParameter: type Identifier;

methodBody: '{' '/*' .*? '*/' '}';

type: 'int' | 'void' | 'String' | 'your_custom_type';

Identifier: [a-zA-Z_] [a-zA-Z_0-9]*;

// 忽略空白字符
WS: [ \t\r\n]+ -> skip;

// 忽略单行注释
LINE_COMMENT: '//' ~[\r\n]* -> skip;

// 忽略多行注释
MULTILINE_COMMENT: '/*' .*? '*/' -> skip;

// 其他语法规则...

