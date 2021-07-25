load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "util",
    srcs = ["daoutil.go"],
    importpath = "github.com/tedsilb/twooter/backend/dao/util",
    visibility = ["//visibility:public"],
    deps = ["@org_golang_google_protobuf//types/known/timestamppb:go_default_library"],
)

go_test(
    name = "util_test",
    srcs = ["daoutil_test.go"],
    embed = [":util"],
)
