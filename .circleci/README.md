# Carbonaut CI/CD Doc

[CircleCI](https://app.circleci.com/pipelines/github/carbonaut-cloud/carbonaut?branch=main) is used as platform to run the CI/CD pipeline. The flowchart below illustrates how the pipeline runs.

<img src="./circleci-logo.png" width="120" height="120" align="right" style="margin-left:32px"/>

## Current pipeline configuration

The CircleCI pipeline config can be found in the file `config.yml`

```mermaid
flowchart
    install --> verify[Verify] -->|branch:main| publish
    publish[Publish] -.- cr[(Container Registry)]

    classDef plain fill:#ddd,stroke:#464a81,stroke-width:4px,color:#000;
    classDef carbonaut fill:#464a81,stroke:#fff,stroke-width:4px,color:#fff;
    class GHCarbonaut,cr,storage plain;
    class install,publish,verify carbonaut;
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
    install[Install] --> verify[Verify] & e2eTest[End to End Test] & api-pen-test[API Security Testing] --> StoreTestArtifacts[Store Test Artifacts]
    StoreTestArtifacts -->|branch:main| SignArtifacts[Sign Artifacts] --> PublishContainerImages[Publish Container Images] --> CreateGitHubRelease[Create GitHub Release]
    PublishContainerImages -.- cr[(Container Registry)]
    StoreTestArtifacts -->|branch:main| GenerateReleaseNotes[Generate Release Notes] --> CreateGitHubRelease -.- GHCarbonaut([GitHub / Carbonaut])
    StoreTestArtifacts -.- storage[(Cloud Storage)]
    CreateGitHubRelease --> PublishAnnouncement[Publish Announcement]

    classDef plain fill:#ddd,stroke:#464a81,stroke-width:4px,color:#000;
    classDef carbonaut fill:#464a81,stroke:#fff,stroke-width:4px,color:#fff;
    class GHCarbonaut,cr,storage plain;
    class install,verify,e2eTest,api-pen-test,StoreTestArtifacts,CreateGitHubRelease,GenerateReleaseNotes,PublishAnnouncement,PublishContainerImages,SignArtifacts carbonaut;
    
```