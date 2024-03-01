# Contributing to the Goverland Platform Events Protocol

**Please note:** At Goverland, we prioritize the security of our platform 
and the trust of our users. If you believe you have discovered a security
vulnerability, please responsibly disclose it by contacting us at security@goverland.xyz.

**First and Foremost:** If you're uncertain or hesitant about anything, 
don't hesitate to ask questions or submit your issue or pull request. 
We welcome all contributions and appreciate your efforts. Our aim is 
to facilitate collaboration without imposing unnecessary barriers.

## Issues

This section outlines what we expect when reporting issues, ensuring smoother 
and quicker resolution.

### Reporting an Issue

*Before reporting, test against the latest released version as the issue may 
  have already been addressed. Testing against the `main` branch is even better,
  as it includes the latest fixes not yet released.

* Provide clear steps to reproduce the issue, along with expected and actual
  results. Whenever possible, include text and/or screenshots to aid in understanding.

* If the issue involves an internal error (e.g., a status code of 5xx), include 
  relevant portions or the entire log, as such errors may not be immediately 
  visible to users but are crucial for diagnosis.

* In case of a panic, create a [gist](https://gist.github.com) of the complete crash
  log for our review, ensuring no sensitive information is included.

* Respond promptly to any inquiries from the Goverland team regarding your reported issue.

### Issue Lifecycle

1. **Report:** The issue is reported.

2. **Verification and Categorization:** A Goverland collaborator verifies and categorizes
   the issue using appropriate tags (e.g., bugs for bug reports).

3. **Triaging and Community Engagement:** Non-critical issues may remain open for a period, 
   allowing community contributors to address them. We value community involvement
   in issue resolution.

4. **Resolution:** The issue is addressed through a pull request or commit, with clear 
   references linking the fix to the reported issue.

5. **Closure:** Upon resolution, the issue is closed.

6. **Stale Issue Handling:** Issues that remain unresolved or unresponsive for an extended 
   period are considered stale. To maintain an organized issue tracker, we may close stale 
   issues after 90 days. However, users are encouraged to reopen them if the issue 
   remains relevant.

## Pull requests

When submitting a pull request (PR), it's essential to reference an existing issue. 
If none exists, create one. Exceptions can be made for trivial PRs like typo fixes.

Creating an issue beforehand can prevent duplication of effort and allow for guidance 
or insights from the community.

Your PR should include a description of its purpose, implementation approach, and 
justification for the chosen approach. Additionally, include unit tests to validate 
correctness, ensuring existing tests pass. Corrections to tests do not require 
a new issue.

PRs undergo initial review to ensure compliance with guidelines outlined in this document. 
Incomplete PRs will be marked for follow-up to address missing requirements.

### Changelog Entries

Include changes from your PR in the CHANGELOG.md under the Unreleased section.
Refer to existing entries or [keepachangelog.com](https://keepachangelog.com/en/1.1.0/) 
for formatting examples.
