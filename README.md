# Cosmos SDK module wiring PoC using google/wire

### Pros:
* we wire up things in a functional way. All dependencies defined in functions

#### Cons:
* not possible to load modules dynamically during the runtime


#### Neutral:
+ Other dependency injection tools for Go like dig or facebookgo/inject are based on reflection. Wire runs as a code generator, which means that the injector works without making calls to a runtime library. This enables easier introspection of initialization and correct cross-references for tooling and editors

+ Wire does not allow multiple providers for one type to exist in the transitive closure of the providers presented to wire.Build, as this is usually a mistake. For legitimate cases where you need multiple dependencies of the same type, you need to invent a new type to call this other dependency.
For the same reason, wire forbids including the same provider multiple times.
