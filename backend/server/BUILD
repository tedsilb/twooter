load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "server_lib",
    srcs = ["server.go"],
    importpath = "github.com/tedsilb/twooter/backend/server",
    visibility = ["//visibility:private"],
    deps = [
        "//backend/dao/twoot",
        "//backend/validate:validate_lib",
        "//proto:twooterpb",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_test(
    name = "server_test",
    srcs = ["server_test.go"],
    embed = [":server_lib"],
    deps = [
        "//backend/dao/twoot",
        "//backend/dao/util",
        "//proto:twooterpb",
        "@org_golang_google_protobuf//types/known/timestamppb:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":server_lib"],
    visibility = ["//visibility:public"],
)
