package keeper

import (
	"testing"

	"github.com/BenWolfaardt/Checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	accountKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	accountTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func setupKeeper(t testing.TB) (*Keeper, sdk.Context) {
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	tkeys := sdk.NewTransientStoreKeys(paramsTypes.TStoreKey)
	blacklistedAddrs := make(map[string]bool)
	maccPerms := map[string][]string{
		accountTypes.FeeCollectorName: nil,
		types.ModuleName:              {accountTypes.Minter},
	}

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	paramsKeeper := paramsKeeper.NewKeeper(cdc, codec.NewLegacyAmino(), storeKey, tkeys[paramsTypes.TStoreKey])
	accountKeeper := accountKeeper.NewAccountKeeper(cdc, storeKey, paramsKeeper.Subspace("auth"), accountTypes.ProtoBaseAccount, maccPerms)
	bankKeeper := bankKeeper.NewBaseKeeper(cdc, storeKey, accountKeeper, paramsKeeper.Subspace("bank"), blacklistedAddrs)

	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	keeper := NewKeeper(
		bankKeeper, // boh, not sure
		cdc,
		storeKey,
		memStoreKey,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return keeper, ctx
}
