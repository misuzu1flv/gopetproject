package sender_test

import (
	"errors"
	mock_producer "homework-8/internal/pkg/kafka/mocks"
	"homework-8/internal/pkg/sender"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	testMsg = sender.Message{EventType: "test", Request: "test", Time: time.Now()}
)

func TestSendMessageSuccess(t *testing.T) {
	t.Parallel()

	// arrange
	ctrl := gomock.NewController(t)
	mockProducer := mock_producer.NewMockProducer(ctrl)
	mockProducer.EXPECT().SendSyncMessage(gomock.Any()).Return(int32(0), int64(0), nil)
	sender := sender.NewKafkaSender(mockProducer, "logs")

	//act
	err := sender.SendMessage(testMsg)
	// assert
	require.Nil(t, err)
}

func TestSendMessageFail(t *testing.T) {
	t.Parallel()

	// arrange
	ctrl := gomock.NewController(t)
	mockProducer := mock_producer.NewMockProducer(ctrl)
	mockProducer.EXPECT().SendSyncMessage(gomock.Any()).Return(int32(0), int64(0), errors.New("Hello"))
	sender := sender.NewKafkaSender(mockProducer, "logs")

	//act
	err := sender.SendMessage(testMsg)
	// assert
	require.NotNil(t, err)
}
