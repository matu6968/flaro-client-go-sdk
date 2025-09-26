# Flaro social media API docs

These API docs allow you to create a client SDK in another language, Go is shown as a reference implementation, but you may use C, Python, etc to implement the binding.

## Sign up endpoint
**URL:** https://sb.flaroapp.pl/auth/v1/signup
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` (optional) : `Bearer insert-obtained-api-key-from-app>`

**Body**:
```json
{
   "email": "<email>",
   "password": "<password>",
   "data": null,
   "gotrue_meta_security": {
     "captcha_token": null
   },
   "code_challenge": "<insert_generated_pkcs_code_challenge>",
   "code_challenge_method": "s256"
}
```

**Response (200)**:
```json
{
  "access_token": "<access_token>",
  "token_type": "bearer",
  "expires_in": 3600,
  "expires_at": 1758463067,
  "refresh_token": "<refresh-token>",
  "user": {
    "id": "<user-id>",
    "aud": "authenticated",
    "role": "authenticated",
    "email": "<user-email>",
    "email_confirmed_at": "2025-09-21T12:57:47.692957761Z",
    "phone": "",
    "last_sign_in_at": "2025-09-21T12:57:47.699863065Z",
    "app_metadata": {
       "provider": "email",
       "providers": [
          "email"
       ]
    },
   "user_metadata": {
   "email": "<user-email>",
   "email_verified": true,
   "phone_verified": false,
   "sub": "<user-id>"
  },
  "identities": [
    {
       "identity_id": "<user-identity_id>",
       "id": "<user-id>",
       "user_id": "<user-id>",
       "identity_data": {
       "email": "<user-email>",
       "email_verified": true,
       "phone_verified": false,
       "sub": "<user-id>"

   },
   "provider": "email",
   "last_sign_in_at": "2025-09-21T12:57:47.685430845Z",
   "created_at": "2025-09-21T12:57:47.685541Z",
   "updated_at": "2025-09-21T12:57:47.685541Z",
    "email": "<user-email>"
   }
 ],
 "created_at": "2025-09-21T12:57:47.67796Z",
 "updated_at": "2025-09-21T12:57:47.703332Z",
 "is_anonymous": false
  }
}
```

**Before proceeding further with this newly made account via the provided steps, do the following:**
## prerequisite step after making the account from the Sign up endpoint after signing in
**URL:** https://sb.flaroapp.pl/rest/v1/users
**HTTP request type:** POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**Body**:
```json
{
   "user_id":"<user-id>",
   "username":"<username>",
   "display_name":"<username>",
   "bio":"",
   "profile_picture":"https://i.postimg.cc/660K5Hrr/pfp.png", // default profile picture from a external source
   "website":null,
   "is_private":false,
   "created_at":"2025-09-21T14:57:47.818532", // current date in ISO 8601 format
   "username_updated_at":null,
   "is_verified":false,
   "last_seen":"2025-09-21T14:57:47.818534", // current date in ISO 8601 format
   "ranks":null,
   "premium_expires":null
}
```

**Response (201)**: No response.

## Sign in endpoint
**URL:** https://sb.flaroapp.pl/auth/v1/token?grant_type=password
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` (optional) : `Bearer insert-obtained-api-key-from-app>`

**Body**:
```json
{
   "email": "<email>",
   "password": "<password>",
   "gotrue_meta_security": {
     "captcha_token": null
   }
}
```

**Response (200)**:
```json
{
  "access_token": "<access_token>",
  "token_type": "bearer",
  "expires_in": 3600,
  "expires_at": 1758463067,
  "refresh_token": "<refresh-token>",
  "user": {
    "id": "<user-id>",
    "aud": "authenticated",
    "role": "authenticated",
    "email": "<user-email>",
    "email_confirmed_at": "2025-09-21T12:57:47.692957761Z",
    "phone": "",
    "last_sign_in_at": "2025-09-21T12:57:47.699863065Z",
    "app_metadata": {
       "provider": "email",
       "providers": [
          "email"
       ]
    },
   "user_metadata": {
   "email": "<user-email>",
   "email_verified": true,
   "phone_verified": false,
   "sub": "<user-id>"
  },
  "identities": [
    {
       "identity_id": "<user-identity_id>",
       "id": "<user-id>",
       "user_id": "<user-id>",
       "identity_data": {
       "email": "<user-email>",
       "email_verified": true,
       "phone_verified": false,
       "sub": "<user-id>"

   },
   "provider": "email",
   "last_sign_in_at": "2025-09-21T12:57:47.685430845Z",
   "created_at": "2025-09-21T12:57:47.685541Z",
   "updated_at": "2025-09-21T12:57:47.685541Z",
    "email": "<user-email>"
   }
 ],
 "created_at": "2025-09-21T12:57:47.67796Z",
 "updated_at": "2025-09-21T12:57:47.703332Z",
 "is_anonymous": false
  }
}
```

## View posts endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/posts?select=%2A&order=created_at.desc.nullslast&offset=0&limit=20
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (by default all posts are selected  (*))
- order: sort by latest comment aka`created_at.asc.nullslast` 
- offset: (likely sets post offset?)
- limit: sets the post limit as a integer (likely max is 100 here but unsure)

**Response (200, 5 post limit)**:
```json
[
  {
     "id": "67b6ad31-42a0-44b8-b13a-06cf4be3248f",
     "creator_id": "3310385e-e6fa-4145-a665-07fa3e9cef41",
     "content": "global announcement and tag",
     "media_urls": [
       "https://sb.flaroapp.pl/storage/v1/object/public/post-images/uploads/1758441303289-0" // images can be accessed without auth (unless it's private then likely a authorization key is needed)
     ],
     "created_at": "2025-09-21T09:55:31.201465+00:00",
     "tags": [],
     "score": 0, // view count (?)
     "boost_ends": null,
     "boost": 1, // post like count (?)
     "is_private": false,
     "location": null,
     "mentions": [],
     "comments": [],
     "likes": [
      "74b809b3-875c-4eb3-81bd-5ba82897589c" // each like is refered to it's user UUID
    ]
 },
 {
    "id": "69fe20af-7cb7-4994-8a30-299315abf5d6",
    "creator_id": "2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c",
    "content": "meow",
    "media_urls": [
        "https://sb.flaroapp.pl/storage/v1/object/public/post-images/uploads/1758400558407-0"
      ],
    "created_at": "2025-09-20T22:35:59.029939+00:00",
    "tags": [],
    "score": 0,
    "boost_ends": null,
    "boost": 1,
    "is_private": false,
    "location": null,
    "mentions": [],
    "comments": [],
    "likes": [
      "74b809b3-875c-4eb3-81bd-5ba82897589c"
    ]
 },
 {
    "id": "33404c58-45ef-4657-9c27-707b926cad41",
    "creator_id": "74b809b3-875c-4eb3-81bd-5ba82897589c",
    "content": "co lubicie se zjeść?",
    "media_urls": [],
    "created_at": "2025-09-20T21:03:43.370296+00:00",
    "tags": [],
    "score": 0,
    "boost_ends": null,
    "boost": 1,
    "is_private": false,
    "location": null,
    "mentions": [],
    "comments": [],
    "likes": []
},
{
    "id": "207819a7-fa35-4b21-b0b0-1d295a57d9ec",
    "creator_id": "69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
    "content": "kup ktos subskrypcje",
    "media_urls": [
       "https://sb.flaroapp.pl/storage/v1/object/public/post-images/uploads/1758391056960-0"
    ],
    "created_at": "2025-09-20T19:57:37.841495+00:00",
    "tags": [],
    "score": 0,
    "boost_ends": null,
    "boost": 1,
    "is_private": false,
    "location": null,
    "mentions": [],
    "comments": [],
    "likes": [
       "74b809b3-875c-4eb3-81bd-5ba82897589c"
     ]
},
{
    "id": "396b962b-1f50-482d-8e15-dc2b9a690288",
    "creator_id": "74b809b3-875c-4eb3-81bd-5ba82897589c",
    "content": "Gratulacje dla Ferrari za dzisiejsze kwalifikacje do GP Azerbejdżanu. Hamilton odpadł w Q2, a Leclerc uderzył w bandę podczas Q3. (przynajmniej tyle że Piastri też się rozbił 🤗)",
    "media_urls": [
       "https://sb.flaroapp.pl/storage/v1/object/public/post-images/uploads/1758389930779-0"
    ],
    "created_at": "2025-09-20T19:38:51.54441+00:00",
    "tags": [],
    "score": 0,
    "boost_ends": null,
    "boost": 1,
    "is_private": false,
    "location": null,
    "mentions": [],
    "comments": [],
    "likes": [
       "69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
       "74b809b3-875c-4eb3-81bd-5ba82897589c",
       "f6a7ab14-9f4e-4eca-9d61-5691c96a1eda",
       "81ae751b-b5ae-4fe9-ae69-28bc363c5cc0"
    ]

}
]
```
## View following endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/follows?select=following_id%2Cusers%21follows_following_id_fkey%28%2A%29&follower_id=eq.d677717a-0b96-4071-b637-c589fbb75d4e
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view all followers based on a a followers user , aka `following_id,users!follows_following_id_fkey(*)`)
- following_id: `eq.` followed by the users ID like `eq.d677717a-0b96-4071-b637-c589fbb75d4e`

**Response (200)**:
```json
[
   {
      "following_id": "bd44d15e-8c39-433a-a3d4-1dfffd9b44bc",
      "users": {
      "bio": "",
      "ranks": null,
      "user_id": "bd44d15e-8c39-433a-a3d4-1dfffd9b44bc",
      "website": null,
      "username": "Flaro",
      "last_seen": "2025-09-14T14:30:37.535305",
      "created_at": "2025-09-14T14:30:37.535295",
      "is_private": false,
      "is_verified": true,
      "display_name": "Flaro",
      "premium_expires": null,
      "profile_picture": "https://i.postimg.cc/660K5Hrr/pfp.png",
      "username_updated_at": null

   }
}
]
```
## View user endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/users?select=%2A&user_id=eq.d677717a-0b96-4071-b637-c589fbb75d4e
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view all followers aka *)
- user_id: `eq.` followed by the users ID like `eq.d677717a-0b96-4071-b637-c589fbb75d4e`

**Response (200)**:
```json
{
   "user_id": "d677717a-0b96-4071-b637-c589fbb75d4e",
   "username": "mateusz6768",
   "display_name": "mateusz6768",
   "bio": "",
   "profile_picture": "https://i.postimg.cc/660K5Hrr/pfp.png",
   "website": null,
   "is_private": false,
   "created_at": "2025-09-21T14:57:47.818532",
   "username_updated_at": null,
   "is_verified": false,
   "last_seen": "2025-09-21T14:57:47.818534",
   "ranks": null,
   "premium_expires": null
}
```
## View posts (from specified user) endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/posts?select=%2A&creator_id=eq.2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view all posts aka *)
- creator_id: `eq.` followed by the users ID like `eq.d677717a-0b96-4071-b637-c589fbb75d4e`

**Response (200)**:
```json
[
   {
      "id":"ef8eed9e-70ff-4623-a712-7ff636822033",
      "creator_id":"2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c",
      "content":"pierwszy furry femboy na flaro?",
      "media_urls":[
         "https://sb.flaroapp.pl/storage/v1/object/public/post-images/uploads/1758313279419-0"
      ],
      "created_at":"2025-09-19T22:21:20.565617+00:00",
      "tags":[
         
      ],
      "score":0,
      "boost_ends":null,
      "boost":1,
      "is_private":false,
      "location":null,
      "mentions":[
         
      ],
      "comments":[
         
      ],
      "likes":[
         "2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c",
         "74b809b3-875c-4eb3-81bd-5ba82897589c"
      ]
   },
   {
      "id":"34d194a2-2993-42c4-ae66-b988268c7a22",
      "creator_id":"2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c",
      "content":"nudy a wogole fajny kotek?",
      "media_urls":[
         "https://sb.flaroapp.pl/storage/v1/object/public/post-images/uploads/1758385253150-0"
      ],
      "created_at":"2025-09-20T18:20:54.455591+00:00",
      "tags":[
         
      ],
      "score":0,
      "boost_ends":null,
      "boost":1,
      "is_private":false,
      "location":null,
      "mentions":[
         
      ],
      "comments":[
         
      ],
      "likes":[
         "6e12c069-275c-4cd0-b394-dce191330b5e",
         "74b809b3-875c-4eb3-81bd-5ba82897589c",
         "69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7"
      ]
   },
   {
      "id":"7204bcf1-d273-4e12-b797-7cfcebac74c5",
      "creator_id":"2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c",
      "content":"bomba",
      "media_urls":[
         
      ],
      "created_at":"2025-09-19T23:34:51.570537+00:00",
      "tags":[
         
      ],
      "score":0,
      "boost_ends":null,
      "boost":1,
      "is_private":false,
      "location":null,
      "mentions":[
         
      ],
      "comments":[
         
      ],
      "likes":[
         "3310385e-e6fa-4145-a665-07fa3e9cef41"
      ]
   },
   {
      "id":"69fe20af-7cb7-4994-8a30-299315abf5d6",
      "creator_id":"2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c",
      "content":"meow",
      "media_urls":[
         "https://sb.flaroapp.pl/storage/v1/object/public/post-images/uploads/1758400558407-0"
      ],
      "created_at":"2025-09-20T22:35:59.029939+00:00",
      "tags":[
         
      ],
      "score":0,
      "boost_ends":null,
      "boost":1,
      "is_private":false,
      "location":null,
      "mentions":[
         
      ],
      "comments":[
         
      ],
      "likes":[
         "74b809b3-875c-4eb3-81bd-5ba82897589c"
      ]
   }
]
```

## Add/remove like to a post endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/posts?id=eq.ef8eed9e-70ff-4623-a712-7ff636822033
**Request type:** PATCH
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- id: `eq.` followed by the post ID like `eq.ef8eed9e-70ff-4623-a712-7ff636822033

**Body**:
```json
{
   "likes":[
      "2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c",
      "74b809b3-875c-4eb3-81bd-5ba82897589c",
      "d677717a-0b96-4071-b637-c589fbb75d4e" // add or remove your user ID here if it exists, do not touch other user ID's other then your user ID
   ]
}
```

**Response (204)**: No response.

## View comments on a user post endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/comments?select=%2A&post_id=eq.33404c58-45ef-4657-9c27-707b926cad41&order=created_at.asc.nullslast
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view all comments aka *)
- post_id: `eq.` followed by the post ID like `eq.33404c58-45ef-4657-9c27-707b926cad41
- order: sort by latest comment aka`created_at.asc.nullslast` 

**Response (200)**:
```json
[
   {
      "id":"27f10349-50e8-4c97-8214-c18076da8c2c",
      "created_at":"2025-09-20T19:04:54.962197+00:00",
      "user_id":"69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
      "content":"pizze",
      "likes":[
         "74b809b3-875c-4eb3-81bd-5ba82897589c"
      ],
      "is_edited":false,
      "parent_id":null,
      "post_id":"33404c58-45ef-4657-9c27-707b926cad41",
      "reel_id":null,
      "reply_to_username":null
   },
   {
      "id":"fab04fd8-c6c6-418d-a397-2c3a8f5b9c4b",
      "created_at":"2025-09-20T19:11:25.723517+00:00",
      "user_id":"74b809b3-875c-4eb3-81bd-5ba82897589c",
      "content":"Klasyczek zawsze na topie",
      "likes":[
         "6e12c069-275c-4cd0-b394-dce191330b5e"
      ],
      "is_edited":false,
      "parent_id":"27f10349-50e8-4c97-8214-c18076da8c2c", // parent_id refers to which comment ID it got replied to
      "post_id":"33404c58-45ef-4657-9c27-707b926cad41",
      "reel_id":null,
      "reply_to_username":null
   },
   {
      "id":"38417da2-8ef1-438a-8115-aeaa98919cb5",
      "created_at":"2025-09-20T20:39:33.71427+00:00",
      "user_id":"2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c",
      "content":"Pizzw",
      "likes":[
         "74b809b3-875c-4eb3-81bd-5ba82897589c"
      ],
      "is_edited":false,
      "parent_id":null,
      "post_id":"33404c58-45ef-4657-9c27-707b926cad41",
      "reel_id":null,
      "reply_to_username":null
   }
]
```

## Get comment count (also known as just ID's) on a user post endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/comments?select=id&post_id=eq.ef8eed9e-70ff-4623-a712-7ff636822033
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view comment ID's aka `id`) 
- post_id:  `eq.` followed by the post ID like `eq.ef8eed9e-70ff-4623-a712-7ff636822033

**Response (200)**:
```json
[
   {
      "id":"c027eeb3-7372-45ab-a11e-06110046e51b"
   },
   {
      "id":"6c1dc281-0a28-4595-8b71-97a94abb7ae0"
   },
   {
      "id":"b014ee87-8489-482f-a46b-daaf3bb8b303"
   },
   {
      "id":"2733e811-a94c-4350-bc9f-07b4c258bd96"
   }
]
```
## Post comment on a user post/reel endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/comments?select=%2A
**Request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view all comments aka *) 

**Body**: 
```json
{
   "post_id":"33404c58-45ef-4657-9c27-707b926cad41", // post_id would be filled with the user post ID, if it's a reel, set it to null
   "reel_id":null, // post_id would be filled with the user reel ID it it's a reel otherwaise set to null
   "user_id":"d677717a-0b96-4071-b637-c589fbb75d4e", // user ID is the user ID of the account
   "content":"tortilia", // actual message content
   "parent_id":null, // replied to message ID, if none use null
   "likes":[
      
   ]
}
```
**Response (201)**:
```json
{
   "id":"7d59fad0-76a7-494e-a994-4aa4d7cf06e5",
   "created_at":"2025-09-21T15:13:07.368964+00:00",
   "user_id":"d677717a-0b96-4071-b637-c589fbb75d4e",
   "content":"tortilia",
   "likes":[
      
   ],
   "is_edited":false,
   "parent_id":null,
   "post_id":"33404c58-45ef-4657-9c27-707b926cad41",
   "reel_id":null,
   "reply_to_username":null
}
```

## Delete comment on a user post endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/comments?id=eq.74c0a8d9-fa29-48f9-b91c-4137f78b2aae
**Request type**: DELETE
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- id:  `eq.` followed by the comment ID like `eq.74c0a8d9-fa29-48f9-b91c-4137f78b2aae`

**Response (204)**: No content.

## Post image on a new user post endpoint
**URL:** https://sb.flaroapp.pl/storage/v1/object/post-images/uploads/[unix-time-with-ms]-[count-from-0-to-5]
**Request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL parameters:**
- unix-time-with-ms: current system time in the Unix epoch with prefixed milliseconds (like 1758469632000)
- count-from-0-to-5: if there are multiple files in the post, you can add a number to specify a image number from 0 to 5.

**Body**: 
```
Multi part data:
<name-field-empty>: image data
CacheControl: specify in seconds the cache busting time of the file (typically 3600 aka 1 hour)
```
**Response (200)**:
```json
{
	"Key": "post-images/uploads/1758469632000-0",
	"Id": "ceb41cbb-250c-4b11-bedf-aa06a2a88671"
}
```

To use this file after in new posts, the URL will be like this: `https://sb.flaroapp.pl/storage/v1/object/public/post-images/uploads/1758469632000-0`

## Make new user post endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/posts
**Request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**Body**: 
```json
{
   "creator_id":"d677717a-0b96-4071-b637-c589fbb75d4e", // creator_id is the current user ID
   "content":"test post", // description of the post
   "media_urls":[
      "https://sb.flaroapp.pl/storage/v1/object/public/post-images/uploads/1758468921858-0" // specify uploaded media url's from the above post image endpoint
   ],
   "created_at":"2025-09-21T17:35:22.323266", //  current date in ISO 8601 format
   "tags":[
      
   ],
   "score":0,
   "boost_ends":null,
   "boost":1,
   "is_private":false,
   "location":null,
   "mentions":[
      
   ],
   "comments":[
      
   ],
   "likes":[
      
   ]
}
```
**Response (201)**: No response.
## Refresh token endpoint
**URL:** https://sb.flaroapp.pl/auth/v1/token?grant_type=refresh_token
**Request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- grant_type: type of operation to do on the token (in this case to renew the expiry date, aka `refresh_token`) 

**Body**: 
```json
{
   "refresh_token": "f2cjd47bsxpi"
}
```
**Response (200):

```json
{
  "access_token": "<access_token>",
  "token_type": "bearer",
  "expires_in": 3600,
  "expires_at": 1758463067,
  "refresh_token": "<refresh-token>",
  "user": {
    "id": "<user-id>",
    "aud": "authenticated",
    "role": "authenticated",
    "email": "<user-email>",
    "email_confirmed_at": "2025-09-21T12:57:47.692957761Z",
    "phone": "",
    "last_sign_in_at": "2025-09-21T12:57:47.699863065Z",
    "app_metadata": {
       "provider": "email",
       "providers": [
          "email"
       ]
    },
   "user_metadata": {
   "email": "<user-email>",
   "email_verified": true,
   "phone_verified": false,
   "sub": "d677717a-0b96-4071-b637-c589fbb75d4e"
  },
  "identities": [
    {
       "identity_id": "<user-identity_id>",
       "id": "<user-id>",
       "user_id": "<user-id>",
       "identity_data": {
       "email": "<user-email>",
       "email_verified": true,
       "phone_verified": false,
       "sub": "<user-id>"

   },
   "provider": "email",
   "last_sign_in_at": "2025-09-21T12:57:47.685430845Z",
   "created_at": "2025-09-21T12:57:47.685541Z",
   "updated_at": "2025-09-21T12:57:47.685541Z",
    "email": "<user-email>"
   }
 ],
 "created_at": "2025-09-21T12:57:47.67796Z",
 "updated_at": "2025-09-21T12:57:47.703332Z",
 "is_anonymous": false
  }
}

```

## Get reels endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/reels?select=%2A&order=created_at.desc.nullslast
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view all reels aka *) 
- order: sort by latest reel aka`created_at.asc.nullslast` 

**Response (200, showing last 5 reels)**:
```json
[
   {
      "id":"94feaeba-a456-4fc1-86f9-c0e3a2d735d8",
      "creator_id":"81ae751b-b5ae-4fe9-ae69-28bc363c5cc0",
      "content":"💩",
      "video":"https://sb.flaroapp.pl/storage/v1/object/public/reel-videos/uploads/1758467573846",
      "created_at":"2025-09-21T18:13:00.733333+00:00",
      "tags":[
         
      ],
      "score":0,
      "boost_ends":null,
      "boost":1,
      "is_private":false,
      "location":null,
      "mentions":[
         
      ],
      "comments":[
         
      ],
      "likes":[
         "6922e95d-4303-4c31-a28d-49eaf67e6192",
         "81ae751b-b5ae-4fe9-ae69-28bc363c5cc0"
      ]
   },
   {
      "id":"61e0b49c-242a-4be7-9d5a-12dcab9fb25a",
      "creator_id":"69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
      "content":"fr",
      "video":"https://sb.flaroapp.pl/storage/v1/object/public/reel-videos/uploads/1758319172354",
      "created_at":"2025-09-19T23:59:34.295129+00:00",
      "tags":[
         
      ],
      "score":0,
      "boost_ends":null,
      "boost":1,
      "is_private":false,
      "location":null,
      "mentions":[
         
      ],
      "comments":[
         
      ],
      "likes":[
         "74b809b3-875c-4eb3-81bd-5ba82897589c",
         "85d6bbb2-95f2-416c-aa59-1bb79f57de57"
      ]
   },
   {
      "id":"2798f642-e0c2-4ac7-a3a5-c58961a32eb9",
      "creator_id":"69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
      "content":"brainrot",
      "video":"https://sb.flaroapp.pl/storage/v1/object/public/reel-videos/uploads/1758319100780",
      "created_at":"2025-09-19T23:58:21.652859+00:00",
      "tags":[
         
      ],
      "score":0,
      "boost_ends":null,
      "boost":1,
      "is_private":false,
      "location":null,
      "mentions":[
         
      ],
      "comments":[
         
      ],
      "likes":[
         "2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c"
      ]
   },
   {
      "id":"2ce65cbe-e025-43e2-809c-f3bbc86420af",
      "creator_id":"2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c",
      "content":"fr",
      "video":"https://sb.flaroapp.pl/storage/v1/object/public/reel-videos/uploads/1758317825528",
      "created_at":"2025-09-19T23:37:05.932564+00:00",
      "tags":[],
      "score":0,
      "boost_ends":null,
      "boost":1,
      "is_private":false,
      "location":null,
      "mentions":[],
      "comments":[],
      "likes":[
         "2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c"
      ]
   },
   {
      "id":"8cbaef07-6e08-41ca-a184-000186406ec4",
      "creator_id":"f6a7ab14-9f4e-4eca-9d61-5691c96a1eda",
      "content":"twórca: @capesy.sm",
      "video":"https://sb.flaroapp.pl/storage/v1/object/public/reel-videos/uploads/1758314508882",
      "created_at":"2025-09-19T22:41:53.128619+00:00",
      "tags":[],
      "score":0,
      "boost_ends":null,
      "boost":1,
      "is_private":false,
      "location":null,
      "mentions":[],
      "comments":[],
      "likes":[
         "f6a7ab14-9f4e-4eca-9d61-5691c96a1eda",
         "69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
         "6e12c069-275c-4cd0-b394-dce191330b5e",
         "85d6bbb2-95f2-416c-aa59-1bb79f57de57",
         "81ae751b-b5ae-4fe9-ae69-28bc363c5cc0"
      ]
   }
]
```

## Get reels (by their ID) endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/reels?select=%2A&id=eq.94feaeba-a456-4fc1-86f9-c0e3a2d735d8
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view all reels aka *) 
- id: `eq.` followed by the reel ID like `eq.94feaeba-a456-4fc1-86f9-c0e3a2d735d8

**Response (200)**:
```json
[
   {
      "id":"94feaeba-a456-4fc1-86f9-c0e3a2d735d8",
      "creator_id":"81ae751b-b5ae-4fe9-ae69-28bc363c5cc0",
      "content":"💩",
      "video":"https://sb.flaroapp.pl/storage/v1/object/public/reel-videos/uploads/1758467573846",
      "created_at":"2025-09-21T18:13:00.733333+00:00",
      "tags":[],
      "score":0,
      "boost_ends":null,
      "boost":1,
      "is_private":false,
      "location":null,
      "mentions":[],
      "comments":[],
      "likes":[
         "6922e95d-4303-4c31-a28d-49eaf67e6192",
         "81ae751b-b5ae-4fe9-ae69-28bc363c5cc0"
      ]
   }
]
```
## View comments on a user reel endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/comments?select=%2A&reel_id=eq.94feaeba-a456-4fc1-86f9-c0e3a2d735d8&order=created_at.asc.nullslast
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view all comments aka *)
- reel_id: `eq.` followed by the reel ID like `eq.94feaeba-a456-4fc1-86f9-c0e3a2d735d8
- order: sort by latest comment aka`created_at.asc.nullslast` 

**Response (200)**:
```json
[
   {
      "id":"03f00738-3cb5-43cd-986c-b00ae1722a82",
      "created_at":"2025-09-21T15:27:52.45411+00:00",
      "user_id":"6922e95d-4303-4c31-a28d-49eaf67e6192",
      "content":"🤣",
      "likes":[
         "81ae751b-b5ae-4fe9-ae69-28bc363c5cc0"
      ],
      "is_edited":false,
      "parent_id":null,
      "post_id":null,
      "reel_id":"94feaeba-a456-4fc1-86f9-c0e3a2d735d8",
      "reply_to_username":null
   }
]
```

## View system messages endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/system_messages?select=%2A&order=created_at.desc.nullslast&limit=1
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view all system messages aka *)
- order: sort by latest system message aka`created_at.asc.nullslast` 
- limit: set the system message limit amount (like `1`)

**Response (200)**:
```json
{
   "id":1,
   "created_at":"2025-09-25T17:33:37+00:00",
   "title":"New Autumn Theme!\r\nImproved UI \r\nGlobal Chat\r\nAnd much more!\r\n",
   "content":"Update!",
   "read_by":[
      "69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
      "f75eedcf-d743-4762-b1b7-2ef57017af16",
      "74b809b3-875c-4eb3-81bd-5ba82897589c",
      "6e12c069-275c-4cd0-b394-dce191330b5e",
      "85d6bbb2-95f2-416c-aa59-1bb79f57de57",
      "f6a7ab14-9f4e-4eca-9d61-5691c96a1eda",
      "b78b6dd0-4cca-49df-8853-0f8b0f829003",
      "d1246357-ab88-45c3-9b6d-79b3c92171bb",
      "630c4184-3d34-42eb-83a8-9f179a07ccf4",
      "d086f7b4-f079-4541-bf6d-1de6f5e744f6",
      "79d0e59d-0ce2-42c0-8ba6-494a63c99d5f",
      "5d7ae3ce-7c4d-48a4-872f-22a54901a682",
      "e2beeaa6-59fa-43d4-902e-66e97fed7bad",
      "9db2491d-05e7-46a3-94dc-7f33624d00ab",
      "fe26f580-0021-4a91-a42f-0e99ccaa84e0",
      "7ff038c0-dd9b-4a87-91da-99e2fdb725a8",
      "0ab5808f-a272-4d9f-a735-68665a650c6f"
   ],
   "image":"https://i.postimg.cc/5ySQfvqz/logo-text-bg.png"
}
```

## Mark system message as read endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/system_messages?id=eq.1
**HTTP request type**: PATCH
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- id: `ed.` followed by the system message ID like `eq.1`

**Body**:

```json
{
   "read_by":[
      "69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
      "f75eedcf-d743-4762-b1b7-2ef57017af16",
      "74b809b3-875c-4eb3-81bd-5ba82897589c",
      "6e12c069-275c-4cd0-b394-dce191330b5e",
      "85d6bbb2-95f2-416c-aa59-1bb79f57de57",
      "f6a7ab14-9f4e-4eca-9d61-5691c96a1eda",
      "b78b6dd0-4cca-49df-8853-0f8b0f829003",
      "d1246357-ab88-45c3-9b6d-79b3c92171bb",
      "630c4184-3d34-42eb-83a8-9f179a07ccf4",
      "d086f7b4-f079-4541-bf6d-1de6f5e744f6",
      "79d0e59d-0ce2-42c0-8ba6-494a63c99d5f",
      "5d7ae3ce-7c4d-48a4-872f-22a54901a682",
      "e2beeaa6-59fa-43d4-902e-66e97fed7bad",
      "9db2491d-05e7-46a3-94dc-7f33624d00ab",
      "fe26f580-0021-4a91-a42f-0e99ccaa84e0",
      "7ff038c0-dd9b-4a87-91da-99e2fdb725a8",
      "0ab5808f-a272-4d9f-a735-68665a650c6f",
      "d677717a-0b96-4071-b637-c589fbb75d4e" // add your user ID to the body to read_by entity, do not modify other user ID's
   ]
}
```

**Response (204)**: No content.
## Searching for users endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/users?select=%2A&username=ilike.%25churro%25
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to search all users aka *)
- username: the username to search for with a `ilike.%<username>%` format  (such as`ilike.%churro%`)

**Response (200)**:
```json
[
   {
      "username":"churro",
      "display_name":"churro",
      "bio":"14, furry femboy, bottom, bomba",
      "profile_picture":"https://sb.flaroapp.pl/storage/v1/object/public/profile-pictures/uploads/2025-09-19T22:22:54.227197",
      "website":null,
      "is_private":false,
      "created_at":"2025-09-19T22:20:58.009966",
      "username_updated_at":null,
      "is_verified":false,
      "last_seen":"2025-09-19T22:20:58.009967",
      "user_id":"2e2bfc78-cba4-433e-ab0f-a12e7cdcf16c",
      "ranks":null,
      // or if there is a rank:
      // "ranks": [
      //   "Beta"
      //],
      "premium_expires":null // if the user has Flaro Premium, this would be replaced by a ISO 8601 compatible date like "3009-01-30T05:12:10+00:00"
   }
]
```

## Editing user details (bio, username or profile picture) endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/users?user_id=eq.d677717a-0b96-4071-b637-c589fbb75d4e
**Request type**: PATCH
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- user_id: `eq.` followed by the user ID like `eq.d677717a-0b96-4071-b637-c589fbb75d4e

**Body**:
```json
{
   "bio": "test",
   "username": "test",
   "profile_picture": "https://sb.flaroapp.pl/storage/v1/object/public/profile-pictures/uploads/2025-09-21T19:49:01.444342" // you can only input either bio, username or profile_picture one at a time in the JSON body
}
```
**Response (204)**: No response.

## Delete user post endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/posts?id=eq.cb61e4ba-6050-45de-8c73-c9f866567bc0
**Request type**: DELETE
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- id:  `eq.` followed by the post ID like `eq.cb61e4ba-6050-45de-8c73-c9f866567bc0

**Response (204)**: No content.

## Reporting a user/post/reel endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/reports
**Request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**Body**:
```json
{
   "created_at":"2025-09-21T19:30:16.412552",
   "user_id":"d677717a-0b96-4071-b637-c589fbb75d4e", // user ID you want to report
   "post_id":"cb61e4ba-6050-45de-8c73-c9f866567bc0", // post ID you want to report, optional (if there isn't a specific post ID you want to report, specify null)
   "reel_id":null, // reel ID you want to report, optional (if there isn't a specific reel ID you want to report, specify null)
   "reason":"Unknown", // optional report reason, you may specify Unknown otherwaise
   "status":"UNREAD",
   "reported_by":"d677717a-0b96-4071-b637-c589fbb75d4e" // must specify your user ID
}
```
**Response (201)**: No response.

## Reporting a problem endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/problems
**Request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**Body**:
```json
{
   "subject":"test report", // report subject here
   "content":"hello", // report description here
   "user_id":"d677717a-0b96-4071-b637-c589fbb75d4e", // your user ID here
   "status":"open",
   "created_at":"2025-09-21T19:36:36.027211" // current time in ISO 8601 format
}
```
**Response (201)**: No response.

## Contacting support endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/support
**Request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**Body**:
```json
{
   "subject":"test report", // support subject here
   "content":"hello", // support description here
   "user_id":"d677717a-0b96-4071-b637-c589fbb75d4e", // your user ID here
   "status":"open",
   "created_at":"2025-09-21T19:36:36.027211" // current time in ISO 8601 format
}
```
**Response (201)**: No response.

## Change password endpoint
**URL:** https://sb.flaroapp.pl/auth/v1/user
**Request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**Body**:
```json
{
   "password": "passwordhere"
}
```
**Response (200)**:
```json
{
   "id":"<user-id>",
   "aud":"authenticated",
   "role":"authenticated",
   "email":"<user-email>",
   "email_confirmed_at":"2025-09-21T12:57:47.692957Z",
   "phone":"",
   "confirmed_at":"2025-09-21T12:57:47.692957Z",
   "last_sign_in_at":"2025-09-21T17:41:18.314693Z",
   "app_metadata":{
      "provider":"email",
      "providers":[
         "email"
      ]
   },
   "user_metadata":{
      "email":"<user-email>",
      "email_verified":true,
      "phone_verified":false,
      "sub":"<user-id>"
   },
   "identities":[
      {
         "identity_id":"<user-identity_id>",
         "id":"<user-id>",
         "user_id":"<user-id>",
         "identity_data":{
            "email":"<user-email>",
            "email_verified":false,
            "phone_verified":false,
            "sub":"<user-id>"
         },
         "provider":"email",
         "last_sign_in_at":"2025-09-21T12:57:47.68543Z",
         "created_at":"2025-09-21T12:57:47.685541Z",
         "updated_at":"2025-09-21T12:57:47.685541Z",
         "email":"<user-email>"
      }
   ],
   "created_at":"2025-09-21T12:57:47.67796Z",
   "updated_at":"2025-09-21T17:41:18.720993Z",
   "is_anonymous":false
}
```

**Response if password is the same (422)**:
```json
{
  "code": "same_password",
  "message": "New password should be different from the old password."
}
```

## Post video on a new user post endpoint
**URL:** https://sb.flaroapp.pl/storage/v1/object/reel-videos/uploads/[unix-time-with-ms]
**Request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL parameters:**
- unix-time-with-ms: current system time in the Unix epoch with prefixed milliseconds (like 1758482692793)

**Body**: 
```
Multi part data:
<name-field-empty>: video data
CacheControl: specify in seconds the cache busting time of the file (typically 3600 aka 1 hour)
```
**Response (200)**:
```json
{
	"Key": "reel-videos/uploads/1758482692793",
	"Id": "ceb41cbb-250c-4b11-bedf-aa06a2a88671"
}
```

To use this file after in new posts, the URL will be like this: `https://sb.flaroapp.pl/storage/v1/object/public/reel-videos/uploads/1758469632000-0`

## Make new user reel endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/reels
**Request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**Body**: 
```json
{
   "creator_id":"d677717a-0b96-4071-b637-c589fbb75d4e", // creator_id is the current user ID
   "content":"test reel", // description of the post
   "media_urls":[
      "`https://sb.flaroapp.pl/storage/v1/object/public/reel-videos/uploads/1758482692793`" // specify uploaded media url's from the above post image endpoint
   ],
   "created_at":"2025-09-21T17:35:22.323266", //  current date in ISO 8601 format
   "tags":[],
   "score":0,
   "boost_ends":null,
   "boost":1,
   "is_private":false,
   "location":null,
   "mentions":[
      
   ],
   "comments":[
      
   ],
   "likes":[
      
   ]
}
```
**Response (201)**: No response.


## WebSocket (live messages) endpoint

**Note:** This isn't fully documented as i haven't documented a response when a new post/reel is detected, any SDK's that add this feature should mark it as experimental and offer a build flag (put websocket implemenation to a seperate file) to disable it

**URL:** wss://sb.flaroapp.pl/realtime/v1/websocket?apikey=insert-obtained-api-key-from-app>&vsn=1.0.0

**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`

**URL prefix parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
- vsn: refers to the app version aka `1.0.0`

**Messages**: 
To initialize a connection (posts):
```json
{
   "topic":"realtime:public:posts:135", // if it's tracking reels then replace posts with reels instead
   "event":"phx_join",
   "payload":{
      "config":{
         "broadcast":{
            "ack":false,
            "self":false
         },
         "presence":{
            "key":""
         },
         "postgres_changes":[
            {
               "event":"*",
               "schema":"public",
               "table":"posts", // if it's tracking reels then put reels instead
               "filter":"creator_id=eq.<user-id>"
            }
         ],
         "private":false
      },
      "access_token":"<user-access-token>"
   },
   "ref":"4", // number is incremeted by one per request
   "join_ref":"4" // number is incremeted by one per request
}
```

To initialize a connection (Global Channel):
```json
{
   "topic":"realtime:public:messages:103",
   "event":"phx_join",
   "payload":{
      "config":{
         "broadcast":{
            "ack":false,
            "self":false
         },
         "presence":{
            "key":""
         },
         "postgres_changes":[
            {
               "event":"*",
               "schema":"public",
               "table":"messages"
            }
         ],
         "private":false
      },
      "access_token":"<user_access_token>"
   },
   "ref":"11", // number is incremeted by one per request
   "join_ref":"11" // number is incremeted by one per request
}
```
When connection has been initialized:
```json
{
   "ref":"1", // number is incremeted by one per request
   "event":"phx_reply",
   "payload":{
      "status":"ok",
      "response":{
         "postgres_changes":[
            {
               "id":23091052, // id for Global Channel will be 72696871 instead
               "event":"*",
               "filter":"creator_id=eq.<user-id>", // won't be present on table messages
               "schema":"public",
               "table":"posts" // if it's tracking reels then it will be reels or if it's tracking a Global Channel then it will be called messages instead
            }
         ]
      }
   },
   "topic":"realtime:public:posts:135"
}
```

If subscription to events has succeeded (Global channel):
```json
{
   "ref":null,
   "event":"system",
   "payload":{
      "message":"Subscribed to PostgreSQL",
      "status":"ok",
      "extension":"postgres_changes",
      "channel":"public:messages:103"
   },
   "topic":"realtime:public:messages:103"
}
```
Heartbeat (make it every around 10 seconds or so):
```json
{
   "topic":"phoenix",
   "event":"heartbeat",
   "payload":{},
   "ref":"5" // number is incremeted by one per request
}
```

Response to heartbeat:

```json
{
   "ref":"5", // number is incremeted by one per request
   "event":"phx_reply",
   "payload":{
      "status":"ok",
      "response":{}
   },
   "topic":"phoenix"
}
```

Response to heartbeat if token is expired:
```json
{
   "ref":"4", // number is incremeted by one per request
   "event":"phx_reply",
   "payload":{
      "status":"error",
      "response":{
         "reason":"Token has expired 129 seconds ago"
      }
   },
   "topic":"realtime:public:posts:135"
}
```

Response if something went wrong (example error message)
```json
{
   "ref":null,  // number is incremeted by one per request or sometimes it can be null if your implemenation is buggy
   "event":"system",
   "payload":{
      "message":"{:error, \"Unable to subscribe to changes with given parameters. Please check Realtime is enabled for the given connect parameters: [event: *, filter: creator_id=eq.<user-id>, schema: public, table: posts]\"}",
      "status":"error",
      "extension":"postgres_changes",
      "channel":"public:posts:135"
   },
   "topic":"realtime:public:posts:135"
}
```

Response to presence state (?)
```json
{
   "ref":null, // number is incremeted by one per request or sometimes it can be null if your implemenation is buggy
   "event":"presence_state",
   "payload":{},
   "topic":"realtime:public:posts:135"
}
```

## Sign off endpoint
**URL:** https://sb.flaroapp.pl/auth/v1/logout?scope=local
**HTTP request type**: POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- scope: likely refers the logout type, if it's set to `local` then only logout from this device, otherwise assume log out from all devices, unsure though

**Body**: Can be a empty JSON body like `{}`

**Response (204)**: No content.

## Read messages from Global Channel endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/messages?select=%2A&order=created_at.asc.nullslast
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**URL prefix parameters:**
- select: (specified to view all messages aka *)
- order: sort by latest system message aka`created_at.asc.nullslast` 

**Response (200, showing last 5 messages)**: 

```json
[
   {
      "id":1,
      "sender_id":"69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
      "content":"siema",
      "created_at":"2025-09-25T16:02:45.52834+00:00"
   },
   {
      "id":2,
      "sender_id":"69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
    "content":"siemasiemsdnaklhdsalkdkasgdsadkasgdjkgfjksdgfjhsgjfkgkjgdhjsfghjsdgfjdsgjfdsgfgsdkfgsjkdgfjkdsgfksdgfksdgkfjsdgkfjsgkjfgkjsdgfksdgfksdgfksgfksgkfjsdgkfsgksfgfgdgkgfsfg",
      "created_at":"2025-09-25T16:17:51.079067+00:00"
   },
   {
      "id":3,
      "sender_id":"69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
      "content":"grajer",
      "created_at":"2025-09-25T16:26:56.137879+00:00"
   },
   {
      "id":4,
      "sender_id":"69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
      "content":"c",
      "created_at":"2025-09-25T16:26:59.059202+00:00"
   },
   {
      "id":5,
      "sender_id":"69fd6e93-4370-42ee-bb59-2ad2cbdf2cf7",
      "content":"c",
      "created_at":"2025-09-25T16:27:00.228039+00:00"
   }
]
```

## Send messages to Global Channel endpoint
**URL:** https://sb.flaroapp.pl/rest/v1/messages
**HTTP request type:** POST
**HTTP header parameters:** 
- `apikey`: `insert-obtained-api-key-from-app>`
  - `authorization` : `Bearer <user_access_token>`

**Body:**

```json
{
   "content":"hej", // your message content
   "sender_id":"d677717a-0b96-4071-b637-c589fbb75d4e", // your user ID
   "created_at":"2025-09-26T12:34:26.264261" // current timestamp is ISO 8601 format
}
```


**Response (201)**:  No content.
