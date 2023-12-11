package test

import (
	"assimp/common/logger"
	"assimp/driver"
	"go.uber.org/zap"
	"testing"
)

func TestAcImporter(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/closedLine.ac")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}

func TestAcImporter1(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/nosurfaces.ac")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}

func TestAcImporter2(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/openLine.ac")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}

func TestAcImporter3(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/sample_subdiv.ac")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}

func TestAcImporter4(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/SphereWithLight.ac")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}

func TestAcImporter5(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/SphereWithLight_UTF8BOM.ac")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}

func TestAcImporter6(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/SphereWithLight_UTF16LE.ac")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}

func TestAcImporter7(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/SphereWithLightUvScaling4X.ac")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}

func TestAcImporter8(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/TestFormatDetection")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}

func TestAcImporter9(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/Wuson.ac")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}

func TestAcImporter10(t *testing.T) {
	im := driver.NewImporter()
	p, err := im.ReadFile("../example_data/models/AC/Wuson.acc")
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	Assert(t, p != nil)
}