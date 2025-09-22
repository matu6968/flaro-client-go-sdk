//go:build realtime

package flaro

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// RealtimeClient is an experimental websocket client for Flaro realtime messages.
// This API is experimental and may change without notice.
type RealtimeClient struct {
	conn   *websocket.Conn
	mu     sync.Mutex
	baseWS string
	apiKey string
	ref    int64
}

// NewRealtimeClient creates a new realtime client. baseWS should be ws(s):// host.
func NewRealtimeClient(apiKey string) *RealtimeClient {
	return &RealtimeClient{
		baseWS: "wss://sb.flaroapp.pl",
		apiKey: apiKey,
		ref:    0,
	}
}

// Connect dials the websocket endpoint.
func (r *RealtimeClient) Connect(ctx context.Context) error {
	u := url.URL{
		Scheme: "wss",
		Host:   "sb.flaroapp.pl",
		Path:   "/realtime/v1/websocket",
	}
	q := u.Query()
	q.Set("apikey", r.apiKey)
	q.Set("vsn", "1.0.0")
	u.RawQuery = q.Encode()

	dialer := websocket.Dialer{}
	headers := http.Header{}
	headers.Set("apikey", r.apiKey)

	conn, _, err := dialer.DialContext(ctx, u.String(), headers)
	if err != nil {
		return fmt.Errorf("failed to connect websocket: %w", err)
	}
	r.conn = conn
	return nil
}

// Close closes the websocket connection.
func (r *RealtimeClient) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.conn != nil {
		return r.conn.Close()
	}
	return nil
}

// nextRef returns a new incremented ref id as string.
func (r *RealtimeClient) nextRef() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// subscribePayload represents the join message per docs.
type subscribePayload struct {
	Topic   string      `json:"topic"`
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
	Ref     string      `json:"ref"`
	JoinRef string      `json:"join_ref"`
}

// SubscribePostsForCreator subscribes to realtime changes for posts of a given creator_id.
func (r *RealtimeClient) SubscribePostsForCreator(accessToken, creatorID string) error {
	if r.conn == nil {
		return fmt.Errorf("websocket not connected")
	}

	p := map[string]interface{}{
		"config": map[string]interface{}{
			"broadcast": map[string]interface{}{"ack": false, "self": false},
			"presence":  map[string]interface{}{"key": ""},
			"postgres_changes": []map[string]interface{}{
				{
					"event":  "*",
					"schema": "public",
					"table":  "posts",
					"filter": fmt.Sprintf("creator_id=eq.%s", creatorID),
				},
			},
			"private": false,
		},
		"access_token": accessToken,
	}

	msg := subscribePayload{
		Topic:   fmt.Sprintf("realtime:public:posts:%s", creatorID),
		Event:   "phx_join",
		Payload: p,
		Ref:     r.nextRef(),
		JoinRef: r.nextRef(),
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	return r.conn.WriteJSON(msg)
}

// StartHeartbeat starts sending phoenix heartbeats every interval until ctx is done.
func (r *RealtimeClient) StartHeartbeat(ctx context.Context, interval time.Duration) error {
	if r.conn == nil {
		return fmt.Errorf("websocket not connected")
	}
	if interval <= 0 {
		interval = 10 * time.Second
	}
	go func() {
		t := time.NewTicker(interval)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				msg := map[string]interface{}{
					"topic":   "phoenix",
					"event":   "heartbeat",
					"payload": map[string]interface{}{},
					"ref":     r.nextRef(),
				}
				r.mu.Lock()
				_ = r.conn.WriteJSON(msg)
				r.mu.Unlock()
			}
		}
	}()
	return nil
}

// ReadRaw blocks and reads a raw frame into v (json). Callers should run this in a goroutine.
func (r *RealtimeClient) ReadRaw(v interface{}) error {
	if r.conn == nil {
		return fmt.Errorf("websocket not connected")
	}
	_, data, err := r.conn.ReadMessage()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// RealtimeEnvelope describes the top-level structure of incoming realtime frames.
type RealtimeEnvelope struct {
	Ref     *string         `json:"ref"`
	Event   string          `json:"event"`
	Payload json.RawMessage `json:"payload"`
	Topic   string          `json:"topic"`
}

// PhxReplyPayload matches payload for event "phx_reply".
type PhxReplyPayload struct {
	Status   string          `json:"status"`
	Response json.RawMessage `json:"response"`
}

// SystemPayload matches payload for event "system" (error messages, etc.).
type SystemPayload struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	Extension string `json:"extension"`
	Channel   string `json:"channel"`
}

// ReadMessage reads a single frame and decodes into RealtimeEnvelope.
func (r *RealtimeClient) ReadMessage() (*RealtimeEnvelope, error) {
	if r.conn == nil {
		return nil, fmt.Errorf("websocket not connected")
	}
	_, data, err := r.conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	var env RealtimeEnvelope
	if err := json.Unmarshal(data, &env); err != nil {
		return nil, fmt.Errorf("failed to decode realtime envelope: %w", err)
	}
	return &env, nil
}

// UnmarshalPayload decodes env.Payload into dst based on the caller's type.
func (env *RealtimeEnvelope) UnmarshalPayload(dst interface{}) error {
	if env == nil {
		return fmt.Errorf("nil envelope")
	}
	if len(env.Payload) == 0 {
		return nil
	}
	return json.Unmarshal(env.Payload, dst)
}
