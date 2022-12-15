The basic theory of polymorphism:

* Polymorphism is the ability of a message to be displayed in more than one form. Polymorphism is considered as one of the important features of Object-Oriented Programming and can be achieved during either at runtime or compile time.

Example:

class French:
    def say_hello(self):
        print("bonjour")

class Germany:
    def say_hello(self):
        print("hallo")

def intro(lang):
    lang.say_hello()

ramesh = French()
suresh = Germany()

intro(ramesh) # output: bonjour
intro(suresh) # output: hallo