load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "root",
    srcs = ["root.go"],
    importpath = "wcnt/cmd/root",
    visibility = ["//visibility:public"],
    deps = [
        "//cmd:count",
        "@com_github_spf13_cobra//:cobra",
    ],
)

go_library(
    name = "count",
    srcs = ["count.go"],
    importpath = "wcnt/cmd/count",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_spf13_cobra//:cobra",
    ],
)
