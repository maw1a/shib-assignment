// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Phantom is ERC721, Ownable {
    uint256 private _tokenIds;

    uint256 public constant MAX_SUPPLY = 100;
    mapping(address => bool) public hasMinted;

    constructor() ERC721("Phantom", "PHNTM") Ownable(msg.sender) {}

    function mint(address to) external {
        require(_tokenIds < MAX_SUPPLY, "MAX_SUPPLY_REACHED");
        require(!hasMinted[to], "WALLET_ALREADY_OWN");
        
        _tokenIds++;
        uint256 newTokenId = _tokenIds;
        
        hasMinted[to] = true;
        _safeMint(to, newTokenId);
    }

    function totalSupply() public view returns (uint256) {
        return _tokenIds;
    }

    function transferFrom(address from, address to, uint256 tokenId) public override {
        require(!hasMinted[to], "WALLET_ALREADY_OWN");

        super.transferFrom(from, to, tokenId);

        if (to != address(0)) {
            hasMinted[from] = false;
            hasMinted[to] = true;
        }
    }

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory _data) public override {
        require(!hasMinted[to], "WALLET_ALREADY_OWN");

        super.safeTransferFrom(from, to, tokenId, _data);

        if (to != address(0)) {
            hasMinted[from] = false;
            hasMinted[to] = true;
        }
    }
}
