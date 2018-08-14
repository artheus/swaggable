# Swaggable

Simple language to generate larger swagger specs for code generation

# Swaggable syntax

## Types

All types specified in OpenApi are available. These are:
* string (this includes dates and files)
* number
* integer
* boolean
* array
* object

As parameters in the swaggable syntax there are two that stands out. It is `array` and `object`

### Arrays

The syntax for creating an array parameter in swaggable is `array<string> stringArray` where `array` is the
parameter type identifier, `string` is the array item type and `stringArray` is the name of the parameter.

### Objects

You will never write `object` as the type of a parameter in Swaggable. You rather use the name of an other
component as the type identifier for the parameter. E.g: `MyComponent parameterName`

## Base components

As in the `example/employers` script, there are two base components. The meaning of having these is to
have extendable components that should not be part of the Swagger specification components.

Other components can extends these bases, and will in the output have all the parameters of the base.

From `example/employers` 
```
base Entity {
  number id `required indexed`
  string created `format.datetime`
  string updated `format.datetime`
  Employee createdBy
  Employee updatedBy
}
```

## Regular components

All of the regular components will turn up in the OpenApi output in `#/components/schemas`

Inheritance can work two ways. Either inherit from a base component, which in the output will only
add the parameters from the base component. Or inherit from a regular component, which till use
the OpenApi type inheritance, because both models will exist in the same specification.

From `example/employers`, a regular component inheriting from a base component:
```
comp Key < Entity {
  Employee employee
  Customer customer `required`
  string position `required`
  array<Note> notes
}
```
