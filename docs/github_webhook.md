# github webhook 

```bash
curl -X POST http://localhost:8080/github/webhook \
-H "Content-Type: application/json" \
-H "X-GitHub-Event: push" \
-H "X-Hub-Signature-256: sha256=123abc456def789..." \
-d '{
  "after": "383a1c1ebd6505eec62ba5e5033d0a26462fb298",
  "base_ref": null,
  "before": "d6f86c8d8bc52386213cfb513e6bc844651fc21a",
  "commits": [
    {
      "added": [],
      "author": {
        "email": "raphaeldias.ti@gmail.com",
        "name": "Raphael Oliveira",
        "username": "raphaelbh"
      },
      "committer": {
        "email": "noreply@github.com",
        "name": "GitHub",
        "username": "web-flow"
      },
      "distinct": true,
      "id": "383a1c1ebd6505eec62ba5e5033d0a26462fb298",
      "message": "Update README.md",
      "modified": [
        "README.md"
      ],
      "removed": [],
      "timestamp": "2024-10-29T23:41:00-03:00",
      "tree_id": "6ab2347b1f6f5cc91f497d4a65fbdf54b3770ae9",
      "url": "https://github.com/DevexPlatform/sample-api/commit/383a1c1ebd6505eec62ba5e5033d0a26462fb298"
    }
  ],
  "compare": "https://github.com/DevexPlatform/sample-api/compare/d6f86c8d8bc5...383a1c1ebd65",
  "created": false,
  "deleted": false,
  "forced": false,
  "head_commit": {
    "added": [],
    "author": {
      "email": "raphaeldias.ti@gmail.com",
      "name": "Raphael Oliveira",
      "username": "raphaelbh"
    },
    "committer": {
      "email": "noreply@github.com",
      "name": "GitHub",
      "username": "web-flow"
    },
    "distinct": true,
    "id": "383a1c1ebd6505eec62ba5e5033d0a26462fb298",
    "message": "Update README.md",
    "modified": [
      "README.md"
    ],
    "removed": [],
    "timestamp": "2024-10-29T23:41:00-03:00",
    "tree_id": "6ab2347b1f6f5cc91f497d4a65fbdf54b3770ae9",
    "url": "https://github.com/DevexPlatform/sample-api/commit/383a1c1ebd6505eec62ba5e5033d0a26462fb298"
  },
  "organization": {
    "avatar_url": "https://avatars.githubusercontent.com/u/186420057?v=4",
    "description": null,
    "events_url": "https://api.github.com/orgs/DevexPlatform/events",
    "hooks_url": "https://api.github.com/orgs/DevexPlatform/hooks",
    "id": 186420057,
    "issues_url": "https://api.github.com/orgs/DevexPlatform/issues",
    "login": "DevexPlatform",
    "members_url": "https://api.github.com/orgs/DevexPlatform/members{/member}",
    "node_id": "O_kgDOCxyLWQ",
    "public_members_url": "https://api.github.com/orgs/DevexPlatform/public_members{/member}",
    "repos_url": "https://api.github.com/orgs/DevexPlatform/repos",
    "url": "https://api.github.com/orgs/DevexPlatform"
  },
  "pusher": {
    "email": "raphaeldias.ti@gmail.com",
    "name": "raphaelbh"
  },
  "ref": "refs/heads/main",
  "repository": {
    "allow_forking": false,
    "archive_url": "https://api.github.com/repos/DevexPlatform/sample-api/{archive_format}{/ref}",
    "archived": false,
    "assignees_url": "https://api.github.com/repos/DevexPlatform/sample-api/assignees{/user}",
    "blobs_url": "https://api.github.com/repos/DevexPlatform/sample-api/git/blobs{/sha}",
    "branches_url": "https://api.github.com/repos/DevexPlatform/sample-api/branches{/branch}",
    "clone_url": "https://github.com/DevexPlatform/sample-api.git",
    "collaborators_url": "https://api.github.com/repos/DevexPlatform/sample-api/collaborators{/collaborator}",
    "comments_url": "https://api.github.com/repos/DevexPlatform/sample-api/comments{/number}",
    "commits_url": "https://api.github.com/repos/DevexPlatform/sample-api/commits{/sha}",
    "compare_url": "https://api.github.com/repos/DevexPlatform/sample-api/compare/{base}...{head}",
    "contents_url": "https://api.github.com/repos/DevexPlatform/sample-api/contents/{+path}",
    "contributors_url": "https://api.github.com/repos/DevexPlatform/sample-api/contributors",
    "created_at": 1729974629,
    "custom_properties": {},
    "default_branch": "main",
    "deployments_url": "https://api.github.com/repos/DevexPlatform/sample-api/deployments",
    "description": null,
    "disabled": false,
    "downloads_url": "https://api.github.com/repos/DevexPlatform/sample-api/downloads",
    "events_url": "https://api.github.com/repos/DevexPlatform/sample-api/events",
    "fork": false,
    "forks": 0,
    "forks_count": 0,
    "forks_url": "https://api.github.com/repos/DevexPlatform/sample-api/forks",
    "full_name": "DevexPlatform/sample-api",
    "git_commits_url": "https://api.github.com/repos/DevexPlatform/sample-api/git/commits{/sha}",
    "git_refs_url": "https://api.github.com/repos/DevexPlatform/sample-api/git/refs{/sha}",
    "git_tags_url": "https://api.github.com/repos/DevexPlatform/sample-api/git/tags{/sha}",
    "git_url": "git://github.com/DevexPlatform/sample-api.git",
    "has_discussions": false,
    "has_downloads": true,
    "has_issues": true,
    "has_pages": false,
    "has_projects": true,
    "has_wiki": false,
    "homepage": null,
    "hooks_url": "https://api.github.com/repos/DevexPlatform/sample-api/hooks",
    "html_url": "https://github.com/DevexPlatform/sample-api",
    "id": 879049086,
    "is_template": false,
    "issue_comment_url": "https://api.github.com/repos/DevexPlatform/sample-api/issues/comments{/number}",
    "issue_events_url": "https://api.github.com/repos/DevexPlatform/sample-api/issues/events{/number}",
    "issues_url": "https://api.github.com/repos/DevexPlatform/sample-api/issues{/number}",
    "keys_url": "https://api.github.com/repos/DevexPlatform/sample-api/keys{/key_id}",
    "labels_url": "https://api.github.com/repos/DevexPlatform/sample-api/labels{/name}",
    "language": null,
    "languages_url": "https://api.github.com/repos/DevexPlatform/sample-api/languages",
    "license": null,
    "master_branch": "main",
    "merges_url": "https://api.github.com/repos/DevexPlatform/sample-api/merges",
    "milestones_url": "https://api.github.com/repos/DevexPlatform/sample-api/milestones{/number}",
    "mirror_url": null,
    "name": "sample-api",
    "node_id": "R_kgDONGU5fg",
    "notifications_url": "https://api.github.com/repos/DevexPlatform/sample-api/notifications{?since,all,participating}",
    "open_issues": 0,
    "open_issues_count": 0,
    "organization": "DevexPlatform",
    "owner": {
      "avatar_url": "https://avatars.githubusercontent.com/u/186420057?v=4",
      "email": null,
      "events_url": "https://api.github.com/users/DevexPlatform/events{/privacy}",
      "followers_url": "https://api.github.com/users/DevexPlatform/followers",
      "following_url": "https://api.github.com/users/DevexPlatform/following{/other_user}",
      "gists_url": "https://api.github.com/users/DevexPlatform/gists{/gist_id}",
      "gravatar_id": "",
      "html_url": "https://github.com/DevexPlatform",
      "id": 186420057,
      "login": "DevexPlatform",
      "name": "DevexPlatform",
      "node_id": "O_kgDOCxyLWQ",
      "organizations_url": "https://api.github.com/users/DevexPlatform/orgs",
      "received_events_url": "https://api.github.com/users/DevexPlatform/received_events",
      "repos_url": "https://api.github.com/users/DevexPlatform/repos",
      "site_admin": false,
      "starred_url": "https://api.github.com/users/DevexPlatform/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/DevexPlatform/subscriptions",
      "type": "Organization",
      "url": "https://api.github.com/users/DevexPlatform",
      "user_view_type": "public"
    },
    "private": true,
    "pulls_url": "https://api.github.com/repos/DevexPlatform/sample-api/pulls{/number}",
    "pushed_at": 1730256061,
    "releases_url": "https://api.github.com/repos/DevexPlatform/sample-api/releases{/id}",
    "size": 0,
    "ssh_url": "git@github.com:DevexPlatform/sample-api.git",
    "stargazers": 0,
    "stargazers_count": 0,
    "stargazers_url": "https://api.github.com/repos/DevexPlatform/sample-api/stargazers",
    "statuses_url": "https://api.github.com/repos/DevexPlatform/sample-api/statuses/{sha}",
    "subscribers_url": "https://api.github.com/repos/DevexPlatform/sample-api/subscribers",
    "subscription_url": "https://api.github.com/repos/DevexPlatform/sample-api/subscription",
    "svn_url": "https://github.com/DevexPlatform/sample-api",
    "tags_url": "https://api.github.com/repos/DevexPlatform/sample-api/tags",
    "teams_url": "https://api.github.com/repos/DevexPlatform/sample-api/teams",
    "topics": [],
    "trees_url": "https://api.github.com/repos/DevexPlatform/sample-api/git/trees{/sha}",
    "updated_at": "2024-10-26T20:30:32Z",
    "url": "https://github.com/DevexPlatform/sample-api",
    "visibility": "private",
    "watchers": 0,
    "watchers_count": 0,
    "web_commit_signoff_required": false
  },
  "sender": {
    "avatar_url": "https://avatars.githubusercontent.com/u/119721?v=4",
    "events_url": "https://api.github.com/users/raphaelbh/events{/privacy}",
    "followers_url": "https://api.github.com/users/raphaelbh/followers",
    "following_url": "https://api.github.com/users/raphaelbh/following{/other_user}",
    "gists_url": "https://api.github.com/users/raphaelbh/gists{/gist_id}",
    "gravatar_id": "",
    "html_url": "https://github.com/raphaelbh",
    "id": 119721,
    "login": "raphaelbh",
    "node_id": "MDQ6VXNlcjExOTcyMQ==",
    "organizations_url": "https://api.github.com/users/raphaelbh/orgs",
    "received_events_url": "https://api.github.com/users/raphaelbh/received_events",
    "repos_url": "https://api.github.com/users/raphaelbh/repos",
    "site_admin": false,
    "starred_url": "https://api.github.com/users/raphaelbh/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/raphaelbh/subscriptions",
    "type": "User",
    "url": "https://api.github.com/users/raphaelbh",
    "user_view_type": "public"
  }
}
'
```