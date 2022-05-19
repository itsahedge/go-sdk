package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

var (
	nftDropContractAddress string
)

var nftDropCmd = &cobra.Command{
	Use:   "nftdrop [command]",
	Short: "Interact with an NFT Drop contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var nftDropGetAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "Get all available nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		nftDrop, err := getNftDrop()
		if err != nil {
			panic(err)
		}

		allNfts, err := nftDrop.GetAll()
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got drop nft with name '%v' and description '%v' and id '%d'\n", nft.Metadata.Name, nft.Metadata.Description, nft.Metadata.Id)
		}
	},
}

var nftDropClaimCmd = &cobra.Command{
	Use:   "claim",
	Short: "Claim an nft",
	Run: func(cmd *cobra.Command, args []string) {
		nftDrop, err := getNftDrop()
		if err != nil {
			panic(err)
		}

		if tx, err := nftDrop.Claim(1); err != nil {
			panic(err)
		} else {
			log.Printf("Claimed nft successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

var nftDropCreateBatchCmd = &cobra.Command{
	Use:   "createBatch",
	Short: "Create a batch of nfts",
	Run: func(cmd *cobra.Command, args []string) {
		nftDrop, err := getNftDrop()
		if err != nil {
			panic(err)
		}

		if tx, err := nftDrop.CreateBatch(
			[]*thirdweb.NFTMetadataInput{
				{
					Name: "Drop NFT 1",
				},
				{
					Name: "Drop NFT 2",
				},
				{
					Name: "Drop NFT 3",
				},
			},
		); err != nil {
			panic(err)
		} else {
			log.Printf("Created batch of nfts successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

func init() {
	nftDropCmd.PersistentFlags().StringVarP(&nftDropContractAddress, "address", "a", "", "nft drop contract address")
	nftDropCmd.AddCommand(nftDropGetAllCmd)
	nftDropCmd.AddCommand(nftDropClaimCmd)
	nftDropCmd.AddCommand(nftDropCreateBatchCmd)
}
