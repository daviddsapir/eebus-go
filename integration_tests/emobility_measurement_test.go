package integrationtests

import (
	"testing"

	"github.com/DerAndereAndi/eebus-go/features"
	"github.com/DerAndereAndi/eebus-go/spine"
	"github.com/DerAndereAndi/eebus-go/spine/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestEmobilityMeasurementSuite(t *testing.T) {
	suite.Run(t, new(EmobilityMeasurementSuite))
}

type EmobilityMeasurementSuite struct {
	suite.Suite
	sut *spine.DeviceLocalImpl

	measurement          *features.Measurement
	electricalconnection *features.ElectricalConnection

	remoteSki string
	readC     chan []byte
	writeC    chan []byte
}

func (s *EmobilityMeasurementSuite) SetupSuite() {
}

func (s *EmobilityMeasurementSuite) BeforeTest(suiteName, testName string) {
	s.sut = spine.NewDeviceLocalImpl("TestBrandName", "TestDeviceModel", "TestSerialNumber", "TestDeviceCode",
		"TestDeviceAddress", model.DeviceTypeTypeEnergyManagementSystem, model.NetworkManagementFeatureSetTypeSmart)
	localEntity := spine.NewEntityLocalImpl(s.sut, model.EntityTypeTypeCEM, spine.NewAddressEntityType([]uint{1}))
	s.sut.AddEntity(localEntity)

	f := spine.NewFeatureLocalImpl(1, localEntity, model.FeatureTypeTypeElectricalConnection, model.RoleTypeClient)
	localEntity.AddFeature(f)
	f = spine.NewFeatureLocalImpl(2, localEntity, model.FeatureTypeTypeMeasurement, model.RoleTypeClient)
	localEntity.AddFeature(f)

	s.remoteSki = "TestRemoteSki"

	s.readC = make(chan []byte, 1)
	s.writeC = make(chan []byte, 1)

	s.sut.AddRemoteDevice(s.remoteSki, s.readC, s.writeC)

	initialCommunication(s.T(), s.readC, s.writeC)
}

func (s *EmobilityMeasurementSuite) AfterTest(suiteName, testName string) {
}

func (s *EmobilityMeasurementSuite) TestGetValuesPerPhaseForScope() {
	remoteEntity := s.sut.RemoteDeviceForSki(s.remoteSki).Entity([]model.AddressEntityType{1, 1})
	assert.NotNil(s.T(), remoteEntity)

	var err error
	s.measurement, err = features.NewMeasurement(model.RoleTypeClient, model.RoleTypeServer, s.sut, remoteEntity)
	assert.Nil(s.T(), err)

	s.electricalconnection, err = features.NewElectricalConnection(model.RoleTypeClient, model.RoleTypeServer, s.sut, remoteEntity)
	assert.Nil(s.T(), err)

	// Act
	s.readC <- loadFileData(s.T(), ec_parameterdescriptionlistdata_recv_reply_file_path)
	waitForAck(s.T(), s.writeC)

	s.readC <- loadFileData(s.T(), m_descriptionListData_recv_reply_file_path)
	waitForAck(s.T(), s.writeC)

	s.readC <- loadFileData(s.T(), m_measurementListData_recv_notify_file_path)
	waitForAck(s.T(), s.writeC)

	resultMap, err := s.measurement.GetValuesPerPhaseForScope(model.ScopeTypeTypeACCurrent, s.electricalconnection)

	// Assert
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resultMap)
	assert.Equal(s.T(), 1, len(resultMap))
	assert.Equal(s.T(), 5.0, resultMap["a"])

	resultMap, err = s.measurement.GetValuesPerPhaseForScope(model.ScopeTypeTypeACPower, s.electricalconnection)

	// Assert
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resultMap)
	assert.Equal(s.T(), 1, len(resultMap))
	assert.Equal(s.T(), 1185.0, resultMap["a"])

	resultMap, err = s.measurement.GetValuesPerPhaseForScope(model.ScopeTypeTypeCharge, s.electricalconnection)

	// Assert
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resultMap)
	assert.Equal(s.T(), 1, len(resultMap))
	assert.Equal(s.T(), 1825.0, resultMap["a"])
}