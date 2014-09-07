# @Graph Family
A graph to describe a family. Families can have several kinds of entities
including people, animals, and neighbors (who are just people, in reality).

## @Node Person
People have some attributes and can own animals.

### @Prop name

@Type str

The full name of the person.

### @Prop age

@Type num

The age of the person.

### @Rel MOTHER_OF

@From Person
@To Person

A person to whom this individual is a mother.

### @Rel FATHER_OF

@From Person
@To Person

A person to whom this individual is a father.

### @Rel BROTHER_OF

@From Person
@To Person

A person to whom this individual is a brother.

### @Rel SISTER_OF

@From Person
@To Person

A person to whom this individual is a sister.

### @Rel NEIGHBOR_OF

@From Person
@To Person

A person to whom this individual is a neighbor.

## @Node Animal
An animal that is a part of the family.

### @Prop name

@Type str

The name of the animal.

### @Prop species

@Type str

The type of animal.

### @Rel OWNED_BY

@From Animal
@To Person

The human who owns this animal.
