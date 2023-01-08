package config

import (
	"fmt"
	"os"
	"testing"

	fruit "github.com/gowhale/go-test-data/pkg/fruits"
	"github.com/stretchr/testify/suite"
)

type configSuite struct {
	suite.Suite

	mockJSONReader *mockReaderJSON
}

func (c *configSuite) SetupTest() {
	c.mockJSONReader = new(mockReaderJSON)
}

func (*configSuite) AfterTest() {

}

func TestTerminalSuite(t *testing.T) {
	suite.Run(t, new(configSuite))
}

func (c *configSuite) Test_loadLyricsImpl_Pass() {
	expectedLyrics := LyricsConfig{}

	c.mockJSONReader.On("Open", "lyrics/Apple").Return(&os.File{}, nil)
	c.mockJSONReader.On("ReadAll", &os.File{}).Return([]byte{}, nil)
	var lyricsConfig LyricsConfig
	c.mockJSONReader.On("Unmarshal", []byte{}, &lyricsConfig).Return(nil)

	l, err := loadLyricsImpl(fruit.Apple, c.mockJSONReader)
	c.Nil(err)
	c.Equal(expectedLyrics, l)
}

func (c *configSuite) Test_loadLyricsImpl_Unmarshall_Fail() {
	expectedLyrics := LyricsConfig{}

	c.mockJSONReader.On("Open", "lyrics/Apple").Return(&os.File{}, nil)
	c.mockJSONReader.On("ReadAll", &os.File{}).Return([]byte{}, nil)
	var lyricsConfig LyricsConfig
	c.mockJSONReader.On("Unmarshal", []byte{}, &lyricsConfig).Return(fmt.Errorf("unmarshall err"))

	l, err := loadLyricsImpl(fruit.Apple, c.mockJSONReader)
	c.EqualError(err, "unmarshall err")
	c.Equal(expectedLyrics, l)
}

func (c *configSuite) Test_loadLyricsImpl_ReadAll_Fail() {
	expectedLyrics := LyricsConfig{}

	var lyricsConfig LyricsConfig
	c.mockJSONReader.On("Unmarshal", []byte{}, &lyricsConfig).Return(nil)
	c.mockJSONReader.On("Open", "lyrics/Apple").Return(&os.File{}, nil)
	c.mockJSONReader.On("ReadAll", &os.File{}).Return([]byte{}, fmt.Errorf("read err"))

	l, err := loadLyricsImpl(fruit.Apple, c.mockJSONReader)
	c.EqualError(err, "read err")
	c.Equal(expectedLyrics, l)
}

func (c *configSuite) Test_loadLyricsImpl_Open_Fail() {
	expectedLyrics := LyricsConfig{}

	var lyricsConfig LyricsConfig
	c.mockJSONReader.On("Unmarshal", []byte{}, &lyricsConfig).Return(nil)
	c.mockJSONReader.On("Open", "lyrics/Apple").Return(&os.File{}, fmt.Errorf("open err"))

	l, err := loadLyricsImpl(fruit.Apple, c.mockJSONReader)
	c.EqualError(err, "open err")
	c.Equal(expectedLyrics, l)
}
