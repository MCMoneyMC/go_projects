#!/bin/bash

go run . "banana" standard abc

go run . "hello" standard | cat -e

go run . "hello world" shadow | cat -e

go run . "nice 2 meet you" thinkertoy | cat -e

go run . "you & me" standard | cat -e

go run . "123" shadow | cat -e

go run . "/(\")" thinkertoy | cat -e

go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow | cat -e

go run . "\"#$%&/()*+,-./" thinkertoy | cat -e

go run . "It's Working" thinkertoy | cat -e

go run . "Single Argument!!!" | cat -e
