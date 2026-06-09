# Git Trending Rank

GitHub 热门仓库趋势排行榜，自动抓取 GitHub Trending 数据并生成静态站点。

[https://git-trending-rank.github.io/](https://git-trending-rank.github.io/)

## 数据

- **Daily / Weekly / Monthly** 三个时间维度的趋势仓库
- 自动更新，每 2 小时抓取一次
- 支持评论（Gitalk），可在页面中对仓库进行讨论

## 技术栈

| 组件 | 技术 |
|---|---|
| 爬虫 | Go + goquery |
| 站点 | Hugo + hugo-theme-stack |
| 自动构建 | GitHub Actions |
| 托管 | GitHub Pages |

## 运行

```bash
# 初始化
make init

# 本地构建
make build

# 本地预览
make server
```

打开 http://localhost:1313 即可查看。
