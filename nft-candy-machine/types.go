// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import ag_solanago "github.com/gagliardetto/solana-go"

type CandyMachineData struct {
	Uuid           string
	Price          uint64
	ItemsAvailable uint64
	GoLiveDate     *int64 `bin:"optional"`
}

type ConfigData struct {
	Uuid                 string
	Symbol               string
	SellerFeeBasisPoints uint16
	Creators             []Creator
	MaxSupply            uint64
	IsMutable            bool
	RetainAuthority      bool
	MaxNumberOfLines     uint32
}

type ConfigLine struct {
	Name string
	Uri  string
}

type Creator struct {
	Address  ag_solanago.PublicKey
	Verified bool
	Share    uint8
}
