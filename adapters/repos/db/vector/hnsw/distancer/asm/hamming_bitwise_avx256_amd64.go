//go:build !noasm && amd64
// AUTO-GENERATED BY GOAT -- DO NOT EDIT

package asm

import "unsafe"

//go:noescape
func popcnt_64bit(src, popcnt_constants unsafe.Pointer)

//go:noescape
func popcnt_lookup_64bit(data, n, lookup64bit, popcnt_constants unsafe.Pointer)

//go:noescape
func popcount(x, lookup64bit unsafe.Pointer)

//go:noescape
func hamming_bitwise_256(a, b, res, len, lookup64bit, popcnt_constants unsafe.Pointer)
