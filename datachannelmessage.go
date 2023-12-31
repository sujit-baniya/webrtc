// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package webrtc

// DataChannelMessage represents a message received from the
// data channel. IsString will be set to true if the incoming
// message is of the string type. Otherwise the message is of
// a binary type.
type DataChannelMessage struct {
	IsString bool
	Data     []byte
}
