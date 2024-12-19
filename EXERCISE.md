## Exercise

We ask you to develop: Two Smart Contracts on the Sepolia Network

### Smart Contract 1 - Mint NFT:

- One NFT per wallet.
- No associated minting cost.
- Total supply: 100 NFTs.
- Transaction fees will be managed by the backend.

### Smart Contract 2 - Mint NFT with Additional Rules:

- Minting allowed only for holders of an NFT from the first smart contract.
- Maximum of 5 NFTs per transaction.
- Total supply: 40 NFTs.
- Minting cost:
- First 15 NFTs: 0.00000005 ETH each.
- Remaining NFTs: 0.0006 ETH each.
- Option to purchase a custom name for the NFT at an additional cost of 0.0004 ETH.

### Backend written in GoLang:

- Exposure of an API for minting NFTs from the first smart contract.
- Management of transaction fees through the backend.
- Listening to on-chain events from the second smart contract to update metadata on Pinata and the smart contract.

### Frontend for Minting and Viewing NFTs:

- Development of a frontend interface for interacting with the backend to mint NFTs.
- Display of purchased NFTs, including details such as name, features, and updated status based on customisations.

### Additional Requirements:

- Pinata Integration: For uploading new JSON files.
- Event Management: Interception and handling of events related to NFT name changes.
  Stack: GoLang, Tailwind, NextJS, Shadcn/ui
  Deadline:
  The completed code must be shared via a GitHub repository under account mentioned in Resume, (2d to develop)
  For simplified version we had 1day, I rewrote Francesco's full version to use GoLang and NextJS as our full-stack devs should be familiar with both.
