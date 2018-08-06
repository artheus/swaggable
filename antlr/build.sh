#!/bin/bash

antlr4 -o parser -Dlanguage=Go SwaggableLexer.g4 SwaggableParser.g4
