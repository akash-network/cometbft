package core

import (
	rpc "github.com/cometbft/cometbft/rpc/jsonrpc/server"
)

// TODO: better system than "unsafe" prefix

type RoutesMap map[string]*rpc.RPCFunc

var Routes RoutesMap = RoutesMap{}

// GetRoutes is a map of available routes.
func (env *Environment) GetRoutes() RoutesMap {
	routes := make(RoutesMap)
	for name, route := range Routes {
		routes[name] = route
	}

	// subscribe/unsubscribe are reserved for websocket events.
	routes["subscribe"] = rpc.NewWSRPCFunc(env.Subscribe, "query")
	routes["unsubscribe"] = rpc.NewWSRPCFunc(env.Unsubscribe, "query")
	routes["unsubscribe_all"] = rpc.NewWSRPCFunc(env.UnsubscribeAll, "")

	// info AP
	routes["health"] = rpc.NewRPCFunc(env.Health, "")
	routes["status"] = rpc.NewRPCFunc(env.Status, "")
	routes["net_info"] = rpc.NewRPCFunc(env.NetInfo, "")
	routes["blockchain"] = rpc.NewRPCFunc(env.BlockchainInfo, "minHeight,maxHeight", rpc.Cacheable())
	routes["genesis"] = rpc.NewRPCFunc(env.Genesis, "", rpc.Cacheable())
	routes["genesis_chunked"] = rpc.NewRPCFunc(env.GenesisChunked, "chunk", rpc.Cacheable())
	routes["block"] = rpc.NewRPCFunc(env.Block, "height", rpc.Cacheable("height"))
	routes["block_by_hash"] = rpc.NewRPCFunc(env.BlockByHash, "hash", rpc.Cacheable())
	routes["block_results"] = rpc.NewRPCFunc(env.BlockResults, "height", rpc.Cacheable("height"))
	routes["commit"] = rpc.NewRPCFunc(env.Commit, "height", rpc.Cacheable("height"))
	routes["header"] = rpc.NewRPCFunc(env.Header, "height", rpc.Cacheable("height"))
	routes["header_by_hash"] = rpc.NewRPCFunc(env.HeaderByHash, "hash", rpc.Cacheable())
	routes["check_tx"] = rpc.NewRPCFunc(env.CheckTx, "tx")
	routes["tx"] = rpc.NewRPCFunc(env.Tx, "hash,prove", rpc.Cacheable())
	routes["tx_search"] = rpc.NewRPCFunc(env.TxSearch, "query,prove,page,per_page,order_by")
	routes["block_search"] = rpc.NewRPCFunc(env.BlockSearch, "query,page,per_page,order_by")
	routes["validators"] = rpc.NewRPCFunc(env.Validators, "height,page,per_page", rpc.Cacheable("height"))
	routes["dump_consensus_state"] = rpc.NewRPCFunc(env.DumpConsensusState, "")
	routes["consensus_state"] = rpc.NewRPCFunc(env.GetConsensusState, "")
	routes["consensus_params"] = rpc.NewRPCFunc(env.ConsensusParams, "height", rpc.Cacheable("height"))
	routes["unconfirmed_txs"] = rpc.NewRPCFunc(env.UnconfirmedTxs, "limit")
	routes["num_unconfirmed_txs"] = rpc.NewRPCFunc(env.NumUnconfirmedTxs, "")

	// tx broadcast API
	routes["broadcast_tx_commit"] = rpc.NewRPCFunc(env.BroadcastTxCommit, "tx")
	routes["broadcast_tx_sync"] = rpc.NewRPCFunc(env.BroadcastTxSync, "tx")
	routes["broadcast_tx_async"] = rpc.NewRPCFunc(env.BroadcastTxAsync, "tx")

	// abci API
	routes["abci_query"] = rpc.NewRPCFunc(env.ABCIQuery, "path,data,height,prove")
	routes["abci_info"] = rpc.NewRPCFunc(env.ABCIInfo, "", rpc.Cacheable())

	// evidence API
	routes["broadcast_evidence"] = rpc.NewRPCFunc(env.BroadcastEvidence, "evidence")

	return routes
}

// AddUnsafeRoutes adds unsafe routes.
func (env *Environment) AddUnsafeRoutes(routes RoutesMap) {
	// control API
	routes["dial_seeds"] = rpc.NewRPCFunc(env.UnsafeDialSeeds, "seeds")
	routes["dial_peers"] = rpc.NewRPCFunc(env.UnsafeDialPeers, "peers,persistent,unconditional,private")
	routes["unsafe_flush_mempool"] = rpc.NewRPCFunc(env.UnsafeFlushMempool, "")
}
