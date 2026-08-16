## SOLID Principles in Go

SOLID is a set of five principles that help you write code that is maintainable, extensible, testable, and loosely coupled.

In Go, SOLID is implemented mainly using interfaces, composition, and dependency injection rather than classes and inheritance.

### 1. S — Single Responsibility Principle

    A struct/type should have one responsibility or one reason to change.

### 2. O — Open/Closed Principle

Open for extension, closed for modification.

### 3. L — Liskov Substitution Principle

An implementation should be safely substitutable for the interface it implements.

### 4. I — Interface Segregation Principle

Don't force a type to implement methods it doesn't need.

### 5. D — Dependency Inversion Principle

High-level business logic should depend on abstractions, not concrete implementations.