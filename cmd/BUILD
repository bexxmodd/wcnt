load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "root",
    srcs = ["root.go"],
    importpath = "wcnt/cmd",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_spf13_cobra//:cobra",
    ],
)