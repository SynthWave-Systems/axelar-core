package app

import (
	"encoding/json"
	"fmt"
	"io"
	stdlog "log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authAnte "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v4/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v4/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v4/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibcante "github.com/cosmos/ibc-go/v4/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
	"github.com/gorilla/mux"
	ibchooks "github.com/osmosis-labs/osmosis/x/ibc-hooks"
	ibchookskeeper "github.com/osmosis-labs/osmosis/x/ibc-hooks/keeper"
	ibchookstypes "github.com/osmosis-labs/osmosis/x/ibc-hooks/types"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"golang.org/x/mod/semver"

	axelarParams "github.com/axelarnetwork/axelar-core/app/params"
	"github.com/axelarnetwork/axelar-core/x/ante"
	"github.com/axelarnetwork/axelar-core/x/axelarnet"
	axelarnetclient "github.com/axelarnetwork/axelar-core/x/axelarnet/client"
	axelarnetKeeper "github.com/axelarnetwork/axelar-core/x/axelarnet/keeper"
	axelarnetTypes "github.com/axelarnetwork/axelar-core/x/axelarnet/types"
	axelarbankkeeper "github.com/axelarnetwork/axelar-core/x/bank/keeper"
	"github.com/axelarnetwork/axelar-core/x/evm"
	evmKeeper "github.com/axelarnetwork/axelar-core/x/evm/keeper"
	evmTypes "github.com/axelarnetwork/axelar-core/x/evm/types"
	"github.com/axelarnetwork/axelar-core/x/multisig"
	multisigKeeper "github.com/axelarnetwork/axelar-core/x/multisig/keeper"
	multisigTypes "github.com/axelarnetwork/axelar-core/x/multisig/types"
	"github.com/axelarnetwork/axelar-core/x/nexus"
	nexusKeeper "github.com/axelarnetwork/axelar-core/x/nexus/keeper"
	nexusTypes "github.com/axelarnetwork/axelar-core/x/nexus/types"
	"github.com/axelarnetwork/axelar-core/x/permission"
	permissionKeeper "github.com/axelarnetwork/axelar-core/x/permission/keeper"
	permissionTypes "github.com/axelarnetwork/axelar-core/x/permission/types"
	"github.com/axelarnetwork/axelar-core/x/reward"
	rewardKeeper "github.com/axelarnetwork/axelar-core/x/reward/keeper"
	rewardTypes "github.com/axelarnetwork/axelar-core/x/reward/types"
	"github.com/axelarnetwork/axelar-core/x/snapshot"
	snapKeeper "github.com/axelarnetwork/axelar-core/x/snapshot/keeper"
	snapTypes "github.com/axelarnetwork/axelar-core/x/snapshot/types"
	"github.com/axelarnetwork/axelar-core/x/tss"
	tssKeeper "github.com/axelarnetwork/axelar-core/x/tss/keeper"
	tssTypes "github.com/axelarnetwork/axelar-core/x/tss/types"
	"github.com/axelarnetwork/axelar-core/x/vote"
	voteKeeper "github.com/axelarnetwork/axelar-core/x/vote/keeper"
	voteTypes "github.com/axelarnetwork/axelar-core/x/vote/types"
	"github.com/axelarnetwork/utils/maps"

	// Override with generated statik docs
	_ "github.com/axelarnetwork/axelar-core/client/docs/statik"
)

// Name is the name of the application
const Name = "axelar"

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification. It is dynamically initialized by GetModuleBasics method.
	ModuleBasics module.BasicManager

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		axelarnetTypes.ModuleName:      {authtypes.Minter, authtypes.Burner},
		rewardTypes.ModuleName:         {authtypes.Minter},
		wasm.ModuleName:                {authtypes.Burner},
	}

	// WasmEnabled indicates whether wasm module is added to the app.
	// "true" setting means it will be, otherwise it won't.
	// This is configured during the build.
	WasmEnabled = ""

	// WasmCapabilities specifies the capabilities of the wasm vm
	// capabilities are detailed here: https://github.com/CosmWasm/cosmwasm/blob/main/docs/CAPABILITIES-BUILT-IN.md
	WasmCapabilities = ""
)

var (
	_ servertypes.Application = (*AxelarApp)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		stdlog.Println("Failed to get home dir %2", err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// AxelarApp defines the axelar Cosmos app that runs all modules
type AxelarApp struct {
	*bam.BaseApp

	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	// necessery keepers for export
	stakingKeeper  stakingkeeper.Keeper
	crisisKeeper   crisiskeeper.Keeper
	distrKeeper    distrkeeper.Keeper
	slashingKeeper slashingkeeper.Keeper

	// keys to access the substores
	keys map[string]*sdk.KVStoreKey

	mm            *module.Manager
	upgradeKeeper upgradekeeper.Keeper
}

// NewAxelarApp is a constructor function for axelar
func NewAxelarApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	homePath string, invCheckPeriod uint, encodingConfig axelarParams.EncodingConfig,
	appOpts servertypes.AppOptions, wasmOpts []wasm.Option, baseAppOptions ...func(*bam.BaseApp)) *AxelarApp {

	appCodec := encodingConfig.Codec
	interfaceRegistry := encodingConfig.InterfaceRegistry

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := createStoreKeys()
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	paramsK := initParamsKeeper(appCodec, encodingConfig.Amino, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])
	// set the BaseApp's parameter store
	bApp.SetParamStore(getSubspace(paramsK, bam.Paramspace))

	// add keepers
	accountK := authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], getSubspace(paramsK, authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	bankK := bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], accountK, getSubspace(paramsK, banktypes.ModuleName),
		maps.Filter(moduleAccountAddrs(), func(addr string, _ bool) bool {
			// we do not rely on internal balance tracking for invariance checks in the axelarnet module
			// (https://github.com/cosmos/cosmos-sdk/issues/12825 for more details on the purpose of the blocked list),
			// but the module address must be able to use ibc transfers,
			// so we exclude this address from the blocked list
			return addr != authtypes.NewModuleAddress(axelarnetTypes.ModuleName).String()
		}),
	)
	stakingK := stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], accountK, bankK, getSubspace(paramsK, stakingtypes.ModuleName),
	)

	mintK := mintkeeper.NewKeeper(
		appCodec, keys[minttypes.StoreKey], getSubspace(paramsK, minttypes.ModuleName), &stakingK,
		accountK, bankK, authtypes.FeeCollectorName,
	)
	distrK := distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], getSubspace(paramsK, distrtypes.ModuleName), accountK, bankK,
		&stakingK, authtypes.FeeCollectorName, moduleAccountAddrs(),
	)
	slashingK := slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &stakingK, getSubspace(paramsK, slashingtypes.ModuleName),
	)
	crisisK := crisiskeeper.NewKeeper(
		getSubspace(paramsK, crisistypes.ModuleName), invCheckPeriod, bankK, authtypes.FeeCollectorName,
	)
	upgradeK := upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, bApp)

	evidenceK := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &stakingK, slashingK,
	)

	feegrantK := feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], accountK)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	stakingK = *stakingK.SetHooks(
		stakingtypes.NewMultiStakingHooks(distrK.Hooks(), slashingK.Hooks()),
	)

	// add capability keeper and ScopeToModule for ibc module
	capabilityK := capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCK := capabilityK.ScopeToModule(ibchost.ModuleName)
	scopedTransferK := capabilityK.ScopeToModule(ibctransfertypes.ModuleName)

	// Create IBC Keeper
	ibcKeeper := ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], getSubspace(paramsK, ibchost.ModuleName), stakingK, upgradeK, scopedIBCK,
	)

	// Custom axelarnet/evm/nexus keepers
	axelarnetK := axelarnetKeeper.NewKeeper(
		appCodec, keys[axelarnetTypes.StoreKey], getSubspace(paramsK, axelarnetTypes.ModuleName), ibcKeeper.ChannelKeeper, feegrantK,
	)

	evmK := evmKeeper.NewKeeper(
		appCodec, keys[evmTypes.StoreKey], paramsK,
	)

	nexusK := nexusKeeper.NewKeeper(
		appCodec, keys[nexusTypes.StoreKey], getSubspace(paramsK, nexusTypes.ModuleName),
	)

	// Setting Router will finalize all routes by sealing router
	// No more routes can be added
	nexusRouter := nexusTypes.NewRouter()
	nexusRouter.AddAddressValidator(evmTypes.ModuleName, evmKeeper.NewAddressValidator()).
		AddAddressValidator(axelarnetTypes.ModuleName, axelarnetKeeper.NewAddressValidator(axelarnetK))
	nexusK.SetRouter(nexusRouter)

	// IBC Transfer Stack: SendPacket
	//
	// Originates from the transferKeeper and goes up the stack
	// transferKeeper.SendPacket -> ibc_hooks.SendPacket -> rateLimiter.SendPacket -> channel.SendPacket
	//
	// After this, the wasm keeper is required to be set on WasmHooks

	// Create IBC rate limiter
	rateLimiter := axelarnet.NewRateLimiter(axelarnetK, ibcKeeper.ChannelKeeper, nexusK)
	var ibcHooksMiddleware ibchooks.ICS4Middleware
	var ics4Wrapper ibctransfertypes.ICS4Wrapper
	var wasmHooks ibchooks.WasmHooks
	ics4Wrapper = rateLimiter

	if IsWasmEnabled() {
		// Configure the IBC hooks keeper to make wasm calls via IBC transfer memo
		ibcHooksKeeper := ibchookskeeper.NewKeeper(
			keys[ibchookstypes.StoreKey],
		)
		accPrefix := sdk.GetConfig().GetBech32AccountAddrPrefix()
		wasmHooks = ibchooks.NewWasmHooks(&ibcHooksKeeper, nil, accPrefix) // The contract keeper needs to be set later
		ibcHooksMiddleware = ibchooks.NewICS4Middleware(
			rateLimiter, // Wrap IBC hooks on top of the rate limit middleware
			&wasmHooks,
		)

		ics4Wrapper = ibcHooksMiddleware
	}

	// Create Transfer Keepers
	transferKeeper := ibctransferkeeper.NewKeeper(
		appCodec, keys[ibctransfertypes.StoreKey], getSubspace(paramsK, ibctransfertypes.ModuleName),
		// Use the IBC middleware stack
		ics4Wrapper,
		ibcKeeper.ChannelKeeper, &ibcKeeper.PortKeeper,
		accountK, bankK, scopedTransferK,
	)

	// IBC Transfer Stack: RecvPacket
	//
	// Packet originates from core IBC and goes down to app, the flow is the other way
	// channel.RecvPacket -> axelarnet.OnRecvPacket (transfer, GMP, and rate limit handler) -> ibc_hooks.OnRecvPacket -> transfer.OnRecvPacket
	var transferStack porttypes.IBCModule = transfer.NewIBCModule(transferKeeper)
	if IsWasmEnabled() {
		transferStack = ibchooks.NewIBCMiddleware(transferStack, &ibcHooksMiddleware)
	}

	ibcK := axelarnetKeeper.NewIBCKeeper(axelarnetK, transferKeeper, ibcKeeper.ChannelKeeper)
	axelarnetModule := axelarnet.NewAppModule(axelarnetK, nexusK, axelarbankkeeper.NewBankKeeper(bankK), accountK, ibcK, transferStack, rateLimiter, logger)

	// Create static IBC router, add axelarnet module as the IBC transfer route, and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, axelarnetModule)

	// axelar custom keepers
	// axelarnet / evm / nexus keepers created above
	rewardK := rewardKeeper.NewKeeper(
		appCodec, keys[rewardTypes.StoreKey], getSubspace(paramsK, rewardTypes.ModuleName), axelarbankkeeper.NewBankKeeper(bankK), distrK, stakingK,
	)
	multisigK := multisigKeeper.NewKeeper(
		appCodec, keys[multisigTypes.StoreKey], getSubspace(paramsK, multisigTypes.ModuleName),
	)

	multisigRouter := multisigTypes.NewSigRouter()
	multisigRouter.AddHandler(evmTypes.ModuleName, evmKeeper.NewSigHandler(appCodec, evmK))
	multisigK.SetSigRouter(multisigRouter)

	tssK := tssKeeper.NewKeeper(
		appCodec, keys[tssTypes.StoreKey], getSubspace(paramsK, tssTypes.ModuleName),
	)
	snapK := snapKeeper.NewKeeper(
		appCodec, keys[snapTypes.StoreKey], getSubspace(paramsK, snapTypes.ModuleName), stakingK, axelarbankkeeper.NewBankKeeper(bankK),
		slashingK,
	)
	votingK := voteKeeper.NewKeeper(
		appCodec, keys[voteTypes.StoreKey], getSubspace(paramsK, voteTypes.ModuleName), snapK, stakingK, rewardK,
	)
	permissionK := permissionKeeper.NewKeeper(
		appCodec, keys[permissionTypes.StoreKey], getSubspace(paramsK, permissionTypes.ModuleName),
	)

	var wasmK wasm.Keeper
	var wasmAnteDecorators []sdk.AnteDecorator

	if IsWasmEnabled() {
		wasmDir := filepath.Join(homePath, "wasm")
		wasmConfig, err := wasm.ReadWasmConfig(appOpts)
		if err != nil {
			panic(fmt.Sprintf("error while reading wasm config: %s", err))
		}

		scopedWasmK := capabilityK.ScopeToModule(wasm.ModuleName)
		// The last arguments can contain custom message handlers, and custom query handlers,
		// if we want to allow any custom callbacks
		wasmOpts = append(wasmOpts, wasmkeeper.WithMessageHandlerDecorator(func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
			return wasmkeeper.NewMessageHandlerChain(old, nexusKeeper.NewMessenger(nexusK))
		}))
		wasmK = wasm.NewKeeper(
			appCodec,
			keys[wasm.StoreKey],
			getSubspace(paramsK, wasm.ModuleName),
			accountK,
			bankK,
			stakingK,
			distrK,
			ibcKeeper.ChannelKeeper,
			ibcKeeper.ChannelKeeper,
			&ibcKeeper.PortKeeper,
			scopedWasmK,
			transferKeeper,
			bApp.MsgServiceRouter(),
			bApp.GRPCQueryRouter(),
			wasmDir,
			wasmConfig,
			WasmCapabilities,
			wasmOpts...,
		)

		wasmAnteDecorators = []sdk.AnteDecorator{
			wasmkeeper.NewLimitSimulationGasDecorator(wasmConfig.SimulationGasLimit),
			wasmkeeper.NewCountTXDecorator(keys[wasm.StoreKey]),
		}

		// Create wasm ibc stack
		var wasmStack porttypes.IBCModule = wasm.NewIBCHandler(wasmK, ibcKeeper.ChannelKeeper, ibcKeeper.ChannelKeeper)
		ibcRouter.AddRoute(wasm.ModuleName, wasmStack)

		// set the contract keeper for the Ics20WasmHooks
		wasmHooks.ContractKeeper = wasmkeeper.NewDefaultPermissionKeeper(wasmK)
	}

	// Finalize the IBC router
	ibcKeeper.SetRouter(ibcRouter)

	// Add governance proposal hooks
	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(paramsK)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(distrK)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(upgradeK)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(ibcKeeper.ClientKeeper)).
		AddRoute(axelarnetTypes.RouterKey, axelarnet.NewProposalHandler(axelarnetK, nexusK, accountK))

	if IsWasmEnabled() {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(wasmK, wasm.EnableAllProposals))
	}

	govK := govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], getSubspace(paramsK, govtypes.ModuleName), accountK, bankK,
		&stakingK, govRouter,
	)
	govK.SetHooks(govtypes.NewMultiGovHooks(axelarnetK.Hooks(nexusK, govK)))

	semverVersion := bApp.Version()
	if !strings.HasPrefix(semverVersion, "v") {
		semverVersion = fmt.Sprintf("v%s", semverVersion)
	}

	upgradeName := semver.MajorMinor(semverVersion)
	if upgradeName == "" {
		panic(fmt.Errorf("invalid app version %s", bApp.Version()))
	}

	// todo: change order of commands so this doesn't have to be defined before initialization
	var configurator module.Configurator
	var mm *module.Manager
	upgradeK.SetUpgradeHandler(
		upgradeName,
		func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			return mm.RunMigrations(ctx, configurator, fromVM)
		},
	)

	upgradeInfo, err := upgradeK.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == upgradeName && !upgradeK.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := store.StoreUpgrades{}

		if IsWasmEnabled() {
			storeUpgrades.Added = append(storeUpgrades.Added, ibchookstypes.StoreKey)
			storeUpgrades.Added = append(storeUpgrades.Added, wasm.ModuleName)
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		bApp.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

	voteRouter := voteTypes.NewRouter()
	voteRouter.AddHandler(evmTypes.ModuleName, evmKeeper.NewVoteHandler(appCodec, evmK, nexusK, rewardK))
	votingK.SetVoteRouter(voteRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	appModules := []module.AppModule{
		genutil.NewAppModule(accountK, stakingK, bApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(appCodec, accountK, nil),
		vesting.NewAppModule(accountK, bankK),
		bank.NewAppModule(appCodec, bankK, accountK),
		crisis.NewAppModule(&crisisK, skipGenesisInvariants),
		gov.NewAppModule(appCodec, govK, accountK, bankK),
		mint.NewAppModule(appCodec, mintK, accountK),
		slashing.NewAppModule(appCodec, slashingK, accountK, bankK, stakingK),
		distr.NewAppModule(appCodec, distrK, accountK, bankK, stakingK),
		staking.NewAppModule(appCodec, stakingK, accountK, bankK),
		upgrade.NewAppModule(upgradeK),
		evidence.NewAppModule(*evidenceK),
		params.NewAppModule(paramsK),
		capability.NewAppModule(appCodec, *capabilityK),
	}

	// wasm module needs to be added in a specific order
	if IsWasmEnabled() {
		appModules = append(
			appModules,
			wasm.NewAppModule(appCodec, &wasmK, stakingK, accountK, bankK),
			ibchooks.NewAppModule(accountK),
		)
	}

	appModules = append(appModules, []module.AppModule{
		evidence.NewAppModule(*evidenceK),
		ibc.NewAppModule(ibcKeeper),
		transfer.NewAppModule(transferKeeper),
		feegrantmodule.NewAppModule(appCodec, accountK, bankK, feegrantK, interfaceRegistry),

		snapshot.NewAppModule(snapK),
		multisig.NewAppModule(multisigK, stakingK, slashingK, snapK, rewardK, nexusK),
		tss.NewAppModule(tssK, snapK, nexusK, stakingK, multisigK),
		vote.NewAppModule(votingK),
		nexus.NewAppModule(nexusK, snapK, slashingK, stakingK, axelarnetK, rewardK),
		evm.NewAppModule(evmK, votingK, nexusK, snapK, stakingK, slashingK, multisigK),
		axelarnetModule,
		reward.NewAppModule(rewardK, nexusK, mintK, stakingK, slashingK, multisigK, snapK, bankK, bApp.MsgServiceRouter(), bApp.Router()),
		permission.NewAppModule(permissionK),
	}...)

	var app = &AxelarApp{
		BaseApp:           bApp,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		upgradeKeeper:     upgradeK,
	}

	mm = module.NewManager(
		appModules...,
	)
	app.mm = mm

	app.mm.SetOrderMigrations(orderMigrations()...)
	app.mm.SetOrderBeginBlockers(orderBeginBlockers()...)
	app.mm.SetOrderEndBlockers(orderEndBlockers()...)
	app.mm.SetOrderInitGenesis(orderModulesForGenesis()...)

	app.mm.RegisterInvariants(&crisisK)

	// register all module routes and module queriers
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(configurator)

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// The baseAnteHandler handles signature verification and transaction pre-processing
	baseAnteHandler, err := authAnte.NewAnteHandler(
		authAnte.HandlerOptions{
			AccountKeeper:   accountK,
			BankKeeper:      bankK,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  feegrantK,
			SigGasConsumer:  authAnte.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		panic(err)
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewAnteHandlerDecorator(baseAnteHandler),
	}

	// enforce wasm limits earlier in the ante handler chain
	if IsWasmEnabled() {
		anteDecorators = append(anteDecorators, wasmAnteDecorators...)
	}

	anteDecorators = append(anteDecorators,
		ante.NewLogMsgDecorator(appCodec),
		ante.NewCheckCommissionRate(stakingK),
		ante.NewUndelegateDecorator(multisigK, nexusK, snapK),
		ante.NewCheckRefundFeeDecorator(app.interfaceRegistry, accountK, stakingK, snapK, rewardK),
		ante.NewCheckProxy(snapK),
		ante.NewRestrictedTx(permissionK),
		ibcante.NewAnteDecorator(ibcKeeper),
	)

	anteHandler := sdk.ChainAnteDecorators(
		anteDecorators...,
	)
	app.SetAnteHandler(anteHandler)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}

		if IsWasmEnabled() {
			ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})

			// Initialize pinned codes in wasmvm as they are not persisted there
			if err := wasmK.InitializePinnedCodes(ctx); err != nil {
				tmos.Exit(fmt.Sprintf("failed initialize pinned codes %s", err))
			}
		}
	}

	/* ==== at this point all stores are fully loaded ==== */

	// we need to ensure that all chain subspaces are loaded at start-up to prevent unexpected consensus failures
	// when the params keeper is used outside the evm module's context
	evmK.InitChains(app.NewContext(true, tmproto.Header{}))

	return app
}

func orderMigrations() []string {
	migrationOrder := []string{
		// auth module needs to go first
		authtypes.ModuleName,
		// sdk modules
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
	}

	// wasm module needs to be added in a specific order
	if IsWasmEnabled() {
		migrationOrder = append(migrationOrder, wasm.ModuleName, ibchookstypes.ModuleName)
	}

	// axelar modules
	migrationOrder = append(migrationOrder,
		multisigTypes.ModuleName,
		tssTypes.ModuleName,
		rewardTypes.ModuleName,
		voteTypes.ModuleName,
		evmTypes.ModuleName,
		nexusTypes.ModuleName,
		permissionTypes.ModuleName,
		snapTypes.ModuleName,
		axelarnetTypes.ModuleName,
	)
	return migrationOrder
}

func orderBeginBlockers() []string {
	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	beginBlockerOrder := []string{
		// upgrades should be run first
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
	}

	// wasm module needs to be added in a specific order
	if IsWasmEnabled() {
		beginBlockerOrder = append(beginBlockerOrder, wasm.ModuleName, ibchookstypes.ModuleName)
	}

	// axelar custom modules
	beginBlockerOrder = append(beginBlockerOrder,
		rewardTypes.ModuleName,
		nexusTypes.ModuleName,
		permissionTypes.ModuleName,
		multisigTypes.ModuleName,
		tssTypes.ModuleName,
		evmTypes.ModuleName,
		snapTypes.ModuleName,
		axelarnetTypes.ModuleName,
		voteTypes.ModuleName,
	)
	return beginBlockerOrder
}

func orderEndBlockers() []string {
	endBlockerOrder := []string{
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		feegrant.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
	}

	// wasm module needs to be added in a specific order
	if IsWasmEnabled() {
		endBlockerOrder = append(endBlockerOrder, wasm.ModuleName, ibchookstypes.ModuleName)
	}

	// axelar custom modules
	endBlockerOrder = append(endBlockerOrder,
		multisigTypes.ModuleName,
		tssTypes.ModuleName,
		evmTypes.ModuleName,
		nexusTypes.ModuleName,
		rewardTypes.ModuleName,
		snapTypes.ModuleName,
		axelarnetTypes.ModuleName,
		permissionTypes.ModuleName,
		voteTypes.ModuleName,
	)
	return endBlockerOrder
}

func orderModulesForGenesis() []string {
	// Sets the order of Genesis - Order matters, genutil is to always come last
	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	genesisOrder := []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibchost.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
	}

	// wasm module needs to be added in a specific order
	if IsWasmEnabled() {
		genesisOrder = append(genesisOrder, wasm.ModuleName, ibchookstypes.ModuleName)
	}

	genesisOrder = append(genesisOrder,
		snapTypes.ModuleName,
		multisigTypes.ModuleName,
		tssTypes.ModuleName,
		evmTypes.ModuleName,
		nexusTypes.ModuleName,
		voteTypes.ModuleName,
		axelarnetTypes.ModuleName,
		rewardTypes.ModuleName,
		permissionTypes.ModuleName,
	)
	return genesisOrder
}

func createStoreKeys() map[string]*sdk.KVStoreKey {
	return sdk.NewKVStoreKeys(authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		minttypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		upgradetypes.StoreKey,
		evidencetypes.StoreKey,
		ibchost.StoreKey,
		ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey,
		feegrant.StoreKey,
		wasm.StoreKey,
		ibchookstypes.StoreKey,
		voteTypes.StoreKey,
		evmTypes.StoreKey,
		snapTypes.StoreKey,
		multisigTypes.StoreKey,
		tssTypes.StoreKey,
		nexusTypes.StoreKey,
		axelarnetTypes.StoreKey,
		rewardTypes.StoreKey,
		permissionTypes.StoreKey)
}

func initParamsKeeper(appCodec codec.Codec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(bam.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable())

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(wasm.ModuleName)

	paramsKeeper.Subspace(snapTypes.ModuleName)
	paramsKeeper.Subspace(multisigTypes.ModuleName)
	paramsKeeper.Subspace(tssTypes.ModuleName)
	paramsKeeper.Subspace(nexusTypes.ModuleName)
	paramsKeeper.Subspace(axelarnetTypes.ModuleName)
	paramsKeeper.Subspace(rewardTypes.ModuleName)
	paramsKeeper.Subspace(voteTypes.ModuleName)
	paramsKeeper.Subspace(permissionTypes.ModuleName)

	return paramsKeeper
}

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState map[string]json.RawMessage

// InitChainer handles the chain initialization from a genesis file
func (app *AxelarApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	app.upgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// BeginBlocker calls the BeginBlock() function of every module at the beginning of a new block
func (app *AxelarApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker calls the EndBlock() function of every module at the end of a block
func (app *AxelarApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// LoadHeight loads the application version at a given height. It will panic if called
// more than once on a running baseapp.
func (app *AxelarApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// AppCodec returns AxelarApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *AxelarApp) AppCodec() codec.Codec {
	return app.appCodec
}

// moduleAccountAddrs returns all the app's module account addresses.
func moduleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *AxelarApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	GetModuleBasics().RegisterRESTRoutes(clientCtx, apiSvr.Router)
	GetModuleBasics().RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(apiSvr.Router)
	}
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticServer))
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *AxelarApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *AxelarApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// GetModuleBasics initializes the module BasicManager is in charge of setting up basic,
// non-dependant module elements, such as codec registration and genesis verification.
// Initialization is dependent on whether wasm is enabled.
func GetModuleBasics() module.BasicManager {
	if ModuleBasics != nil {
		return ModuleBasics
	}

	var wasmProposals []govclient.ProposalHandler
	if IsWasmEnabled() {
		wasmProposals = wasmclient.ProposalHandlers
	}

	managers := []module.AppModuleBasic{
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			append(
				wasmProposals,
				paramsclient.ProposalHandler,
				distrclient.ProposalHandler,
				upgradeclient.ProposalHandler,
				upgradeclient.CancelProposalHandler,
				ibcclientclient.UpdateClientProposalHandler,
				ibcclientclient.UpgradeProposalHandler,
				axelarnetclient.ProposalHandler,
			)...,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		vesting.AppModuleBasic{},
	}

	if IsWasmEnabled() {
		managers = append(managers, wasm.AppModuleBasic{}, ibchooks.AppModuleBasic{})
	}

	managers = append(managers,
		ibc.AppModuleBasic{},
		transfer.AppModuleBasic{},

		multisig.AppModuleBasic{},
		tss.AppModuleBasic{},
		vote.AppModuleBasic{},
		evm.AppModuleBasic{},
		snapshot.AppModuleBasic{},
		nexus.AppModuleBasic{},
		axelarnet.AppModuleBasic{},
		reward.AppModuleBasic{},
		permission.AppModuleBasic{},
	)

	ModuleBasics = module.NewBasicManager(managers...)

	return ModuleBasics
}

func getSubspace(keeper paramskeeper.Keeper, moduleName string) paramstypes.Subspace {
	subspace, _ := keeper.GetSubspace(moduleName)
	return subspace
}

// IsWasmEnabled returns whether wasm is enabled
func IsWasmEnabled() bool {
	return WasmEnabled != ""
}
