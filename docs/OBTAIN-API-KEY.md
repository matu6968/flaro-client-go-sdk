# How to obtain an API key for Flaro

> [!WARNING]
> Do not attempt to obtain any API keys on devices you do not own.

> [!NOTE]
> Any obtained keys should NOT BE DISTRBUTED other then for your own personal/private use.
> Flaro may at any time prevent the following instructions on obtaining the API key from working.

**Prerequisites**:

- Rooted Android phone
- HTTP Toolkit running on a host PC
- ADB debugging enabled (wireless or wired)

1. Download the Flaro app from the Play Store
2. Open HTTP Toolkit on the host PC and connect your phone to it via ADB
3. Open the Flaro app and sign in/sign up to a account
4. If done right, you should see requests coming from `sb.flaroapp.pl` and there should be a `apikey` field in the HTTP  headers per request.
5. Copy the `apikey` header field from the request.

**Timeline:**
- 2025-09-22 10:47: Notified developer about hard coded API keys in their beta testing program.
- 2025-09-22 13:55: Developer introduced anti abuse protection to the posting API's in open beta, likely breaks posting within this Go SDK.