parser grammar SwaggableParser;

options {
  tokenVocab=SwaggableLexer;
}

program
  : modelElements? EOF
  ;

baseDeclaration
  : BaseInit Identifier objTail
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
  : Identifier (Comma Identifier)*
  ;

extendsStatement
  : (LessThan identifierList)
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
  : ArrayType LessThan typeName MoreThan
  ;

mapType
  : MapType LessThan typeName Comma typeName MoreThan
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
  : typeName Identifier (Tick annotationType+ Tick)?
  ;

objElement
  : parameterDefinition
  ;

modelElement
  : objDeclaration
  | enumDeclaration
  | baseDeclaration
  ;

modelElements
  : modelElement+
  ;
