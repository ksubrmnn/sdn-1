package hcsshim

import (
	"crypto/sha1"
	"path/filepath"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/Microsoft/hcsshim/internal/wclayer"
)

func layerPath(info *DriverInfo, id string) string {
	return filepath.Join(info.HomeDir, id)
}

func ActivateLayer(info DriverInfo, id string) error {
	return wclayer.ActivateLayer(layerPath(&info, id))
}
func CreateLayer(info DriverInfo, id, parent string) error {
	return wclayer.CreateLayer(layerPath(&info, id), parent)
}

// New clients should use CreateScratchLayer instead. Kept in to preserve API compatibility.
func CreateSandboxLayer(info DriverInfo, layerId, parentId string, parentLayerPaths []string) error {
	return wclayer.CreateScratchLayer(layerPath(&info, layerId), parentLayerPaths)
}
func CreateScratchLayer(info DriverInfo, layerId, parentId string, parentLayerPaths []string) error {
	return wclayer.CreateScratchLayer(layerPath(&info, layerId), parentLayerPaths)
}
func DeactivateLayer(info DriverInfo, id string) error {
	return wclayer.DeactivateLayer(layerPath(&info, id))
}
func DestroyLayer(info DriverInfo, id string) error {
	return wclayer.DestroyLayer(layerPath(&info, id))
}

// New clients should use ExpandScratchSize instead. Kept in to preserve API compatibility.
func ExpandSandboxSize(info DriverInfo, layerId string, size uint64) error {
	return wclayer.ExpandScratchSize(layerPath(&info, layerId), size)
}
func ExpandScratchSize(info DriverInfo, layerId string, size uint64) error {
	return wclayer.ExpandScratchSize(layerPath(&info, layerId), size)
}
func ExportLayer(info DriverInfo, layerId string, exportFolderPath string, parentLayerPaths []string) error {
	return wclayer.ExportLayer(layerPath(&info, layerId), exportFolderPath, parentLayerPaths)
}
func GetLayerMountPath(info DriverInfo, id string) (string, error) {
	return wclayer.GetLayerMountPath(layerPath(&info, id))
}
func GetSharedBaseImages() (imageData string, err error) {
	return wclayer.GetSharedBaseImages()
}
func ImportLayer(info DriverInfo, layerID string, importFolderPath string, parentLayerPaths []string) error {
	return wclayer.ImportLayer(layerPath(&info, layerID), importFolderPath, parentLayerPaths)
}
func LayerExists(info DriverInfo, id string) (bool, error) {
	return wclayer.LayerExists(layerPath(&info, id))
}
func PrepareLayer(info DriverInfo, layerId string, parentLayerPaths []string) error {
	return wclayer.PrepareLayer(layerPath(&info, layerId), parentLayerPaths)
}
func ProcessBaseLayer(path string) error {
	return wclayer.ProcessBaseLayer(path)
}
func ProcessUtilityVMImage(path string) error {
	return wclayer.ProcessUtilityVMImage(path)
}
func UnprepareLayer(info DriverInfo, layerId string) error {
	return wclayer.UnprepareLayer(layerPath(&info, layerId))
}

type DriverInfo struct {
	Flavour int
	HomeDir string
}

type GUID [16]byte

func NameToGuid(name string) (id GUID, err error) {
	g, err := wclayer.NameToGuid(name)
	return g.ToWindowsArray(), err
}

func NewGUID(source string) *GUID {
	h := sha1.Sum([]byte(source))
	var g GUID
	copy(g[0:], h[0:16])
	return &g
}

func (g *GUID) ToString() string {
	return guid.FromWindowsArray(*g).String()
}

type LayerReader = wclayer.LayerReader

func NewLayerReader(info DriverInfo, layerID string, parentLayerPaths []string) (LayerReader, error) {
	return wclayer.NewLayerReader(layerPath(&info, layerID), parentLayerPaths)
}

type LayerWriter = wclayer.LayerWriter

func NewLayerWriter(info DriverInfo, layerID string, parentLayerPaths []string) (LayerWriter, error) {
	return wclayer.NewLayerWriter(layerPath(&info, layerID), parentLayerPaths)
}

type WC_LAYER_DESCRIPTOR = wclayer.WC_LAYER_DESCRIPTOR
