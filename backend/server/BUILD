load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "server_lib",
    srcs = ["server.go"],
    importpath = "github.com/tedsilb/twooter/backend/server",
    visibility = ["//visibility:private"],
    deps = [
        "//proto:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":server_lib"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "server_test",
    srcs = ["server_test.go"],
    embed = [":server_lib"],
)
