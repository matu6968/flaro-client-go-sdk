package flaro

import "time"

// SignUpRequest represents the request body for user sign up
type SignUpRequest struct {
	Email               string             `json:"email"`
	Password            string             `json:"password"`
	Data                interface{}        `json:"data"`
	GotrueMetaSecurity  GotrueMetaSecurity `json:"gotrue_meta_security"`
	CodeChallenge       string             `json:"code_challenge"`
	CodeChallengeMethod string             `json:"code_challenge_method"`
}

// SignInRequest represents the request body for user sign in
type SignInRequest struct {
	Email              string             `json:"email"`
	Password           string             `json:"password"`
	GotrueMetaSecurity GotrueMetaSecurity `json:"gotrue_meta_security"`
}

// GotrueMetaSecurity represents security metadata
type GotrueMetaSecurity struct {
	CaptchaToken *string `json:"captcha_token"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresAt    int64  `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

// User represents user information
type User struct {
	ID               string       `json:"id"`
	Aud              string       `json:"aud"`
	Role             string       `json:"role"`
	Email            string       `json:"email"`
	EmailConfirmedAt *time.Time   `json:"email_confirmed_at"`
	Phone            string       `json:"phone"`
	LastSignInAt     *time.Time   `json:"last_sign_in_at"`
	AppMetadata      AppMetadata  `json:"app_metadata"`
	UserMetadata     UserMetadata `json:"user_metadata"`
	Identities       []Identity   `json:"identities"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	IsAnonymous      bool         `json:"is_anonymous"`
}

// AppMetadata represents application metadata
type AppMetadata struct {
	Provider  string   `json:"provider"`
	Providers []string `json:"providers"`
}

// UserMetadata represents user metadata
type UserMetadata struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PhoneVerified bool   `json:"phone_verified"`
	Sub           string `json:"sub"`
}

// Identity represents user identity information
type Identity struct {
	IdentityID   string       `json:"identity_id"`
	ID           string       `json:"id"`
	UserID       string       `json:"user_id"`
	IdentityData IdentityData `json:"identity_data"`
	Provider     string       `json:"provider"`
	LastSignInAt *time.Time   `json:"last_sign_in_at"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Email        string       `json:"email"`
}

// IdentityData represents identity data
type IdentityData struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PhoneVerified bool   `json:"phone_verified"`
	Sub           string `json:"sub"`
}

// Post represents a social media post
type Post struct {
	ID        string     `json:"id"`
	CreatorID string     `json:"creator_id"`
	Content   string     `json:"content"`
	MediaURLs []string   `json:"media_urls"`
	CreatedAt time.Time  `json:"created_at"`
	Tags      []string   `json:"tags"`
	Score     int        `json:"score"`
	BoostEnds *time.Time `json:"boost_ends"`
	Boost     int        `json:"boost"`
	IsPrivate bool       `json:"is_private"`
	Location  *string    `json:"location"`
	Mentions  []string   `json:"mentions"`
	Comments  []string   `json:"comments"`
	Likes     []string   `json:"likes"`
}

// UserProfile represents a user profile
type UserProfile struct {
	UserID            string      `json:"user_id"`
	Username          string      `json:"username"`
	DisplayName       string      `json:"display_name"`
	Bio               string      `json:"bio"`
	ProfilePicture    string      `json:"profile_picture"`
	Website           *string     `json:"website"`
	IsPrivate         bool        `json:"is_private"`
	CreatedAt         string      `json:"created_at"`          // Changed to string to handle parsing
	UsernameUpdatedAt *string     `json:"username_updated_at"` // Changed to string to handle parsing
	IsVerified        bool        `json:"is_verified"`
	LastSeen          string      `json:"last_seen"`       // Changed to string to handle parsing
	Ranks             interface{} `json:"ranks"`           // Changed to interface{} to handle both null and array
	PremiumExpires    *string     `json:"premium_expires"` // Changed to string to handle parsing
}

// Follow represents a follow relationship
type Follow struct {
	FollowingID string      `json:"following_id"`
	Users       UserProfile `json:"users"`
}

// PostsQueryParams represents query parameters for posts endpoint
type PostsQueryParams struct {
	Select string
	Order  string
	Offset int
	Limit  int
}

// UserQueryParams represents query parameters for user endpoint
type UserQueryParams struct {
	Select string
	UserID string
}

// FollowsQueryParams represents query parameters for follows endpoint
type FollowsQueryParams struct {
	Select     string
	FollowerID string
}

// Comment represents a comment on a post
type Comment struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UserID          string    `json:"user_id"`
	Content         string    `json:"content"`
	Likes           []string  `json:"likes"`
	IsEdited        bool      `json:"is_edited"`
	ParentID        *string   `json:"parent_id"`
	PostID          string    `json:"post_id"`
	ReelID          *string   `json:"reel_id"`
	ReplyToUsername *string   `json:"reply_to_username"`
}

// CommentID represents just a comment ID (for counting)
type CommentID struct {
	ID string `json:"id"`
}

// LikePostRequest represents the request body for liking/unliking a post
type LikePostRequest struct {
	Likes []string `json:"likes"`
}

// PostCommentRequest represents the request body for posting a comment
type PostCommentRequest struct {
	PostID   string   `json:"post_id"`
	ReelID   *string  `json:"reel_id"`
	UserID   string   `json:"user_id"`
	Content  string   `json:"content"`
	ParentID *string  `json:"parent_id"`
	Likes    []string `json:"likes"`
}

// CommentsQueryParams represents query parameters for comments endpoint
type CommentsQueryParams struct {
	Select string
	PostID string
	Order  string
}

// ImageUploadResponse represents the response from image upload
type ImageUploadResponse struct {
	Key string `json:"Key"`
	ID  string `json:"Id"`
}

// CreatePostRequest represents the request body for creating a new post
type CreatePostRequest struct {
	CreatorID string   `json:"creator_id"`
	Content   string   `json:"content"`
	MediaURLs []string `json:"media_urls"`
	CreatedAt string   `json:"created_at"`
	Tags      []string `json:"tags"`
	Score     int      `json:"score"`
	BoostEnds *string  `json:"boost_ends"`
	Boost     int      `json:"boost"`
	IsPrivate bool     `json:"is_private"`
	Location  *string  `json:"location"`
	Mentions  []string `json:"mentions"`
	Comments  []string `json:"comments"`
	Likes     []string `json:"likes"`
}

// RefreshTokenRequest represents the request body for refreshing a token
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// Reel represents a reel (video post)
type Reel struct {
	ID        string    `json:"id"`
	CreatorID string    `json:"creator_id"`
	Content   string    `json:"content"`
	Video     string    `json:"video"`
	CreatedAt time.Time `json:"created_at"`
	Tags      []string  `json:"tags"`
	Score     int       `json:"score"`
	BoostEnds *string   `json:"boost_ends"`
	Boost     int       `json:"boost"`
	IsPrivate bool      `json:"is_private"`
	Location  *string   `json:"location"`
	Mentions  []string  `json:"mentions"`
	Comments  []string  `json:"comments"`
	Likes     []string  `json:"likes"`
}

// SystemMessage represents a system message
type SystemMessage struct {
	ID        string    `json:"id"`
	UserID    *string   `json:"user_id"` // null for global messages
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// SearchUser represents a user found in search results
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

// ReelsQueryParams represents query parameters for reels endpoint
type ReelsQueryParams struct {
	Select string
	Order  string
	ID     string
}

// ReelCommentsQueryParams represents query parameters for reel comments endpoint
type ReelCommentsQueryParams struct {
	Select string
	ReelID string
	Order  string
}

// SystemMessagesQueryParams represents query parameters for system messages endpoint
type SystemMessagesQueryParams struct {
	Select string
	UserID string
	Order  string
}

// SearchUsersQueryParams represents query parameters for user search endpoint
type SearchUsersQueryParams struct {
	Select   string
	Username string
}

// UpdateUserDetailsRequest represents the request body for updating user details
type UpdateUserDetailsRequest struct {
	Bio            *string `json:"bio,omitempty"`             // Optional: update bio
	Username       *string `json:"username,omitempty"`        // Optional: update username
	ProfilePicture *string `json:"profile_picture,omitempty"` // Optional: update profile picture (can only update one at a time)
}

// ReportRequest represents the request body for reporting users/posts/reels
type ReportRequest struct {
	CreatedAt  string  `json:"created_at"`
	UserID     string  `json:"user_id"`     // User ID to report
	PostID     *string `json:"post_id"`     // Optional: Post ID to report
	ReelID     *string `json:"reel_id"`     // Optional: Reel ID to report
	Reason     string  `json:"reason"`      // Report reason
	Status     string  `json:"status"`      // Report status
	ReportedBy string  `json:"reported_by"` // User ID of the reporter
}

// ProblemReportRequest represents the request body for reporting problems
type ProblemReportRequest struct {
	Subject   string `json:"subject"`    // Problem subject
	Content   string `json:"content"`    // Problem description
	UserID    string `json:"user_id"`    // User ID reporting the problem
	Status    string `json:"status"`     // Problem status
	CreatedAt string `json:"created_at"` // Current time in ISO 8601 format
}

// SupportRequest represents the request body for contacting support
type SupportRequest struct {
	Subject   string `json:"subject"`    // Support subject
	Content   string `json:"content"`    // Support description
	UserID    string `json:"user_id"`    // User ID contacting support
	Status    string `json:"status"`     // Support status
	CreatedAt string `json:"created_at"` // Current time in ISO 8601 format
}

// ChangePasswordRequest represents the request body for changing password
type ChangePasswordRequest struct {
	Password string `json:"password"` // New password
}

// ChangePasswordResponse represents the response from changing password
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

// VideoUploadResponse represents the response from uploading a video
type VideoUploadResponse struct {
	Key string `json:"Key"` // The key/path of the uploaded video
	Id  string `json:"Id"`  // The ID of the uploaded video
}

// CreateReelRequest represents the request body for creating a reel
type CreateReelRequest struct {
	CreatorID string   `json:"creator_id"` // Creator ID is the current user ID
	Content   string   `json:"content"`    // Description of the reel
	Video     string   `json:"video"`      // Video URL from upload
	CreatedAt string   `json:"created_at"` // Current date in ISO 8601 format
	Tags      []string `json:"tags"`       // Tags for the reel
	Score     int      `json:"score"`      // Score (typically 0)
	BoostEnds *string  `json:"boost_ends"` // Boost end time (typically null)
	Boost     int      `json:"boost"`      // Boost value (typically 1)
	IsPrivate bool     `json:"is_private"` // Whether the reel is private
	Location  *string  `json:"location"`   // Location (typically null)
	Mentions  []string `json:"mentions"`   // Mentions (typically empty)
	Comments  []string `json:"comments"`   // Comments (typically empty)
	Likes     []string `json:"likes"`      // Likes (typically empty)
}

// APIError represents an API error response
type APIError struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}
