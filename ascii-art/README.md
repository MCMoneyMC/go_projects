```
  _____                        _     _                                   _______                                 _                
 / ____|                      | |   (_)                                 |__   __|                               | |               
| |  __   _ __    ___    ___  | |_   _   _ __     __ _   ___               | |     _ __    __ _  __   __   ___  | |   ___   _ __  
| | |_ | | '__|  / _ \  / _ \ | __| | | | '_ \   / _` | / __|              | |    | '__|  / _` | \ \ / /  / _ \ | |  / _ \ | '__| 
| |__| | | |    |  __/ |  __/ \ |_  | | | | | | | (_| | \__ \  _           | |    | |    | (_| |  \ V /  |  __/ | | |  __/ | |    
 \_____| |_|     \___|  \___|  \__| |_| |_| |_|  \__, | |___/ ( )          |_|    |_|     \__,_|   \_/    \___| |_|  \___| |_|    
                                                  __/ |       |/                                                                  
                                                 |___/                                                                            
```
## Task Description & Audit Prompts
https://github.com/01-edu/public/tree/master/subjects/ascii-art

## Usage
#### 1. Download and `cd` into the repository

#### 2. Use `go run . <text-to-asciify> <OPTIONAL-write-to-file>` or `cd src/` and run `go test` to try some pre-written unit-tests
- At least one argument is required, but up to 2 are accepted.
- If the second argument has been declared, the output will be written to `output.txt` in the repository's root folder.
- This has been added for ease of use, since you need at least 4000px of width to comfortably display panagrams and other longer strings from the terminal.

### Warning!
This has been designed and tested on Ubuntu, using it on other distros or operating systems may have different results.