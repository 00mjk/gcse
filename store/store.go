// Package store handlings all the storage in GCSE backend.
package store

import (
	"log"
	"time"

	"github.com/golangplus/bytes"
	"github.com/golangplus/errors"

	"github.com/daviddengcn/bolthelper"
	"github.com/daviddengcn/gcse/configs"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	sppb "github.com/daviddengcn/gcse/proto/spider"
	stpb "github.com/daviddengcn/gcse/proto/store"
)

var (
	// pkgs
	//   - <site>
	//     - <path> -> PackageInfo
	// persons
	//   - <site>
	//     - <id> -> PersonInfo
	// history
	//  - pkgs
	//     - <site/path> -> HistoryInfo
	//  - persons
	//     - <site/id> -> HistoryInfo
	pkgsRoot    = []byte("pkgs")
	personsRoot = []byte("persons")
	historyRoot = []byte("history")
)

var box = &bh.RefCountBox{
	DataPath: func() string {
		return configs.DataRoot.Join("store.bolt").S()
	},
}

func RepoInfoAge(r *sppb.RepoInfo) time.Duration {
	t, _ := ptypes.Timestamp(r.LastCrawled)
	return time.Now().Sub(t)
}

// Returns an empty (non-nil) PackageInfo if not found.
func ReadPackage(site, path string) (*stpb.PackageInfo, error) {
	info := &stpb.PackageInfo{}
	if err := box.View(func(tx bh.Tx) error {
		return tx.Value([][]byte{pkgsRoot, []byte(site), []byte(path)}, func(bs bytesp.Slice) error {
			if err := errorsp.WithStacksAndMessage(proto.Unmarshal(bs, info), "Unmarshal %d bytes failed", len(bs)); err != nil {
				log.Printf("Unmarshal failed: %v", err)
				*info = stpb.PackageInfo{}
			}
			return nil
		})
	}); err != nil {
		return nil, err
	}
	return info, nil
}

func UpdatePackage(site, path string, f func(*stpb.PackageInfo) error) error {
	return box.Update(func(tx bh.Tx) error {
		b, err := tx.CreateBucketIfNotExists([][]byte{pkgsRoot, []byte(site)})
		if err != nil {
			return err
		}
		info := &stpb.PackageInfo{}
		if err := b.Value([][]byte{[]byte(path)}, func(bs bytesp.Slice) error {
			if err := errorsp.WithStacksAndMessage(proto.Unmarshal(bs, info), "Unmarshal %d bytes", len(bs)); err != nil {
				log.Printf("Unmarshaling failed: %v", err)
				*info = stpb.PackageInfo{}
			}
			return nil
		}); err != nil {
			return err
		}
		if err := errorsp.WithStacks(f(info)); err != nil {
			return err
		}
		bs, err := proto.Marshal(info)
		if err != nil {
			return errorsp.WithStacksAndMessage(err, "marshaling %v failed: %v", info, err)
		}
		return b.Put([][]byte{[]byte(path)}, bs)
	})
}

func DeletePackage(site, path string) error {
	return box.Update(func(tx bh.Tx) error {
		return tx.Delete([][]byte{pkgsRoot, []byte(site), []byte(path)})
	})
}

func ReadPerson(site, id string) (*stpb.PersonInfo, error) {
	info := &stpb.PersonInfo{}
	if err := box.View(func(tx bh.Tx) error {
		return tx.Value([][]byte{personsRoot, []byte(site), []byte(id)}, func(bs bytesp.Slice) error {
			if err := errorsp.WithStacksAndMessage(proto.Unmarshal(bs, info), "Unmarshal %d bytes failed", len(bs)); err != nil {
				log.Printf("Unmarshal failed: %v", err)
				*info = stpb.PersonInfo{}
			}
			return nil
		})
	}); err != nil {
		return nil, err
	}
	return info, nil
}

func UpdatePerson(site, id string, f func(*stpb.PersonInfo) error) error {
	return box.Update(func(tx bh.Tx) error {
		b, err := tx.CreateBucketIfNotExists([][]byte{personsRoot, []byte(site)})
		if err != nil {
			return err
		}
		info := &stpb.PersonInfo{}
		if err := b.Value([][]byte{[]byte(id)}, func(bs bytesp.Slice) error {
			err := errorsp.WithStacksAndMessage(proto.Unmarshal(bs, info), "Unmarshal %d bytes", len(bs))
			if err != nil {
				log.Printf("Unmarshaling failed: %v", err)
				*info = stpb.PersonInfo{}
			}
			return nil
		}); err != nil {
			return err
		}
		if err := errorsp.WithStacks(f(info)); err != nil {
			return err
		}
		bs, err := proto.Marshal(info)
		if err != nil {
			return errorsp.WithStacksAndMessage(err, "marshaling %v failed: %v", info, err)
		}
		return b.Put([][]byte{[]byte(id)}, bs)
	})
}

func DeletePerson(site, id string) error {
	return box.Update(func(tx bh.Tx) error {
		return tx.Delete([][]byte{personsRoot, []byte(site), []byte(id)})
	})
}
