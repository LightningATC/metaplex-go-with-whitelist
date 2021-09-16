// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Helpful method that isn't necessary to use for main users of the app, but allows one to create/update
// existing external price account fields if they are signers of this account.
// Useful for testing purposes, and the CLI makes use of it as well so that you can verify logic.
type UpdateExternalPriceAccount struct {
	Args *ExternalPriceAccount

	// [0] = [WRITE] externalPrice
	// ··········· External price account
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewUpdateExternalPriceAccountInstructionBuilder creates a new `UpdateExternalPriceAccount` instruction builder.
func NewUpdateExternalPriceAccountInstructionBuilder() *UpdateExternalPriceAccount {
	nd := &UpdateExternalPriceAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *UpdateExternalPriceAccount) SetArgs(args ExternalPriceAccount) *UpdateExternalPriceAccount {
	inst.Args = &args
	return inst
}

// SetExternalPriceAccount sets the "externalPrice" account.
// External price account
func (inst *UpdateExternalPriceAccount) SetExternalPriceAccount(externalPrice ag_solanago.PublicKey) *UpdateExternalPriceAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(externalPrice).WRITE()
	return inst
}

// GetExternalPriceAccount gets the "externalPrice" account.
// External price account
func (inst *UpdateExternalPriceAccount) GetExternalPriceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

func (inst UpdateExternalPriceAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_UpdateExternalPriceAccount),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateExternalPriceAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateExternalPriceAccount) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.ExternalPrice is not set")
		}
	}
	return nil
}

func (inst *UpdateExternalPriceAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateExternalPriceAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("externalPrice", inst.AccountMetaSlice[0]))
					})
				})
		})
}

func (obj UpdateExternalPriceAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateExternalPriceAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateExternalPriceAccountInstruction declares a new UpdateExternalPriceAccount instruction with the provided parameters and accounts.
func NewUpdateExternalPriceAccountInstruction(
	// Parameters:
	args ExternalPriceAccount,
	// Accounts:
	externalPrice ag_solanago.PublicKey) *UpdateExternalPriceAccount {
	return NewUpdateExternalPriceAccountInstructionBuilder().
		SetArgs(args).
		SetExternalPriceAccount(externalPrice)
}
