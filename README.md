# ndlvr

## WIP
This package is work-in-progress and is not production ready yet.

What started as another LIVR go library turned out to be this, since I've seen how strange stuff can LIVR do.
This library was created because [go-livr](https://github.com/k33nice/go-livr) can't do struct validation, which is like 99% of web app use case in go apps, since languages like Go, Rust, Java, Kotlin unlike JS, Python PHP tend to deserialize JSON(and other formats often) to predefined data structures rather than to language
type, which can contain any data.


## What's NDLVR?

Not-dynamic-langauge validation rules is set of improvements over LIVR that I've designed.

NDLVR is like LIVR but better. It:
* Is like LIVR
* But does no type juggling for example, in LIVR eq can change value from string to number and vice versa(!!)
* Has better support for non-dynamic languages like go or java
* Has no support for modification of incoming data, and thus does not remove not validated fields from data, even when it could be possible.
  The reason is quite simple: you can't remove struct field or java's class field, and sometimes modified value may not fit struct, so no modifications allowed. One should already sent valid data or preprocess it on it's own.
* Makes filters like `max_length` do not work on numbers and FP numbers and return error if not-number is passed
  In general makes filters work only with types they were designed to work with. Regex on numbers does not work as well.
* Has JOI-style builders for rules, so JSON does not have to be written by hand, which is nice.
* Some other minor stuff
* Has single reference implementation(go one), and all other implementations should mimic behaviour of reference one, this way bugs can be found
  by doing fuzzing and checking if results differ.

Aside from changes above it's compatibile with LIVR and for most use cases can be just used with existing LIVR rules and libraries.