#!/bin/bash

set -e

# 下载 Alist 二进制文件
echo "Downloading Alist binary..."
wget -O alist.tar.gz https://github.com/alist-org/alist/releases/download/v3.35.0/alist-linux-amd64.tar.gz
tar -zxvf alist.tar.gz


# 确保二进制文件有执行权限
chmod +x alist

# 创建 dist 目录并移动二进制文件
mkdir -p dist
mv alist dist/
