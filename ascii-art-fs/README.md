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

## You can find the task's description and audit prompts here:
### https://github.com/01-edu/public/tree/master/subjects/ascii-art

## Usage
### 1. Download and cd into the repository

### 2. Use `go run . <text-to-asciify> <OPTIONAL-banner>` or run `autoTester.sh` to display the output for generated with prompts taken from the audit guide mentioned above.
- At least one argument is required, but up to 2 are accepted.
- The flag `<OPTIONAL-banner>` needs to have a value of `standard`, `shadow` or `thinkertoy`. This flag decides which "font" is used to display the given text.

## Warning!
This has been designed and tested on Ubuntu 20.04 LTS, using it on other distros or operating systems may have different results