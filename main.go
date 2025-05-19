package main

import (
	"context"
	"fmt"
	"log"

	cosmosMath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types"
	amm "github.com/elys-network/elys/x/amm/types"
	"google.golang.org/grpc"
)

func main() {
	grpcAddress := "elys-grpc.polkachu.com:22090"

	// Connect to gRPC server
	grpcClient, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcClient.Close()

	tokenInAmount := cosmosMath.NewInt(1000000)
	tokenInDenom := "uelys"
	tokenOutDenom := "ibc/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9"

	ammClient := amm.NewQueryClient(grpcClient)

	queryPool, err := ammClient.SwapEstimationByDenom(context.Background(), &amm.QuerySwapEstimationByDenomRequest{
		DenomIn:  tokenInDenom,
		DenomOut: tokenOutDenom,
		Amount: types.Coin{
			Denom:  tokenInDenom,
			Amount: tokenInAmount,
		},
		Address: "elys1qyqszqgpqyqszqgpqyqszqgpqyqszqgpqyqszqgp",
	})
	if err != nil {
		log.Fatalf("Error fetching pool for swap estimation: %v", err)
	}

	fmt.Printf("Token Out Amount: %s\n", queryPool.Amount.Amount.String())
	fmt.Printf("Slippage: %f\n", queryPool.Slippage.MustFloat64())
}
