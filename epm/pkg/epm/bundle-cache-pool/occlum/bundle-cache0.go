package occlum

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	cache_metadata "github.com/alibaba/inclavare-containers/epm/pkg/metadata"

	cache_manager "github.com/alibaba/inclavare-containers/epm/pkg/epm"
	"github.com/alibaba/inclavare-containers/epm/pkg/epm-api/v1alpha1"
	"github.com/alibaba/inclavare-containers/epm/pkg/epm/bundle-cache-pool/occlum/constants"
	"github.com/alibaba/inclavare-containers/epm/pkg/epm/bundle-cache-pool/occlum/types"
	"github.com/alibaba/inclavare-containers/epm/pkg/utils"
)

type BundleCach0Manager struct {
	cache_manager.DefaultEnclavePool
}

func NewBundleCach0Manager(root string, metadata *cache_metadata.Metadata) *BundleCach0Manager {
	return &BundleCach0Manager{
		DefaultEnclavePool: cache_manager.DefaultEnclavePool{
			Root:          root,
			Type:          string(types.BundleCache0PoolType),
			CacheMetadata: metadata,
		}}
}

func (d *BundleCach0Manager) GetPoolType() string {
	return d.Type
}

func (d *BundleCach0Manager) SaveCache(sourcePath string, cache *v1alpha1.Cache) error {
	savePath, err := d.BuildCacheSavePath(d.Root, cache)
	if err != nil {
		return err
	}
	if err := os.RemoveAll(savePath); err != nil {
		return nil
	}
	if err := os.MkdirAll(savePath, 0755); err != nil {
		return err
	}

	sourceDirs := []string{"build/mount/"}
	for _, dir := range sourceDirs {
		srcDir := filepath.Join(sourcePath, dir)
		destDir := filepath.Join(savePath, dir)
		src, err := os.Stat(srcDir)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(destDir, src.Mode()); err != nil {
			return err
		}
		if err := utils.CopyDirectory(srcDir, destDir); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(filepath.Join(savePath, constants.OcclumStatusFileName), []byte(types.ImageBuilt), 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(savePath, constants.OcclumSGXModeFileName), []byte("HW"), 0644); err != nil {
		return err
	}

	size, err := utils.DirSize(savePath)
	if err != nil {
		return err
	}

	cache.SavePath = savePath
	cache.Size = size
	cache.Created = time.Now().Unix()
	return d.CacheMetadata.SaveCache(d.GetPoolType(), cache.ID, cache)
}
