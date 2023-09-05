# wild-workouts-go-ddd-walkthrough
Three Dots Labs 《Too modern》系列文章项目手抄练习

源项目: [Wild Workouts](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example)

系列文章在源项目 Readme 中有链接，并附有文章对应代码版本 tag

---
原版的前端代码跑不起来，花了几天时间，使用 vue3 + bootstrap5 + typescript 重新实现了。
Toast + dialog 使用了 element-plus 的组件。

实现过程中发现这个项目有时区处理不明确的问题。前端展示使用，必须转换为西四时区时间。通过引入 Day.js 来解决。

---
项目代码还有一些其他小毛病，但这个项目主要是讲架构的，请先不清理了。

后面有时间，把对应文章翻译下发 blog，再进再从头实现一遍，作为翻译文档的配套代码。