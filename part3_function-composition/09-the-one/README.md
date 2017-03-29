## Style #9 - "The One"

### Recap

This style is a special case of function composition. It makes use of an abstraction (the *TFTheOne* structure
in this case) which wraps a value and takes a function to be applied on that value. After every function 
application, the abstraction holds on the new value and returns a reference (a pointer) to itself. So what you
get back is the same object or structure, but it's holding the new value depending on the procedure which was 
provided.  
The TFTheOne structure provides the following functions as its interface:

 * A "constructor" which takes the value that is wrapped by the abstraction (note that Golang has no
notion of classes and objects).
 * A *bind* method which takes an arbitrary function, applies it on the current value, stores the
  result of the call as the new value and returns a pointer to the abstraction struct.
  
    ```golang
    func (theOne *TFTheOne) bind(f func(arg interface{}) interface{}) *TFTheOne {
  	    theOne.value = f(theOne.value)
  	    return theOne
    }
    ```
  
 * A *printMe* method which publishes the abstraction's internal value by printing it to standard
  output.
  
    ```golang
    func (theOne *TFTheOne) printMe() {
    	fmt.Println(theOne.value)
    }
    ```
  
A very nice feature of this style is that it allows you to chain functions in a way that's very easy
to read and understand. If you remember the *Pipeline Style*, we ended up with something like this:

```golang 
  printAll(sortPairs(countFrequencies(removeStopWords(scan(filterCharsAndNormalize(readFile("input.txt")))))))
``` 

The order in which the functions are invoked and processed is from right to left, which makes it hard
to see what's going on at first sight. However, using this style, we can string together our procedures
in a very clear and well-arranged way:

```golang
  theOne := TFTheOne{value:"input.txt"}
  theOne.bind(readFile).
       .
       .
       .
  bind(top25Frequencies).
  printMe()
```


### Where does this style come from?

This style is derived from *Haskell*, a purely functional programming language which doesn't allow
functions to have any side effects (because: no assignments = no side effects). In the 1990's, the
concept of *monads* was brought to Haskell in order to integrate the idea of side effects into an
entirely functional programming language.    
A monad usually provides the following operations:

* A constructor which takes a value to be wrapped by the monad.
* A *bind* procedure that takes a function which is applied to the monad's internal value and
whose return value is captured by the monad. Finally, another monad (e.g. the current instance itself)
is returned.
* An additional method which publishes the monad's internal value in some way (e.g. print).
 
If you compare these specs to the implementation, you can see that this interface is pretty much
what the *TFTheOne* struct tries to mimic.


### Sources
* Lopes, Christina Videira. (2014). *Exercises in Programming Style*. Boca Raton: CRC Press, pp.74-78. 