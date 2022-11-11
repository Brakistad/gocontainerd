// +build tools

// Package tools records tool dependencies. It cannot actually be compiled.
package tools

import _ "github.com/cosmtrek/air"

/*
2022-11-11 01:46:23.564Z: #12 0.995 TARGET_GO_VERSION=1.19.3
2022-11-11 01:46:23.564Z: #12 1.104 Go already installed. Skipping.
2022-11-11 01:46:23.730Z: #12 1.104 Installing common Go tools...
2022-11-11 01:46:24.329Z: #12 1.795 go: downloading golang.org/x/tools/gopls v0.10.1
2022-11-11 01:46:24.329Z: #12 1.847 go: downloading golang.org/x/tools v0.3.0
2022-11-11 01:46:24.920Z: #12 2.343 go: downloading golang.org/x/tools v0.2.1-0.20221101170700-b5bc717366b2
2022-11-11 01:46:25.365Z: #12 2.810 go: downloading github.com/sergi/go-diff v1.1.0
2022-11-11 01:46:25.365Z: #12 2.813 go: downloading honnef.co/go/tools v0.3.3
2022-11-11 01:46:25.507Z: #12 2.966 go: downloading mvdan.cc/gofumpt v0.3.1
2022-11-11 01:46:25.507Z: #12 2.966 go: downloading mvdan.cc/xurls/v2 v2.4.0
2022-11-11 01:46:25.507Z: #12 3.049 go: downloading golang.org/x/mod v0.6.0
2022-11-11 01:46:25.675Z: #12 3.091 go: downloading golang.org/x/text v0.4.0
2022-11-11 01:46:25.675Z: #12 3.096 go: downloading golang.org/x/exp/typeparams v0.0.0-20220722155223-a9213eeb770e
2022-11-11 01:46:25.675Z: #12 3.133 go: downloading golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e
2022-11-11 01:46:25.786Z: #12 3.285 go: downloading github.com/google/go-cmp v0.5.8
2022-11-11 01:46:25.786Z: #12 3.330 go: downloading golang.org/x/sys v0.1.0
2022-11-11 01:46:26.234Z: #12 3.639 go: downloading golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
2022-11-11 01:46:26.234Z: #12 3.657 go: downloading golang.org/x/vuln v0.0.0-20221010193109-563322be2ea9
2022-11-11 01:46:26.431Z: #12 3.795 go: downloading github.com/BurntSushi/toml v1.2.0
2022-11-11 01:46:27.298Z: #12 4.706 github.com/sergi/go-diff/diffmatchpatch
2022-11-11 01:46:27.298Z: #12 4.716 golang.org/x/tools/gopls/internal/lsp/safetoken
2022-11-11 01:46:27.582Z: #12 5.050 golang.org/x/tools/internal/bug
2022-11-11 01:46:27.582Z: #12 5.115 golang.org/x/tools/gopls/internal/span
2022-11-11 01:46:27.834Z: #12 5.295 golang.org/x/tools/internal/event/label
2022-11-11 01:46:27.834Z: #12 5.378 golang.org/x/tools/internal/event/keys
2022-11-11 01:46:27.950Z: #12 5.481 golang.org/x/tools/internal/event/core
2022-11-11 01:46:28.089Z: #12 5.489 golang.org/x/tools/internal/event/tag
2022-11-11 01:46:28.089Z: #12 5.544 golang.org/x/tools/internal/event
2022-11-11 01:46:28.089Z: #12 5.575 golang.org/x/tools/internal/event/export
2022-11-11 01:46:28.257Z: #12 5.636 golang.org/x/tools/internal/xcontext
2022-11-11 01:46:28.534Z: #12 5.964 golang.org/x/tools/internal/jsonrpc2
2022-11-11 01:46:28.534Z: #12 5.974 golang.org/x/tools/internal/jsonrpc2_v2
2022-11-11 01:46:28.680Z: #12 6.218 golang.org/x/mod/internal/lazyregexp
2022-11-11 01:46:28.818Z: #12 6.289 golang.org/x/mod/semver
2022-11-11 01:46:28.818Z: #12 6.317 golang.org/x/tools/gopls/internal/lsp/protocol
2022-11-11 01:46:28.818Z: #12 6.355 golang.org/x/mod/module
2022-11-11 01:46:28.953Z: #12 6.499 golang.org/x/mod/modfile
2022-11-11 01:46:29.579Z: #12 7.050 golang.org/x/text/unicode/runenames
2022-11-11 01:46:29.881Z: #12 7.268 golang.org/x/tools/internal/analysisinternal
2022-11-11 01:46:29.990Z: #12 7.411 golang.org/x/tools/go/analysis
2022-11-11 01:46:29.990Z: #12 7.515 golang.org/x/tools/go/analysis/passes/internal/analysisutil
2022-11-11 01:46:30.129Z: #12 7.539 golang.org/x/tools/internal/typeparams
2022-11-11 01:46:30.129Z: #12 7.568 golang.org/x/tools/go/analysis/passes/asmdecl
2022-11-11 01:46:30.277Z: #12 7.749 golang.org/x/tools/go/ast/inspector
2022-11-11 01:46:30.419Z: #12 7.832 golang.org/x/tools/go/analysis/passes/inspect
2022-11-11 01:46:30.420Z: #12 7.865 golang.org/x/tools/go/analysis/passes/buildtag
2022-11-11 01:46:30.420Z: #12 7.876 golang.org/x/tools/go/analysis/passes/assign
2022-11-11 01:46:30.420Z: #12 7.959 golang.org/x/tools/go/analysis/passes/atomic
2022-11-11 01:46:30.586Z: #12 7.963 golang.org/x/tools/go/analysis/passes/atomicalign
2022-11-11 01:46:30.586Z: #12 8.035 golang.org/x/tools/go/analysis/passes/bools
2022-11-11 01:46:30.586Z: #12 8.039 golang.org/x/tools/go/analysis/passes/cgocall
2022-11-11 01:46:30.694Z: #12 8.113 golang.org/x/tools/go/analysis/passes/composite
2022-11-11 01:46:30.694Z: #12 8.185 golang.org/x/tools/go/analysis/passes/copylock
2022-11-11 01:46:30.694Z: #12 8.216 golang.org/x/tools/go/ast/astutil
2022-11-11 01:46:30.809Z: #12 8.349 golang.org/x/tools/go/analysis/passes/fieldalignment
2022-11-11 01:46:30.918Z: #12 8.462 golang.org/x/tools/go/analysis/passes/httpresponse
2022-11-11 01:46:31.057Z: #12 8.586 golang.org/x/tools/go/analysis/passes/ifaceassert
2022-11-11 01:46:31.173Z: #12 8.609 golang.org/x/tools/go/types/typeutil
2022-11-11 01:46:31.173Z: #12 8.699 golang.org/x/tools/go/cfg
2022-11-11 01:46:31.314Z: #12 8.842 golang.org/x/tools/go/analysis/passes/deepequalerrors
2022-11-11 01:46:31.449Z: #12 8.924 golang.org/x/tools/go/analysis/passes/errorsas
2022-11-11 01:46:31.449Z: #12 8.926 golang.org/x/tools/go/analysis/passes/loopclosure
2022-11-11 01:46:31.617Z: #12 9.003 golang.org/x/tools/go/analysis/passes/ctrlflow
2022-11-11 01:46:31.618Z: #12 9.015 golang.org/x/tools/go/analysis/passes/nilfunc
2022-11-11 01:46:31.618Z: #12 9.087 golang.org/x/tools/go/ssa
2022-11-11 01:46:31.618Z: #12 9.092 golang.org/x/tools/go/analysis/passes/lostcancel
2022-11-11 01:46:31.758Z: #12 9.220 golang.org/x/tools/go/analysis/passes/printf
2022-11-11 01:46:32.066Z: #12 9.504 golang.org/x/tools/go/analysis/passes/shadow
2022-11-11 01:46:32.205Z: #12 9.598 golang.org/x/tools/go/analysis/passes/shift
2022-11-11 01:46:32.205Z: #12 9.694 golang.org/x/tools/go/analysis/passes/sortslice
2022-11-11 01:46:32.374Z: #12 9.800 golang.org/x/tools/go/analysis/passes/stdmethods
2022-11-11 01:46:32.463Z: #12 9.916 golang.org/x/tools/go/analysis/passes/stringintconv
2022-11-11 01:46:32.463Z: #12 9.998 golang.org/x/tools/go/analysis/passes/structtag
2022-11-11 01:46:32.571Z: #12 10.10 golang.org/x/tools/go/analysis/passes/testinggoroutine
2022-11-11 01:46:32.882Z: #12 10.26 golang.org/x/tools/go/analysis/passes/tests
2022-11-11 01:46:33.018Z: #12 10.54 golang.org/x/tools/go/analysis/passes/timeformat
2022-11-11 01:46:33.133Z: #12 10.65 golang.org/x/tools/go/analysis/passes/unmarshal
2022-11-11 01:46:33.249Z: #12 10.79 golang.org/x/tools/go/analysis/passes/unreachable
2022-11-11 01:46:33.389Z: #12 10.93 golang.org/x/tools/go/analysis/passes/unsafeptr
2022-11-11 01:46:33.530Z: #12 11.07 golang.org/x/tools/go/analysis/passes/unusedresult
2022-11-11 01:46:33.697Z: #12 11.16 golang.org/x/sys/execabs
2022-11-11 01:46:33.839Z: #12 11.24 golang.org/x/tools/go/internal/pkgbits
2022-11-11 01:46:33.950Z: #12 11.38 golang.org/x/tools/go/analysis/passes/buildssa
2022-11-11 01:46:33.950Z: #12 11.45 golang.org/x/tools/internal/goroot
2022-11-11 01:46:33.950Z: #12 11.49 golang.org/x/tools/go/analysis/passes/nilness
2022-11-11 01:46:34.062Z: #12 11.50 golang.org/x/tools/go/analysis/passes/unusedwrite
2022-11-11 01:46:34.062Z: #12 11.61 golang.org/x/tools/go/internal/gcimporter
2022-11-11 01:46:34.230Z: #12 11.65 golang.org/x/tools/internal/gocommand
2022-11-11 01:46:34.373Z: #12 11.82 golang.org/x/tools/go/internal/packagesdriver
2022-11-11 01:46:34.373Z: #12 11.86 golang.org/x/tools/internal/packagesinternal
2022-11-11 01:46:34.373Z: #12 11.89 golang.org/x/tools/internal/typesinternal
2022-11-11 01:46:34.486Z: #12 11.96 golang.org/x/tools/gopls/internal/govulncheck/semver
2022-11-11 01:46:34.487Z: #12 11.99 golang.org/x/vuln/internal
2022-11-11 01:46:34.487Z: #12 12.02 golang.org/x/vuln/internal/derrors
2022-11-11 01:46:34.605Z: #12 12.05 golang.org/x/vuln/internal/web
2022-11-11 01:46:34.606Z: #12 12.10 golang.org/x/vuln/internal/semver
2022-11-11 01:46:34.606Z: #12 12.13 golang.org/x/vuln/osv
2022-11-11 01:46:35.050Z: #12 12.57 golang.org/x/tools/go/gcexportdata
2022-11-11 01:46:35.162Z: #12 12.64 golang.org/x/tools/go/packages
2022-11-11 01:46:35.162Z: #12 12.69 golang.org/x/vuln/client
2022-11-11 01:46:35.413Z: #12 12.95 golang.org/x/exp/constraints
2022-11-11 01:46:35.531Z: #12 12.98 golang.org/x/exp/slices
2022-11-11 01:46:35.531Z: #12 13.05 golang.org/x/tools/go/callgraph
2022-11-11 01:46:35.817Z: #12 13.21 golang.org/x/tools/go/buildutil
2022-11-11 01:46:35.817Z: #12 13.23 golang.org/x/tools/go/internal/cgo
2022-11-11 01:46:35.818Z: #12 13.33 golang.org/x/tools/go/callgraph/vta/internal/trie
2022-11-11 01:46:35.961Z: #12 13.39 golang.org/x/tools/go/loader
2022-11-11 01:46:36.106Z: #12 13.53 golang.org/x/tools/go/callgraph/vta
2022-11-11 01:46:36.242Z: #12 13.64 golang.org/x/tools/go/ssa/ssautil
2022-11-11 01:46:36.383Z: #12 13.79 golang.org/x/tools/go/callgraph/cha
2022-11-11 01:46:36.384Z: #12 13.91 golang.org/x/vuln/vulncheck/internal/gosym
2022-11-11 01:46:36.530Z: #12 13.97 golang.org/x/tools/gopls/internal/lsp/analysis/embeddirective
2022-11-11 01:46:36.530Z: #12 14.05 golang.org/x/tools/internal/fuzzy
2022-11-11 01:46:36.609Z: #12 14.15 golang.org/x/vuln/vulncheck/internal/binscan
2022-11-11 01:46:36.752Z: #12 14.18 golang.org/x/tools/gopls/internal/lsp/analysis/fillreturns
2022-11-11 01:46:36.752Z: #12 14.30 golang.org/x/tools/gopls/internal/lsp/analysis/fillstruct
2022-11-11 01:46:36.921Z: #12 14.33 golang.org/x/vuln/vulncheck
2022-11-11 01:46:37.041Z: #12 14.46 golang.org/x/tools/gopls/internal/lsp/analysis/infertypeargs
2022-11-11 01:46:37.041Z: #12 14.57 golang.org/x/tools/gopls/internal/lsp/analysis/nonewvars
2022-11-11 01:46:37.182Z: #12 14.66 golang.org/x/tools/gopls/internal/lsp/analysis/noresultvalues
2022-11-11 01:46:37.182Z: #12 14.71 golang.org/x/tools/gopls/internal/govulncheck
2022-11-11 01:46:37.295Z: #12 14.73 golang.org/x/tools/gopls/internal/lsp/analysis/simplifycompositelit
2022-11-11 01:46:37.295Z: #12 14.82 golang.org/x/tools/gopls/internal/lsp/analysis/simplifyrange
2022-11-11 01:46:37.410Z: #12 14.84 golang.org/x/tools/gopls/internal/lsp/analysis/simplifyslice
2022-11-11 01:46:37.410Z: #12 14.91 golang.org/x/tools/gopls/internal/lsp/analysis/stubmethods
2022-11-11 01:46:37.410Z: #12 14.93 golang.org/x/tools/gopls/internal/lsp/analysis/undeclaredname
2022-11-11 01:46:37.550Z: #12 15.08 golang.org/x/tools/gopls/internal/lsp/analysis/unusedparams
2022-11-11 01:46:37.662Z: #12 15.10 golang.org/x/tools/gopls/internal/lsp/analysis/unusedvariable
2022-11-11 01:46:37.662Z: #12 15.20 golang.org/x/tools/gopls/internal/lsp/analysis/useany
2022-11-11 01:46:37.774Z: #12 15.27 golang.org/x/tools/gopls/internal/lsp/command
2022-11-11 01:46:37.775Z: #12 15.31 golang.org/x/tools/gopls/internal/lsp/lsppos
2022-11-11 01:46:37.886Z: #12 15.37 golang.org/x/tools/internal/diff/lcs
2022-11-11 01:46:37.886Z: #12 15.43 golang.org/x/tools/internal/fastwalk
2022-11-11 01:46:38.198Z: #12 15.58 golang.org/x/tools/internal/gopathwalk
2022-11-11 01:46:38.198Z: #12 15.60 golang.org/x/tools/internal/diff
2022-11-11 01:46:38.198Z: #12 15.66 golang.org/x/tools/internal/imports
2022-11-11 01:46:38.198Z: #12 15.72 golang.org/x/tools/internal/diff/myers
2022-11-11 01:46:38.353Z: #12 15.78 golang.org/x/tools/refactor/satisfy
2022-11-11 01:46:38.632Z: #12 16.02 honnef.co/go/tools/analysis/lint
2022-11-11 01:46:38.798Z: #12 16.17 golang.org/x/exp/typeparams
2022-11-11 01:46:38.916Z: #12 16.44 honnef.co/go/tools/analysis/facts/generated
2022-11-11 01:46:39.022Z: #12 16.55 honnef.co/go/tools/go/ast/astutil
2022-11-11 01:46:39.162Z: #12 16.70 honnef.co/go/tools/go/types/typeutil
2022-11-11 01:46:39.304Z: #12 16.73 golang.org/x/tools/gopls/internal/lsp/source
2022-11-11 01:46:39.305Z: #12 16.83 honnef.co/go/tools/go/ir
2022-11-11 01:46:41.861Z: #12 19.38 honnef.co/go/tools/analysis/facts/tokenfile
2022-11-11 01:46:42.001Z: #12 19.46 honnef.co/go/tools/pattern
2022-11-11 01:46:42.282Z: #12 19.70 honnef.co/go/tools/go/ir/irutil
2022-11-11 01:46:42.592Z: #12 20.01 honnef.co/go/tools/internal/passes/buildir
2022-11-11 01:46:42.592Z: #12 20.11 honnef.co/go/tools/analysis/facts/purity
2022-11-11 01:46:42.709Z: #12 20.23 honnef.co/go/tools/analysis/report
2022-11-11 01:46:42.870Z: #12 20.32 honnef.co/go/tools/analysis/code
2022-11-11 01:46:42.870Z: #12 20.33 honnef.co/go/tools/analysis/edit
2022-11-11 01:46:42.870Z: #12 20.38 honnef.co/go/tools/knowledge
2022-11-11 01:46:43.125Z: #12 20.56 honnef.co/go/tools/internal/sharedcheck
2022-11-11 01:46:43.125Z: #12 20.56 honnef.co/go/tools/analysis/facts/deprecated
2022-11-11 01:46:43.125Z: #12 20.65 honnef.co/go/tools/analysis/facts/nilness
2022-11-11 01:46:43.269Z: #12 20.68 honnef.co/go/tools/quickfix
2022-11-11 01:46:43.409Z: #12 20.88 honnef.co/go/tools/simple
2022-11-11 01:46:43.549Z: #12 20.97 honnef.co/go/tools/analysis/facts/typedness
2022-11-11 01:46:43.549Z: #12 21.09 honnef.co/go/tools/printf
2022-11-11 01:46:43.674Z: #12 21.14 honnef.co/go/tools/staticcheck/fakereflect
2022-11-11 01:46:43.675Z: #12 21.22 honnef.co/go/tools/staticcheck/fakejson
2022-11-11 01:46:43.982Z: #12 21.39 honnef.co/go/tools/staticcheck/fakexml
2022-11-11 01:46:44.122Z: #12 21.56 github.com/BurntSushi/toml/internal
2022-11-11 01:46:44.122Z: #12 21.57 honnef.co/go/tools/staticcheck
2022-11-11 01:46:44.122Z: #12 21.58 github.com/BurntSushi/toml
2022-11-11 01:46:44.893Z: #12 22.41 honnef.co/go/tools/config
2022-11-11 01:46:45.174Z: #12 22.59 honnef.co/go/tools/stylecheck
2022-11-11 01:46:45.621Z: #12 23.05 github.com/google/go-cmp/cmp/internal/flags
2022-11-11 01:46:45.621Z: #12 23.05 github.com/google/go-cmp/cmp/internal/diff
2022-11-11 01:46:45.621Z: #12 23.10 github.com/google/go-cmp/cmp/internal/function
2022-11-11 01:46:45.621Z: #12 23.14 github.com/google/go-cmp/cmp/internal/value
2022-11-11 01:46:45.730Z: #12 23.15 mvdan.cc/gofumpt/internal/version
2022-11-11 01:46:45.730Z: #12 23.19 mvdan.cc/xurls/v2
2022-11-11 01:46:45.730Z: #12 23.26 golang.org/x/sync/errgroup
2022-11-11 01:46:45.841Z: #12 23.28 github.com/google/go-cmp/cmp
2022-11-11 01:46:45.841Z: #12 23.30 golang.org/x/tools/gopls/internal/lsp/progress
2022-11-11 01:46:45.841Z: #12 23.38 golang.org/x/tools/internal/memoize
2022-11-11 01:46:45.956Z: #12 23.45 golang.org/x/tools/internal/persistent
2022-11-11 01:46:45.956Z: #12 23.49 golang.org/x/tools/gopls/internal/lsp/cache
2022-11-11 01:46:46.860Z: #12 24.27 mvdan.cc/gofumpt/format
2022-11-11 01:46:47.166Z: #12 24.61 golang.org/x/tools/gopls/internal/hooks
2022-11-11 01:46:47.449Z: #12 24.85 golang.org/x/tools/gopls/internal/lsp/debug/log
2022-11-11 01:46:47.449Z: #12 24.89 golang.org/x/tools/internal/event/export/metric
2022-11-11 01:46:47.626Z: #12 25.08 golang.org/x/tools/internal/event/export/ocagent/wire
2022-11-11 01:46:47.766Z: #12 25.15 golang.org/x/tools/internal/event/export/ocagent
2022-11-11 01:46:47.906Z: #12 25.39 golang.org/x/tools/internal/event/export/prometheus
2022-11-11 01:46:48.078Z: #12 25.53 golang.org/x/tools/gopls/internal/lsp/mod
2022-11-11 01:46:48.330Z: #12 25.82 golang.org/x/tools/gopls/internal/lsp/snippet
2022-11-11 01:46:48.330Z: #12 25.86 golang.org/x/tools/gopls/internal/lsp/source/completion
2022-11-11 01:46:48.469Z: #12 25.88 golang.org/x/tools/gopls/internal/lsp/debug
2022-11-11 01:46:49.086Z: #12 26.49 golang.org/x/tools/gopls/internal/lsp/template
2022-11-11 01:46:49.533Z: #12 26.98 golang.org/x/tools/gopls/internal/lsp/work
2022-11-11 01:46:49.533Z: #12 27.06 golang.org/x/exp/maps
2022-11-11 01:46:49.646Z: #12 27.09 golang.org/x/vuln/internal/govulncheck
2022-11-11 01:46:49.646Z: #12 27.15 golang.org/x/tools/gopls/internal/lsp/browser
2022-11-11 01:46:49.646Z: #12 27.19 golang.org/x/tools/internal/fakenet
2022-11-11 01:46:49.815Z: #12 27.27 golang.org/x/tools/internal/tool
2022-11-11 01:46:49.926Z: #12 27.39 golang.org/x/vuln/exp/govulncheck
2022-11-11 01:46:49.926Z: #12 27.40 golang.org/x/tools/gopls/internal/vulncheck
2022-11-11 01:46:49.926Z: #12 27.47 golang.org/x/tools/gopls/internal/lsp
2022-11-11 01:46:50.705Z: #12 28.22 golang.org/x/tools/gopls/internal/lsp/lsprpc
2022-11-11 01:46:50.985Z: #12 28.46 golang.org/x/tools/gopls/internal/lsp/cmd
2022-11-11 01:46:51.586Z: #12 29.05 golang.org/x/tools/gopls
2022-11-11 01:46:54.720Z: #12 32.25 go: downloading golang.org/x/tools v0.1.11-0.20220513221640-090b14e8501f
2022-11-11 01:46:55.418Z: #12 32.86 go: downloading golang.org/x/exp/typeparams v0.0.0-20220218215828-6cf2b201936e
2022-11-11 01:46:55.418Z: #12 32.95 go: downloading golang.org/x/sys v0.0.0-20211019181941-9d821ace8654
2022-11-11 01:46:55.542Z: #12 32.95 go: downloading github.com/BurntSushi/toml v0.4.1
2022-11-11 01:46:55.542Z: #12 33.08 go: downloading golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4
2022-11-11 01:46:56.137Z: #12 33.53 honnef.co/go/tools/internal/sync
2022-11-11 01:46:56.137Z: #12 33.55 honnef.co/go/tools/sarif
2022-11-11 01:46:56.707Z: #12 34.18 golang.org/x/tools/internal/lsp/fuzzy
2022-11-11 01:46:56.707Z: #12 34.24 golang.org/x/tools/go/buildutil
2022-11-11 01:46:56.868Z: #12 34.26 golang.org/x/tools/internal/typeparams
2022-11-11 01:46:57.007Z: #12 34.40 golang.org/x/sys/execabs
2022-11-11 01:46:57.007Z: #12 34.44 golang.org/x/tools/go/ast/astutil
2022-11-11 01:46:57.007Z: #12 34.45 golang.org/x/tools/go/internal/gcimporter
2022-11-11 01:46:57.303Z: #12 34.82 golang.org/x/tools/internal/analysisinternal
2022-11-11 01:46:57.414Z: #12 34.96 golang.org/x/tools/go/analysis
2022-11-11 01:46:57.630Z: #12 35.06 golang.org/x/mod/semver
2022-11-11 01:46:57.697Z: #12 35.13 golang.org/x/tools/internal/event/label
2022-11-11 01:46:57.697Z: #12 35.22 golang.org/x/tools/internal/event/keys
2022-11-11 01:46:57.959Z: #12 35.41 golang.org/x/tools/internal/event/core
2022-11-11 01:46:57.959Z: #12 35.46 golang.org/x/tools/go/gcexportdata
2022-11-11 01:46:57.959Z: #12 35.47 golang.org/x/tools/internal/event
2022-11-11 01:46:58.069Z: #12 35.54 golang.org/x/tools/internal/typesinternal
2022-11-11 01:46:58.069Z: #12 35.55 golang.org/x/tools/internal/gocommand
2022-11-11 01:46:58.069Z: #12 35.61 honnef.co/go/tools/analysis/lint
2022-11-11 01:46:58.181Z: #12 35.69 golang.org/x/tools/go/internal/packagesdriver
2022-11-11 01:46:58.181Z: #12 35.71 golang.org/x/tools/internal/packagesinternal
2022-11-11 01:46:58.325Z: #12 35.74 github.com/BurntSushi/toml/internal
2022-11-11 01:46:58.325Z: #12 35.74 golang.org/x/tools/go/packages
2022-11-11 01:46:58.325Z: #12 35.77 github.com/BurntSushi/toml
2022-11-11 01:46:59.086Z: #12 36.53 honnef.co/go/tools/go/buildid
2022-11-11 01:46:59.225Z: #12 36.62 honnef.co/go/tools/config
2022-11-11 01:46:59.225Z: #12 36.66 golang.org/x/exp/typeparams
2022-11-11 01:46:59.334Z: #12 36.81 honnef.co/go/tools/internal/robustio
2022-11-11 01:46:59.334Z: #12 36.82 golang.org/x/tools/go/types/objectpath
2022-11-11 01:46:59.334Z: #12 36.84 honnef.co/go/tools/internal/renameio
2022-11-11 01:46:59.334Z: #12 36.88 honnef.co/go/tools/lintcmd/cache
2022-11-11 01:46:59.474Z: #12 37.02 honnef.co/go/tools/analysis/facts/generated
2022-11-11 01:46:59.665Z: #12 37.03 honnef.co/go/tools/go/loader
2022-11-11 01:46:59.665Z: #12 37.10 honnef.co/go/tools/go/ast/astutil
2022-11-11 01:46:59.772Z: #12 37.17 golang.org/x/tools/go/ast/inspector
2022-11-11 01:46:59.772Z: #12 37.25 golang.org/x/tools/go/analysis/passes/inspect
2022-11-11 01:46:59.772Z: #12 37.25 honnef.co/go/tools/analysis/report
2022-11-11 01:46:59.772Z: #12 37.30 golang.org/x/tools/go/types/typeutil
2022-11-11 01:46:59.906Z: #12 37.39 golang.org/x/tools/go/internal/cgo
2022-11-11 01:47:00.046Z: #12 37.48 honnef.co/go/tools/go/types/typeutil
2022-11-11 01:47:00.046Z: #12 37.49 golang.org/x/tools/go/loader
2022-11-11 01:47:00.046Z: #12 37.57 honnef.co/go/tools/go/ir
2022-11-11 01:47:00.326Z: #12 37.73 honnef.co/go/tools/analysis/facts/tokenfile
2022-11-11 01:47:00.326Z: #12 37.79 honnef.co/go/tools/pattern
2022-11-11 01:47:01.198Z: #12 38.74 honnef.co/go/tools/analysis/facts/directives
2022-11-11 01:47:01.366Z: #12 38.81 honnef.co/go/tools/lintcmd/version
2022-11-11 01:47:01.506Z: #12 38.89 honnef.co/go/tools/analysis/edit
2022-11-11 01:47:01.506Z: #12 38.99 honnef.co/go/tools/analysis/facts/deprecated
2022-11-11 01:47:01.645Z: #12 39.11 honnef.co/go/tools/staticcheck/fakejson
2022-11-11 01:47:01.818Z: #12 39.28 honnef.co/go/tools/staticcheck/fakexml
2022-11-11 01:47:02.543Z: #12 40.07 honnef.co/go/tools/go/ir/irutil
2022-11-11 01:47:02.682Z: #12 40.09 honnef.co/go/tools/internal/passes/buildir
2022-11-11 01:47:02.682Z: #12 40.21 honnef.co/go/tools/analysis/facts/nilness
2022-11-11 01:47:02.795Z: #12 40.28 honnef.co/go/tools/analysis/facts/purity
2022-11-11 01:47:02.795Z: #12 40.34 honnef.co/go/tools/analysis/facts/typedness
2022-11-11 01:47:02.961Z: #12 40.40 honnef.co/go/tools/analysis/code
2022-11-11 01:47:03.102Z: #12 40.53 honnef.co/go/tools/unused
2022-11-11 01:47:03.102Z: #12 40.53 honnef.co/go/tools/internal/sharedcheck
2022-11-11 01:47:03.246Z: #12 40.67 honnef.co/go/tools/quickfix
2022-11-11 01:47:03.527Z: #12 40.99 honnef.co/go/tools/simple
2022-11-11 01:47:03.527Z: #12 41.05 honnef.co/go/tools/lintcmd/runner
2022-11-11 01:47:04.113Z: #12 41.52 honnef.co/go/tools/lintcmd
2022-11-11 01:47:04.281Z: #12 41.70 honnef.co/go/tools/staticcheck
2022-11-11 01:47:04.566Z: #12 42.05 honnef.co/go/tools/stylecheck
2022-11-11 01:47:05.470Z: #12 42.92 honnef.co/go/tools/cmd/staticcheck
2022-11-11 01:47:06.339Z: #12 43.79 go: downloading golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
2022-11-11 01:47:06.339Z: #12 43.88 go: downloading golang.org/x/tools v0.0.0-20200130002326-2f3ba24bd6e7
2022-11-11 01:47:06.962Z: #12 44.35 golang.org/x/tools/go/ast/astutil
2022-11-11 01:47:06.962Z: #12 44.37 golang.org/x/tools/go/internal/gcimporter
2022-11-11 01:47:07.546Z: #12 44.99 golang.org/x/tools/go/gcexportdata
2022-11-11 01:47:07.546Z: #12 45.03 golang.org/x/lint
2022-11-11 01:47:07.685Z: #12 45.23 golang.org/x/lint/golint
2022-11-11 01:47:08.582Z: #12 46.11 go: downloading github.com/mgechev/revive v1.2.4
2022-11-11 01:47:08.696Z: #12 46.19 go: downloading github.com/fatih/color v1.13.0
2022-11-11 01:47:08.696Z: #12 46.19 go: downloading github.com/mitchellh/go-homedir v1.1.0
2022-11-11 01:47:08.696Z: #12 46.22 go: downloading github.com/mgechev/dots v0.0.0-20210922191527-e955255bf517
2022-11-11 01:47:08.780Z: #12 46.22 go: downloading github.com/pkg/errors v0.9.1
2022-11-11 01:47:08.780Z: #12 46.24 go: downloading github.com/mattn/go-colorable v0.1.9
2022-11-11 01:47:08.780Z: #12 46.25 go: downloading github.com/mattn/go-isatty v0.0.14
2022-11-11 01:47:08.780Z: #12 46.28 go: downloading github.com/chavacava/garif v0.0.0-20220630083739-93517212f375
2022-11-11 01:47:08.780Z: #12 46.29 go: downloading github.com/olekukonko/tablewriter v0.0.5
2022-11-11 01:47:08.780Z: #12 46.32 go: downloading github.com/fatih/structtag v1.2.0
2022-11-11 01:47:08.949Z: #12 46.34 go: downloading golang.org/x/tools v0.1.12
2022-11-11 01:47:08.949Z: #12 46.37 go: downloading golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f
2022-11-11 01:47:09.227Z: #12 46.72 go: downloading github.com/mattn/go-runewidth v0.0.9
2022-11-11 01:47:09.699Z: #12 47.22 golang.org/x/sys/internal/unsafeheader
2022-11-11 01:47:09.782Z: #12 47.26 github.com/chavacava/garif
2022-11-11 01:47:09.782Z: #12 47.26 golang.org/x/sys/unix
2022-11-11 01:47:09.782Z: #12 47.32 github.com/mgechev/revive/internal/typeparams
2022-11-11 01:47:09.953Z: #12 47.40 github.com/mgechev/revive/lint
2022-11-11 01:47:10.205Z: #12 47.64 github.com/mattn/go-runewidth
2022-11-11 01:47:10.205Z: #12 47.73 github.com/olekukonko/tablewriter
2022-11-11 01:47:10.535Z: #12 47.98 github.com/fatih/structtag
2022-11-11 01:47:10.653Z: #12 48.09 golang.org/x/tools/internal/typeparams
2022-11-11 01:47:10.770Z: #12 48.32 golang.org/x/tools/go/ast/astutil
2022-11-11 01:47:11.190Z: #12 48.64 github.com/mattn/go-isatty
2022-11-11 01:47:11.190Z: #12 48.69 github.com/mattn/go-colorable
2022-11-11 01:47:11.191Z: #12 48.72 github.com/mgechev/revive/rule
2022-11-11 01:47:11.330Z: #12 48.73 github.com/fatih/color
2022-11-11 01:47:11.330Z: #12 48.86 github.com/mgechev/revive/formatter
2022-11-11 01:47:11.638Z: #12 49.05 github.com/mgechev/dots
2022-11-11 01:47:11.782Z: #12 49.22 github.com/mgechev/revive/logging
2022-11-11 01:47:11.782Z: #12 49.24 github.com/pkg/errors
2022-11-11 01:47:11.922Z: #12 49.35 github.com/mitchellh/go-homedir
2022-11-11 01:47:12.483Z: #12 49.94 github.com/mgechev/revive/config
2022-11-11 01:47:12.484Z: #12 49.99 github.com/mgechev/revive/revivelib
2022-11-11 01:47:12.484Z: #12 50.03 github.com/mgechev/revive/cli
2022-11-11 01:47:12.651Z: #12 50.07 github.com/mgechev/revive
2022-11-11 01:47:13.211Z: #12 50.71 go: downloading github.com/uudashr/gopkgs/v2 v2.1.2
2022-11-11 01:47:13.211Z: #12 50.76 go: downloading github.com/uudashr/gopkgs v2.0.1+incompatible
2022-11-11 01:47:13.352Z: #12 50.81 go: downloading github.com/karrick/godirwalk v1.12.0
2022-11-11 01:47:13.352Z: #12 50.81 go: downloading github.com/pkg/errors v0.8.1
2022-11-11 01:47:13.352Z: #12 50.89 github.com/pkg/errors
2022-11-11 01:47:13.491Z: #12 50.89 github.com/karrick/godirwalk
2022-11-11 01:47:13.491Z: #12 51.01 github.com/uudashr/gopkgs/v2
2022-11-11 01:47:13.630Z: #12 51.08 github.com/uudashr/gopkgs/v2/cmd/gopkgs
2022-11-11 01:47:14.081Z: #12 51.50 go: downloading github.com/ramya-rao-a/go-outline v0.0.0-20210608161538-9736a4bde949
2022-11-11 01:47:14.222Z: #12 51.67 go: downloading golang.org/x/tools v0.1.1
2022-11-11 01:47:14.642Z: #12 52.19 golang.org/x/tools/go/buildutil
2022-11-11 01:47:14.758Z: #12 52.30 github.com/ramya-rao-a/go-outline
2022-11-11 01:47:15.233Z: #12 52.70 go: downloading github.com/go-delve/delve v1.9.1
2022-11-11 01:47:17.002Z: #12 54.54 go: downloading github.com/sirupsen/logrus v1.6.0
2022-11-11 01:47:17.114Z: #12 54.54 go: downloading github.com/spf13/cobra v1.1.3
2022-11-11 01:47:17.114Z: #12 54.58 go: downloading github.com/mattn/go-isatty v0.0.3
2022-11-11 01:47:17.114Z: #12 54.59 go: downloading golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab
2022-11-11 01:47:17.114Z: #12 54.60 go: downloading gopkg.in/yaml.v2 v2.4.0
2022-11-11 01:47:17.114Z: #12 54.62 go: downloading github.com/cosiner/argv v0.1.0
2022-11-11 01:47:17.114Z: #12 54.64 go: downloading github.com/derekparker/trie v0.0.0-20200317170641-1fdf38b7b0e9
2022-11-11 01:47:17.229Z: #12 54.65 go: downloading github.com/go-delve/liner v1.2.3-0.20220127212407-d32d89dd2a5d
2022-11-11 01:47:17.229Z: #12 54.70 go: downloading github.com/google/go-dap v0.6.0
2022-11-11 01:47:17.229Z: #12 54.72 go: downloading github.com/cpuguy83/go-md2man/v2 v2.0.0
2022-11-11 01:47:17.229Z: #12 54.73 go: downloading github.com/spf13/pflag v1.0.5
2022-11-11 01:47:17.229Z: #12 54.76 go: downloading go.starlark.net v0.0.0-20220816155156-cfacd8902214
2022-11-11 01:47:17.366Z: #12 54.83 go: downloading github.com/mattn/go-runewidth v0.0.13
2022-11-11 01:47:17.367Z: #12 54.84 go: downloading github.com/hashicorp/golang-lru v0.5.4
2022-11-11 01:47:17.367Z: #12 54.85 go: downloading golang.org/x/arch v0.0.0-20190927153633-4e8777c89be4
2022-11-11 01:47:17.536Z: #12 54.96 go: downloading github.com/russross/blackfriday/v2 v2.0.1
2022-11-11 01:47:17.536Z: #12 54.96 go: downloading github.com/rivo/uniseg v0.2.0
2022-11-11 01:47:17.537Z: #12 54.99 go: downloading github.com/cilium/ebpf v0.7.0
2022-11-11 01:47:17.646Z: #12 55.07 go: downloading github.com/shurcooL/sanitized_anchor_name v1.0.0
2022-11-11 01:47:17.646Z: #12 55.15 golang.org/x/sys/internal/unsafeheader
2022-11-11 01:47:17.647Z: #12 55.18 golang.org/x/sys/unix
2022-11-11 01:47:17.786Z: #12 55.20 github.com/go-delve/delve/pkg/dwarf/util
2022-11-11 01:47:17.786Z: #12 55.31 github.com/go-delve/delve/pkg/dwarf/op
2022-11-11 01:47:18.066Z: #12 55.51 github.com/go-delve/delve/pkg/dwarf/godwarf
2022-11-11 01:47:18.514Z: #12 55.96 github.com/go-delve/delve/pkg/astutil
2022-11-11 01:47:18.515Z: #12 55.98 github.com/go-delve/delve/pkg/dwarf/frame
2022-11-11 01:47:18.626Z: #12 56.17 github.com/go-delve/delve/pkg/dwarf/line
2022-11-11 01:47:18.907Z: #12 56.38 github.com/go-delve/delve/pkg/dwarf/loclist
2022-11-11 01:47:18.907Z: #12 56.44 github.com/go-delve/delve/pkg/dwarf/reader
2022-11-11 01:47:19.046Z: #12 56.54 github.com/go-delve/delve/pkg/dwarf/regnum
2022-11-11 01:47:19.046Z: #12 56.59 github.com/go-delve/delve/pkg/elfwriter
2022-11-11 01:47:19.187Z: #12 56.67 github.com/go-delve/delve/pkg/proc/debuginfod
2022-11-11 01:47:19.187Z: #12 56.71 github.com/go-delve/delve/pkg/version
2022-11-11 01:47:19.328Z: #12 56.71 github.com/sirupsen/logrus
2022-11-11 01:47:19.329Z: #12 56.79 github.com/cilium/ebpf/internal/unix
2022-11-11 01:47:19.470Z: #12 56.88 github.com/cilium/ebpf/asm
2022-11-11 01:47:19.723Z: #12 57.24 github.com/go-delve/delve/pkg/logflags
2022-11-11 01:47:19.723Z: #12 57.27 github.com/cilium/ebpf/internal
2022-11-11 01:47:19.894Z: #12 57.36 github.com/go-delve/delve/pkg/goversion
2022-11-11 01:47:20.041Z: #12 57.45 github.com/hashicorp/golang-lru/simplelru
2022-11-11 01:47:20.041Z: #12 57.51 github.com/cilium/ebpf/internal/btf
2022-11-11 01:47:20.041Z: #12 57.51 golang.org/x/arch/arm64/arm64asm
2022-11-11 01:47:20.793Z: #12 58.17 github.com/cilium/ebpf
2022-11-11 01:47:20.930Z: #12 58.37 golang.org/x/arch/x86/x86asm
2022-11-11 01:47:21.490Z: #12 59.03 github.com/cilium/ebpf/link
2022-11-11 01:47:21.801Z: #12 59.22 github.com/cilium/ebpf/ringbuf
2022-11-11 01:47:21.942Z: #12 59.33 gopkg.in/yaml.v2
2022-11-11 01:47:21.942Z: #12 59.40 github.com/go-delve/delve/pkg/proc/internal/ebpf
2022-11-11 01:47:22.110Z: #12 59.52 github.com/go-delve/delve/pkg/proc
2022-11-11 01:47:23.154Z: #12 60.67 github.com/cosiner/argv
2022-11-11 01:47:23.398Z: #12 60.84 github.com/derekparker/trie
2022-11-11 01:47:23.399Z: #12 60.94 github.com/go-delve/delve/pkg/terminal/colorize
2022-11-11 01:47:23.549Z: #12 61.07 github.com/go-delve/delve/pkg/proc/macutil
2022-11-11 01:47:23.679Z: #12 61.13 github.com/mattn/go-isatty
2022-11-11 01:47:23.679Z: 
2022-11-11 01:47:23.846Z: #12 61.29 github.com/rivo/uniseg
2022-11-11 01:47:24.011Z: #12 61.41 github.com/mattn/go-runewidth
2022-11-11 01:47:24.133Z: #12 61.54 github.com/go-delve/liner
2022-11-11 01:47:24.590Z: #12 62.00 go.starlark.net/internal/spell
2022-11-11 01:47:24.591Z: #12 62.03 go.starlark.net/syntax
2022-11-11 01:47:25.178Z: #12 62.64 go.starlark.net/resolve
2022-11-11 01:47:25.178Z: #12 62.70 github.com/go-delve/delve/service/api
2022-11-11 01:47:25.430Z: #12 62.87 github.com/go-delve/delve/pkg/proc/amd64util
2022-11-11 01:47:25.430Z: #12 62.97 github.com/go-delve/delve/pkg/proc/winutil
2022-11-11 01:47:25.598Z: #12 63.01 github.com/go-delve/delve/pkg/config
2022-11-11 01:47:25.598Z: #12 63.05 github.com/go-delve/delve/pkg/locspec
2022-11-11 01:47:25.709Z: #12 63.12 github.com/go-delve/delve/pkg/gobuild
2022-11-11 01:47:25.710Z: #12 63.19 github.com/go-delve/delve/pkg/proc/core/minidump
2022-11-11 01:47:25.710Z: #12 63.23 github.com/go-delve/delve/pkg/proc/linutil
2022-11-11 01:47:25.850Z: #12 63.32 go.starlark.net/internal/compile
2022-11-11 01:47:25.994Z: #12 63.44 github.com/go-delve/delve/pkg/proc/core
2022-11-11 01:47:26.309Z: #12 63.73 github.com/go-delve/delve/pkg/proc/gdbserial
2022-11-11 01:47:26.310Z: #12 63.74 github.com/go-delve/delve/pkg/proc/native
2022-11-11 01:47:26.754Z: #12 64.23 go.starlark.net/starlark
2022-11-11 01:47:27.017Z: #12 64.54 github.com/go-delve/delve/service/debugger
2022-11-11 01:47:27.746Z: #12 65.22 github.com/go-delve/delve/service
2022-11-11 01:47:27.913Z: #12 65.35 github.com/go-delve/delve/service/rpc2
2022-11-11 01:47:28.342Z: #12 65.81 github.com/go-delve/delve/service/internal/sameuser
2022-11-11 01:47:28.343Z: #12 65.89 github.com/google/go-dap
2022-11-11 01:47:28.510Z: #12 66.03 github.com/go-delve/delve/pkg/terminal/starbind
2022-11-11 01:47:28.790Z: #12 66.26 github.com/go-delve/delve/service/dap
2022-11-11 01:47:29.450Z: #12 66.86 github.com/go-delve/delve/pkg/terminal
2022-11-11 01:47:29.842Z: #12 67.27 github.com/go-delve/delve/service/rpc1
2022-11-11 01:47:30.023Z: #12 67.54 github.com/go-delve/delve/service/rpccommon
2022-11-11 01:47:30.266Z: #12 67.74 github.com/spf13/pflag
2022-11-11 01:47:30.266Z: #12 67.79 github.com/shurcooL/sanitized_anchor_name
2022-11-11 01:47:30.406Z: #12 67.82 github.com/russross/blackfriday/v2
2022-11-11 01:47:31.162Z: #12 68.61 github.com/cpuguy83/go-md2man/v2/md2man
2022-11-11 01:47:31.302Z: #12 68.79 github.com/spf13/cobra
2022-11-11 01:47:31.611Z: #12 69.07 github.com/go-delve/delve/cmd/dlv/cmds
2022-11-11 01:47:31.611Z: #12 69.07 github.com/spf13/cobra/doc
2022-11-11 01:47:31.918Z: #12 69.34 github.com/go-delve/delve/cmd/dlv
2022-11-11 01:47:33.406Z: #12 70.84 Installing golangci-lint latest...
2022-11-11 01:47:33.406Z: #12 70.84 sh: 0: getcwd() failed: No such file or directory
2022-11-11 01:47:33.406Z: #12 70.88 golangci/golangci-lint info checking GitHub for latest tag
2022-11-11 01:47:33.884Z: #12 71.27 golangci/golangci-lint info found version: 1.50.1 for v1.50.1/linux/amd64
2022-11-11 01:47:34.922Z: #12 72.38 golangci/golangci-lint info installed /go/bin/golangci-lint
2022-11-11 01:51:55.806Z: #12 333.4 Done!
*/