# Unofficial Flaro Client Go SDK

A Go SDK for interacting with the Flaro social media app API.

> [!NOTE]
> This SDK is an independently developed, reverse-engineered client that replicates the publicly observable HTTP(S) behavior of Flaro for interoperability and developer convenience. 
> It is not endorsed, sponsored, or supported by the Flaro developers and may lack features compared to the offcial Flaro app.
>
> Use this SDK only with accounts you own and in compliance with Flaro’s Terms of Service and applicable law. This project 
> does not provide any API keys, nor does it bypass server authorizations or paid features and may break at any time, you have been warned.

## Features

- User authentication (sign up and sign in)
- PKCS code challenge generation for secure authentication
- Type-safe API responses
- Comprehensive error handling

## Planned features

- [x] WebSocket support (experimental, behind build tag `realtime`)

## Installation

```bash
go get github.com/matu6968/flaro-client-go-sdk
```

## Usage

### Basic Setup
> [!NOTE]
> The SDK requires an API key to be provided. 
> This project does not provide such keys, you must obtain the API key from your app to use this client.
> You can either pass it directly or use an environment variable.
> Instructions to obtain such API key are in [OBTAIN-API-KEY.md](docs/OBTAIN-API-KEY.md) 
> In this file it also includes any disclosure request timelines.

#### Option 1: Using Environment Variable (Recommended)

Set the `FLARO_API_KEY` environment variable:

```bash
export FLARO_API_KEY="your-api-key-here"
```

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/matu6968/flaro-client-go-sdk"
)

func main() {
    // Create a new Flaro client using API key from environment variable
    client, err := flaro.NewClientFromEnv()
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
    
    // Sign up a new user
    authResp, err := client.SignUp("user@example.com", "password123")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("User signed up: %s\n", authResp.User.Email)
}
```

#### Option 2: Passing API Key Directly

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/matu6968/flaro-client-go-sdk"
)

func main() {
    // Create a new Flaro client with API key
    client := flaro.NewClient("your-api-key-here")
    
    // Sign up a new user
    authResp, err := client.SignUp("user@example.com", "password123")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("User signed up: %s\n", authResp.User.Email)
}
```

#### Option 3: Custom Configuration

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/matu6968/flaro-client-go-sdk"
)

func main() {
    // Create a new Flaro client with custom base URL and API key
    client := flaro.NewClientWithOptions("https://custom-api-url.com", "your-api-key-here")
    
    // Sign up a new user
    authResp, err := client.SignUp("user@example.com", "password123")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("User signed up: %s\n", authResp.User.Email)
}
```

### Authentication

#### Sign Up

```go
authResp, err := client.SignUp("user@example.com", "password123")
if err != nil {
    log.Printf("Sign up failed: %v", err)
    return
}

fmt.Printf("Access Token: %s\n", authResp.AccessToken)
fmt.Printf("User ID: %s\n", authResp.User.ID)
```

#### Sign In

```go
authResp, err := client.SignIn("user@example.com", "password123")
if err != nil {
    log.Printf("Sign in failed: %v", err)
    return
}

fmt.Printf("Access Token: %s\n", authResp.AccessToken)
```

#### Get Posts

```go
// Get posts with default pagination (20 posts, offset 0)
posts, err := client.GetPosts(accessToken, nil)
if err != nil {
    log.Fatal(err)
}

// Get posts with custom pagination
posts, err := client.GetPosts(accessToken, &flaro.PostsQueryParams{
    Select: "*",
    Order:  "created_at.desc.nullslast",
    Offset: 0,
    Limit:  10,
})
```

#### Get User Profile

```go
// Get a user's profile by user ID
user, err := client.GetUser(accessToken, "user-id-here")
if err != nil {
    log.Fatal(err)
}
```

#### Get User Posts

```go
// Get posts from a specific user
userPosts, err := client.GetUserPosts(accessToken, "user-id-here")
if err != nil {
    log.Fatal(err)
}

for _, post := range userPosts {
    fmt.Printf("Post: %s\n", post.Content)
}
```

#### Get Following

```go
// Get users that a specific user follows
following, err := client.GetFollowing(accessToken, "user-id-here")
if err != nil {
    log.Fatal(err)
}

for _, follow := range following {
    fmt.Printf("Following: %s (@%s)\n", follow.Users.DisplayName, follow.Users.Username)
}
```

#### Like/Unlike Posts

```go
// Like a post
err := client.LikePost(accessToken, "post-id-here", "user-id-here", true)
if err != nil {
    log.Fatal(err)
}

// Unlike a post
err = client.LikePost(accessToken, "post-id-here", "user-id-here", false)
if err != nil {
    log.Fatal(err)
}
```

#### Comments

```go
// Get comments for a post
comments, err := client.GetComments(accessToken, "post-id-here")
if err != nil {
    log.Fatal(err)
}

// Get comment count
count, err := client.GetCommentCount(accessToken, "post-id-here")
if err != nil {
    log.Fatal(err)
}

// Post a new comment
comment, err := client.PostComment(accessToken, "post-id-here", "user-id-here", "Great post!", nil)
if err != nil {
    log.Fatal(err)
}

// Reply to a comment
reply, err := client.PostComment(accessToken, "post-id-here", "user-id-here", "Nice comment!", &parentCommentID)
if err != nil {
    log.Fatal(err)
}
```

#### Delete Comments

```go
// Delete a comment
err := client.DeleteComment(accessToken, "comment-id-here")
if err != nil {
    log.Fatal(err)
}
```

#### Upload Images

```go
// Read image file
imageData, err := os.ReadFile("image.jpg")
if err != nil {
    log.Fatal(err)
}

// Upload image
uploadResp, err := client.UploadImage(accessToken, imageData, 3600) // 1 hour cache
if err != nil {
    log.Fatal(err)
}

// Use the uploaded image URL in posts
imageURL := fmt.Sprintf("https://sb.flaroapp.pl/storage/v1/object/public/%s", uploadResp.Key)
```

#### Create Posts

```go
// Create a text post
err := client.CreatePost(accessToken, "user-id-here", "Hello world!", []string{})
if err != nil {
    log.Fatal(err)
}

// Create a post with images
err = client.CreatePost(accessToken, "user-id-here", "Check out this image!", []string{imageURL})
if err != nil {
    log.Fatal(err)
}
```

### Custom Configuration

```go
// Create client with custom options
client := flaro.NewClientWithOptions(
    "https://custom-api.example.com",
    "your-api-key"
)
```

## API Reference

### Client

#### `NewClient(apiKey string) *Client`
Creates a new Flaro API client with the provided API key.

#### `NewClientFromEnv() (*Client, error)`
Creates a new Flaro API client using API key from the `FLARO_API_KEY` environment variable. Returns an error if the environment variable is not set.

#### `NewClientWithOptions(baseURL, apiKey string) *Client`
Creates a new Flaro API client with custom base URL and API key.

### Authentication Methods

#### `SignUp(email, password string) (*AuthResponse, error)`
Creates a new user account. Automatically generates PKCS code challenge for secure authentication.

#### `SignIn(email, password string) (*AuthResponse, error)`
Authenticates an existing user and returns access tokens.

### Social Media Methods

#### `GetPosts(accessToken string, params *PostsQueryParams) ([]Post, error)`
Retrieves posts with optional pagination. If params is nil, uses default values (limit: 20, offset: 0).

#### `GetFollowing(accessToken, followerID string) ([]Follow, error)`
Retrieves users that a specific user follows.

#### `GetUser(accessToken, userID string) (*UserProfile, error)`
Retrieves a specific user's profile by user ID.

#### `GetUserPosts(accessToken, userID string) ([]Post, error)`
Retrieves posts from a specific user.

#### `LikePost(accessToken, postID, userID string, isLiked bool) error`
Adds or removes a like from a post.

#### `GetComments(accessToken, postID string) ([]Comment, error)`
Retrieves comments for a specific post.

#### `GetCommentCount(accessToken, postID string) (int, error)`
Retrieves the count of comments for a specific post.

#### `PostComment(accessToken, postID, userID, content string, parentID *string) (*Comment, error)`
Creates a new comment on a post.

#### `DeleteComment(accessToken, commentID string) error`
Deletes a comment by its ID.

#### `UploadImage(accessToken string, imageData []byte, cacheControl int) (*ImageUploadResponse, error)`
Uploads an image for use in posts.

#### `CreatePost(accessToken, userID, content string, mediaURLs []string) error`
Creates a new post.

#### `RefreshToken(refreshToken string) (*AuthResponse, error)`
Refreshes an access token using a refresh token.

#### `GetReels(accessToken string) ([]Reel, error)`
Retrieves all reels (video posts).

#### `GetReelByID(accessToken, reelID string) ([]Reel, error)`
Retrieves a specific reel by its ID.

#### `GetReelComments(accessToken, reelID string) ([]Comment, error)`
Retrieves comments for a specific reel.

#### `GetSystemMessages(accessToken, userID string) ([]SystemMessage, error)`
Retrieves system messages for a user (includes both user-specific and global messages).

#### `SearchUsers(accessToken, username string) ([]SearchUser, error)`
Searches for users by username using partial matching.

#### `UpdateUserDetails(accessToken, userID string, bio, username, profilePicture *string) error`
Updates a user's bio, username, or profile picture. Note: Only one field can be updated at a time.

#### `DeletePost(accessToken, postID string) error`
Deletes a post by its ID.

#### `ReportUser(accessToken, userID, reportedBy string, postID, reelID *string, reason string) error`
Reports a user, post, or reel.

#### `ReportProblem(accessToken, userID, subject, content string) error`
Reports a problem to the system.

#### `ContactSupport(accessToken, userID, subject, content string) error`
Contacts support with a message.

#### `ChangePassword(accessToken, newPassword string) (*ChangePasswordResponse, error)`
Changes a user's password.

#### `UploadVideo(accessToken string, videoData []byte, cacheControl int) (*VideoUploadResponse, error)`
Uploads a video for use in reels.

#### `CreateReel(accessToken, userID, content, videoURL string) error`
Creates a new reel with the provided video URL.

### Realtime (Experimental)

This feature is experimental and disabled by default. Build with the `realtime` tag to enable it:

```bash
go build -tags realtime ./...
```

#### `NewRealtimeClient(apiKey string) *RealtimeClient`
Creates a realtime websocket client.

#### `(*RealtimeClient) Connect(ctx context.Context) error`
Connects to the realtime websocket endpoint.

#### `(*RealtimeClient) SubscribePostsForCreator(accessToken, creatorID string) error`
Subscribes to realtime events for a creator's posts.

#### `(*RealtimeClient) StartHeartbeat(ctx context.Context, interval time.Duration) error`
Sends phoenix heartbeats periodically (e.g., every 10s).

#### `(*RealtimeClient) ReadMessage() (*RealtimeEnvelope, error)`
Reads and decodes a single frame into a typed envelope.

#### `RealtimeEnvelope`
```go
type RealtimeEnvelope struct {
    Ref    *string         `json:"ref"`
    Event  string          `json:"event"`
    Payload json.RawMessage `json:"payload"`
    Topic  string          `json:"topic"`
}
```

#### `PhxReplyPayload`
```go
type PhxReplyPayload struct {
    Status   string          `json:"status"`
    Response json.RawMessage `json:"response"`
}
```

#### `SystemPayload`
```go
type SystemPayload struct {
    Message   string `json:"message"`
    Status    string `json:"status"`
    Extension string `json:"extension"`
    Channel   string `json:"channel"`
}
```

Usage example:
```go
env, err := rtc.ReadMessage()
if err != nil { /* handle */ }
switch env.Event {
case "phx_reply":
    var pr PhxReplyPayload
    _ = env.UnmarshalPayload(&pr)
    // inspect pr.Status/pr.Response
case "system":
    var sp SystemPayload
    _ = env.UnmarshalPayload(&sp)
    // inspect sp.Message
}
```

If built without the tag, all methods return an error stating realtime is disabled.

## Response Types

### AuthResponse
```go
type AuthResponse struct {
    AccessToken  string    `json:"access_token"`
    TokenType    string    `json:"token_type"`
    ExpiresIn    int       `json:"expires_in"`
    ExpiresAt    int64     `json:"expires_at"`
    RefreshToken string    `json:"refresh_token"`
    User         User      `json:"user"`
}
```

### User
```go
type User struct {
    ID                string                 `json:"id"`
    Email             string                 `json:"email"`
    EmailConfirmedAt  *time.Time             `json:"email_confirmed_at"`
    // ... other fields
}
```

### Post
```go
type Post struct {
    ID          string    `json:"id"`
    CreatorID   string    `json:"creator_id"`
    Content     string    `json:"content"`
    MediaURLs   []string  `json:"media_urls"`
    CreatedAt   time.Time `json:"created_at"`
    Tags        []string  `json:"tags"`
    Score       int       `json:"score"`
    Boost       int       `json:"boost"`
    IsPrivate   bool      `json:"is_private"`
    Likes       []string  `json:"likes"`
    // ... other fields
}
```

### UserProfile
```go
type UserProfile struct {
    UserID            string     `json:"user_id"`
    Username          string     `json:"username"`
    DisplayName       string     `json:"display_name"`
    Bio               string     `json:"bio"`
    ProfilePicture    string     `json:"profile_picture"`
    IsVerified        bool       `json:"is_verified"`
    // ... other fields
}
```

### Follow
```go
type Follow struct {
    FollowingID string      `json:"following_id"`
    Users       UserProfile `json:"users"`
}
```

### Comment
```go
type Comment struct {
    ID                string    `json:"id"`
    CreatedAt         time.Time `json:"created_at"`
    UserID            string    `json:"user_id"`
    Content           string    `json:"content"`
    Likes             []string  `json:"likes"`
    IsEdited          bool      `json:"is_edited"`
    ParentID          *string   `json:"parent_id"`
    PostID            string    `json:"post_id"`
    ReelID            *string   `json:"reel_id"`
    ReplyToUsername   *string   `json:"reply_to_username"`
}
```

### ImageUploadResponse
```go
type ImageUploadResponse struct {
    Key string `json:"Key"`
    ID  string `json:"Id"`
}
```

### Reel
```go
type Reel struct {
    ID          string    `json:"id"`
    CreatorID   string    `json:"creator_id"`
    Content     string    `json:"content"`
    Video       string    `json:"video"`
    CreatedAt   time.Time `json:"created_at"`
    Tags        []string  `json:"tags"`
    Score       int       `json:"score"`
    BoostEnds   *string   `json:"boost_ends"`
    Boost       int       `json:"boost"`
    IsPrivate   bool      `json:"is_private"`
    Location    *string   `json:"location"`
    Mentions    []string  `json:"mentions"`
    Comments    []string  `json:"comments"`
    Likes       []string  `json:"likes"`
}
```

### SystemMessage
```go
type SystemMessage struct {
    ID        string     `json:"id"`
    UserID    *string    `json:"user_id"` // null for global messages
    Content   string     `json:"content"`
    CreatedAt time.Time  `json:"created_at"`
}
```

### SearchUser
```go
type SearchUser struct {
    Username          string      `json:"username"`
    DisplayName       string      `json:"display_name"`
    Bio               string      `json:"bio"`
    ProfilePicture    string      `json:"profile_picture"`
    Website           *string     `json:"website"`
    IsPrivate         bool        `json:"is_private"`
    CreatedAt         string      `json:"created_at"`
    UsernameUpdatedAt *string     `json:"username_updated_at"`
    IsVerified        bool        `json:"is_verified"`
    LastSeen          string      `json:"last_seen"`
    UserID            string      `json:"user_id"`
    Ranks             interface{} `json:"ranks"`
    PremiumExpires    *string     `json:"premium_expires"`
}
```

### UpdateUserDetailsRequest
```go
type UpdateUserDetailsRequest struct {
    Bio            *string `json:"bio,omitempty"`             // Optional: update bio
    Username       *string `json:"username,omitempty"`        // Optional: update username
    ProfilePicture *string `json:"profile_picture,omitempty"` // Optional: update profile picture (can only update one at a time)
}
```

### ReportRequest
```go
type ReportRequest struct {
    CreatedAt  string  `json:"created_at"`
    UserID     string  `json:"user_id"`     // User ID to report
    PostID     *string `json:"post_id"`     // Optional: Post ID to report
    ReelID     *string `json:"reel_id"`     // Optional: Reel ID to report
    Reason     string  `json:"reason"`      // Report reason
    Status     string  `json:"status"`      // Report status
    ReportedBy string  `json:"reported_by"` // User ID of the reporter
}
```

### ProblemReportRequest
```go
type ProblemReportRequest struct {
    Subject   string `json:"subject"`   // Problem subject
    Content   string `json:"content"`   // Problem description
    UserID    string `json:"user_id"`   // User ID reporting the problem
    Status    string `json:"status"`    // Problem status
    CreatedAt string `json:"created_at"` // Current time in ISO 8601 format
}
```

### SupportRequest
```go
type SupportRequest struct {
    Subject   string `json:"subject"`   // Support subject
    Content   string `json:"content"`   // Support description
    UserID    string `json:"user_id"`   // User ID contacting support
    Status    string `json:"status"`    // Support status
    CreatedAt string `json:"created_at"` // Current time in ISO 8601 format
}
```

### ChangePasswordRequest
```go
type ChangePasswordRequest struct {
    Password string `json:"password"` // New password
}
```

### ChangePasswordResponse
```go
type ChangePasswordResponse struct {
    ID               string       `json:"id"`
    Aud              string       `json:"aud"`
    Role             string       `json:"role"`
    Email            string       `json:"email"`
    EmailConfirmedAt string       `json:"email_confirmed_at"`
    Phone            string       `json:"phone"`
    ConfirmedAt      string       `json:"confirmed_at"`
    LastSignInAt     string       `json:"last_sign_in_at"`
    AppMetadata      AppMetadata  `json:"app_metadata"`
    UserMetadata     UserMetadata `json:"user_metadata"`
    Identities       []Identity   `json:"identities"`
    CreatedAt        string       `json:"created_at"`
    UpdatedAt        string       `json:"updated_at"`
    IsAnonymous      bool         `json:"is_anonymous"`
}
```

### VideoUploadResponse
```go
type VideoUploadResponse struct {
    Key string `json:"Key"` // The key/path of the uploaded video
    Id  string `json:"Id"`  // The ID of the uploaded video
}
```

### CreateReelRequest
```go
type CreateReelRequest struct {
    CreatorID  string   `json:"creator_id"`  // Creator ID is the current user ID
    Content    string   `json:"content"`     // Description of the reel
    Video      string   `json:"video"`       // Video URL from upload
    CreatedAt  string   `json:"created_at"`  // Current date in ISO 8601 format
    Tags       []string `json:"tags"`        // Tags for the reel
    Score      int      `json:"score"`       // Score (typically 0)
    BoostEnds  *string  `json:"boost_ends"`  // Boost end time (typically null)
    Boost      int      `json:"boost"`       // Boost value (typically 1)
    IsPrivate  bool     `json:"is_private"`  // Whether the reel is private
    Location   *string  `json:"location"`    // Location (typically null)
    Mentions   []string `json:"mentions"`    // Mentions (typically empty)
    Comments   []string `json:"comments"`    // Comments (typically empty)
    Likes      []string `json:"likes"`       // Likes (typically empty)
}
```

## Error Handling

The SDK returns structured errors that implement the `error` interface:

```go
authResp, err := client.SignUp("user@example.com", "password123")
if err != nil {
    if apiErr, ok := err.(*flaro.APIError); ok {
        fmt.Printf("API Error: %s (Code: %s)\n", apiErr.Message, apiErr.Code)
    } else {
        fmt.Printf("Other error: %v\n", err)
    }
}
```

## Examples

See the `example/` directory for complete working examples.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Takedown notices

If you believe that there are security concerns within their API, contact support@flaroapp.pl with "Security: PoC" in the subject or visit the Flaro app with a logged in account --> profile icon --> Report a Problem and set subject to "Security: PoC" and explain any vulnerabilities found. If required by the Flaro developers I will take down the SDK upon request.