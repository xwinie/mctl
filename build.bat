gox -osarch="windows/amd64" -ldflags "-s -w " -gcflags -m -output gobatisctl
upx -f -9 gobatisctl.exe