# Reproduction steps (Reproduced in this repo)

#
1. Install the Elys pkg using `go get github.com/elys-network/elys`
2. Git clone the elys repo with `git clone https://github.com/elys-network/elys`
3. Inside of your current workding dir inside of go.mod replace elys with the local one 
`github.com/elys-network/elys => ./elys`
   - You might have to do all three of these to fix any issues
   ```replace (
	cosmossdk.io/core => cosmossdk.io/core v0.11.1
	github.com/elys-network/elys => ./elys
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1 // Common cosmos override
    )```

4. `go mod tidy`
