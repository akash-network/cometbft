package node

import (
	"errors"
	"fmt"
)

var (
	// ErrNonEmptyBlockStore is returned when the blockstore is not empty and the node is trying to initialize non empty state.
	ErrNonEmptyBlockStore = errors.New("blockstore not empty, trying to initialize non empty state")
	// ErrNonEmptyState is returned when the state is not empty and the node is trying to initialize non empty state.
	ErrNonEmptyState = errors.New("state not empty, trying to initialize non empty state")
	// ErrSwitchStateSync is returned when the blocksync reactor does not support switching from state sync.
	ErrSwitchStateSync = errors.New("this blocksync reactor does not support switching from state sync")
	// ErrGenesisHashDecode is returned when the genesis hash provided by the operator cannot be decoded.
	ErrGenesisHashDecode = errors.New("genesis hash provided by operator cannot be decoded")
	// ErrPassedGenesisHashMismatch is returned when the genesis doc hash in the database does not match the passed --genesis_hash value.
	ErrPassedGenesisHashMismatch = errors.New("genesis doc hash in db does not match passed --genesis_hash value")
	// ErrLoadedGenesisDocHashMismatch is returned when the genesis doc hash in the database does not match the loaded genesis doc.
	ErrLoadedGenesisDocHashMismatch = errors.New("genesis doc hash in db does not match loaded genesis doc")
)

// ErrGenesisDoc is returned when the node fails to load the genesis doc.
type ErrGenesisDoc struct {
	Err error
}

func (e ErrGenesisDoc) Error() string {
	return fmt.Sprintf("error in genesis doc: %v", e.Err)
}

func (e ErrGenesisDoc) Unwrap() error {
	return e.Err
}

// ErrRetrieveGenesisDocHash is returned when the node fails to retrieve the genesis doc hash from the database.
type ErrRetrieveGenesisDocHash struct {
	Err error
}

func (e ErrRetrieveGenesisDocHash) Error() string {
	return fmt.Sprintf("error retrieving genesis doc hash: %v", e.Err)
}

func (e ErrRetrieveGenesisDocHash) Unwrap() error {
	return e.Err
}

// ErrSaveGenesisDocHash is returned when the node fails to save the genesis doc hash to the database.
type ErrSaveGenesisDocHash struct {
	Err error
}

func (e ErrSaveGenesisDocHash) Error() string {
	return fmt.Sprintf("failed to save genesis doc hash to db: %v", e.Err)
}

func (e ErrSaveGenesisDocHash) Unwrap() error {
	return e.Err
}
