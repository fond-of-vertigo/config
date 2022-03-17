# Config library

This library loads a json file and marshals it into a given `struct`.

By default, it loads `app.json` from sub-folder config. If you define a profile then this profile is added to file name.
E.g. staging: `app-staging.json`
You can set the profile via the environment variable `CONFIG_PROFILE` or load a specific profile with the function 
`LoadFromWithProfile`.

**DO NOT ENTER SECRETS INTO JSON FILE** Use environment variables instead.

## Example

```go
package myPackage

import "github.com/fond-of/config.go/config"

// serviceConfig defines a configuration struct
type serviceConfig struct { 
	AttributeA string
	AttributeB int
}

func main() {
	s := &serviceConfig{}
	
	// load config into predefined struct from the default profile app.json
	err := config.Load(s) 
	if err != nil {
		panic(err)
	}

	// fetch environment variables for secrets
	apiPassword := config.MustGetEnv("API_PASSWORD") 
}
```