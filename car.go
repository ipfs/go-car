package car

import (
	"bufio"
	"context"
	"io"

	cid "github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	ipldcar "github.com/ipld/go-car"
)

type Store = ipldcar.Store
type CarHeader = ipldcar.CarHeader
type CarReader = ipldcar.CarReader
type WalkFunc = ipldcar.WalkFunc

func WriteCar(ctx context.Context, ds format.DAGService, roots []cid.Cid, w io.Writer) error {
	return ipldcar.WriteCar(ctx, ds, roots, w)
}

func WriteCarWithWalker(ctx context.Context, ds format.DAGService, roots []cid.Cid, w io.Writer, walk WalkFunc) error {
	return ipldcar.WriteCarWithWalker(ctx, ds, roots, w, walk)
}

func DefaultWalkFunc(nd format.Node) ([]*format.Link, error) {
	return ipldcar.DefaultWalkFunc(nd)
}

func ReadHeader(br *bufio.Reader) (*CarHeader, error) {
	return ipldcar.ReadHeader(br)
}

func NewCarReader(r io.Reader) (*CarReader, error) {
	return ipldcar.NewCarReader(r)
}

func LoadCar(s Store, r io.Reader) (*CarHeader, error) {
	return ipldcar.LoadCar(s, r)
}
