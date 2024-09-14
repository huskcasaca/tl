package main

import (
	"context"
	"fmt"
	tl "github.com/xelaj/tl"
)

type GenericType interface {
	tl.Any
	_GenericType()
}

var (
	_ GenericType = (*GenericsTypePredict)(nil)
)

type GenericsTypePredict[X tl.Any] struct {
	X X `tl:"x"`
}

func (*GenericsTypePredict[X]) CRC() uint32 {
	return 0x12345678
}
func (*GenericsTypePredict[X]) _GenericType() {}

type Update interface {
	tl.Any
	_Update()
}

var (
	_ Update = (*GenericsContainerTypePredict)(nil)
)

type GenericsContainerTypePredict struct {
	Var [][]GenericType `tl:"var"`
}

func (*GenericsContainerTypePredict) CRC() uint32 {
	return 0x12345678
}
func (*GenericsContainerTypePredict) _Update() {}

type Requester interface {
	MakeRequest(ctx context.Context, msg []byte) ([]byte, error)
}

func request[IN any, OUT any](ctx context.Context, m Requester, in *IN, out *OUT) error {
	if msg, err := tl.Marshal(in); err != nil {
		return fmt.Errorf("marshaling: %w", err)
	} else if respRaw, err := m.MakeRequest(ctx, msg); err != nil {
		return fmt.Errorf("sending: %w", err)
	} else if err := tl.Unmarshal(respRaw, out); err != nil {
		return fmt.Errorf("got invalid response type: %w", err)
	}
	return nil
}
