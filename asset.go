/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// Asset stores a value
type Asset struct {
	// Value string `json:"value"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
