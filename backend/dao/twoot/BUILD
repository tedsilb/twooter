load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "twoot",
    srcs = ["twootdao.go"],
    importpath = "github.com/tedsilb/twooter/backend/dao/twoot",
    visibility = ["//visibility:public"],
    deps = [
        "//backend/dao/util",
        "//proto:twooterpb",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)

go_test(
    name = "twoot_test",
    srcs = ["twootdao_test.go"],
    embed = [":twoot"],
    deps = [
        "//proto:twooterpb",
        "@org_golang_google_protobuf//proto:go_default_library",
    ],
)
