bazel help

bazel build //:hello-world

bazel query  --notool_deps --noimplicit_deps "deps(//:hello-world)" --output graph

bazel clean --expunge


