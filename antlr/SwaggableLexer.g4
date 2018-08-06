lexer grammar SwaggableLexer;

channels { ERROR }

// operators
LessThan:     '<';
MoreThan:     '>';
OpenParen:    '(';
CloseParen:   ')';
OpenBrace:    '{';
CloseBrace:   '}';
Dot:          '.';
Comma:        ',';
Underscore:   '_';

// keywords
ComponentInit:    'comp';
EnumInit:         'enum';
//SecSchemeInit:    'sec';

StringType:       'string' | 'str';
NumberType:       'number' | 'num';
ArrayType:        'array' | 'arr';
BooleanType:      'boolean' | 'bool';
MapType:          'map';

// annotations
Required:     'required' | 'req';
Searchable:   'searchable' | 'search';
Indexed:      'indexed' | 'index' | 'idx';
Format:       'format' | 'form';

// openapi format types
PasswordFormat:     'format.password';
DateFormat:         'format.date';
DateTimeFormat:     'format.datetime';
ByteFormat:         'format.byte';
BinaryFormat:       'format.binary';

// other available format types
EmailFormat:        'format.email';
UuidFormat:         'format.uuid';
UriFormat:          'format.uri';
HostnameFormat:     'format.hostname';
Ipv4Format:         'format.ipv4';
Ipv6Format:         'format.ipv6';

// identifier
Identifier:         Letter LetterOrDigit*;

// fragments
fragment LetterOrDigit
    : Letter
    | [0-9]
    ;

fragment Letter
    : [a-zA-Z$_]
    | ~[\u0000-\u007F\uD800-\uDBFF]
    | [\uD800-\uDBFF] [\uDC00-\uDFFF]
    ;

// lexer rules
WS:                 [ \t\r\n\u000C]+ -> channel(HIDDEN);
COMMENT:            '/*' .*? '*/'    -> channel(HIDDEN);
LINE_COMMENT:       '//' ~[\r\n]*    -> channel(HIDDEN);
