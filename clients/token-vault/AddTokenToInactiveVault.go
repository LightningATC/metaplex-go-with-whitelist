// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Add a token to a inactive token vault
type AddTokenToInactiveVault struct {
	Args *AmountArgs

	// [0] = [WRITE] uninitializedSafetyDepositBoxAccount
	// ··········· Uninitialized safety deposit box account address (will be created and allocated by this endpoint)
	// ··········· Address should be pda with seed of [PREFIX, vault_address, token_mint_address]
	//
	// [1] = [WRITE] initializedToken
	// ··········· Initialized Token account
	//
	// [2] = [WRITE] initializedTokenStore
	// ··········· Initialized Token store account with authority of this program, this will get set on the safety deposit box
	//
	// [3] = [WRITE] initializedInactiveFractionalizedTokenVault
	// ··········· Initialized inactive fractionalized token vault
	//
	// [4] = [SIGNER] vaultAuthority
	// ··········· Authority on the vault
	//
	// [5] = [SIGNER] payer
	// ··········· Payer
	//
	// [6] = [SIGNER] transferAuthority
	// ··········· Transfer Authority to move desired token amount from token account to safety deposit
	//
	// [7] = [] tokenProgram
	// ··········· Token program
	//
	// [8] = [] rentSysvar
	// ··········· Rent sysvar
	//
	// [9] = [] systemAccountSysvar
	// ··········· System account sysvar
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewAddTokenToInactiveVaultInstructionBuilder creates a new `AddTokenToInactiveVault` instruction builder.
func NewAddTokenToInactiveVaultInstructionBuilder() *AddTokenToInactiveVault {
	nd := &AddTokenToInactiveVault{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *AddTokenToInactiveVault) SetArgs(args AmountArgs) *AddTokenToInactiveVault {
	inst.Args = &args
	return inst
}

// SetUninitializedSafetyDepositBoxAccount sets the "uninitializedSafetyDepositBoxAccount" account.
// Uninitialized safety deposit box account address (will be created and allocated by this endpoint)
// Address should be pda with seed of [PREFIX, vault_address, token_mint_address]
func (inst *AddTokenToInactiveVault) SetUninitializedSafetyDepositBoxAccount(uninitializedSafetyDepositBoxAccount ag_solanago.PublicKey) *AddTokenToInactiveVault {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(uninitializedSafetyDepositBoxAccount).WRITE()
	return inst
}

// GetUninitializedSafetyDepositBoxAccount gets the "uninitializedSafetyDepositBoxAccount" account.
// Uninitialized safety deposit box account address (will be created and allocated by this endpoint)
// Address should be pda with seed of [PREFIX, vault_address, token_mint_address]
func (inst *AddTokenToInactiveVault) GetUninitializedSafetyDepositBoxAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetInitializedTokenAccount sets the "initializedToken" account.
// Initialized Token account
func (inst *AddTokenToInactiveVault) SetInitializedTokenAccount(initializedToken ag_solanago.PublicKey) *AddTokenToInactiveVault {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(initializedToken).WRITE()
	return inst
}

// GetInitializedTokenAccount gets the "initializedToken" account.
// Initialized Token account
func (inst *AddTokenToInactiveVault) GetInitializedTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetInitializedTokenStoreAccount sets the "initializedTokenStore" account.
// Initialized Token store account with authority of this program, this will get set on the safety deposit box
func (inst *AddTokenToInactiveVault) SetInitializedTokenStoreAccount(initializedTokenStore ag_solanago.PublicKey) *AddTokenToInactiveVault {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(initializedTokenStore).WRITE()
	return inst
}

// GetInitializedTokenStoreAccount gets the "initializedTokenStore" account.
// Initialized Token store account with authority of this program, this will get set on the safety deposit box
func (inst *AddTokenToInactiveVault) GetInitializedTokenStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetInitializedInactiveFractionalizedTokenVaultAccount sets the "initializedInactiveFractionalizedTokenVault" account.
// Initialized inactive fractionalized token vault
func (inst *AddTokenToInactiveVault) SetInitializedInactiveFractionalizedTokenVaultAccount(initializedInactiveFractionalizedTokenVault ag_solanago.PublicKey) *AddTokenToInactiveVault {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(initializedInactiveFractionalizedTokenVault).WRITE()
	return inst
}

// GetInitializedInactiveFractionalizedTokenVaultAccount gets the "initializedInactiveFractionalizedTokenVault" account.
// Initialized inactive fractionalized token vault
func (inst *AddTokenToInactiveVault) GetInitializedInactiveFractionalizedTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetVaultAuthorityAccount sets the "vaultAuthority" account.
// Authority on the vault
func (inst *AddTokenToInactiveVault) SetVaultAuthorityAccount(vaultAuthority ag_solanago.PublicKey) *AddTokenToInactiveVault {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(vaultAuthority).SIGNER()
	return inst
}

// GetVaultAuthorityAccount gets the "vaultAuthority" account.
// Authority on the vault
func (inst *AddTokenToInactiveVault) GetVaultAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *AddTokenToInactiveVault) SetPayerAccount(payer ag_solanago.PublicKey) *AddTokenToInactiveVault {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *AddTokenToInactiveVault) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetTransferAuthorityAccount sets the "transferAuthority" account.
// Transfer Authority to move desired token amount from token account to safety deposit
func (inst *AddTokenToInactiveVault) SetTransferAuthorityAccount(transferAuthority ag_solanago.PublicKey) *AddTokenToInactiveVault {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(transferAuthority).SIGNER()
	return inst
}

// GetTransferAuthorityAccount gets the "transferAuthority" account.
// Transfer Authority to move desired token amount from token account to safety deposit
func (inst *AddTokenToInactiveVault) GetTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *AddTokenToInactiveVault) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *AddTokenToInactiveVault {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *AddTokenToInactiveVault) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *AddTokenToInactiveVault) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *AddTokenToInactiveVault {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *AddTokenToInactiveVault) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetSystemAccountSysvarAccount sets the "systemAccountSysvar" account.
// System account sysvar
func (inst *AddTokenToInactiveVault) SetSystemAccountSysvarAccount(systemAccountSysvar ag_solanago.PublicKey) *AddTokenToInactiveVault {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemAccountSysvar)
	return inst
}

// GetSystemAccountSysvarAccount gets the "systemAccountSysvar" account.
// System account sysvar
func (inst *AddTokenToInactiveVault) GetSystemAccountSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

func (inst AddTokenToInactiveVault) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_AddTokenToInactiveVault),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddTokenToInactiveVault) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddTokenToInactiveVault) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.UninitializedSafetyDepositBoxAccount is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.InitializedToken is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.InitializedTokenStore is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.InitializedInactiveFractionalizedTokenVault is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.VaultAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TransferAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemAccountSysvar is not set")
		}
	}
	return nil
}

func (inst *AddTokenToInactiveVault) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddTokenToInactiveVault")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("              uninitializedSafetyDepositBox", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("                           initializedToken", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("                      initializedTokenStore", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("initializedInactiveFractionalizedTokenVault", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("                             vaultAuthority", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("                                      payer", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("                          transferAuthority", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("                               tokenProgram", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("                                 rentSysvar", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("                        systemAccountSysvar", inst.AccountMetaSlice[9]))
					})
				})
		})
}

func (obj AddTokenToInactiveVault) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddTokenToInactiveVault) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewAddTokenToInactiveVaultInstruction declares a new AddTokenToInactiveVault instruction with the provided parameters and accounts.
func NewAddTokenToInactiveVaultInstruction(
	// Parameters:
	args AmountArgs,
	// Accounts:
	uninitializedSafetyDepositBoxAccount ag_solanago.PublicKey,
	initializedToken ag_solanago.PublicKey,
	initializedTokenStore ag_solanago.PublicKey,
	initializedInactiveFractionalizedTokenVault ag_solanago.PublicKey,
	vaultAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	transferAuthority ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey,
	systemAccountSysvar ag_solanago.PublicKey) *AddTokenToInactiveVault {
	return NewAddTokenToInactiveVaultInstructionBuilder().
		SetArgs(args).
		SetUninitializedSafetyDepositBoxAccount(uninitializedSafetyDepositBoxAccount).
		SetInitializedTokenAccount(initializedToken).
		SetInitializedTokenStoreAccount(initializedTokenStore).
		SetInitializedInactiveFractionalizedTokenVaultAccount(initializedInactiveFractionalizedTokenVault).
		SetVaultAuthorityAccount(vaultAuthority).
		SetPayerAccount(payer).
		SetTransferAuthorityAccount(transferAuthority).
		SetTokenProgramAccount(tokenProgram).
		SetRentSysvarAccount(rentSysvar).
		SetSystemAccountSysvarAccount(systemAccountSysvar)
}
