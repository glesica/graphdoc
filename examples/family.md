# ::Graph Family::
A graph to describe a family. Families can have several kinds of entities
including people, animals, and neighbors (who are just people, in reality).

## ::Node Person::
People have some attributes and can own animals.

### ::Rel MOTHER_OF -> Person::
A person to whom this individual is a mother.

### ::Rel FATHER_OF -> Person::
A person to whom this individual is a father.

### ::Rel BROTHER_OF -> Person::
A person to whom this individual is a brother.

### ::Rel SISTER_OF -> Person::
A person to whom this individual is a sister.

### ::Rel NEIGHBOR_OF -> Person::
A person to whom this individual is a neighbor.

### ::Prop name:str::
The full name of the person.

### ::Prop age:num::
The age of the person.

## ::Node Animal::
An animal that is a part of the family.

### ::Prop name:str::
The name of the animal.

### ::Prop species:str::
The type of animal.

### ::Rel OWNED_BY->Person::
The human who owns this animal.
