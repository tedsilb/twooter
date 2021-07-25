load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "validate_lib",
    srcs = ["validate.go"],
    importpath = "github.com/tedsilb/twooter/backend/validate",
    visibility = ["//backend:__subpackages__"],
    deps = [
        "//proto:twooterpb",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)
