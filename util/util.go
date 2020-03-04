package util

import (
	"bufio"
	"io"

	cid "github.com/ipfs/go-cid"
	ipldcarutil "github.com/ipld/go-car/util"
)

type BytesReader = ipldcarutil.BytesReader

func ReadCid(buf []byte) (cid.Cid, int, error) {
	return ipldcarutil.ReadCid(buf)
}

func ReadNode(br *bufio.Reader) (cid.Cid, []byte, error) {
	return ipldcarutil.ReadNode(br)
}

func LdWrite(w io.Writer, d ...[]byte) error {
	return ipldcarutil.LdWrite(w, d...)
}

func LdRead(r *bufio.Reader) ([]byte, error) {
	return ipldcarutil.LdRead(r)
}
