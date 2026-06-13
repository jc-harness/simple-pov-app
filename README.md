# simple-pov-app 

A deliberately tiny Go HTTP service (one file, no dependencies) with a hardened
multi-stage Dockerfile. Its only job is to be a *real* container for the
Build-and-Push template to build and push to ECR.

```
main.go         tiny net/http server on :8080 ("/" and "/healthz")
go.mod          module def, no deps
Dockerfile      multi-stage: golang -> distroless/static (nonroot)
.dockerignore   keep the build context small
```
