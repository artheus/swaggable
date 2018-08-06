parser grammar SwaggableParser;

options {
  tokenVocab=SwaggableLexer;
}

program
  : modelElements? EOF
  ;

objDeclaration
  : ComponentInit Identifier objTail
  ;

enumDeclaration
  : EnumInit Identifier '<' typeName '>' enumBlock
  ;

enumBlock
  : '{' Identifier* '}'
  ;

identifierList
  : Identifier (',' Identifier)*
  ;

extendsStatement
  : ('<' identifierList)
  ;

objTail
  : extendsStatement? '{' objElement* '}'
  ;

nativeType
  : StringType
  | NumberType
  | BooleanType
  | arrayType
  | mapType
  ;

arrayType
  : ArrayType '<' typeName '>'
  ;

mapType
  : MapType '<' typeName ',' typeName '>'
  ;

typeName
  : Identifier
  | nativeType
  ;

formatIdentifier
  : Format '.' Identifier
  ;

formatType
  : PasswordFormat
  | DateFormat
  | DateTimeFormat
  | ByteFormat
  | BinaryFormat
  | EmailFormat
  | UuidFormat
  | UriFormat
  | HostnameFormat
  | Ipv4Format
  | Ipv6Format
  | formatIdentifier
  ;

formatAnnotation
  : formatType
  ;

annotationType
  : Required
  | Indexed
  | Searchable
  | formatAnnotation
  ;

parameterDefinition
  : typeName Identifier ('(' annotationType+ ')')?
  ;

objElement
  : parameterDefinition
  ;

modelElement
  : objDeclaration
  | enumDeclaration
  ;

modelElements
  : modelElement+
  ;
