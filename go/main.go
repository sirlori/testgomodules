package main

import (
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    // UNIX Time is faster and smaller than most timestamps
    // If you set zerolog.TimeFieldFormat to an empty string,
    // logs will write with UNIX time
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    log.Print("hello world")
}

// Output: {"time":1516134303,"level":"debug","message":"hello world"}
