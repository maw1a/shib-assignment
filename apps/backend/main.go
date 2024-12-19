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
	router.GET("/nft", CORSMiddleware(), api.OwnedNFT(client, ph))
	router.OPTIONS("/nft/mint", CORSMiddleware())
	router.POST("/nft/mint", CORSMiddleware(), api.MintNFT(client, ph))

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
