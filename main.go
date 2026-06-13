// =============================================================================
// Tiny demo HTTP service — the "app" Scenario 1 builds & pushes to ECR.
// =============================================================================
// Deliberately minimal: no external dependencies, one file. The point is to give
// the Build-and-Push template something real to containerize, not to showcase Go.
// VERSION is injected at build time (-ldflags) or via env, so the running image
// can echo which build it came from — handy when we later deploy it.
// =============================================================================
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// version is overridable at build time:  -ldflags "-X main.version=<+pipeline.sequenceId>"
var version = "dev"

func main() {
	// Allow env override too (e.g. set from the image tag at deploy time).
	if v := os.Getenv("APP_VERSION"); v != "" {
		version = v
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Santander PoV app — build %s\n", version)
	})

	addr := ":8080"
	log.Printf("santander-pov-app version=%s listening on %s", version, addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
