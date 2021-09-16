// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Puff a Metadata - make all of it's variable length fields (name/uri/symbol) a fixed length using a null character
// so that it can be found using offset searches by the RPC to make client lookups cheaper.
type PuffMetadata struct {

	// [0] = [WRITE] metadataAccount
	// ··········· Metadata account
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewPuffMetadataInstructionBuilder creates a new `PuffMetadata` instruction builder.
func NewPuffMetadataInstructionBuilder() *PuffMetadata {
	nd := &PuffMetadata{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetMetadataAccount sets the "metadataAccount" account.
// Metadata account
func (inst *PuffMetadata) SetMetadataAccount(metadataAccount ag_solanago.PublicKey) *PuffMetadata {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadataAccount).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadataAccount" account.
// Metadata account
func (inst *PuffMetadata) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

func (inst PuffMetadata) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_PuffMetadata),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PuffMetadata) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PuffMetadata) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.MetadataAccount is not set")
		}
	}
	return nil
}

func (inst *PuffMetadata) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PuffMetadata")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("metadata", inst.AccountMetaSlice[0]))
					})
				})
		})
}

func (obj PuffMetadata) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *PuffMetadata) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewPuffMetadataInstruction declares a new PuffMetadata instruction with the provided parameters and accounts.
func NewPuffMetadataInstruction(
	// Accounts:
	metadataAccount ag_solanago.PublicKey) *PuffMetadata {
	return NewPuffMetadataInstructionBuilder().
		SetMetadataAccount(metadataAccount)
}
