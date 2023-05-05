# SuperCalculator

SuperCalculator is a CLI calculator app that can do various useful things :
  - Compute simple and more complicated calculations
  - Store values in custom variables
  - Solve equations
  - Determine the derivative of a function
  - And more...

Ultimately, I want to make this app not only a CLI app, but also a library to maybe
make a GUI for it.

One other goal is to not use any (or a very few) libraries, to learn new things and
build up an experience.


## Roadmap
  - [x] Math expression parser
  - [x] Computing calculations
  - [x] Custom variable and constants (such as PI)
  - [x] Derivative of a function
  - [ ] Math expression simplifier
  - [ ] Solve linear and quadratic equations
  - [ ] Solve certain integrals


## Syntax
- Use 
  - `+` for addition
  - `-` for substraction
  - `*` for multiplication
  - `/` for division
  - `^` for exponent. 
  
  *Example : `-(2 + 3)^2 = -25`*
- To define a custom variable, type : `[variable name] = [variable value]`. You can't define already defined constants such as `PI` or `e`, and you can't use `x` as a variable name.
- Most common functions are implemented. *Examples : `exp 1 ≈ 2.718`, or `cos(PI/2) = 0`*
- To get the derivative of a function in terms of x, type `derive [expression]`. *Example : `derive cos x + x^2 = -sin x + 2x`*


## Contributing 
This is a personal project more than anything, so I'll most likely not accept pull requests.

If you have a suggestion, or if you encountered a bug, feel free to open an [issue](https://github.com/leo-210/SuperCalculator/issues).
 
