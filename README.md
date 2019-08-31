# findexec
A go library to find paths of executables

- [Usage](#usage)
- [License](#License)
- [Credits](#Credits)

## Usage


```go

package main

import "github.com/jonhadfield/findexec"

func main() {
    // find an executable called "diff" without specifying paths which will force
    // searching of the system paths found in environment variable 'PATH'
    _ = findexec.Find("diff", "")
    
    // find an executable called "bash" in specific paths
    _ = findexec.Find("bash", "/home/bin:/bin:/usr/local/bin")
}
```  


## License

The source code is made available under the terms of the Unlicense License, as stated in the file `LICENSE`.

## Credits

This is rewrite of the [find_executable()](https://docs.python.org/2/distutils/apiref.html#module-distutils.spawn) function provided in the python 2 standard library.

