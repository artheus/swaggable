base Entity {
  number id `required indexed`
  string created `format.datetime`
  string updated `format.datetime`
  Employee createdBy
  Employee updatedBy
}

base NamedEntity < Entity {
  string name `required indexed`
}

comp Employee < NamedEntity {
  string email `required indexed format.email`
  string password `required format.password`
  string phone `required`
  array<Address> address `required`
  array<string> area
  array<Key> keys
  array<Note> notes
}

comp Customer < NamedEntity {
  string email `required indexed format.email`
  string phone `required`
  array<Address> address `required`
  array<Contact> contacts `required`
  Key key
  array<Note> notes
}

comp Contact < NamedEntity {
  string email `required indexed format.email`
  string phone `required`
  string personalNumber `required`
}

comp Address {
  string street `required`
  string zip `required`
  string type `required`
  string city
}

comp Key < Entity {
  Employee employee
  Customer customer `required`
  string position `required`
  array<Note> notes
}

comp Note < Entity {
  string message `required`
  string type
}

enum Area<string> {
  NORTH
  SOUTH
}
