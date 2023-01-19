## What is Design Patterns
  * Design patterns identify solutions to common programming problems, but they are not actual code or algorithims themselves.
  * Patterns are form of best practices for approaching a problem
  * Usage of patterns can be significantly influenced by the language
    Ex: patterns are useful in oops programming languages, may not work in non-oops  

#### Golang features that affect design patterns
* No support for traditional OOP classes like java, c++ etc.
    * No static members or constructors--> affetcs patterns like Singleton


* No supports for class based inheritance or method overloading
    * Affects patterns like Visitor


* No support for exceptions-error handling i.e via return values
    * Affects patterns like Builder


* No support for Abstract class
    * Affects patterns like Abstract Factory and Bridge


### Desgin Patterns Categories
  There are three Categories:
  1. Creational
  2. Structural
  3. Behavioral

  #### Creational Patterns:
    The creational patterns are primarily concerned with, as the category name implies, how objects are created during the lifetime of a program.

  * Abstract Factory
  * Builder
  * Factory
  * Lazy Initialization
  * Multiton
  * Object Pool
  * Prototype
  * Singleton

#### Structural Patterns:
  The structural patterns are used to organize the various parts of a program into larger, more complex structures, and to provide new functionality based on that organization
  * Adapter
  * Bridge
  * Composite
  * Decorator
  * Facade
  * Flyweight

#### Behavioral Patterns:
The behavioral patterns identify common patterns of communication and interaction between objects within a program and are intended to increase flexibility and efficiency in carrying out those communications and interactions.
  * Chain of responsibility
  * Command
  * Interpreter
  * Iterator
  * Mediator
  * Memento
  * Observer
  * Strategy
  * Visitor


### Details of design patterns with Example:

#### Builder design patterns
Using this patterns, all parameters is not required to give to a function always.

https://github.com/LinkedInLearning/go-design-patterns-2880139/tree/main/Finished/Creational/Builder

https://www.youtube.com/watch?v=KbIdk5BRn0w
