package api

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gin-gonic/gin"

	"github.com/maw1a/shib-backend/contract"
	"github.com/maw1a/shib-backend/utils"
	"github.com/maw1a/shib-backend/utils/logger"
)

type MintRequestBody struct {
	To string `json:"to"`
}

type OwnedRequestBody struct {
	By string `json:"by"`
}

func MintNFT(client *ethclient.Client, ph *contract.Phantom) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := &MintRequestBody{}

		if err := ctx.BindJSON(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}

		signer, err := utils.GetSigner(client, utils.WALLET_PRIVATE_KEY)
		if err != nil {
			logger.Debug("❗️ Failed to get Signer ❗️ : with error: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to mint NFT", "err": err.Error()})
			return
		}

		t, err := ph.Mint(signer, common.HexToAddress(body.To))
		if err != nil {
			logger.Debug("❗️ Failed to mint NFT ❗️ : with error: %v", err.Error())

			msg := "❗️ failed to mint NFT"
			status := http.StatusInternalServerError

			if strings.Contains(err.Error(), "MAX_SUPPLY_REACHED") {
				msg = "🚨 maximum number of tokens for this contract have been minted"
				status = http.StatusForbidden
			} else if strings.Contains(err.Error(), "WALLET_ALREADY_OWN") {
				msg = "🧩 wallet already owns token from this contract"
				status = http.StatusForbidden
			}

			ctx.JSON(status, gin.H{"message": msg})
			return
		}

		fmt.Printf("🎉 NFT minted for wallet %s 🎉 : with hash: %s & cost: %d\n", body.To, t.Hash().String(), t.Cost())

		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("🎉 NFT minted for wallet %s 🎉", body.To)})
	}
}

func OwnedNFT(client *ethclient.Client, ph *contract.Phantom) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Owner := ctx.Query("owner")

		if len(Owner) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "no owner provided"})
			return
		}

		signer, err := utils.GetSigner(client, utils.WALLET_PRIVATE_KEY)
		if err != nil {
			logger.Debug("❗️ Failed to get Signer ❗️ : with error: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to mint NFT", "err": err.Error()})
			return
		}

		callOpts := bind.CallOpts{From: signer.From, Context: context.Background()}
		tokenId := int64(0)
		failed := []int64{}

		for tokenId < 100 {
			if address, err := ph.OwnerOf(&callOpts, big.NewInt(tokenId)); err != nil {
				logger.Error("failed to get owner of token ID: %d", tokenId)
				failed = append(failed, tokenId)
			} else if address.Hex() == Owner {
				break
			}

			tokenId++
		}

		symbol, _ := ph.Symbol(&callOpts)

		if tokenId < 100 {
			logger.Info("👻 PHNTM#%d is owned by address %s", tokenId, Owner)

			ctx.JSON(http.StatusOK, gin.H{
				"message":  fmt.Sprintf("👻 PHNTM#%d is owned by address %s", tokenId, Owner),
				"symbol":   symbol,
				"token_id": tokenId,
				"owner":    Owner,
			})
			return
		}

		logger.Info("❌ no tokens found for the address %s", Owner)
		ctx.JSON(http.StatusNotFound, gin.H{
			"message":    fmt.Sprintf("❌ no tokens found for the address %s", Owner),
			"symbol":     symbol,
			"failed_ids": failed,
		})
	}
}
