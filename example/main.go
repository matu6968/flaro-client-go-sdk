package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matu6968/flaro-client-go-sdk"
)

func main() {
	// Create a new Flaro client using API key from environment variable
	client, err := flaro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example: Sign up a new user
	//fmt.Println("=== Sign Up Example ===")
	email := os.Getenv("FLARO_EMAIL")
	password := os.Getenv("FLARO_PASSWORD")

	//authResp, err := client.SignUp(email, password)
	//if err != nil {
	//	log.Printf("Sign up failed: %v", err)
	//} else {
	//	fmt.Printf("Sign up successful!\n")
	//	fmt.Printf("Access Token: %s\n", authResp.AccessToken)
	//	fmt.Printf("User ID: %s\n", authResp.User.ID)
	//	fmt.Printf("User Email: %s\n", authResp.User.Email)
	//}

	// Example: Sign in an existing user
	fmt.Println("\n=== Sign In Example ===")
	authResp, err := client.SignIn(email, password)
	if err != nil {
		log.Printf("Sign in failed: %v", err)
		return
	}

	fmt.Printf("Sign in successful!\n")
	fmt.Printf("Access Token: %s\n", authResp.AccessToken)
	fmt.Printf("User ID: %s\n", authResp.User.ID)
	fmt.Printf("User Email: %s\n", authResp.User.Email)

	// Example: Get current user profile (using user ID from auth response)
	fmt.Println("\n=== Get Current User Example ===")
	currentUser, err := client.GetUser(authResp.AccessToken, authResp.User.ID)
	if err != nil {
		log.Printf("Get current user failed: %v", err)
	} else {
		fmt.Printf("Current User: %s (@%s)\n", currentUser.DisplayName, currentUser.Username)
		fmt.Printf("Bio: %s\n", currentUser.Bio)
		fmt.Printf("Verified: %t\n", currentUser.IsVerified)
	}

	// Example: Get posts
	fmt.Println("\n=== Get Posts Example ===")
	posts, err := client.GetPosts(authResp.AccessToken, &flaro.PostsQueryParams{
		Select: "*",
		Order:  "created_at.desc.nullslast",
		Offset: 0,
		Limit:  5,
	})
	if err != nil {
		log.Printf("Get posts failed: %v", err)
	} else {
		fmt.Printf("Retrieved %d posts:\n", len(posts))
		for i, post := range posts {
			// Get username for the creator
			creator, err := client.GetUser(authResp.AccessToken, post.CreatorID)
			creatorName := post.CreatorID // fallback to ID if username fetch fails
			if err == nil {
				creatorName = "@" + creator.Username
			}

			content := post.Content
			if len(content) > 50 {
				content = content[:50] + "..."
			}

			fmt.Printf("  %d. %s (by %s)\n", i+1, content, creatorName)
			if len(post.MediaURLs) > 0 {
				fmt.Printf("     Media: %d image(s)\n", len(post.MediaURLs))
			}
			fmt.Printf("     Likes: %d, Boost: %d\n", len(post.Likes), post.Boost)
		}
	}

	// Example: Get following
	fmt.Println("\n=== Get Following Example ===")
	following, err := client.GetFollowing(authResp.AccessToken, authResp.User.ID)
	if err != nil {
		log.Printf("Get following failed: %v", err)
	} else {
		fmt.Printf("Following %d users:\n", len(following))
		for i, follow := range following {
			fmt.Printf("  %d. %s (@%s)\n", i+1, follow.Users.DisplayName, follow.Users.Username)
			if follow.Users.IsVerified {
				fmt.Printf("     ✓ Verified\n")
			}
		}
	}

	// Example: Get posts from a specific user (using the first post creator as example)
	if len(posts) > 0 {
		fmt.Println("\n=== Get User Posts Example ===")
		userPosts, err := client.GetUserPosts(authResp.AccessToken, posts[0].CreatorID)
		if err != nil {
			log.Printf("Get user posts failed: %v", err)
		} else {
			fmt.Printf("Posts from user %s:\n", posts[0].CreatorID)
			for i, post := range userPosts {
				content := post.Content
				if len(content) > 50 {
					content = content[:50] + "..."
				}
				fmt.Printf("  %d. %s\n", i+1, content)
				fmt.Printf("     Likes: %d, Boost: %d\n", len(post.Likes), post.Boost)
			}
		}
	}

	// Example: Get comments for the first post
	if len(posts) > 0 {
		fmt.Println("\n=== Get Comments Example ===")
		comments, err := client.GetComments(authResp.AccessToken, posts[0].ID)
		if err != nil {
			log.Printf("Get comments failed: %v", err)
		} else {
			fmt.Printf("Comments on post %s:\n", posts[0].ID)
			for i, comment := range comments {
				content := comment.Content
				if len(content) > 50 {
					content = content[:50] + "..."
				}
				fmt.Printf("  %d. %s (by %s)\n", i+1, content, comment.UserID)
				fmt.Printf("     Likes: %d\n", len(comment.Likes))
			}
		}

		// Example: Get comment count
		fmt.Println("\n=== Get Comment Count Example ===")
		commentCount, err := client.GetCommentCount(authResp.AccessToken, posts[0].ID)
		if err != nil {
			log.Printf("Get comment count failed: %v", err)
		} else {
			fmt.Printf("Comment count for post %s: %d\n", posts[0].ID, commentCount)
		}

		// Example: Post a comment
		fmt.Println("\n=== Post Comment Example ===")
		newComment, err := client.PostComment(authResp.AccessToken, posts[0].ID, authResp.User.ID, "Great post!", nil)
		if err != nil {
			log.Printf("Post comment failed: %v", err)
		} else {
			fmt.Printf("Posted comment: %s\n", newComment.Content)
			fmt.Printf("Comment ID: %s\n", newComment.ID)
		}

		// Example: Like/Unlike a post
		fmt.Println("\n=== Like Post Example ===")
		err = client.LikePost(authResp.AccessToken, posts[0].ID, authResp.User.ID, true)
		if err != nil {
			log.Printf("Like post failed: %v", err)
		} else {
			fmt.Printf("Successfully liked post %s\n", posts[0].ID)
		}

		// Example: Delete a comment (if we posted one and got a valid ID)
		if newComment != nil && newComment.ID != "" {
			fmt.Println("\n=== Delete Comment Example ===")
			err = client.DeleteComment(authResp.AccessToken, newComment.ID)
			if err != nil {
				log.Printf("Delete comment failed: %v", err)
			} else {
				fmt.Printf("Successfully deleted comment %s\n", newComment.ID)
			}
		} else {
			fmt.Println("\n=== Delete Comment Example ===")
			fmt.Println("Skipping delete comment example - no valid comment ID received from API")

			// Try to delete an existing comment if there are any
			if len(comments) > 0 {
				fmt.Println("Attempting to delete an existing comment for demonstration...")
				err = client.DeleteComment(authResp.AccessToken, comments[0].ID)
				if err != nil {
					log.Printf("Delete existing comment failed: %v", err)
				} else {
					fmt.Printf("Successfully deleted existing comment %s\n", comments[0].ID)
				}
			}
		}

		// Example: Create a new post
		fmt.Println("\n=== Create Post Example ===")
		err = client.CreatePost(authResp.AccessToken, authResp.User.ID, "Hello from Go SDK!", []string{})
		if err != nil {
			log.Printf("Create post failed: %v", err)
		} else {
			fmt.Printf("Successfully created new post!\n")
		}

		// Example: Refresh token
		fmt.Println("\n=== Refresh Token Example ===")
		newAuthResp, err := client.RefreshToken(authResp.RefreshToken)
		if err != nil {
			log.Printf("Refresh token failed: %v", err)
		} else {
			fmt.Printf("Successfully refreshed token! New expires at: %d\n", newAuthResp.ExpiresAt)
		}

		// Example: Get reels
		fmt.Println("\n=== Get Reels Example ===")
		reels, err := client.GetReels(authResp.AccessToken)
		if err != nil {
			log.Printf("Get reels failed: %v", err)
		} else {
			fmt.Printf("Found %d reels\n", len(reels))
			if len(reels) > 0 {
				fmt.Printf("Latest reel: %s by %s\n", reels[0].Content, reels[0].CreatorID)
			}
		}

		// Example: Get specific reel by ID (if we have reels)
		if len(reels) > 0 {
			fmt.Println("\n=== Get Reel By ID Example ===")
			specificReels, err := client.GetReelByID(authResp.AccessToken, reels[0].ID)
			if err != nil {
				log.Printf("Get reel by ID failed: %v", err)
			} else {
				fmt.Printf("Found reel: %s\n", specificReels[0].Content)
			}

			// Example: Get reel comments
			fmt.Println("\n=== Get Reel Comments Example ===")
			reelComments, err := client.GetReelComments(authResp.AccessToken, reels[0].ID)
			if err != nil {
				log.Printf("Get reel comments failed: %v", err)
			} else {
				fmt.Printf("Found %d comments on reel\n", len(reelComments))
				if len(reelComments) > 0 {
					fmt.Printf("Latest comment: %s\n", reelComments[0].Content)
				}
			}
		}

		// Example: Get system messages
		fmt.Println("\n=== Get System Messages Example ===")
		systemMessages, err := client.GetSystemMessages(authResp.AccessToken, authResp.User.ID)
		if err != nil {
			log.Printf("Get system messages failed: %v", err)
		} else {
			fmt.Printf("Found %d system messages\n", len(systemMessages))
			if len(systemMessages) > 0 {
				fmt.Printf("Latest message: %s\n", systemMessages[0].Content)
			}
		}

		// Example: Search users
		fmt.Println("\n=== Search Users Example ===")
		searchResults, err := client.SearchUsers(authResp.AccessToken, "churro")
		if err != nil {
			log.Printf("Search users failed: %v", err)
		} else {
			fmt.Printf("Found %d users matching 'churro'\n", len(searchResults))
			if len(searchResults) > 0 {
				fmt.Printf("Found user: %s (@%s)\n", searchResults[0].DisplayName, searchResults[0].Username)
			}
		}

		// Example: Update user details
		fmt.Println("\n=== Update User Details Example ===")
		newBio := "Updated bio from Go SDK!"
		err = client.UpdateUserDetails(authResp.AccessToken, authResp.User.ID, &newBio, nil, nil)
		if err != nil {
			log.Printf("Update user bio failed: %v", err)
		} else {
			fmt.Println("Successfully updated user bio!")
		}

		// Example: Update username (commented out to avoid changing username)
		// newUsername := "newusername"
		// err = client.UpdateUserDetails(authResp.AccessToken, authResp.User.ID, nil, &newUsername, nil)
		// if err != nil {
		//     log.Printf("Update username failed: %v", err)
		// } else {
		//     fmt.Println("Successfully updated username!")
		// }

		// Example: Update profile picture (commented out to avoid changing profile picture)
		// newProfilePicture := "https://example.com/new-profile-picture.jpg"
		// err = client.UpdateUserDetails(authResp.AccessToken, authResp.User.ID, nil, nil, &newProfilePicture)
		// if err != nil {
		//     log.Printf("Update profile picture failed: %v", err)
		// } else {
		//     fmt.Println("Successfully updated profile picture!")
		// }

		// Example: Delete a post (if we have posts)
		if len(posts) > 0 {
			fmt.Println("\n=== Delete Post Example ===")
			err = client.DeletePost(authResp.AccessToken, posts[0].ID)
			if err != nil {
				log.Printf("Delete post failed: %v", err)
			} else {
				fmt.Printf("Successfully deleted post %s\n", posts[0].ID)
			}
		}

		// Example: Report a user
		fmt.Println("\n=== Report User Example ===")
		err = client.ReportUser(authResp.AccessToken, authResp.User.ID, authResp.User.ID, nil, nil, "Unknown")
		if err != nil {
			log.Printf("Report user failed: %v", err)
		} else {
			fmt.Println("Successfully reported user!")
		}

		// Example: Report a problem
		fmt.Println("\n=== Report Problem Example ===")
		err = client.ReportProblem(authResp.AccessToken, authResp.User.ID, "Test Problem", "This is a test problem report from Go SDK")
		if err != nil {
			log.Printf("Report problem failed: %v", err)
		} else {
			fmt.Println("Successfully reported problem!")
		}

		// Example: Contact support
		fmt.Println("\n=== Contact Support Example ===")
		err = client.ContactSupport(authResp.AccessToken, authResp.User.ID, "Test Support Request", "This is a test support request from Go SDK")
		if err != nil {
			log.Printf("Contact support failed: %v", err)
		} else {
			fmt.Println("Successfully contacted support!")
		}

		// Example: Change password (commented out to avoid changing password)
		// fmt.Println("\n=== Change Password Example ===")
		// changeResp, err := client.ChangePassword(authResp.AccessToken, "newpassword123")
		// if err != nil {
		//     log.Printf("Change password failed: %v", err)
		// } else {
		//     fmt.Printf("Successfully changed password! User ID: %s\n", changeResp.ID)
		// }

		// Example: Upload video (commented out - requires actual video file)
		// fmt.Println("\n=== Upload Video Example ===")
		// videoData := []byte("fake video data") // In real usage, this would be actual video file data
		// videoResp, err := client.UploadVideo(authResp.AccessToken, videoData, 3600)
		// if err != nil {
		//     log.Printf("Upload video failed: %v", err)
		// } else {
		//     fmt.Printf("Successfully uploaded video! Key: %s, ID: %s\n", videoResp.Key, videoResp.Id)
		//
		//     // Example: Create reel with uploaded video
		//     fmt.Println("\n=== Create Reel Example ===")
		//     videoURL := fmt.Sprintf("https://sb.flaroapp.pl/storage/v1/object/public/%s", videoResp.Key)
		//     err = client.CreateReel(authResp.AccessToken, authResp.User.ID, "Test reel from Go SDK!", videoURL)
		//     if err != nil {
		//         log.Printf("Create reel failed: %v", err)
		//     } else {
		//         fmt.Println("Successfully created reel!")
		//     }
		// }
	}
}
