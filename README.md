# Config library

This library loads a json file and marshals it into a given `struct`.

By default, it loads `app.json` from sub-folder config. If you define a profile then this profile is added to file name.
E.g. staging: `app-staging.json`
You can set the profile via the environment variable `CONFIG_PROFILE` or load a specific profile with the function
`LoadFromWithProfile`.

**DO NOT ENTER SECRETS DIRECTLY INTO JSON FILE** Use environment variables instead.

## Special types

### EnvVariable

You can use this type if you want to configure something via environment variable or just not want to put the value
directly into json (secrets). If the variable is not set an error will be raised.

## Example

```json
{
  "attributeA": "Hello world!",
  "attributeB": 42,
  "secret": "{ENV:MY_SECRET}",
  "nested": {
    "value": true
  }
}
```

```go
package myPackage

import "github.com/fond-of/config.go/config"

// serviceConfig defines a configuration struct
type serviceConfig struct {
	AttributeA string
	AttributeB int
	Secret     config.EnvVariable // will have the value of env variable MY_SECRET
	Nested     struct {
		Value bool
	}
}

func main() {
	s := &serviceConfig{}

	// load config into predefined struct from the default profile app.json
	err := config.Load(s)
	if err != nil {
		panic(err)
	}
}
```