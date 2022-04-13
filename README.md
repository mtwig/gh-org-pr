# Org PR 

The primary purpose of this extension is to show PRs in an org, where you are requested to review.

This extension shows is intended to support my own workflow.


## Mock output

```bash
$ gh org-pr --org ${ORG_NAME} --repo-filter ${REPO_NAME_REGEX} 
[ mockrepo ]
  - ✨ Add some awseome feature
    https://github.com/orgname/reponame/pull/10
    added by @tag (First Last)
    created 2.0 hours ago, modified 30 minutes ago
```