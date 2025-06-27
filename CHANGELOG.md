## 4. 使用说明

### 4.1 其他项目安装使用
其他Go项目可以通过以下方式安装你的包：

```bash
# 安装最新版本
go get github.com/cloud-evan/gocreem

# 安装特定版本
go get github.com/cloud-evan/gocreem@v1.0.0
```

### 4.2 在其他项目中使用
```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/cloud-evan/gocreem"
)

func main() {
    client, err := creem.NewClient("your_api_key", "your_secret_key", true)
    if err != nil {
        log.Fatal(err)
    }
    
    ctx := context.Background()
    
    // 使用SDK...
}
```

## 5. 后续维护

### 5.1 版本更新流程
1. 修改代码
2. 更新CHANGELOG.md
3. 提交更改
4. 创建新标签：`git tag -a v1.1.0 -m "Release version 1.1.0"`
5. 推送标签：`git push origin v1.1.0`
6. 在GitHub上创建新的Release

### 5.2 持续集成建议
考虑添加GitHub Actions来自动化测试和发布流程。

## 注意事项

1. **模块路径**：确保go.mod中的模块路径与GitHub仓库路径一致
2. **版本管理**：使用语义化版本控制（Semantic Versioning）
3. **文档更新**：每次发布新版本时更新README和CHANGELOG
4. **测试覆盖**：确保所有功能都有相应的测试
5. **依赖管理**：定期更新依赖包版本

完成这些步骤后，你的Creem Go SDK就可以作为第三方包被其他Go项目安装和使用了！ 