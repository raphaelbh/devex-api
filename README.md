# devex-api

## webhook request example
```bash
curl -X POST http://localhost:8080/github/webhook \
-H "Content-Type: application/json" \
-H "X-GitHub-Event: push" \
-H "X-Hub-Signature-256: sha256=123abc456def789..." \
-d '{
  "ref": "refs/heads/main",
  "before": "9c3aadebbc2f47394e523ffad1bbad52d51a599d",
  "after": "b6589fc6ab0dc82cf12099d1c2d40ab994e8410c",
  "commits": [
    {
      "id": "b6589fc6ab0dc82cf12099d1c2d40ab994e8410c",
      "message": "Adicionando nova feature",
      "timestamp": "2024-10-24T14:00:00Z",
      "url": "https://github.com/org/repo/commit/b6589fc6ab0dc82cf12099d1c2d40ab994e8410c",
      "author": {
        "name": "usuário",
        "email": "user@example.com"
      },
      "committer": {
        "name": "usuário",
        "email": "user@example.com"
      },
      "added": ["file1.txt"],
      "removed": [],
      "modified": ["file2.txt"]
    }
  ],
  "repository": {
    "id": 123456,
    "name": "repo",
    "full_name": "org/repo",
    "private": false,
    "owner": {
      "login": "org",
      "id": 789123,
      "url": "https://api.github.com/orgs/org"
    },
    "html_url": "https://github.com/org/repo",
    "default_branch": "main"
  },
  "pusher": {
    "name": "usuário",
    "email": "user@example.com"
  },
  "sender": {
    "login": "usuário",
    "id": 789456,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "url": "https://api.github.com/users/usuário"
  }
}'
```