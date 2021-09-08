// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeCandyMachine is the `initializeCandyMachine` instruction.
type InitializeCandyMachine struct {
	Bump *uint8
	Data *CandyMachineData

	// [0] = [WRITE] candyMachine
	//
	// [1] = [] wallet
	//
	// [2] = [] config
	//
	// [3] = [SIGNER] authority
	//
	// [4] = [WRITE, SIGNER] payer
	//
	// [5] = [] systemProgram
	//
	// [6] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitializeCandyMachineInstructionBuilder creates a new `InitializeCandyMachine` instruction builder.
func NewInitializeCandyMachineInstructionBuilder() *InitializeCandyMachine {
	nd := &InitializeCandyMachine{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetBump sets the "bump" parameter.
func (inst *InitializeCandyMachine) SetBump(bump uint8) *InitializeCandyMachine {
	inst.Bump = &bump
	return inst
}

// SetData sets the "data" parameter.
func (inst *InitializeCandyMachine) SetData(data CandyMachineData) *InitializeCandyMachine {
	inst.Data = &data
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *InitializeCandyMachine) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *InitializeCandyMachine) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetWalletAccount sets the "wallet" account.
func (inst *InitializeCandyMachine) SetWalletAccount(wallet ag_solanago.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(wallet)
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *InitializeCandyMachine) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetConfigAccount sets the "config" account.
func (inst *InitializeCandyMachine) SetConfigAccount(config ag_solanago.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *InitializeCandyMachine) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *InitializeCandyMachine) SetAuthorityAccount(authority ag_solanago.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *InitializeCandyMachine) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetPayerAccount sets the "payer" account.
func (inst *InitializeCandyMachine) SetPayerAccount(payer ag_solanago.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *InitializeCandyMachine) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeCandyMachine) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeCandyMachine) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetRentAccount sets the "rent" account.
func (inst *InitializeCandyMachine) SetRentAccount(rent ag_solanago.PublicKey) *InitializeCandyMachine {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitializeCandyMachine) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

func (inst InitializeCandyMachine) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeCandyMachine,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeCandyMachine) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeCandyMachine) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Bump == nil {
			return errors.New("Bump parameter is not set")
		}
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *InitializeCandyMachine) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeCandyMachine")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Bump", *inst.Bump))
						paramsBranch.Child(ag_format.Param("Data", *inst.Data))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" candyMachine", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("       wallet", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("        payer", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("         rent", inst.AccountMetaSlice[6]))
					})
				})
		})
}

func (obj InitializeCandyMachine) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeCandyMachine) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeCandyMachineInstruction declares a new InitializeCandyMachine instruction with the provided parameters and accounts.
func NewInitializeCandyMachineInstruction(
	// Parameters:
	bump uint8,
	data CandyMachineData,
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	wallet ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *InitializeCandyMachine {
	return NewInitializeCandyMachineInstructionBuilder().
		SetBump(bump).
		SetData(data).
		SetCandyMachineAccount(candyMachine).
		SetWalletAccount(wallet).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
