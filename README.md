# ndlvr

What started as another LIVR go library turned out to be this, since I've seen how strange stuff can LIVR do.
This library was created because [go-livr](https://github.com/k33nice/go-livr) can't do struct validation, which is like 99.(9)% of web app use case in go apps.

## What's NDLVR?

Not-dynamic-langauge validation rules is set of improvements over LIVR that I've designed.

NDLVR is like LIVR but better. It:
* Is like LIVR
* But does no type juggling for example, in LIVR eq can change value from string to number and vice versa(!!)
* Has better support for non-dynamic languages like go or java
* Has split stages of modification and validation(note: for now modification is NIY, you should send valid data instead IMHO; also it's sometimes awkward to do so in strict-types environment when modified value may not fit field type that you're using, which results in inconsistent behaviour between languages and is just bad)
* Changes format a bit. Rules are intended to be built via pieces of code rather than using hand-written JSON.
  Code can validate rules, minfy them and in general, it's just better solution.
* Makes filters like `max_length` do not work on numbers and FP numbers and return error if not-number is passed
* Has more JOI-style APIs and validation
* Has more consistent rule serialization format, which is nice when you have to write parser for it(NIY, for now LIVR's is used)
* Some other minor sttuff
* Has single reference implementation, and all other implementation should mimic behaviour of reference one, this way bugs can be found
  by doing fuzzing and checking if results differ.