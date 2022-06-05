# Carbonaut CI/CD Doc

[CircleCI](https://app.circleci.com/pipelines/github/carbonaut-cloud/carbonaut?branch=main) is used as platform to run the CI/CD pipeline. The flowchart below illustrates how the pipeline runs.

<img src="./circleci-logo.png" width="120" height="120" align="right" style="margin-left:32px"/>

## Current pipeline configuration

The CircleCI pipeline config can be found in the file `config.yml`

```mermaid
flowchart
    install --> verify-go-mod & verify-git & verify-linting & verify-unit-tests & docker-lint -->|branch:main| publish
    install -->|branch:main| verify-go-build -->|branch:main| publish
    publish -.- cr[(Container Registry)]
```

## Planned pipeline configuration

The current pipeline configuration mainly implements the verification part that insures that changes made to the project are well tested. Everything behind that, basically cutting the release and pushing artifacts, are not configured. The planned pipeline configuration shown below configures these phases too.

**Phases in short**
1. Install dependencies
2. Verify (build, linting, unit tests)
3. Advanced testing (end to end tests, pen testing)
4. Push test artifacts 
5. Push signed container images
6. Cut release in GitHub
7. Announce new release

```mermaid
flowchart
    install --> verify --> e2eTest & api-pen-test --> StoreTestArtifacts
    StoreTestArtifacts -->|branch:main| SignArtifacts --> PublishContainerImages --> CreateGitHubRelease
    PublishContainerImages -.- cr[(Container Registry)]
    StoreTestArtifacts -->|branch:main| GenerateReleaseNotes --> CreateGitHubRelease -.- GitHub/Carbonaut
    StoreTestArtifacts -.- storage[(Cloud Storage)]
    CreateGitHubRelease --> PublishAnnouncement
```