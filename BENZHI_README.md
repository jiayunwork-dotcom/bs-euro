Go 欧式期权定价命令行工具：按现货、执行价、到期期限、无风险利率和波动率计算 Black-Scholes 价格、d1、d2 以及 Delta/Gamma/Vega/Theta/Rho；命令行用 price / greeks 读 JSON，也可用 serve 在 :8080 提供 /api/price 与 /api/greeks。

## 构建 / 运行 / 测试

```text
go build ./...
go run . price --table example/atm-1y.json
go run . serve --http :8080
go test ./...
```

## 评测镜像

本目录评测专用文件（勿覆盖项目自带 Dockerfile/README）：

- `benzhi.Dockerfile`
- `build_benzhi_docker.sh`
- `BENZHI_README.md`（本文件）

两种架构都要构建并进容器验证：

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh <image-name> linux/arm64
./build_benzhi_docker.sh <image-name> linux/amd64
docker run -it <image-name>:latest
```
