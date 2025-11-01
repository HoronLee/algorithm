param(
    # 添加一个 -Clean 开关参数来模拟 'clean' target
    [switch]$Clean,

    # 添加一个 -Build 开关参数来模拟 'build' target (或设为默认)
    [switch]$Build
)

# 清理函数
function Invoke-Clean {
    Write-Host "Cleaning build artifacts..."
    Get-ChildItem -Path . -Recurse -File -Include "*.o", "*.out", "main.exe" | Remove-Item -Force -ErrorAction SilentlyContinue
    Write-Host "Clean complete."
}

# 构建函数
function Invoke-Build {
    Write-Host "Building project..."
    # ... 在这里放置您的构建命令 (例如 go build) ...
}

# --- 脚本执行逻辑 ---

if ($Clean) {
    Invoke-Clean
}
elseif ($Build) {
    Invoke-Build
}
else {
    # 默认行为 (如果没有指定开关, 就像运行 'make')
    Invoke-Build
}
