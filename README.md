# bazel-projects
Testing out bazel builds
touch .bazelversion
vim .bazelversion 
touch WORKSPACE.bazel
bazel --version
bazel build //...
touch BUILD.bazel
bazel build //...