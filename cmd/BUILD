load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "root",
    srcs = ["root.go"],
    visibility = ["//visibility:public"],
    deps = [
        ":count",
        "@com_github_spf13_cobra//:cobra",
    ],
)

go_library(
    name = "count",
    srcs = ["count.go"],
    visibility = ["//visibility:public"],
    importpath = "github.com/bexxmodd/wcnt/blob/master/cmd/count",
    deps = [
        "@com_github_spf13_cobra//:cobra",
    ],
)

go_test(
    name = "count_test",
    srcs = ["count_test.go"],
    deps = [
        ":count",
    ]
)
