/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

type Auxiliaries interface {
	Get(string) Auxiliary
	GetList() []Auxiliary
}
