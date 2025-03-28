package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Resource 定义资源结构体
type Resource struct {
	URI         string `json:"uri"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
}

// 这里定义一个 URL 资源，指向外部帮助指南
var resources = []Resource{
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/app-info.html",
		Name:        "🌈 了解阅读记录app",
		Description: "阅读记录App使用方法",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/reading-suggest.html",
		Name:        "💁🏻‍♂️ 如何培养阅读习惯",
		Description: "如何培养阅读习惯",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/add-delete-edit-book.html",
		Name:        "📖 添加/编辑书籍",
		Description: "添加或编辑书籍信息",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/delete-edit-book.html",
		Name:        "⛔️ 删除书籍",
		Description: "删除书籍操作指南",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/note_review.html",
		Name:        "📕书摘每日回顾",
		Description: "书摘每日回顾功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/wxread_import.html",
		Name:        "📲 [微信读书] 授权导入指南",
		Description: "微信读书笔记导入指南",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/auto_complete_record.html",
		Name:        "🛠️ 自动补全阅读进度",
		Description: "自动补全阅读进度功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/creat-delete-bookshelf.html",
		Name:        "🗄 创建/删除书架",
		Description: "创建或删除书架操作指南",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/add_widget_android.html",
		Name:        "📱 安卓如何添加小组件",
		Description: "安卓小组件添加指南",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/desktop-window-ios.html",
		Name:        "🖼️ iOS-桌面窗口",
		Description: "iOS桌面窗口功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/desktop-window.html",
		Name:        "🖼️ 安卓-悬浮窗",
		Description: "安卓悬浮窗功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/suit-guide-intro.html",
		Name:        "📚 套装书说明",
		Description: "套装书功能说明",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/suit-page-guide-intro.html",
		Name:        "📝 套装书页码说明",
		Description: "套装书页码功能说明",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/bookInfo_read_data.html",
		Name:        "📊 阅读数据说明",
		Description: "阅读数据功能说明",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/statistics-read_words.html",
		Name:        "📝 字数统计说明",
		Description: "字数统计功能说明",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/update_read_words.html",
		Name:        "📊 更新阅读字数信息",
		Description: "更新阅读字数信息功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/book_progress_intro.html",
		Name:        "🚫 特殊情况下【进度信息】说明",
		Description: "特殊情况下的进度信息说明",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/book-detail-analyse-intro.html",
		Name:        "📈 数据分析说明",
		Description: "数据分析功能说明",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/widget.html",
		Name:        "📱 小组件",
		Description: "小组件功能说明",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/book-tag.html",
		Name:        "🏷️ 使用标签",
		Description: "标签功能说明",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/delete-change-record.html",
		Name:        "📋 删除/修改阅读记录",
		Description: "删除或修改阅读记录操作指南",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/manually-input-record.html",
		Name:        "📝 手动输入记录",
		Description: "手动输入记录功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/note_export.html",
		Name:        "📤 导出书摘",
		Description: "导出书摘功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/wechat-note-import.html",
		Name:        "💬 导入微信读书笔记",
		Description: "导入微信读书笔记功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/kindle-note-import.html",
		Name:        "📥 导入kindle笔记",
		Description: "导入Kindle笔记功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/author-publisher.html",
		Name:        "✍️ 作者/出版社管理",
		Description: "作者或出版社管理功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/vipFunComparison.html",
		Name:        "✍👑 会员权益对比",
		Description: "会员权益对比功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/iwatch.html",
		Name:        "⌚️ 手表相关",
		Description: "手表相关功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/pay-FAQ.html",
		Name:        "💳 支付/会员相关",
		Description: "支付或会员相关问题",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/android-notification.html",
		Name:        "📣 安卓通知提醒问题",
		Description: "安卓通知提醒问题",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/wrong-operation-did.html",
		Name:        "🙅🏻‍♂️ 误操作导致已读完",
		Description: "误操作导致已读完问题",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/Invite-cashback.html",
		Name:        "💰 邀请返现",
		Description: "邀请返现功能",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/user-agreement.html",
		Name:        "👤 用户协议",
		Description: "用户协议详细内容",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/ydjl_privacy.html",
		Name:        "🔑 隐私政策",
		Description: "隐私政策详细内容",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/cancel_account.html",
		Name:        "🙋🏻‍♂️ 注销须知",
		Description: "用户注销流程",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/about_me.html",
		Name:        "💁🏻‍♂️ 关于我们",
		Description: "关于我们",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/product_story.html",
		Name:        "📲 产品故事",
		Description: "产品故事",
		MimeType:    "text/html",
	},
	{
		URI:         "https://use-guide.yidiansz.com/pages/used-method/contact-us.html",
		Name:        "☎️ 联系我们",
		Description: "联系方式",
		MimeType:    "text/html",
	},
}

func main() {
	// 创建 MCP 服务器
	s := server.NewMCPServer("Reading Support", "1.0.0")

	// 注册列出资源的工具
	listTool := mcp.NewTool("get_help_articles",
		mcp.WithDescription("获取所有帮助文章列表"),
	)
	s.AddTool(listTool, listResourcesHandler)

	// 注册读取资源内容的工具，要求传入参数 uri
	readTool := mcp.NewTool("open_article",
		mcp.WithDescription("根据ID打开帮助文章"),
		mcp.WithString("uri", mcp.Required(), mcp.Description("帮助文章的唯一URL")),
	)
	s.AddTool(readTool, readResourceHandler)

	// 启动 stdio 服务器（实际部署可使用网络传输）
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("服务器错误: %v\n", err)
	}
}

// listResourcesHandler 用于列出所有资源信息
func listResourcesHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var result string
	for _, res := range resources {
		result += fmt.Sprintf("URI: %s\n名称: %s\n描述: %s\n\n", res.URI, res.Name, res.Description)
	}
	return mcp.NewToolResultText(result), nil
}

// readResourceHandler 根据传入的 URI 读取资源内容
func readResourceHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	uri, ok := req.Params.Arguments["uri"].(string)
	if !ok {
		return nil, fmt.Errorf("参数 uri 无效")
	}

	// 查找资源是否存在
	var target *Resource
	for _, res := range resources {
		if res.URI == uri {
			target = &res
			break
		}
	}
	if target == nil {
		return nil, fmt.Errorf("未找到资源: %s", uri)
	}

	// 如果 URI 以 "http" 开头，则通过 HTTP GET 请求获取内容
	if len(uri) >= 4 && uri[:4] == "http" {
		resp, err := http.Get(uri)
		if err != nil {
			return nil, fmt.Errorf("获取资源失败: %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("读取响应失败: %v", err)
		}
		result := fmt.Sprintf("资源名称: %s\n内容:\n%s", target.Name, string(body))
		return mcp.NewToolResultText(result), nil
	}

	// 其它情况：返回错误提示（这里仅处理 URL 类型）
	return nil, fmt.Errorf("不支持的 URI 类型: %s", uri)
}
