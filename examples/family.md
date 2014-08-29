# ::Graph Family::
A graph to describe a family.

## ::Node Person::
People have some attributes and can own animals.

### ::Rel BROTHER -> Person::
Another person who is the brother of this person.

### ::Prop name:str::
The name of the person.

### ::Prop age:num::
The age of the person.

## ::Node Animal::
Animals can belong to people.

### ::Prop name:str::
Animals can have names.

### ::Prop species:str::
The type of animal.

### ::Rel OWNED_BY->Person::
A human can own an animal.
