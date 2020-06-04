# bdpass

百度网盘秒传链接生成工具，支持梦姬标准/PanDownload/BaiduPCS-Go三种格式，配合 **[油猴脚本](https://greasyfork.org/zh-CN/scripts/397324-秒传链接提取)** 使用。

[![Actions Status](https://img.shields.io/github/workflow/status/winterssy/bdpass/Build/master?logo=appveyor)](https://github.com/winterssy/bdpass/actions)

## 快速开始

```sh
$ bdpass 文件路径
# 批量生成
$ bdpass 文件路径A 文件路径B 文件路径C...
```

## 可选参数

- -f (--format)

指定输出的编码格式，`std` 为梦姬标准（默认），`pdl` 为 PanDownload，`pcs` 为 BaiduPCS-Go，如：

```sh
# 输出为 PanDownload 格式
$ bdpass -f pdl 文件路径
# 输出为 BaiduPCS-Go 格式
$ bdpass -f pcs 文件路径
```

## License

GPLv3。