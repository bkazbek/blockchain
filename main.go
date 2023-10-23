package main

import (
	"blockchain/internal"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	blockchain := internal.NewBlockchain()

	app := fiber.New()

	app.Get("/mine_block", func(c *fiber.Ctx) error {
		previousBlock := blockchain.GetPreviousBlock()
		previousProof := previousBlock.Proof
		proof := blockchain.ProofOfWork(previousProof)

		block := blockchain.CreateBlock(proof, blockchain.Hash(previousBlock))
		return c.Status(fiber.StatusOK).JSON(block)
	})

	app.Get("/get_chain", func(c *fiber.Ctx) error {
		chain := blockchain.Chain
		return c.Status(fiber.StatusOK).JSON(chain)
	})

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("got err on listen:", err)
	}
}
