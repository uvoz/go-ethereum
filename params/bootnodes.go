// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes MODIFIED for test
	"enode://05b06f88d6236f0518a7f954519997fc9ace0b5edfb8edfdc6a0d8f203b84974fb3ba9c38eb251839446093dfa5760707cbc4e71be0f9e2acb24771c570a4c5a@3.1.4.1:30301",
	"enode://3a1060874a7bcd53a3c5037da28e7c366941f3915602881b16a1e5fd69691959ef482a67cecfb6fb5942ad42b11f1ddb30e6b964e6347ea73b0a530f44438eb2@104.214.217.23:30301"}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network. // //MODIFIED for test
var TestnetBootnodes = []string{
	"enode://05b06f88d6236f0518a7f954519997fc9ace0b5edfb8edfdc6a0d8f203b84974fb3ba9c38eb251839446093dfa5760707cbc4e71be0f9e2acb24771c570a4c5a@3.1.4.1:30301"}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://05b06f88d6236f0518a7f954519997fc9ace0b5edfb8edfdc6a0d8f203b84974fb3ba9c38eb251839446093dfa5760707cbc4e71be0f9e2acb24771c570a4c5a@3.1.4.1:30301"}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://05b06f88d6236f0518a7f954519997fc9ace0b5edfb8edfdc6a0d8f203b84974fb3ba9c38eb251839446093dfa5760707cbc4e71be0f9e2acb24771c570a4c5a@3.1.4.1:30301"}
