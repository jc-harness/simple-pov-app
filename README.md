# santander-pov-app — the thing Scenario 1 builds & pushes

A deliberately tiny Go HTTP service (one file, no dependencies) with a hardened
multi-stage Dockerfile. Its only job is to be a *real* container for the
Build-and-Push template to build and push to ECR.

```
main.go         tiny net/http server on :8080 ("/" and "/healthz")
go.mod          module def, no deps
Dockerfile      multi-stage: golang -> distroless/static (nonroot)
.dockerignore   keep the build context small
```

## Push it to your GitHub so the pipeline can clone it
The pipeline's codebase connector is `account.santander_github`. Create a repo
under your GitHub org and push **the contents of this `app/` folder as the repo
root** (so the `Dockerfile` sits at the repo root — the template uses the default
`context: .` / `dockerfile: Dockerfile`).

```bash
# from this app/ directory:
git init -b main
git add .
git commit -m "santander-pov-app: tiny Go service + hardened Dockerfile"
git remote add origin https://github.com/<YOUR_GH_ORG>/santander-pov-app.git
git push -u origin main
```

Then, when you run the **App - Build and Push** pipeline, supply:
- `repoName` = `santander-pov-app` (or `<YOUR_GH_ORG>/santander-pov-app` per your connector)
- `build`    = `main`

The image lands at:
`759984737373.dkr.ecr.eu-west-1.amazonaws.com/santander-pov-app:<build#>`

## Verify the push afterwards
```bash
aws ecr list-images --repository-name santander-pov-app --region eu-west-1 \
  --profile AWSPowerUserAccess-759984737373
```
