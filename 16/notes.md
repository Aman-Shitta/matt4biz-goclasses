# Go does OOP

The essesntial elements od OOP;
    - abstrcation
    - encapsulation
    - polymorphism
    - inheritance

Sometimes last two items are combined or confused

Go's approach to OOP is similar but different

- Abstraction: decoupling behavious from the implementation details

The Unix file system API is great example of effective abstraction

Roughly five basic functions hide all messy details:
    - open
    - close
    - read
    - write
    - ioctl

many different OS things can be trated like files.

- Encapsulation: hiding implementation from misuse

Its hard to maintain an abstraction if detals are exposd:
    - the internals may be maniupulated in ways contrary to the concept behind the abstraction
    - users od the abstracton may come to depend on the internal details- but those might change.

Encapsulation usually means controlling the visibilty of names ("private" variables)



- Polymorphism: literally means "many shape" - multiple types behind a single interface

    Three main types are:
        - ad-hoc: typically found in funciton/operator overlaoding.
        - parameteric: commnnly known as "generic programming".
        - subtype: subclasses substituting for superclass.

    "Protocol-oriented" programming uses explicit interface types,
    now supported in many popular languages (an ad-hoc method)

    In this case, `behaviour is completely seperate from implement`,
    which is good for abstrcation


- Inheritance has conflicting meanings:
    - substitution (subtype) polymorphism.
    - structual sharing of implementation details

    In theodr, inheritance shoulde always imply subtypoing:
    the subclass should be "kind of" superclass

    See the Liskov substitution principle (https://en.wikipedia.org/wiki/Liskov_substitution_principle)

    Theories about substitution can be pretty messy.

# Why would inheritance be bad ?

It injects a dependence in the superclass into the subclass:
    - What is superclass changes behaviuor?
    - What is the abstract concept is leaky?

Not having Inhertcane means better encapsulation and isolation.

"Interfaces will force you to think in term of communication between objects"
    -- Nicolo Pignatelli in "Inheritance is evil"

See also Composition over inheritance and Inheritance Tax (Pragmatic)


# One More View

"Object Oriented Programming to me means only messaging,
local retention and protection and hiding of state-process,
and extreme late-binsing of all things"     --- Alan Kay


He wrote this to:
    - de-emphasize inheritance heirarchies as a key part of OOP
    - emphasize the idea of self-contained objects sending messages to each other.
    - emphasize polymorphism in behavior

# OO in GO

Go offers four main supports for OO Programming:
    - encapsul;ation using the package of visibility contorl
    - abstraction & polymorphism using interface types
    - enhanced "composition" to provide struture sharing

Go does not offer inheritance or substitutability based on types

Substitutability is based only on interfaces; purely a function of abstract behavior

See : https://go.dev/talks/2014/go4gophers.slide
