package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gin-gonic/gin"

	"github.com/maw1a/shib-backend/api"
	"github.com/maw1a/shib-backend/contract"
	"github.com/maw1a/shib-backend/utils"
)

func main() {
	client, err := ethclient.Dial(utils.ETH_ADDRESS)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress(utils.PHNTM_ADDRESS)

	ph, err := contract.NewPhantom(address, client)
	if err != nil {
		panic(err)
	}

	if err := getRouter(client, ph).Run(":8080"); err != nil {
		panic(err)
	}
}

func getRouter(client *ethclient.Client, ph *contract.Phantom) *gin.Engine {
	router := gin.Default()

	router.GET("/", api.Healthcheck)
	router.POST("/nft/mint", api.MintNFT(client, ph))
	router.GET("/nft", api.OwnedNFT(client, ph))

	return router
}
