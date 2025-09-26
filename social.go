package flaro

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// GetPosts retrieves posts with optional pagination
func (c *Client) GetPosts(accessToken string, params *PostsQueryParams) ([]Post, error) {
	// Set default parameters if not provided
	if params == nil {
		params = &PostsQueryParams{
			Select: "*",
			Order:  "created_at.desc.nullslast",
			Offset: 0,
			Limit:  20,
		}
	}

	// Build query parameters
	queryParams := url.Values{}
	queryParams.Set("select", params.Select)
	queryParams.Set("order", params.Order)
	queryParams.Set("offset", strconv.Itoa(params.Offset))
	queryParams.Set("limit", strconv.Itoa(params.Limit))

	endpoint := "/rest/v1/posts?" + queryParams.Encode()

	// Make the request
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to make get posts request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get posts failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, &apiErr
	}

	// Parse successful response
	var posts []Post
	if err := json.Unmarshal(body, &posts); err != nil {
		return nil, fmt.Errorf("failed to parse posts response: %w", err)
	}

	return posts, nil
}

// GetFollowing retrieves users that a specific user follows
func (c *Client) GetFollowing(accessToken, followerID string) ([]Follow, error) {
	// Build query parameters
	queryParams := url.Values{}
	queryParams.Set("select", "following_id,users!follows_following_id_fkey(*)")
	queryParams.Set("follower_id", "eq."+followerID)

	endpoint := "/rest/v1/follows?" + queryParams.Encode()

	// Make the request
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to make get following request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get following failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, &apiErr
	}

	// Parse successful response
	var follows []Follow
	if err := json.Unmarshal(body, &follows); err != nil {
		return nil, fmt.Errorf("failed to parse following response: %w", err)
	}

	return follows, nil
}

// GetUser retrieves a specific user's profile by user ID
func (c *Client) GetUser(accessToken, userID string) (*UserProfile, error) {
	// Build query parameters
	queryParams := url.Values{}
	queryParams.Set("select", "*")
	queryParams.Set("user_id", "eq."+userID)

	endpoint := "/rest/v1/users?" + queryParams.Encode()

	// Make the request
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to make get user request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get user failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, &apiErr
	}

	// Parse successful response - it returns an array with one user
	var userProfiles []UserProfile
	if err := json.Unmarshal(body, &userProfiles); err != nil {
		return nil, fmt.Errorf("failed to parse user response: %w", err)
	}

	if len(userProfiles) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &userProfiles[0], nil
}

// GetUserPosts retrieves posts from a specific user
func (c *Client) GetUserPosts(accessToken, userID string) ([]Post, error) {
	// Build query parameters
	queryParams := url.Values{}
	queryParams.Set("select", "*")
	queryParams.Set("creator_id", "eq."+userID)

	endpoint := "/rest/v1/posts?" + queryParams.Encode()

	// Make the request
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to make get user posts request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get user posts failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, &apiErr
	}

	// Parse successful response
	var posts []Post
	if err := json.Unmarshal(body, &posts); err != nil {
		return nil, fmt.Errorf("failed to parse user posts response: %w", err)
	}

	return posts, nil
}

// LikePost adds or removes a like from a post
func (c *Client) LikePost(accessToken, postID string, userID string, isLiked bool) error {
	// First, get the current post to see existing likes
	posts, err := c.GetPosts(accessToken, &PostsQueryParams{
		Select: "*",
		Order:  "created_at.desc.nullslast",
		Offset: 0,
		Limit:  1000, // Get more posts to find the specific one
	})
	if err != nil {
		return fmt.Errorf("failed to get posts: %w", err)
	}

	// Find the specific post
	var currentPost *Post
	for _, post := range posts {
		if post.ID == postID {
			currentPost = &post
			break
		}
	}
	if currentPost == nil {
		return fmt.Errorf("post not found")
	}

	// Create new likes array
	newLikes := make([]string, 0, len(currentPost.Likes))
	userFound := false

	// Copy existing likes, removing user if they already liked
	for _, like := range currentPost.Likes {
		if like != userID {
			newLikes = append(newLikes, like)
		} else {
			userFound = true
		}
	}

	// Add user if they want to like and weren't already liking
	if isLiked && !userFound {
		newLikes = append(newLikes, userID)
	}

	// Build request
	req := LikePostRequest{
		Likes: newLikes,
	}

	// Build endpoint
	endpoint := "/rest/v1/posts?id=eq." + postID

	// Make the request
	resp, err := c.makeAuthenticatedRequest("PATCH", endpoint, req, accessToken)
	if err != nil {
		return fmt.Errorf("failed to make like post request: %w", err)
	}
	defer resp.Body.Close()

	// Check for errors
	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("like post failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// GetComments retrieves comments for a specific post
func (c *Client) GetComments(accessToken, postID string) ([]Comment, error) {
	// Build query parameters
	queryParams := url.Values{}
	queryParams.Set("select", "*")
	queryParams.Set("post_id", "eq."+postID)
	queryParams.Set("order", "created_at.asc.nullslast")

	endpoint := "/rest/v1/comments?" + queryParams.Encode()

	// Make the request
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to make get comments request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get comments failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, &apiErr
	}

	// Parse successful response
	var comments []Comment
	if err := json.Unmarshal(body, &comments); err != nil {
		return nil, fmt.Errorf("failed to parse comments response: %w", err)
	}

	return comments, nil
}

// GetCommentCount retrieves the count of comments for a specific post
func (c *Client) GetCommentCount(accessToken, postID string) (int, error) {
	// Build query parameters
	queryParams := url.Values{}
	queryParams.Set("select", "id")
	queryParams.Set("post_id", "eq."+postID)

	endpoint := "/rest/v1/comments?" + queryParams.Encode()

	// Make the request
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return 0, fmt.Errorf("failed to make get comment count request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return 0, fmt.Errorf("get comment count failed with status %d: %s", resp.StatusCode, string(body))
		}
		return 0, &apiErr
	}

	// Parse successful response
	var commentIDs []CommentID
	if err := json.Unmarshal(body, &commentIDs); err != nil {
		return 0, fmt.Errorf("failed to parse comment count response: %w", err)
	}

	return len(commentIDs), nil
}

// PostComment creates a new comment on a post
func (c *Client) PostComment(accessToken, postID, userID, content string, parentID *string) (*Comment, error) {
	// Build request
	req := PostCommentRequest{
		PostID:   postID,
		ReelID:   nil, // Set to nil for posts
		UserID:   userID,
		Content:  content,
		ParentID: parentID,
		Likes:    []string{}, // Start with empty likes
	}

	// Build endpoint
	endpoint := "/rest/v1/comments?select=*"

	// Make the request
	resp, err := c.makeAuthenticatedRequest("POST", endpoint, req, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to make post comment request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 201 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("post comment failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, &apiErr
	}

	// Check if response body is empty
	if len(body) == 0 {
		// If the API returns empty body but 201 status, the comment was created successfully
		// We can return a minimal comment object or just return success
		return &Comment{
			PostID:   postID,
			UserID:   userID,
			Content:  content,
			ParentID: parentID,
			Likes:    []string{},
		}, nil
	}

	// Parse successful response
	var comment Comment
	if err := json.Unmarshal(body, &comment); err != nil {
		return nil, fmt.Errorf("failed to parse post comment response: %w", err)
	}

	return &comment, nil
}

// DeleteComment deletes a comment by its ID
func (c *Client) DeleteComment(accessToken, commentID string) error {
	// Build endpoint
	endpoint := "/rest/v1/comments?id=eq." + commentID

	// Make the request
	resp, err := c.makeAuthenticatedRequest("DELETE", endpoint, nil, accessToken)
	if err != nil {
		return fmt.Errorf("failed to make delete comment request: %w", err)
	}
	defer resp.Body.Close()

	// Check for errors
	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("delete comment failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// UploadImage uploads an image for use in posts
func (c *Client) UploadImage(accessToken string, imageData []byte, cacheControl int) (*ImageUploadResponse, error) {
	// Generate timestamp for filename
	timestamp := time.Now().UnixMilli()
	filename := fmt.Sprintf("%d-0", timestamp)

	// Build endpoint
	endpoint := fmt.Sprintf("/storage/v1/object/post-images/uploads/%s", filename)

	// Create multipart form data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add image data
	part, err := writer.CreateFormFile("", "image")
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}
	_, err = part.Write(imageData)
	if err != nil {
		return nil, fmt.Errorf("failed to write image data: %w", err)
	}

	// Add cache control
	err = writer.WriteField("CacheControl", fmt.Sprintf("%d", cacheControl))
	if err != nil {
		return nil, fmt.Errorf("failed to write cache control: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	// Create request
	req, err := http.NewRequest("POST", c.baseURL+endpoint, &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("apikey", c.apiKey)
	req.Header.Set("authorization", "Bearer "+accessToken)

	// Make the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make upload image request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("upload image failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse successful response
	var uploadResp ImageUploadResponse
	if err := json.Unmarshal(body, &uploadResp); err != nil {
		return nil, fmt.Errorf("failed to parse upload response: %w", err)
	}

	return &uploadResp, nil
}

// CreatePost creates a new post
func (c *Client) CreatePost(accessToken, userID, content string, mediaURLs []string) error {
	// Build request
	req := CreatePostRequest{
		CreatorID: userID,
		Content:   content,
		MediaURLs: mediaURLs,
		CreatedAt: time.Now().Format("2006-01-02T15:04:05.000000"),
		Tags:      []string{},
		Score:     0,
		BoostEnds: nil,
		Boost:     1,
		IsPrivate: false,
		Location:  nil,
		Mentions:  []string{},
		Comments:  []string{},
		Likes:     []string{},
	}

	// Build endpoint
	endpoint := "/rest/v1/posts?select=*"

	// Make the request
	resp, err := c.makeAuthenticatedRequest("POST", endpoint, req, accessToken)
	if err != nil {
		return fmt.Errorf("failed to make create post request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 201 {
		return fmt.Errorf("create post failed with status %d: %s", resp.StatusCode, string(body))
	}

	// The API returns no response body for successful post creation
	return nil
}

// GetReels retrieves all reels
func (c *Client) GetReels(accessToken string) ([]Reel, error) {
	params := ReelsQueryParams{
		Select: "*",
		Order:  "created_at.desc.nullslast",
	}

	queryParams := url.Values{}
	if params.Select != "" {
		queryParams.Add("select", params.Select)
	}
	if params.Order != "" {
		queryParams.Add("order", params.Order)
	}

	endpoint := "/rest/v1/reels?" + queryParams.Encode()
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get reels: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get reels failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("get reels failed: %s", apiErr.Message)
	}

	var reels []Reel
	if err := json.Unmarshal(body, &reels); err != nil {
		return nil, fmt.Errorf("failed to parse reels response: %w", err)
	}

	return reels, nil
}

// GetReelByID retrieves a specific reel by its ID
func (c *Client) GetReelByID(accessToken, reelID string) ([]Reel, error) {
	params := ReelsQueryParams{
		Select: "*",
		ID:     reelID,
	}

	queryParams := url.Values{}
	if params.Select != "" {
		queryParams.Add("select", params.Select)
	}
	if params.ID != "" {
		queryParams.Add("id", "eq."+params.ID)
	}

	endpoint := "/rest/v1/reels?" + queryParams.Encode()
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get reel: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get reel failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("get reel failed: %s", apiErr.Message)
	}

	var reels []Reel
	if err := json.Unmarshal(body, &reels); err != nil {
		return nil, fmt.Errorf("failed to parse reel response: %w", err)
	}

	return reels, nil
}

// GetReelComments retrieves comments for a specific reel
func (c *Client) GetReelComments(accessToken, reelID string) ([]Comment, error) {
	params := ReelCommentsQueryParams{
		Select: "*",
		ReelID: reelID,
		Order:  "created_at.asc.nullslast",
	}

	queryParams := url.Values{}
	if params.Select != "" {
		queryParams.Add("select", params.Select)
	}
	if params.ReelID != "" {
		queryParams.Add("reel_id", "eq."+params.ReelID)
	}
	if params.Order != "" {
		queryParams.Add("order", params.Order)
	}

	endpoint := "/rest/v1/comments?" + queryParams.Encode()
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get reel comments: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get reel comments failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("get reel comments failed: %s", apiErr.Message)
	}

	var comments []Comment
	if err := json.Unmarshal(body, &comments); err != nil {
		return nil, fmt.Errorf("failed to parse reel comments response: %w", err)
	}

	return comments, nil
}

// GetSystemMessages retrieves system messages for a user
func (c *Client) GetSystemMessages(accessToken, _ string) ([]SystemMessage, error) {
	// Updated schema: no user_id column on system_messages. Fetch all, newest first.
	query := url.Values{}
	query.Add("select", "*")
	query.Add("order", "created_at.desc.nullslast")

	endpoint := "/rest/v1/system_messages?" + query.Encode()
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get system messages: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get system messages failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("get system messages failed: %s", apiErr.Message)
	}

	// Unmarshal into detailed form then map to legacy SystemMessage type
	var detailed []SystemMessageDetail
	if err := json.Unmarshal(body, &detailed); err != nil {
		return nil, fmt.Errorf("failed to parse system messages response: %w", err)
	}

	out := make([]SystemMessage, 0, len(detailed))
	for _, d := range detailed {
		// created_at example: 2025-09-25T17:33:37+00:00
		t, err := time.Parse(time.RFC3339, d.CreatedAt)
		if err != nil {
			// try layout with explicit colon offset already covered by RFC3339; if it fails, ignore parse
			t = time.Time{}
		}
		out = append(out, SystemMessage{
			ID:        strconv.Itoa(d.ID),
			UserID:    nil,
			Content:   d.Content,
			CreatedAt: t,
		})
	}
	return out, nil
}

// GetLatestSystemMessages fetches latest system messages with optional limit
func (c *Client) GetLatestSystemMessages(accessToken string, limit int) ([]SystemMessageDetail, error) {
	query := url.Values{}
	query.Add("select", "*")
	query.Add("order", "created_at.desc.nullslast")
	if limit > 0 {
		query.Add("limit", strconv.Itoa(limit))
	}

	endpoint := "/rest/v1/system_messages?" + query.Encode()
	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest system messages: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get latest system messages failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("get latest system messages failed: %s", apiErr.Message)
	}

	// API returns an array
	var messages []SystemMessageDetail
	if err := json.Unmarshal(body, &messages); err != nil {
		return nil, fmt.Errorf("failed to parse latest system messages response: %w", err)
	}

	return messages, nil
}

// SearchUsers searches for users by username
func (c *Client) SearchUsers(accessToken, username string) ([]SearchUser, error) {
	params := SearchUsersQueryParams{
		Select:   "*",
		Username: username,
	}

	// Build the query string manually to avoid double encoding
	var queryParts []string
	if params.Select != "" {
		queryParts = append(queryParts, "select="+url.QueryEscape(params.Select))
	}
	if params.Username != "" {
		// Build the ilike query with proper URL encoding: ilike.%username%
		// The % characters need to be encoded as %25 in the URL
		usernameQuery := "ilike.%25" + params.Username + "%25"
		queryParts = append(queryParts, "username="+usernameQuery)
	}

	var endpoint string
	if len(queryParts) > 0 {
		endpoint = "/rest/v1/users?" + strings.Join(queryParts, "&")
	} else {
		endpoint = "/rest/v1/users"
	}

	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to search users: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("search users failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("search users failed: %s", apiErr.Message)
	}

	var users []SearchUser
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, fmt.Errorf("failed to parse search users response: %w", err)
	}

	return users, nil
}

// UpdateUserDetails updates a user's bio, username, or profile picture
func (c *Client) UpdateUserDetails(accessToken, userID string, bio, username, profilePicture *string) error {
	// Validate that only one field is being updated at a time
	fieldCount := 0
	if bio != nil {
		fieldCount++
	}
	if username != nil {
		fieldCount++
	}
	if profilePicture != nil {
		fieldCount++
	}

	if fieldCount == 0 {
		return fmt.Errorf("must provide either bio, username, or profile_picture to update")
	}
	if fieldCount > 1 {
		return fmt.Errorf("cannot update multiple fields at the same time")
	}

	req := UpdateUserDetailsRequest{
		Bio:            bio,
		Username:       username,
		ProfilePicture: profilePicture,
	}

	endpoint := "/rest/v1/users?user_id=eq." + userID
	resp, err := c.makeAuthenticatedRequest("PATCH", endpoint, req, accessToken)
	if err != nil {
		return fmt.Errorf("failed to update user details: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 204 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("update user details failed with status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("update user details failed: %s", apiErr.Message)
	}

	// The API returns no response body for successful updates (204 status)
	return nil
}

// DeletePost deletes a post by its ID
func (c *Client) DeletePost(accessToken, postID string) error {
	endpoint := "/rest/v1/posts?id=eq." + postID
	resp, err := c.makeAuthenticatedRequest("DELETE", endpoint, nil, accessToken)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 204 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("delete post failed with status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("delete post failed: %s", apiErr.Message)
	}

	// The API returns no response body for successful deletion (204 status)
	return nil
}

// ReportUser reports a user, post, or reel
func (c *Client) ReportUser(accessToken, userID, reportedBy string, postID, reelID *string, reason string) error {
	req := ReportRequest{
		CreatedAt:  time.Now().Format("2006-01-02T15:04:05.000000"),
		UserID:     userID,
		PostID:     postID,
		ReelID:     reelID,
		Reason:     reason,
		Status:     "UNREAD",
		ReportedBy: reportedBy,
	}

	endpoint := "/rest/v1/reports"
	resp, err := c.makeAuthenticatedRequest("POST", endpoint, req, accessToken)
	if err != nil {
		return fmt.Errorf("failed to report user: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 201 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("report user failed with status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("report user failed: %s", apiErr.Message)
	}

	// The API returns no response body for successful report creation (201 status)
	return nil
}

// MarkSystemMessageAsRead appends the caller userID to read_by for a system message
func (c *Client) MarkSystemMessageAsRead(accessToken string, systemMessageID int, currentReadBy []string, userID string) error {
	// ensure userID is included exactly once
	exists := false
	for _, id := range currentReadBy {
		if id == userID {
			exists = true
			break
		}
	}
	updated := currentReadBy
	if !exists {
		updated = append(updated, userID)
	}

	req := MarkSystemMessageReadRequest{ReadBy: updated}
	endpoint := "/rest/v1/system_messages?id=eq." + strconv.Itoa(systemMessageID)
	resp, err := c.makeAuthenticatedRequest("PATCH", endpoint, req, accessToken)
	if err != nil {
		return fmt.Errorf("failed to mark system message as read: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 204 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("mark system message as read failed with status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("mark system message as read failed: %s", apiErr.Message)
	}
	return nil
}

// ReportProblem reports a problem to the system
func (c *Client) ReportProblem(accessToken, userID, subject, content string) error {
	req := ProblemReportRequest{
		Subject:   subject,
		Content:   content,
		UserID:    userID,
		Status:    "open",
		CreatedAt: time.Now().Format("2006-01-02T15:04:05.000000"),
	}

	endpoint := "/rest/v1/problems"
	resp, err := c.makeAuthenticatedRequest("POST", endpoint, req, accessToken)
	if err != nil {
		return fmt.Errorf("failed to report problem: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 201 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("report problem failed with status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("report problem failed: %s", apiErr.Message)
	}

	// The API returns no response body for successful problem report (201 status)
	return nil
}

// ContactSupport contacts support with a message
func (c *Client) ContactSupport(accessToken, userID, subject, content string) error {
	req := SupportRequest{
		Subject:   subject,
		Content:   content,
		UserID:    userID,
		Status:    "open",
		CreatedAt: time.Now().Format("2006-01-02T15:04:05.000000"),
	}

	endpoint := "/rest/v1/support"
	resp, err := c.makeAuthenticatedRequest("POST", endpoint, req, accessToken)
	if err != nil {
		return fmt.Errorf("failed to contact support: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 201 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("contact support failed with status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("contact support failed: %s", apiErr.Message)
	}

	// The API returns no response body for successful support contact (201 status)
	return nil
}

// CreateUserProfile creates a profile row for a newly created account
func (c *Client) CreateUserProfile(accessToken, userID, username string) error {
	now := time.Now().Format("2006-01-02T15:04:05.000000")
	req := CreateUserProfileRequest{
		UserID:            userID,
		Username:          username,
		DisplayName:       username,
		Bio:               "",
		ProfilePicture:    "https://i.postimg.cc/660K5Hrr/pfp.png",
		Website:           nil,
		IsPrivate:         false,
		CreatedAt:         now,
		UsernameUpdatedAt: nil,
		IsVerified:        false,
		LastSeen:          now,
		Ranks:             nil,
		PremiumExpires:    nil,
	}

	endpoint := "/rest/v1/users"
	resp, err := c.makeAuthenticatedRequest("POST", endpoint, req, accessToken)
	if err != nil {
		return fmt.Errorf("failed to create user profile: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 201 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("create user profile failed with status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("create user profile failed: %s", apiErr.Message)
	}

	return nil
}

// UploadVideo uploads a video for use in reels
func (c *Client) UploadVideo(accessToken string, videoData []byte, cacheControl int) (*VideoUploadResponse, error) {
	// Generate timestamp for filename
	timestamp := time.Now().UnixMilli()

	// Create multipart form data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add video data field (empty name field as per API docs)
	part, err := writer.CreateFormFile("", "video")
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	_, err = part.Write(videoData)
	if err != nil {
		return nil, fmt.Errorf("failed to write video data: %w", err)
	}

	// Add CacheControl field
	err = writer.WriteField("CacheControl", fmt.Sprintf("%d", cacheControl))
	if err != nil {
		return nil, fmt.Errorf("failed to write CacheControl field: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Create request
	endpoint := fmt.Sprintf("/storage/v1/object/reel-videos/uploads/%d", timestamp)
	req, err := http.NewRequest("POST", c.baseURL+endpoint, &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("apikey", c.apiKey)
	req.Header.Set("user-agent", "Dart/3.9 (dart:io)") // previously was flaro-go-sdk/1.0.0
	req.Header.Set("x-client-info", "supabase-flutter/2.10.1")
	req.Header.Set("authorization", "Bearer "+accessToken)

	// Make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to upload video: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("upload video failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("upload video failed: %s", apiErr.Message)
	}

	var uploadResp VideoUploadResponse
	if err := json.Unmarshal(body, &uploadResp); err != nil {
		return nil, fmt.Errorf("failed to parse upload video response: %w", err)
	}

	return &uploadResp, nil
}

// CreateReel creates a new reel
func (c *Client) CreateReel(accessToken, userID, content, videoURL string) error {
	req := CreateReelRequest{
		CreatorID: userID,
		Content:   content,
		Video:     videoURL,
		CreatedAt: time.Now().Format("2006-01-02T15:04:05.000000"),
		Tags:      []string{},
		Score:     0,
		BoostEnds: nil,
		Boost:     1,
		IsPrivate: false,
		Location:  nil,
		Mentions:  []string{},
		Comments:  []string{},
		Likes:     []string{},
	}

	endpoint := "/rest/v1/reels"
	resp, err := c.makeAuthenticatedRequest("POST", endpoint, req, accessToken)
	if err != nil {
		return fmt.Errorf("failed to create reel: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 201 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("create reel failed with status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("create reel failed: %s", apiErr.Message)
	}

	// The API returns no response body for successful reel creation (201 status)
	return nil
}

// SendGlobalMessage sends a message to the Global Channel
func (c *Client) SendGlobalMessage(accessToken, senderID, content string) error {
	req := SendGlobalMessageRequest{
		Content:   content,
		SenderID:  senderID,
		CreatedAt: time.Now().Format("2006-01-02T15:04:05.000000"),
	}

	endpoint := "/rest/v1/messages"
	resp, err := c.makeAuthenticatedRequest("POST", endpoint, req, accessToken)
	if err != nil {
		return fmt.Errorf("failed to send global message: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 201 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("send global message failed with status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("send global message failed: %s", apiErr.Message)
	}
	return nil
}

// GetGlobalMessages retrieves messages from the Global Channel
func (c *Client) GetGlobalMessages(accessToken string) ([]GlobalMessage, error) {
	query := url.Values{}
	query.Add("select", "*")
	query.Add("order", "created_at.asc.nullslast")
	endpoint := "/rest/v1/messages?" + query.Encode()

	resp, err := c.makeAuthenticatedRequest("GET", endpoint, nil, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get global messages: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("get global messages failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("get global messages failed: %s", apiErr.Message)
	}

	var messages []GlobalMessage
	if err := json.Unmarshal(body, &messages); err != nil {
		return nil, fmt.Errorf("failed to parse global messages response: %w", err)
	}

	return messages, nil
}
