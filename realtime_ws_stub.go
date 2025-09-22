//go:build !realtime

package flaro

import (
	"context"
	"fmt"
	"time"
)

// RealtimeClient is unavailable without the 'realtime' build tag.
type RealtimeClient struct{}

func NewRealtimeClient(apiKey string) *RealtimeClient { return &RealtimeClient{} }
func (r *RealtimeClient) Connect(_ context.Context) error {
	return fmt.Errorf("realtime disabled: build with -tags realtime")
}
func (r *RealtimeClient) Close() error { return nil }
func (r *RealtimeClient) SubscribePostsForCreator(_, _ string) error {
	return fmt.Errorf("realtime disabled: build with -tags realtime")
}
func (r *RealtimeClient) StartHeartbeat(_ context.Context, _ time.Duration) error {
	return fmt.Errorf("realtime disabled: build with -tags realtime")
}
func (r *RealtimeClient) ReadRaw(_ interface{}) error {
	return fmt.Errorf("realtime disabled: build with -tags realtime")
}

type RealtimeEnvelope struct{}
type PhxReplyPayload struct{}
type SystemPayload struct{}

func (r *RealtimeClient) ReadMessage() (*RealtimeEnvelope, error) {
	return nil, fmt.Errorf("realtime disabled: build with -tags realtime")
}
func (env *RealtimeEnvelope) UnmarshalPayload(_ interface{}) error {
	return fmt.Errorf("realtime disabled: build with -tags realtime")
}
