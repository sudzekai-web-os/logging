module github.com/sudzekai/web-os-api/packages/logging

go 1.26.8

require (
    github.com/sudzekai/web-os-api/packages/abstractions v0.0.0 //direct
    github.com/sudzekai/web-os-api/packages/types v0.0.0 //direct
)

replace github.com/sudzekai/web-os-api/packages/types => ../types
replace github.com/sudzekai/web-os-api/packages/abstractions => ../abstractions