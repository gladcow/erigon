// Code generated by fastssz. DO NOT EDIT.
// Hash: 78616528ff009030b8c5c8324875c6bcfe04bd54faa9ee90c8e216ab17271c73
// Version: 0.1.2
package lightrpc

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the MetadataV1 object
func (m *MetadataV1) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MetadataV1 object to a target array
func (m *MetadataV1) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SeqNumber'
	dst = ssz.MarshalUint64(dst, m.SeqNumber)

	// Field (1) 'Attnets'
	dst = ssz.MarshalUint64(dst, m.Attnets)

	return
}

// UnmarshalSSZ ssz unmarshals the MetadataV1 object
func (m *MetadataV1) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'SeqNumber'
	m.SeqNumber = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Attnets'
	m.Attnets = ssz.UnmarshallUint64(buf[8:16])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MetadataV1 object
func (m *MetadataV1) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the MetadataV1 object
func (m *MetadataV1) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MetadataV1 object with a hasher
func (m *MetadataV1) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SeqNumber'
	hh.PutUint64(m.SeqNumber)

	// Field (1) 'Attnets'
	hh.PutUint64(m.Attnets)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the MetadataV1 object
func (m *MetadataV1) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}

// MarshalSSZ ssz marshals the MetadataV2 object
func (m *MetadataV2) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MetadataV2 object to a target array
func (m *MetadataV2) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SeqNumber'
	dst = ssz.MarshalUint64(dst, m.SeqNumber)

	// Field (1) 'Attnets'
	dst = ssz.MarshalUint64(dst, m.Attnets)

	// Field (2) 'Syncnets'
	dst = ssz.MarshalUint64(dst, m.Syncnets)

	return
}

// UnmarshalSSZ ssz unmarshals the MetadataV2 object
func (m *MetadataV2) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 24 {
		return ssz.ErrSize
	}

	// Field (0) 'SeqNumber'
	m.SeqNumber = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Attnets'
	m.Attnets = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'Syncnets'
	m.Syncnets = ssz.UnmarshallUint64(buf[16:24])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MetadataV2 object
func (m *MetadataV2) SizeSSZ() (size int) {
	size = 24
	return
}

// HashTreeRoot ssz hashes the MetadataV2 object
func (m *MetadataV2) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MetadataV2 object with a hasher
func (m *MetadataV2) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SeqNumber'
	hh.PutUint64(m.SeqNumber)

	// Field (1) 'Attnets'
	hh.PutUint64(m.Attnets)

	// Field (2) 'Syncnets'
	hh.PutUint64(m.Syncnets)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the MetadataV2 object
func (m *MetadataV2) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(m)
}