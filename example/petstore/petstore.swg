comp NewPet {
  string name (required indexed)
}

comp Pet < NewPet {
  number id (required indexed)
}

comp MapComponent {
  map<string,Note> noteMap (required)
}
